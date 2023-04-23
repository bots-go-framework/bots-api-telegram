// Package tgbotapi has functions and types used for interacting with
// the Telegram Bot API.
package tgbotapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"github.com/strongo/log"
	"github.com/technoweenie/multipartstreamer"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

// BotAPI allows you to interact with the Telegram Bot API.
type BotAPI struct {
	Token  string          `json:"token"`
	Self   User            `json:"-"`
	Client *http.Client    `json:"-"`
	c      context.Context // TODO: Wrong? read docs on Context class
}

// EnableDebug enables debugging
func (bot *BotAPI) EnableDebug(c context.Context) {
	bot.c = c
}

// NewBotAPI creates a new BotAPI instance.
//
// It requires a token, provided by @BotFather on Telegram.
func NewBotAPI(token string) *BotAPI {
	return NewBotAPIWithClient(token, &http.Client{})
}

// NewBotAPIWithClient creates a new BotAPI instance
// and allows you to pass a http.Client.
//
// It requires a token, provided by @BotFather on Telegram.
func NewBotAPIWithClient(token string, client *http.Client) *BotAPI {
	if strings.TrimSpace(token) == "" {
		panic("token must not be empty")
	}
	return &BotAPI{
		Token:  token,
		Client: client,
	}
}

// MakeRequestFromChattable makes request from chattable TODO: Is duplicate of Send()?
func (bot *BotAPI) MakeRequestFromChattable(m Chattable) (resp APIResponse, err error) { //
	values, err := m.Values()
	if err != nil {
		return resp, err
	}
	return bot.MakeRequest(m.method(), values)
}

// MakeRequest makes a request to a specific endpoint with our token.
func (bot *BotAPI) MakeRequest(endpoint string, params url.Values) (apiResp APIResponse, err error) {
	method := fmt.Sprintf(APIEndpoint, bot.Token, endpoint)

	var resp *http.Response

	var hadDeadlineExceeded bool

	for i := 1; i <= 2; i++ { // TODO: Should this be in bots framework?
		if resp, err = bot.Client.PostForm(method, params); err != nil {
			if strings.Contains(err.Error(), "DEADLINE_EXCEEDED") {
				hadDeadlineExceeded = true
				log.Warningf(bot.c, "#%v fail to send POST due to DEADLINE_EXCEEDED to %v, will retry: %v", i, method, err)
				continue
			}
		}
		break
	}
	if resp != nil && resp.Body != nil {
		defer func() {
			if err := resp.Body.Close(); err != nil {
				log.Warningf(bot.c, "failed to close response body: %v", err)
			}
		}()
	}

	if err != nil {
		log.Errorf(bot.c, "Failed to send POST to %v: %v", method, err.Error())
		return APIResponse{Ok: false}, fmt.Errorf("%v: %s: %w", "POST", method, err)
	}

	var body []byte
	if resp.ContentLength > 0 {
		var readerErr error
		if body, readerErr = io.ReadAll(resp.Body); err != nil {
			log.Errorf(bot.c, "Failed to read response.body: %v", readerErr)
			if err == nil {
				err = readerErr
			}
		}
	}
	apiResp = APIResponse{
		Result: json.RawMessage(body),
	}
	if resp.StatusCode >= 300 {
		apiResp.ErrorCode = resp.StatusCode
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return apiResp, fmt.Errorf("%s %+v unauthorized: %s", method, params, string(body))
	case http.StatusForbidden:
		return apiResp, &ErrAPIForbidden{}
	}

	if err != nil {
		return APIResponse{Ok: false, Result: json.RawMessage(body)}, fmt.Errorf("%v: %s: %w", "POST", method, err)
	}

	logRequestAndResponse := func() {
		if bot.c != nil {
			log.Debugf(bot.c, "Request to Telegram API: %v => %v", endpoint, params)
			log.Debugf(bot.c, "Telegram API response: %v", string(body))
		}
	}

	//logRequestAndResponse()

	if err = ffjson.Unmarshal(body, &apiResp); err != nil {
		logRequestAndResponse()
		return apiResp, fmt.Errorf("telegram API returned non JSON response or unknown JSON: %w:\n%s", err, string(body))
	} else if !apiResp.Ok {
		logRequestAndResponse()
		if hadDeadlineExceeded && apiResp.ErrorCode == 400 && strings.Contains(apiResp.Description, "message is not modified") {
			return apiResp, nil
		}
		return apiResp, apiResp
	}
	return apiResp, nil
}

// makeMessageRequest makes a request to a method that returns a Message.
func (bot *BotAPI) makeMessageRequest(endpoint string, params url.Values) (Message, error) {
	resp, err := bot.MakeRequest(endpoint, params)
	var message Message

	if err != nil {
		return message, err
	}

	if !resp.Ok || resp.ErrorCode != 0 {
		return message, resp
	}

	if string(resp.Result) != "true" { // TODO: This is a workaround for "answerCallbackQuery" that returns just "true".
		if err = ffjson.Unmarshal(resp.Result, &message); err != nil {
			return message, fmt.Errorf("failed to call json.Unmarshal(s): %w: s=%s", err, string(resp.Result))
		}
	}
	return message, err
}

// UploadFile makes a request to the API with a file.
//
// Requires the parameter to hold the file not be in the params.
// File should be a string to a file path, a FileBytes struct,
// or a FileReader struct.
//
// Note that if your FileReader has a size set to -1, it will read
// the file into memory to calculate a size.
func (bot *BotAPI) UploadFile(endpoint string, params map[string]string, fieldname string, file interface{}) (APIResponse, error) {
	ms := multipartstreamer.New()
	err := ms.WriteFields(params)
	if err != nil {
		return APIResponse{}, err
	}

	switch f := file.(type) {
	case string:
		fileHandle, err := os.Open(f)
		if err != nil {
			return APIResponse{}, err
		}
		defer func() {
			_ = fileHandle.Close()
		}()

		fi, err := os.Stat(f)
		if err != nil {
			return APIResponse{}, err
		}

		if err = ms.WriteReader(fieldname, fileHandle.Name(), fi.Size(), fileHandle); err != nil {
			return APIResponse{}, err
		}
	case FileBytes:
		buf := bytes.NewBuffer(f.Bytes)
		if err = ms.WriteReader(fieldname, f.Name, int64(len(f.Bytes)), buf); err != nil {
			return APIResponse{}, err
		}
	case FileReader:
		if f.Size != -1 {
			if err = ms.WriteReader(fieldname, f.Name, f.Size, f.Reader); err != nil {
				return APIResponse{}, err
			}
			break
		}

		data, err := io.ReadAll(f.Reader)
		if err != nil {
			return APIResponse{}, err
		}

		buf := bytes.NewBuffer(data)

		if err = ms.WriteReader(fieldname, f.Name, int64(len(data)), buf); err != nil {
			return APIResponse{}, err
		}
	default:
		return APIResponse{}, errors.New(ErrBadFileType)
	}

	method := fmt.Sprintf(APIEndpoint, bot.Token, endpoint)

	req, err := http.NewRequest("POST", method, nil)
	if err != nil {
		return APIResponse{}, err
	}

	ms.SetupRequest(req)

	res, err := bot.Client.Do(req)
	if err != nil {
		return APIResponse{}, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return APIResponse{}, err
	}

	if bot.c != nil {
		log.Debugf(bot.c, string(body))
	}

	var apiResp APIResponse
	if err = ffjson.Unmarshal(body, &apiResp); err != nil {
		return apiResp, nil
	}

	if !apiResp.Ok {
		return APIResponse{}, errors.New(apiResp.Description)
	}

	return apiResp, nil
}

// GetFileDirectURL returns direct URL to file
//
// It requires the FileID.
func (bot *BotAPI) GetFileDirectURL(fileID string) (string, error) {
	file, err := bot.GetFile(FileConfig{fileID})

	if err != nil {
		return "", err
	}

	return file.Link(bot.Token), nil
}

// GetMe fetches the currently authenticated bot.
//
// This method is called upon creation to validate the token,
// and so you may get this data from BotAPI.Self without the need for
// another request.
func (bot *BotAPI) GetMe() (User, error) {
	var user User

	resp, err := bot.MakeRequest("getMe", nil)
	if err != nil {
		return user, err
	}

	if err = ffjson.Unmarshal(resp.Result, &user); err != nil {
		return user, err
	}

	bot.debugLog("getMe", nil, user)

	return user, nil
}

func (bot *BotAPI) GetChat(chatID string) (Chat, error) {
	var chat Chat

	resp, err := bot.MakeRequest("getChat", url.Values{"chat_id": []string{chatID}})
	if err != nil {
		return chat, err
	}

	if err = ffjson.Unmarshal(resp.Result, &chat); err != nil {
		return chat, err
	}

	bot.debugLog("getChat", nil, chat)

	return chat, nil
}

// IsMessageToMe returns true if message directed to this bot.
//
// It requires the Message.
func (bot *BotAPI) IsMessageToMe(message Message) bool {
	return strings.Contains(message.Text, "@"+bot.Self.UserName)
}

// Send will send a Chattable item to Telegram.
//
// It requires the Chattable to send.
func (bot *BotAPI) Send(c Chattable) (Message, error) {
	switch t := c.(type) {
	case Fileable:
		return bot.sendFile(t)
	default:
		return bot.sendChattable(t)
	}
}

// debugLog checks if the bot is currently running in debug mode, and if
// so will display information about the request and response in the
// debug log.
func (bot *BotAPI) debugLog(context string, v url.Values, message interface{}) {
	if bot.c != nil {
		log.Debugf(bot.c, "%s req : %+v\n", context, v)
		log.Debugf(bot.c, "%s resp: %+v\n", context, message)
	}
}

// sendExisting will send a Message with an existing file to Telegram.
func (bot *BotAPI) sendExisting(method string, config Fileable) (Message, error) {
	v, err := config.Values()

	if err != nil {
		return Message{}, err
	}

	message, err := bot.makeMessageRequest(method, v)
	if err != nil {
		return Message{}, err
	}

	return message, nil
}

// uploadAndSend will send a Message with a new file to Telegram.
func (bot *BotAPI) uploadAndSend(method string, config Fileable) (Message, error) {
	var message Message

	params, err := config.params()
	if err != nil {
		return message, err
	}

	file := config.getFile()

	resp, err := bot.UploadFile(method, params, config.name(), file)
	if err != nil {
		return message, err
	}

	if err = ffjson.Unmarshal(resp.Result, &message); err != nil {
		return message, err
	}

	bot.debugLog(method, nil, message)

	return message, nil
}

// sendFile determines if the file is using an existing file or uploading
// a new file, then sends it as needed.
func (bot *BotAPI) sendFile(config Fileable) (Message, error) {
	if config.useExistingFile() {
		return bot.sendExisting(config.method(), config)
	}

	return bot.uploadAndSend(config.method(), config)
}

// sendChattable sends a Chattable.
func (bot *BotAPI) sendChattable(config Chattable) (Message, error) {
	v, err := config.Values()
	if err != nil {
		return Message{}, err
	}

	return bot.makeMessageRequest(config.method(), v)
}

// GetUserProfilePhotos gets a user's profile photos.
//
// It requires UserID.
// Offset and Limit are optional.
func (bot *BotAPI) GetUserProfilePhotos(config UserProfilePhotosConfig) (UserProfilePhotos, error) {
	var profilePhotos UserProfilePhotos

	v := url.Values{}
	v.Add("user_id", strconv.Itoa(config.UserID))
	if config.Offset != 0 {
		v.Add("offset", strconv.Itoa(config.Offset))
	}
	if config.Limit != 0 {
		v.Add("limit", strconv.Itoa(config.Limit))
	}

	resp, err := bot.MakeRequest("getUserProfilePhotos", v)
	if err != nil {
		return profilePhotos, err
	}

	if err = ffjson.Unmarshal(resp.Result, &profilePhotos); err != nil {
		return profilePhotos, err
	}

	bot.debugLog("GetUserProfilePhoto", v, profilePhotos)

	return profilePhotos, nil
}

// GetFile returns a File which can download a file from Telegram.
//
// Requires FileID.
func (bot *BotAPI) GetFile(config FileConfig) (File, error) {
	var file File

	v := url.Values{}
	v.Add("file_id", config.FileID)

	resp, err := bot.MakeRequest("getFile", v)
	if err != nil {
		return file, err
	}

	if err = ffjson.Unmarshal(resp.Result, &file); err != nil {
		return file, err
	}

	bot.debugLog("GetFile", v, file)

	return file, nil
}

// GetUpdates fetches updates.
// If a WebHook is set, this will not return any data!
//
// Offset, Limit, and Timeout are optional.
// To avoid stale items, set Offset to one higher than the previous item.
// Set Timeout to a large number to reduce requests so you can get updates
// instantly instead of having to wait between requests.
func (bot *BotAPI) GetUpdates(config *UpdateConfig) ([]Update, error) {
	var updates []Update

	v := url.Values{}
	if config.Offset > 0 {
		v.Add("offset", strconv.Itoa(config.Offset))
	}
	if config.Limit > 0 {
		v.Add("limit", strconv.Itoa(config.Limit))
	}
	if config.Timeout > 0 {
		v.Add("timeout", strconv.Itoa(config.Timeout))
	}

	resp, err := bot.MakeRequest("getUpdates", v)
	if err != nil {
		return updates, err
	}

	if err = ffjson.Unmarshal(resp.Result, &updates); err != nil {
		return updates, err
	}

	bot.debugLog("getUpdates", v, updates)

	return updates, nil
}

// RemoveWebhook unsets the webhook.
func (bot *BotAPI) RemoveWebhook() (APIResponse, error) {
	return bot.MakeRequest("setWebhook", url.Values{})
}

// SetWebhook sets a webhook.
//
// If this is set, GetUpdates will not get any data!
//
// If you do not have a legitimate TLS certificate, you need to include your self-signed certificate with the config.
func (bot *BotAPI) SetWebhook(config WebhookConfig) (APIResponse, error) {
	if config.Certificate == nil {
		params, err := config.Values()
		if err != nil {
			return APIResponse{}, err
		}
		return bot.MakeRequest("setWebhook", params)
	} else {
		var apiResp APIResponse
		resp, err := bot.UploadFile("setWebhook", map[string]string{"url": config.URL.String()}, "certificate", config.Certificate)
		if err != nil {
			return apiResp, err
		}

		if err = ffjson.Unmarshal(resp.Result, &apiResp); err != nil {
			return apiResp, err
		}

		if bot.c != nil {
			log.Debugf(bot.c, "setWebhook resp: %+v\n", apiResp)
		}

		return apiResp, nil
	}
}

// GetUpdatesChan starts and returns a channel for getting updates.
func (bot *BotAPI) GetUpdatesChan(config *UpdateConfig) (<-chan Update, error) {
	updatesChan := make(chan Update, 100)

	go func() {
		for {
			updates, err := bot.GetUpdates(config)
			if err != nil {
				//log.Println(err)
				//log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}

			for _, update := range updates {
				if update.UpdateID >= config.Offset {
					config.Offset = update.UpdateID + 1
					updatesChan <- update
				}
			}
		}
	}()

	return updatesChan, nil
}

// ListenForWebhook registers a http handler for a webhook.
func (bot *BotAPI) ListenForWebhook(pattern string) <-chan Update {
	updatesChan := make(chan Update, 100)

	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)

		var update Update
		if err := ffjson.Unmarshal(body, &update); err != nil {
			log.Errorf(context.Background(), fmt.Errorf("failed to unmarshal update JSON: %w", err).Error())
			return
		}

		updatesChan <- update
	})

	return updatesChan
}

// AnswerInlineQuery sends a response to an inline query.
//
// Note that you must respond to an inline query within 30 seconds.
func (bot *BotAPI) AnswerInlineQuery(config InlineConfig) (APIResponse, error) {
	v := url.Values{}

	v.Add("inline_query_id", config.InlineQueryID)
	v.Add("cache_time", strconv.Itoa(config.CacheTime))
	v.Add("is_personal", strconv.FormatBool(config.IsPersonal))
	v.Add("next_offset", config.NextOffset)
	data, err := ffjson.Marshal(config.Results)
	if err != nil {
		ffjson.Pool(data)
		return APIResponse{}, err
	}
	v.Add("results", string(data))
	ffjson.Pool(data)
	v.Add("switch_pm_text", config.SwitchPMText)
	v.Add("switch_pm_parameter", config.SwitchPMParameter)

	bot.debugLog("answerInlineQuery", v, nil)

	return bot.MakeRequest("answerInlineQuery", v)
}

// KickChatMember kicks a user from a chat. Note that this only will work
// in supergroups, and requires the bot to be an admin. Also note they
// will be unable to rejoin until they are unbanned.
func (bot *BotAPI) KickChatMember(config ChatMemberConfig) (APIResponse, error) {
	v := url.Values{}

	if config.SuperGroupUsername == "" {
		v.Add("chat_id", strconv.FormatInt(config.ChatID, 10))
	} else {
		v.Add("chat_id", config.SuperGroupUsername)
	}
	v.Add("user_id", strconv.Itoa(config.UserID))

	bot.debugLog("kickChatMember", v, nil)

	return bot.MakeRequest("kickChatMember", v)
}

// UnbanChatMember unbans a user from a chat. Note that this only will work
// in supergroups, and requires the bot to be an admin.
func (bot *BotAPI) UnbanChatMember(config ChatMemberConfig) (APIResponse, error) {
	v := url.Values{}

	if config.SuperGroupUsername == "" {
		v.Add("chat_id", strconv.FormatInt(config.ChatID, 10))
	} else {
		v.Add("chat_id", config.SuperGroupUsername)
	}
	v.Add("user_id", strconv.Itoa(config.UserID))

	bot.debugLog("unbanChatMember", v, nil)

	return bot.MakeRequest("unbanChatMember", v)
}
