// Package tgbotapi has functions and types used for interacting with
// the Telegram Bot API.
package tgbotapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/strongo/logus"
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

// telegramRequestError preserves the transport error for errors.Is/errors.As
// without exposing the token-bearing Telegram endpoint through Error().
type telegramRequestError struct {
	method string
	err    error
}

func (e telegramRequestError) Error() string {
	return fmt.Sprintf("Telegram API request %q failed", e.method)
}

func (e telegramRequestError) Unwrap() error {
	return e.err
}

// telegramResponseIOError preserves read failures for errors.Is/errors.As while
// preventing response readers from placing secrets in the public error string.
type telegramResponseIOError struct {
	method string
	err    error
}

func (e telegramResponseIOError) Error() string {
	return fmt.Sprintf("failed to read Telegram API response for method %q", e.method)
}

func (e telegramResponseIOError) Unwrap() error {
	return e.err
}

// telegramProviderError keeps the structured APIResponse available through
// errors.As without exposing Telegram's free-form Description by default.
type telegramProviderError struct {
	method   string
	response APIResponse
}

// TelegramProviderErrorDetails is the diagnostic subset of a Telegram API
// provider error. It deliberately excludes the bot token, request URL,
// request parameters, and raw response body, so it is suitable for structured
// operational logs.
//
// Description is supplied by Telegram. It is intentionally not included in
// Error(), which remains safe for ordinary error logs.
type TelegramProviderErrorDetails struct {
	Method      string
	ErrorCode   int
	Description string
}

// TelegramProviderErrorDetailsFrom extracts safe, structured diagnostics from
// an error returned by the Telegram API. It also works when that error has
// been wrapped with fmt.Errorf("...: %w", err).
//
// It returns false for transport, decoding, and non-Telegram provider errors.
func TelegramProviderErrorDetailsFrom(err error) (details TelegramProviderErrorDetails, ok bool) {
	var providerErr telegramProviderError
	if !errors.As(err, &providerErr) {
		return TelegramProviderErrorDetails{}, false
	}
	return TelegramProviderErrorDetails{
		Method:      providerErr.method,
		ErrorCode:   providerErr.response.ErrorCode,
		Description: providerErr.response.Description,
	}, true
}

func (e telegramProviderError) Error() string {
	return fmt.Sprintf(
		"Telegram API method %q failed with error code %d",
		e.method,
		e.response.ErrorCode,
	)
}

func (e telegramProviderError) Unwrap() error {
	return e.response
}

// EnableDebug enables metadata-only debugging. Request parameters, bot tokens,
// response bodies, and decoded provider objects are intentionally never logged.
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

// MakeRequestFromMessageWithValues makes request from WithValues
func (bot *BotAPI) MakeRequestFromMessageWithValues(method string, m WithValues) (resp APIResponse, err error) { //
	var values url.Values
	if values, err = m.Values(); err != nil {
		return resp, err
	}
	return bot.MakeRequest(method, values)
}

// MakeRequestFromChattable makes request from chattable TODO: Is duplicate of Send()?
func (bot *BotAPI) MakeRequestFromChattable(m Sendable) (resp APIResponse, err error) { //
	return bot.MakeRequestFromMessageWithValues(m.TelegramMethod(), m)
}

// SendRequest sends a request to a specific endpoint with our token and reads response.
func (bot *BotAPI) MakeRequest(telegramMethod string, params url.Values) (apiResp APIResponse, err error) {
	endpointURL := fmt.Sprintf(APIEndpoint, bot.Token, telegramMethod)

	var hadDeadlineExceeded bool
	var resp *http.Response

	for i := 1; i <= 2; i++ { // TODO: Should this be in bots framework?
		if resp, err = bot.Client.PostForm(endpointURL, params); err != nil {
			if strings.Contains(err.Error(), "DEADLINE_EXCEEDED") {
				hadDeadlineExceeded = true
				logus.Warningf(
					bot.c,
					"Telegram API request retry: attempt=%d, method=%q, error_type=%T",
					i,
					telegramMethod,
					err,
				)
				continue
			}
		}
		break
	}
	if resp != nil && resp.Body != nil {
		defer func() {
			if closeErr := resp.Body.Close(); closeErr != nil {
				logus.Warningf(bot.c, "failed to close Telegram API response body: error_type=%T", closeErr)
			}
		}()
	}

	if err != nil {
		logus.Errorf(bot.c, "Telegram API request failed: method=%q, error_type=%T", telegramMethod, err)
		return APIResponse{Ok: false}, telegramRequestError{method: telegramMethod, err: err}
	}

	var body []byte
	{
		var readerErr error
		if body, readerErr = io.ReadAll(resp.Body); readerErr != nil {
			logus.Errorf(
				bot.c,
				"Failed to read Telegram API response body: method=%q, error_type=%T",
				telegramMethod,
				readerErr,
			)
			err = telegramResponseIOError{method: telegramMethod, err: readerErr}
		}
	}
	apiResp = APIResponse{
		Result: body,
	}
	if resp.StatusCode >= 300 {
		apiResp.ErrorCode = resp.StatusCode
	}

	switch resp.StatusCode {
	case http.StatusUnauthorized:
		return apiResp, fmt.Errorf("telegram API method %q returned %s", telegramMethod, http.StatusText(resp.StatusCode))
	case http.StatusForbidden:
		return apiResp, &ErrAPIForbidden{}
	}

	if err != nil {
		return APIResponse{Ok: false, Result: body}, err
	}

	logRequestAndResponse := func() {
		if bot.c != nil {
			logus.Debugf(
				bot.c,
				"Telegram API request: method=%q, parameter_count=%d",
				telegramMethod,
				len(params),
			)
			logus.Debugf(
				bot.c,
				"Telegram API response: method=%q, status=%d, body_bytes=%d, ok=%t, error_code=%d",
				telegramMethod,
				resp.StatusCode,
				len(apiResp.Result),
				apiResp.Ok,
				apiResp.ErrorCode,
			)
		}
	}

	if err = json.Unmarshal(apiResp.Result, &apiResp); err != nil {
		logRequestAndResponse()
		return apiResp, fmt.Errorf("telegram API method %q returned invalid JSON: %w", telegramMethod, err)
	} else if !apiResp.Ok {
		logRequestAndResponse()
		if hadDeadlineExceeded && apiResp.ErrorCode == 400 && strings.Contains(apiResp.Description, "message is not modified") {
			return apiResp, nil
		}
		return apiResp, telegramProviderError{method: telegramMethod, response: apiResp}
	}

	return apiResp, nil
}

func (bot *BotAPI) DeleteMessage(chatID string, messageID int) (apiResp APIResponse, err error) {
	return bot.MakeRequest("deleteMessage", url.Values{"chat_id": {chatID}, "message_id": {strconv.Itoa(messageID)}})
}

// makeMessageRequest makes a request to a TelegramMethod that returns a Message.
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
		if err = json.Unmarshal(resp.Result, &message); err != nil {
			return message, fmt.Errorf("failed to decode Telegram API response for method %q: %w", endpoint, err)
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
func (bot *BotAPI) UploadFile(endpoint string, params map[string]string, fieldname string, file interface{}) (apiResp APIResponse, err error) {
	ms := multipartstreamer.New()
	if err = ms.WriteFields(params); err != nil {
		return
	}

	switch f := file.(type) {
	case string:
		var fileHandle *os.File
		if fileHandle, err = os.Open(f); err != nil {
			return
		}
		defer func() {
			_ = fileHandle.Close()
		}()

		var fi os.FileInfo
		if fi, err = os.Stat(f); err != nil {
			return
		}

		if err = ms.WriteReader(fieldname, fileHandle.Name(), fi.Size(), fileHandle); err != nil {
			return
		}
	case FileBytes:
		buf := bytes.NewBuffer(f.Bytes)
		if err = ms.WriteReader(fieldname, f.Name, int64(len(f.Bytes)), buf); err != nil {
			return
		}
	case FileReader:
		if f.Size != -1 {
			if err = ms.WriteReader(fieldname, f.Name, f.Size, f.Reader); err != nil {
				return
			}
			break
		}

		var data []byte
		if data, err = io.ReadAll(f.Reader); err != nil {
			return
		}

		buf := bytes.NewBuffer(data)

		if err = ms.WriteReader(fieldname, f.Name, int64(len(data)), buf); err != nil {
			return
		}
	default:
		err = ErrBadFileType
		return
	}

	method := fmt.Sprintf(APIEndpoint, bot.Token, endpoint)

	var req *http.Request
	if req, err = http.NewRequest("POST", method, nil); err != nil {
		return
	}

	ms.SetupRequest(req)

	var res *http.Response
	if res, err = bot.Client.Do(req); err != nil {
		return apiResp, telegramRequestError{method: endpoint, err: err}
	}
	defer func() {
		_ = res.Body.Close()
	}()

	var body []byte
	if body, err = io.ReadAll(res.Body); err != nil {
		return apiResp, telegramResponseIOError{method: endpoint, err: err}
	}

	if bot.c != nil {
		logus.Debugf(
			bot.c,
			"Telegram API upload response: method=%q, status=%d, body_bytes=%d",
			endpoint,
			res.StatusCode,
			len(body),
		)
	}

	if err = json.Unmarshal(body, &apiResp); err != nil {
		return
	}

	if !apiResp.Ok {
		return apiResp, telegramProviderError{method: endpoint, response: apiResp}
	}
	return
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
// This TelegramMethod is called upon creation to validate the token,
// and so you may get this data from BotAPI.Self without the need for
// another request.
func (bot *BotAPI) GetMe() (User, error) {
	var user User

	resp, err := bot.MakeRequest("getMe", nil)
	if err != nil {
		return user, err
	}

	if err = json.Unmarshal(resp.Result, &user); err != nil {
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

	if err = json.Unmarshal(resp.Result, &chat); err != nil {
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

// Send will send a Sendable item to Telegram.
//
// It requires the Sendable to send.
func (bot *BotAPI) Send(c Sendable) (Message, error) {
	switch t := c.(type) {
	case Fileable:
		return bot.sendFile(t)
	default:
		return bot.sendChattable(t)
	}
}

// debugLog emits only operational metadata. Provider payloads can contain bot
// tokens, user data, message text, callback data, and payment information.
func (bot *BotAPI) debugLog(operation string, v url.Values, message interface{}) {
	if bot.c != nil {
		logus.Debugf(
			bot.c,
			"Telegram API operation: method=%q, parameter_count=%d, response_type=%T",
			operation,
			len(v),
			message,
		)
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

	if err = json.Unmarshal(resp.Result, &message); err != nil {
		return message, err
	}

	bot.debugLog(method, nil, message)

	return message, nil
}

// sendFile determines if the file is using an existing file or uploading
// a new file, then sends it as needed.
func (bot *BotAPI) sendFile(config Fileable) (Message, error) {
	if config.useExistingFile() {
		return bot.sendExisting(config.TelegramMethod(), config)
	}

	return bot.uploadAndSend(config.TelegramMethod(), config)
}

// sendChattable sends a Sendable.
func (bot *BotAPI) sendChattable(config Sendable) (Message, error) {
	v, err := config.Values()
	if err != nil {
		return Message{}, err
	}

	return bot.makeMessageRequest(config.TelegramMethod(), v)
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

	if err = json.Unmarshal(resp.Result, &profilePhotos); err != nil {
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

	if err = json.Unmarshal(resp.Result, &file); err != nil {
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

	if err = json.Unmarshal(resp.Result, &updates); err != nil {
		return updates, err
	}

	bot.debugLog("getUpdates", v, updates)

	return updates, nil
}

// RemoveWebhook unsets the webhook.
func (bot *BotAPI) RemoveWebhook() (APIResponse, error) {
	return bot.MakeRequest("removeWebhook", url.Values{})
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

		if err = json.Unmarshal(resp.Result, &apiResp); err != nil {
			return apiResp, err
		}

		if bot.c != nil {
			logus.Debugf(
				bot.c,
				"Telegram API operation: method=%q, ok=%t, error_code=%d, result_bytes=%d",
				"setWebhook",
				apiResp.Ok,
				apiResp.ErrorCode,
				len(apiResp.Result),
			)
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
				//logus.Println(err)
				//logus.Println("Failed to get updates, retrying in 3 seconds...")
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
		if err := json.Unmarshal(body, &update); err != nil {
			logus.Errorf(context.Background(), fmt.Errorf("failed to unmarshal update JSON: %w", err).Error())
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
	if config.CacheTime > 0 {
		v.Add("cache_time", strconv.Itoa(config.CacheTime))
	}
	if config.IsPersonal {
		v.Add("is_personal", strconv.FormatBool(config.IsPersonal))
	}
	if config.NextOffset != "" {
		v.Add("next_offset", config.NextOffset)
	}

	data, err := encodeToJson(config.Results)
	if err != nil {
		return APIResponse{}, err
	}
	v.Add("results", string(data))

	if config.Button != nil {
		if data, err = encodeToJson(config.Button); err != nil {
			return APIResponse{}, err
		}
		v.Add("button", string(data))
	}

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

func (bot *BotAPI) SetDescription(config SetMyDescription) (APIResponse, error) {
	return bot.MakeRequestFromChattable(config)
}

func (bot *BotAPI) SetShortDescription(config SetMyShortDescription) (APIResponse, error) {
	return bot.MakeRequestFromChattable(config)
}

func (bot *BotAPI) SetCommands(config SetMyCommandsConfig) (APIResponse, error) {
	return bot.MakeRequestFromChattable(config)
}

func (bot *BotAPI) GetCommands(ctx context.Context, config GetMyCommandsConfig) (commands []TelegramBotCommand, err error) {
	err = bot.SendCustomMessage(ctx, config, &commands)
	return
}

// GetManagedBotToken returns the token of a managed bot.
//
// https://core.telegram.org/bots/api#getmanagedbottoken
func (bot *BotAPI) GetManagedBotToken(userID int64) (token string, err error) {
	v := url.Values{}
	v.Add("user_id", strconv.FormatInt(userID, 10))

	resp, err := bot.MakeRequest("getManagedBotToken", v)
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(resp.Result, &token); err != nil {
		return "", err
	}

	bot.debugLog("getManagedBotToken", v, token)

	return token, nil
}

// ReplaceManagedBotToken revokes the current token of a managed bot and generates a new one.
//
// https://core.telegram.org/bots/api#replacemanagedbottoken
func (bot *BotAPI) ReplaceManagedBotToken(userID int64) (token string, err error) {
	v := url.Values{}
	v.Add("user_id", strconv.FormatInt(userID, 10))

	resp, err := bot.MakeRequest("replaceManagedBotToken", v)
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(resp.Result, &token); err != nil {
		return "", err
	}

	bot.debugLog("replaceManagedBotToken", v, token)

	return token, nil
}

// SavePreparedKeyboardButton stores a keyboard button that can be used by a user within a Mini App.
//
// The button must be of type request_users, request_chat, or request_managed_bot.
//
// https://core.telegram.org/bots/api#savepreparedkeyboardbutton
func (bot *BotAPI) SavePreparedKeyboardButton(userID int64, button KeyboardButton) (prepared PreparedKeyboardButton, err error) {
	v := url.Values{}
	v.Add("user_id", strconv.FormatInt(userID, 10))

	data, err := encodeToJson(button)
	if err != nil {
		return prepared, err
	}
	v.Add("button", string(data))

	resp, err := bot.MakeRequest("savePreparedKeyboardButton", v)
	if err != nil {
		return prepared, err
	}

	if err = json.Unmarshal(resp.Result, &prepared); err != nil {
		return prepared, err
	}

	bot.debugLog("savePreparedKeyboardButton", v, prepared)

	return prepared, nil
}

// GetManagedBotAccessSettings returns the current access settings of a bot managed by the current bot.
//
// https://core.telegram.org/bots/api#getmanagedbotaccesssettings
func (bot *BotAPI) GetManagedBotAccessSettings(userID int64) (settings BotAccessSettings, err error) {
	v := url.Values{}
	v.Add("user_id", strconv.FormatInt(userID, 10))

	resp, err := bot.MakeRequest("getManagedBotAccessSettings", v)
	if err != nil {
		return settings, err
	}

	if err = json.Unmarshal(resp.Result, &settings); err != nil {
		return settings, err
	}

	bot.debugLog("getManagedBotAccessSettings", v, settings)

	return settings, nil
}

// SetManagedBotAccessSettings updates the access settings of a bot managed by the current bot.
//
// https://core.telegram.org/bots/api#setmanagedbotaccesssettings
func (bot *BotAPI) SetManagedBotAccessSettings(userID int64, isAccessRestricted bool, addedUserIDs []int64) (APIResponse, error) {
	if userID == 0 {
		return APIResponse{}, errors.New("user_id is required")
	}
	if len(addedUserIDs) > 10 {
		return APIResponse{}, errors.New("added_user_ids supports at most 10 users")
	}
	v := url.Values{}
	v.Add("user_id", strconv.FormatInt(userID, 10))
	v.Add("is_access_restricted", strconv.FormatBool(isAccessRestricted))
	if len(addedUserIDs) > 0 {
		data, err := encodeToJson(addedUserIDs)
		if err != nil {
			return APIResponse{}, err
		}
		v.Add("added_user_ids", string(data))
	}

	bot.debugLog("setManagedBotAccessSettings", v, nil)

	return bot.MakeRequest("setManagedBotAccessSettings", v)
}

// GetUserPersonalChatMessages returns recent messages posted to a user's personal chat, as shown on
// their profile page.
//
// https://core.telegram.org/bots/api#getuserpersonalchatmessages
func (bot *BotAPI) GetUserPersonalChatMessages(userID int64, requestedLimit ...int) (messages []Message, err error) {
	limit := 20
	if len(requestedLimit) > 1 {
		return nil, errors.New("only one limit may be specified")
	}
	if len(requestedLimit) == 1 {
		limit = requestedLimit[0]
	}
	if userID == 0 {
		return nil, errors.New("user_id is required")
	}
	if limit < 1 || limit > 20 {
		return nil, errors.New("limit must be between 1 and 20")
	}
	v := url.Values{}
	v.Add("user_id", strconv.FormatInt(userID, 10))
	v.Add("limit", strconv.Itoa(limit))

	resp, err := bot.MakeRequest("getUserPersonalChatMessages", v)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(resp.Result, &messages); err != nil {
		return nil, err
	}

	bot.debugLog("getUserPersonalChatMessages", v, messages)

	return messages, nil
}

// AnswerGuestQuery sends a reply, on behalf of the bot, to a message received via Guest Mode in a chat
// the bot is not a member of.
//
// https://core.telegram.org/bots/api#answerguestquery
func (bot *BotAPI) AnswerGuestQuery(guestQueryID string, result InlineQueryResult) (sent SentGuestMessage, err error) {
	if guestQueryID == "" {
		return sent, errors.New("guest_query_id is required")
	}
	if result == nil {
		return sent, errors.New("result is nil")
	}
	if err = result.Validate(); err != nil {
		return sent, fmt.Errorf("invalid inline query result: %w", err)
	}
	v := url.Values{}
	v.Add("guest_query_id", guestQueryID)
	resultJSON, err := encodeToJson(result)
	if err != nil {
		return sent, fmt.Errorf("failed to marshal inline query result: %w", err)
	}
	v.Add("result", string(resultJSON))

	resp, err := bot.MakeRequest("answerGuestQuery", v)
	if err != nil {
		return sent, err
	}

	if err = json.Unmarshal(resp.Result, &sent); err != nil {
		return sent, err
	}

	bot.debugLog("answerGuestQuery", v, sent)

	return sent, nil
}

// GetChatAdministrators returns chat administrators. Set returnBots to include
// administrator bots other than the current bot.
//
// https://core.telegram.org/bots/api#getchatadministrators
func (bot *BotAPI) GetChatAdministrators(chatID string, includeBots ...bool) (members []ChatMember, err error) {
	if chatID == "" {
		return nil, errors.New("chat_id is required")
	}
	if len(includeBots) > 1 {
		return nil, errors.New("only one return_bots value may be specified")
	}
	returnBots := len(includeBots) == 1 && includeBots[0]
	v := url.Values{}
	v.Add("chat_id", chatID)
	if returnBots {
		v.Add("return_bots", "true")
	}
	resp, err := bot.MakeRequest("getChatAdministrators", v)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(resp.Result, &members); err != nil {
		return nil, err
	}
	bot.debugLog("getChatAdministrators", v, members)
	return members, nil
}

// DeleteAllMessageReactions removes all reactions from a message. Requires the can_restrict_members
// administrator right.
//
// https://core.telegram.org/bots/api#deleteallmessagereactions
func (bot *BotAPI) DeleteAllMessageReactions(chatID int64, messageID int) (APIResponse, error) {
	v := url.Values{}
	v.Add("chat_id", strconv.FormatInt(chatID, 10))
	v.Add("message_id", strconv.Itoa(messageID))

	bot.debugLog("deleteAllMessageReactions", v, nil)

	return bot.MakeRequest("deleteAllMessageReactions", v)
}

// DeleteMessageReaction removes a specific user's reaction from a message. Requires the
// can_restrict_members administrator right.
//
// https://core.telegram.org/bots/api#deletemessagereaction
func (bot *BotAPI) DeleteMessageReaction(chatID int64, messageID int, userID int64) (APIResponse, error) {
	v := url.Values{}
	v.Add("chat_id", strconv.FormatInt(chatID, 10))
	v.Add("message_id", strconv.Itoa(messageID))
	v.Add("user_id", strconv.FormatInt(userID, 10))

	bot.debugLog("deleteMessageReaction", v, nil)

	return bot.MakeRequest("deleteMessageReaction", v)
}

// AnswerChatJoinRequestQuery processes a received chat join request query. Bot API 10.1+
//
// https://core.telegram.org/bots/api#answerchatjoinrequestquery
func (bot *BotAPI) AnswerChatJoinRequestQuery(chatJoinRequestQueryID string, result ChatJoinRequestQueryResult) (APIResponse, error) {
	if chatJoinRequestQueryID == "" {
		return APIResponse{}, errors.New("chat_join_request_query_id is required")
	}
	switch result {
	case ChatJoinRequestQueryResultApprove, ChatJoinRequestQueryResultDecline, ChatJoinRequestQueryResultQueue:
	default:
		return APIResponse{}, fmt.Errorf("invalid chat join request result %q", result)
	}
	v := url.Values{}
	v.Add("chat_join_request_query_id", chatJoinRequestQueryID)
	v.Add("result", string(result))

	bot.debugLog("answerChatJoinRequestQuery", v, nil)

	return bot.MakeRequest("answerChatJoinRequestQuery", v)
}

// SendChatJoinRequestWebApp processes a received chat join request query by showing a Mini App to the
// user before deciding the outcome. Call AnswerChatJoinRequestQuery to resolve the join request query
// based on the user interaction with the Mini App. Bot API 10.1+
//
// https://core.telegram.org/bots/api#sendchatjoinrequestwebapp
func (bot *BotAPI) SendChatJoinRequestWebApp(chatJoinRequestQueryID, webAppURL string) (APIResponse, error) {
	if chatJoinRequestQueryID == "" {
		return APIResponse{}, errors.New("chat_join_request_query_id is required")
	}
	if webAppURL == "" {
		return APIResponse{}, errors.New("web_app_url is required")
	}
	v := url.Values{}
	v.Add("chat_join_request_query_id", chatJoinRequestQueryID)
	v.Add("web_app_url", webAppURL)

	bot.debugLog("sendChatJoinRequestWebApp", v, nil)

	return bot.MakeRequest("sendChatJoinRequestWebApp", v)
}

// EditEphemeralMessageText edits an ephemeral text message. Bot API 10.2+
//
// https://core.telegram.org/bots/api#editephemeralmessagetext
func (bot *BotAPI) EditEphemeralMessageText(config EditEphemeralMessageTextConfig) (APIResponse, error) {
	return bot.MakeRequestFromChattable(config)
}

// EditEphemeralMessageMedia edits the media of an ephemeral message. Bot API 10.2+
//
// https://core.telegram.org/bots/api#editephemeralmessagemedia
func (bot *BotAPI) EditEphemeralMessageMedia(config EditEphemeralMessageMediaConfig) (APIResponse, error) {
	return bot.MakeRequestFromChattable(config)
}

// EditEphemeralMessageCaption edits the caption of an ephemeral message. Bot API 10.2+
//
// https://core.telegram.org/bots/api#editephemeralmessagecaption
func (bot *BotAPI) EditEphemeralMessageCaption(config EditEphemeralMessageCaptionConfig) (APIResponse, error) {
	return bot.MakeRequestFromChattable(config)
}

// EditEphemeralMessageReplyMarkup edits only the reply markup of an ephemeral message. Bot API 10.2+
//
// https://core.telegram.org/bots/api#editephemeralmessagereplymarkup
func (bot *BotAPI) EditEphemeralMessageReplyMarkup(config EditEphemeralMessageReplyMarkupConfig) (APIResponse, error) {
	return bot.MakeRequestFromChattable(config)
}

// DeleteEphemeralMessage deletes an ephemeral message. Bot API 10.2+
//
// https://core.telegram.org/bots/api#deleteephemeralmessage
func (bot *BotAPI) DeleteEphemeralMessage(config DeleteEphemeralMessageConfig) (APIResponse, error) {
	return bot.MakeRequestFromChattable(config)
}

func (bot *BotAPI) SendCustomMessage(ctx context.Context, config Sendable, result any) (err error) {
	var values url.Values
	if values, err = config.Values(); err != nil {
		return
	}
	telegramMethod := config.TelegramMethod()
	var apiResponse APIResponse
	apiResponse, err = bot.MakeRequest(telegramMethod, values)
	if err != nil {
		return
	}
	if err = json.Unmarshal(apiResponse.Result, &result); err != nil {
		err = fmt.Errorf("failed to decode Telegram API response for method %q into type %T: %w", telegramMethod, result, err)
		return
	}
	return
}
