package tgbotapi

//go:generate ffjson $GOFILE

import (
	"errors"
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// Telegram constants
const (
	// APIEndpoint is the endpoint for all API methods,
	// with formatting for Sprintf.
	APIEndpoint = "https://api.telegram.org/bot%s/%s"
	// FileEndpoint is the endpoint for downloading a file from Telegram.
	FileEndpoint = "https://api.telegram.org/file/bot%s/%s"
)

// Constant values for ChatActions
const (
	// ChatTyping is chat action
	ChatTyping = "typing"

	// ChatUploadPhoto is chat action
	ChatUploadPhoto = "upload_photo"

	// ChatRecordVideo is chat action
	ChatRecordVideo = "record_video"

	// ChatUploadVideo is chat action
	ChatUploadVideo = "upload_video"

	// ChatRecordAudio is chat action
	ChatRecordAudio = "record_audio"

	// ChatUploadAudio is chat action
	ChatUploadAudio = "upload_audio"

	// ChatUploadDocument is chat action
	ChatUploadDocument = "upload_document"

	// ChatFindLocation is chat action
	ChatFindLocation = "find_location"
)

// API errors
//const (
// ErrAPIForbidden happens when a token is bad
//ErrAPIForbidden = "forbidden"
//)
//var ErrAPIForbidden = errors.New("forbidden")  // happens when a token is bad or user deleted chat

// ErrAPIForbidden is for 'forbidden' API response
type ErrAPIForbidden struct {
}

// Error implements error interface
//
//goland:noinspection GoMixedReceiverTypes
func (err ErrAPIForbidden) Error() string {
	return "forbidden"
}

// IsForbidden indicates is forbidden
//
//goland:noinspection GoMixedReceiverTypes
func (err ErrAPIForbidden) IsForbidden() bool {
	return true
}

// Constant values for ParseMode in MessageConfig
const (
	// ModeMarkdown indicates markdown mode
	ModeMarkdown = "Markdown"

	// ModeHTML indicates HTML mode
	ModeHTML = "HTML"
)

// Library errors
var (
	// ErrBadFileType happens when you pass an unknown type
	ErrBadFileType = errors.New("bad file type")

	// ErrBadURL indicates bad or empty URL
	ErrBadURL = errors.New("bad or empty URL")
)

// Chattable is any config type that can be sent.
type Chattable interface {
	Values() (url.Values, error)
	method() string
}

// Fileable is any config type that can be sent that includes a file.
type Fileable interface {
	Chattable
	params() (map[string]string, error)
	name() string
	getFile() interface{}
	useExistingFile() bool
}

// BaseChat is base type for all chat config types.
type BaseChat struct {
	ChatID              int64       `json:"chat_id,omitempty"`
	ChannelUsername     string      `json:"channel_username,omitempty"`
	ReplyToMessageID    int         `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         interface{} `json:"reply_markup,omitempty"`
	DisableNotification bool        `json:"disable_notification,omitempty"`
}

// Values returns url.Values representation of BaseChat
//
//goland:noinspection GoMixedReceiverTypes
func (j BaseChat) Values() (url.Values, error) {
	values := url.Values{}
	if j.ChannelUsername != "" {
		values.Add("chat_id", j.ChannelUsername)
	} else {
		values.Add("chat_id", strconv.FormatInt(j.ChatID, 10))
	}

	if j.ReplyToMessageID != 0 {
		values.Add("reply_to_message_id", strconv.Itoa(j.ReplyToMessageID))
	}

	if j.ReplyMarkup != nil {
		data, err := ffjson.Marshal(j.ReplyMarkup)
		if err != nil {
			ffjson.Pool(data)
			return values, err
		}
		if string(data) == "null" {
			panic(fmt.Sprintf("string(data) == null, BaseChat: %v", j))
		}

		values.Add("reply_markup", string(data))
		ffjson.Pool(data)
	}

	values.Add("disable_notification", strconv.FormatBool(j.DisableNotification))

	return values, nil
}

// BaseFile is a base type for all file config types.
type BaseFile struct {
	BaseChat
	File        interface{}
	FileID      string
	UseExisting bool
	MimeType    string
	FileSize    int
}

// params returns a map[string]string representation of BaseFile.
//
//goland:noinspection GoMixedReceiverTypes
func (file BaseFile) params() (map[string]string, error) {
	params := make(map[string]string)

	if file.ChannelUsername != "" {
		params["chat_id"] = file.ChannelUsername
	} else {
		params["chat_id"] = strconv.FormatInt(file.ChatID, 10)
	}

	if file.ReplyToMessageID != 0 {
		params["reply_to_message_id"] = strconv.Itoa(file.ReplyToMessageID)
	}

	if file.ReplyMarkup != nil {
		data, err := ffjson.Marshal(file.ReplyMarkup)
		if err != nil {
			ffjson.Pool(data)
			return params, err
		}

		params["reply_markup"] = string(data)
		ffjson.Pool(data)
	}

	if file.MimeType != "" {
		params["mime_type"] = file.MimeType
	}

	if file.FileSize > 0 {
		params["file_size"] = strconv.Itoa(file.FileSize)
	}

	params["disable_notification"] = strconv.FormatBool(file.DisableNotification)

	return params, nil
}

// getFile returns the file.
//
//goland:noinspection GoMixedReceiverTypes
func (file BaseFile) getFile() interface{} {
	return file.File
}

// useExistingFile returns if the BaseFile has already been uploaded.
//
//goland:noinspection GoMixedReceiverTypes
func (file BaseFile) useExistingFile() bool {
	return file.UseExisting
}

type chatEdit struct {
	ChatID    int64 `json:"chat_id,omitempty"`
	MessageID int   `json:"message_id,omitempty"`
}

// NewChatMessageEdit returns BaseEdit
func NewChatMessageEdit(chatID int64, messageID int) BaseEdit {
	return BaseEdit{chatEdit: chatEdit{ChatID: chatID, MessageID: messageID}}
}

// BaseEdit is base type of all chat edits.
type BaseEdit struct {
	chatEdit
	ChannelUsername string                `json:",omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	ReplyMarkup     *InlineKeyboardMarkup `json:",omitempty"`
}

// Values returns URL values
//
//goland:noinspection GoMixedReceiverTypes
func (v BaseEdit) Values() (url.Values, error) {
	values := url.Values{}

	if v.ChannelUsername != "" {
		values.Add("chat_id", v.ChannelUsername)
	}
	if v.ChatID != 0 {
		values.Add("chat_id", strconv.FormatInt(v.ChatID, 10))
	}
	if v.MessageID != 0 {
		values.Add("message_id", strconv.Itoa(v.MessageID))
	}
	if v.InlineMessageID != "" {
		values.Add("inline_message_id", v.InlineMessageID)
	}

	if v.ReplyMarkup != nil {
		data, err := ffjson.Marshal(v.ReplyMarkup)
		if err != nil {
			ffjson.Pool(data)
			return values, err
		}
		values.Add("reply_markup", string(data))
		ffjson.Pool(data)
	}

	return values, nil
}

// MessageConfig contains information about a SendMessage request.
type MessageConfig struct {
	BaseChat
	Text                  string
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview,omitempty"`
}

// Values returns url.Values representation of MessageConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v MessageConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()
	values.Add("text", v.Text)
	values.Add("disable_web_page_preview", strconv.FormatBool(v.DisableWebPagePreview))
	if v.ParseMode != "" {
		values.Add("parse_mode", v.ParseMode)
	}

	return values, nil
}

// method returns Telegram API method name for sending Message.
//
//goland:noinspection GoMixedReceiverTypes
func (v MessageConfig) method() string {
	return "sendMessage"
}

// ForwardConfig contains information about a ForwardMessage request.
type ForwardConfig struct {
	BaseChat
	FromChatID          int64 // required
	FromChannelUsername string
	MessageID           int // required
}

// Values returns url.Values representation of ForwardConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v ForwardConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()
	values.Add("from_chat_id", strconv.FormatInt(v.FromChatID, 10))
	values.Add("message_id", strconv.Itoa(v.MessageID))
	return values, nil
}

// method returns Telegram API method name for sending Forward.
//
//goland:noinspection GoMixedReceiverTypes
func (v ForwardConfig) method() string {
	return "forwardMessage"
}

type chatMethod struct {
	ChatID string `json:"chat_id"`
}

//goland:noinspection GoMixedReceiverTypes
func (v chatMethod) Values() (url.Values, error) {
	if v.ChatID == "" {
		return nil, ErrNoChatID
	}
	return url.Values{"chat_id": []string{v.ChatID}}, nil
}

// LeaveChatConfig is message command for leaving chat
type LeaveChatConfig struct {
	chatMethod
}

//goland:noinspection GoMixedReceiverTypes
func (LeaveChatConfig) method() string {
	return "leaveChat"
}

// ExportChatInviteLink is message command for exporting chat link
type ExportChatInviteLink struct {
	chatMethod
}

//goland:noinspection GoMixedReceiverTypes
func (ExportChatInviteLink) method() string {
	return "exportChatInviteLink"
}

// ErrNoChatID is error when chat_id is missing
var ErrNoChatID = errors.New("missing chat_id")

// PhotoConfig contains information about a SendPhoto request.
type PhotoConfig struct {
	BaseFile
	Caption string
}

// Params returns a map[string]string representation of PhotoConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v PhotoConfig) params() (map[string]string, error) {
	params, _ := v.BaseFile.params()

	if v.Caption != "" {
		params["caption"] = v.Caption
	}

	return params, nil
}

// Values returns url.Values representation of PhotoConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v PhotoConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()

	values.Add(v.name(), v.FileID)
	if v.Caption != "" {
		values.Add("caption", v.Caption)
	}
	return values, nil
}

// name returns the field name for the Photo.
//
//goland:noinspection GoMixedReceiverTypes
func (v PhotoConfig) name() string {
	return "photo"
}

// method returns Telegram API method name for sending Photo.
//
//goland:noinspection GoMixedReceiverTypes
func (v PhotoConfig) method() string {
	return "sendPhoto"
}

// AudioConfig contains information about a SendAudio request.
type AudioConfig struct {
	BaseFile
	Duration  int
	Performer string
	Title     string
}

// Values returns url.Values representation of AudioConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (j AudioConfig) Values() (url.Values, error) {
	values, _ := j.BaseChat.Values()

	values.Add(j.name(), j.FileID)
	if j.Duration != 0 {
		values.Add("duration", strconv.Itoa(j.Duration))
	}

	if j.Performer != "" {
		values.Add("performer", j.Performer)
	}
	if j.Title != "" {
		values.Add("title", j.Title)
	}

	return values, nil
}

// params returns a map[string]string representation of AudioConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (j AudioConfig) params() (map[string]string, error) {
	params, _ := j.BaseFile.params()

	if j.Duration != 0 {
		params["duration"] = strconv.Itoa(j.Duration)
	}

	if j.Performer != "" {
		params["performer"] = j.Performer
	}
	if j.Title != "" {
		params["title"] = j.Title
	}

	return params, nil
}

// name returns the field name for the Audio.
//
//goland:noinspection GoMixedReceiverTypes
func (j AudioConfig) name() string {
	return "audio"
}

// method returns Telegram API method name for sending Audio.
//
//goland:noinspection GoMixedReceiverTypes
func (j AudioConfig) method() string {
	return "sendAudio"
}

// DocumentConfig contains information about a SendDocument request.
type DocumentConfig struct {
	BaseFile
}

// Values returns url.Values representation of DocumentConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v DocumentConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()

	values.Add(v.name(), v.FileID)

	return values, nil
}

// params returns a map[string]string representation of DocumentConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v DocumentConfig) params() (map[string]string, error) {
	params, _ := v.BaseFile.params()

	return params, nil
}

// name returns the field name for the Document.
//
//goland:noinspection GoMixedReceiverTypes
func (v DocumentConfig) name() string {
	return "document"
}

// method returns Telegram API method name for sending Document.
//
//goland:noinspection GoMixedReceiverTypes
func (v DocumentConfig) method() string {
	return "sendDocument"
}

// StickerConfig contains information about a SendSticker request.
type StickerConfig struct {
	BaseFile
}

// Values returns url.Values representation of StickerConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v StickerConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()

	values.Add(v.name(), v.FileID)

	return values, nil
}

// params returns a map[string]string representation of StickerConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v StickerConfig) params() (map[string]string, error) {
	params, _ := v.BaseFile.params()

	return params, nil
}

// name returns the field name for the Sticker.
//
//goland:noinspection GoMixedReceiverTypes
func (v StickerConfig) name() string {
	return "sticker"
}

// method returns Telegram API method name for sending Sticker.
//
//goland:noinspection GoMixedReceiverTypes
func (v StickerConfig) method() string {
	return "sendSticker"
}

// VideoConfig contains information about a SendVideo request.
type VideoConfig struct {
	BaseFile
	Duration int
	Caption  string
}

// Values returns a url.Values representation of VideoConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v VideoConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()

	values.Add(v.name(), v.FileID)
	if v.Duration != 0 {
		values.Add("duration", strconv.Itoa(v.Duration))
	}
	if v.Caption != "" {
		values.Add("caption", v.Caption)
	}

	return values, nil
}

// params returns a map[string]string representation of VideoConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v VideoConfig) params() (map[string]string, error) {
	params, _ := v.BaseFile.params()

	return params, nil
}

// name returns the field name for the Video.
//
//goland:noinspection GoMixedReceiverTypes
func (v VideoConfig) name() string {
	return "video"
}

// method returns Telegram API method name for sending Video.
//
//goland:noinspection GoMixedReceiverTypes
func (v VideoConfig) method() string {
	return "sendVideo"
}

// VoiceConfig contains information about a SendVoice request.
type VoiceConfig struct {
	BaseFile
	Duration int
}

// Values returns a url.Values representation of VoiceConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v VoiceConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()

	values.Add(v.name(), v.FileID)
	if v.Duration != 0 {
		values.Add("duration", strconv.Itoa(v.Duration))
	}

	return values, nil
}

// params returns a map[string]string representation of VoiceConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v VoiceConfig) params() (map[string]string, error) {
	params, _ := v.BaseFile.params()

	if v.Duration != 0 {
		params["duration"] = strconv.Itoa(v.Duration)
	}

	return params, nil
}

// name returns the field name for the Voice.
//
//goland:noinspection GoMixedReceiverTypes
func (v VoiceConfig) name() string {
	return "voice"
}

// method returns Telegram API method name for sending Voice.
//
//goland:noinspection GoMixedReceiverTypes
func (v VoiceConfig) method() string {
	return "sendVoice"
}

// LocationConfig contains information about a SendLocation request.
type LocationConfig struct {
	BaseChat
	Latitude  float64 // required
	Longitude float64 // required
}

// Values returns url.Values representation of LocationConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (j LocationConfig) Values() (url.Values, error) {
	values, _ := j.BaseChat.Values()

	values.Add("latitude", strconv.FormatFloat(j.Latitude, 'f', 6, 64))
	values.Add("longitude", strconv.FormatFloat(j.Longitude, 'f', 6, 64))

	return values, nil
}

// method returns Telegram API method name for sending Location.
//
//goland:noinspection GoMixedReceiverTypes
func (j LocationConfig) method() string {
	return "sendLocation"
}

// VenueConfig contains information about a SendVenue request.
type VenueConfig struct {
	BaseChat
	Latitude     float64 // required
	Longitude    float64 // required
	Title        string  // required
	Address      string  // required
	FoursquareID string
}

// Values returns URL values representation of VenueConfig
//
//goland:noinspection GoMixedReceiverTypes
func (v VenueConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()

	values.Add("latitude", strconv.FormatFloat(v.Latitude, 'f', 6, 64))
	values.Add("longitude", strconv.FormatFloat(v.Longitude, 'f', 6, 64))
	values.Add("title", v.Title)
	values.Add("address", v.Address)
	if v.FoursquareID != "" {
		values.Add("foursquare_id", v.FoursquareID)
	}

	return values, nil
}

//goland:noinspection GoMixedReceiverTypes
func (v VenueConfig) method() string {
	return "sendVenue"
}

// ContactConfig allows you to send a contact.
type ContactConfig struct {
	BaseChat
	PhoneNumber string
	FirstName   string
	LastName    string
}

// Values returns URL values representation of ContactConfig
//
//goland:noinspection GoMixedReceiverTypes
func (j ContactConfig) Values() (url.Values, error) {
	values, _ := j.BaseChat.Values()

	values.Add("phone_number", j.PhoneNumber)
	values.Add("first_name", j.FirstName)
	values.Add("last_name", j.LastName)

	return values, nil
}

//goland:noinspection GoMixedReceiverTypes
func (j ContactConfig) method() string {
	return "sendContact"
}

// ChatActionConfig contains information about a SendChatAction request.
type ChatActionConfig struct {
	BaseChat
	Action string // required
}

// Values returns url.Values representation of ChatActionConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (config ChatActionConfig) Values() (url.Values, error) {
	values, _ := config.BaseChat.Values()
	values.Add("action", config.Action)
	return values, nil
}

// method returns Telegram API method name for sending ChatAction.
//
//goland:noinspection GoMixedReceiverTypes
func (config ChatActionConfig) method() string {
	return "sendChatAction"
}

// DeleteMessage is a command to delete a message.
// It should not be used with SendMessage()
// Instead use BotAPI.DeleteMessage(chatID string, messageID int)
type DeleteMessage chatEdit

//goland:noinspection GoMixedReceiverTypes
func (*DeleteMessage) method() string {
	return "deleteMessage"
}

// Values returns URL values representation of DeleteMessage
//
//goland:noinspection GoMixedReceiverTypes
func (m DeleteMessage) Values() (url.Values, error) {
	return url.Values{
		"chat_id":    []string{strconv.FormatInt(m.ChatID, 10)},
		"message_id": []string{strconv.Itoa(m.MessageID)},
	}, nil
}

var _ Chattable = (*DeleteMessage)(nil)

// EditMessageTextConfig allows you to modify the text in a message.
type EditMessageTextConfig struct {
	BaseEdit
	Text                  string
	ParseMode             string
	DisableWebPagePreview bool
}

// Values returns URL values representation of EditMessageTextConfig
//
//goland:noinspection GoMixedReceiverTypes
func (j EditMessageTextConfig) Values() (url.Values, error) {
	v, _ := j.BaseEdit.Values()

	v.Add("text", j.Text)
	if j.ParseMode != "" {
		v.Add("parse_mode", j.ParseMode)
	}
	if j.DisableWebPagePreview {
		v.Add("disable_web_page_preview", strconv.FormatBool(j.DisableWebPagePreview))
	}

	return v, nil
}

//goland:noinspection GoMixedReceiverTypes
func (j EditMessageTextConfig) method() string {
	return "editMessageText"
}

// EditMessageCaptionConfig allows you to modify the caption of a message.
type EditMessageCaptionConfig struct {
	BaseEdit
	Caption string
}

// Values returns URL values representation of EditMessageCaptionConfig
//
//goland:noinspection GoMixedReceiverTypes
func (j EditMessageCaptionConfig) Values() (url.Values, error) {
	v, _ := j.BaseEdit.Values()

	v.Add("caption", j.Caption)

	return v, nil
}

//goland:noinspection GoMixedReceiverTypes
func (j EditMessageCaptionConfig) method() string {
	return "editMessageCaption"
}

// EditMessageReplyMarkupConfig allows you to modify the reply markup
// of a message.
type EditMessageReplyMarkupConfig struct {
	BaseEdit
}

// Values returns URL values representation of EditMessageReplyMarkupConfig
//
//goland:noinspection GoMixedReceiverTypes
func (config EditMessageReplyMarkupConfig) Values() (url.Values, error) {
	return config.BaseEdit.Values()
}

//goland:noinspection GoMixedReceiverTypes
func (config EditMessageReplyMarkupConfig) method() string {
	return "editMessageReplyMarkup"
}

// UserProfilePhotosConfig contains information about a
// GetUserProfilePhotos request.
type UserProfilePhotosConfig struct {
	UserID int
	Offset int
	Limit  int
}

// FileConfig has information about a file hosted on Telegram.
type FileConfig struct {
	FileID string
}

// UpdateConfig contains information about a GetUpdates request.
type UpdateConfig struct {
	Offset  int
	Limit   int
	Timeout int
}

// WebhookConfig contains information about a SetWebhook request.
type WebhookConfig struct {

	// URL - HTTPS url to send updates to. Use an empty string to remove webhook integration
	URL *url.URL `json:"url"` // REQUIRED!

	// Certificate - 	Upload your public key certificate so that the root certificate in use can be checked.
	// See https://core.telegram.org/bots/self-signed guide for details.
	Certificate interface{} `json:"certificate,omitempty"`

	// IPAddress - The fixed IP address which will be used to send webhook requests instead of the IP address resolved through DNS
	IPAddress string `json:"ip_address,omitempty"`

	// MaxConnections - 	The maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100.
	// Defaults to 40. Use lower values to limit the load on your bot's server, and higher values to increase your bot's throughput.
	MaxConnections int `json:"max_connections,omitempty"`

	// AllowedUpdates - A JSON-serialized list of the update types you want your bot to receive.
	// For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types.
	// See Update for a complete list of available update types.
	// Specify an empty list to receive all update types except chat_member (default).
	// If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`

	// DropPendingUpdates - Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`

	// SecretToken - A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token” in every webhook request, 1-256 characters.
	// Only characters A-Z, a-z, 0-9, _ and - are allowed. The header is useful to ensure that the request comes from a webhook set by you.
	SecretToken string `json:"secret_token,omitempty"`
}

// Values returns url.Values representation of WebhookConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (j WebhookConfig) Values() (url.Values, error) {
	if j.URL == nil {
		return nil, errors.New("URL is nil")
	}
	values := make(url.Values, 6)
	if webhookUrl := j.URL.String(); webhookUrl == "" {
		return nil, errors.New("URL is empty string")
	} else {
		values.Add("url", webhookUrl)
	}
	if j.IPAddress != "" {
		values.Add("ip_address", j.IPAddress)
	}
	if j.MaxConnections > 0 {
		values.Add("max_connections", strconv.Itoa(j.MaxConnections))
	}
	if len(j.AllowedUpdates) > 0 {
		values.Add("allowed_updates", `["`+strings.Join(j.AllowedUpdates, `","`)+`"]`)
	}
	if j.DropPendingUpdates {
		values.Add("drop_pending_updates", "True")
	}
	if j.SecretToken != "" {
		values.Add("secret_token", j.SecretToken)
	}
	return values, nil
}

// Validate returns an error if the WebhookConfig struct is invalid.
//
//goland:noinspection GoMixedReceiverTypes
func (j WebhookConfig) Validate() error {
	if j.URL == nil || j.URL.String() == "" {
		return errors.New("URL is required")
	}
	return nil
}

// FileBytes contains information about a set of bytes to upload
// as a File.
type FileBytes struct {
	Name  string
	Bytes []byte
}

// FileReader contains information about a reader to upload as a File.
// If Size is -1, it will read the entire Reader into memory to
// calculate a Size.
type FileReader struct {
	Name   string
	Reader io.Reader
	Size   int64
}

// InlineConfig contains information on making an InlineQuery response.
type InlineConfig struct {
	InlineQueryID     string        `json:"inline_query_id"`
	Results           []interface{} `json:"results,omitempty"`
	CacheTime         int           `json:"cache_time"`
	IsPersonal        bool          `json:"is_personal,omitempty"`
	NextOffset        string        `json:"next_offset,omitempty"`
	SwitchPMText      string        `json:"switch_pm_text,omitempty"`
	SwitchPMParameter string        `json:"switch_pm_parameter,omitempty"`
}

//goland:noinspection GoMixedReceiverTypes
func (config InlineConfig) method() string {
	return "answerInlineQuery"
}

// Values returns URL values representation of InlineConfig
//
//goland:noinspection GoMixedReceiverTypes
func (config InlineConfig) Values() (url.Values, error) {
	v := url.Values{}

	v.Add("inline_query_id", config.InlineQueryID)
	if config.CacheTime != 0 {
		v.Add("cache_time", strconv.Itoa(config.CacheTime))
	}
	if config.IsPersonal {
		v.Add("is_personal", strconv.FormatBool(config.IsPersonal))
	}
	if config.NextOffset != "" {
		v.Add("next_offset", config.NextOffset)
	}
	if config.SwitchPMText != "" {
		v.Add("switch_pm_text", config.SwitchPMText)
	}
	if config.SwitchPMParameter != "" {
		v.Add("switch_pm_parameter", config.SwitchPMParameter)
	}

	data, err := ffjson.Marshal(config.Results)
	if err != nil {
		ffjson.Pool(data)
		return v, err
	}
	v.Add("results", string(data))
	ffjson.Pool(data)

	return v, nil
}

// AnswerCallbackQueryConfig contains information on making a CallbackQuery response.
type AnswerCallbackQueryConfig struct {
	CallbackQueryID string `json:"callback_query_id,"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
	CacheTime       int    `json:"cache_time,omitempty"`
}

var _ Chattable = (*AnswerCallbackQueryConfig)(nil)

//goland:noinspection GoMixedReceiverTypes
func (j AnswerCallbackQueryConfig) method() string {
	return "answerCallbackQuery"
}

// Values returns URL values representation of AnswerCallbackQueryConfig
//
//goland:noinspection GoMixedReceiverTypes
func (j AnswerCallbackQueryConfig) Values() (url.Values, error) {
	values := make(url.Values, 3) // if removed causes nil pointer exception
	values.Add("callback_query_id", j.CallbackQueryID)
	if j.Text != "" && j.URL != "" {
		return nil, errors.New("both j.Text && j.URL supplied")
	}
	if j.Text != "" {
		values.Add("text", j.Text)
		if j.ShowAlert {
			values.Add("show_alert", "true")
		}
	} else if j.URL != "" {
		values.Add("url", j.URL)
	}
	return values, nil
}

// ChatMemberConfig contains information about a user in a chat for use
// with administrative functions such as kicking or unbanning a user.
type ChatMemberConfig struct {
	ChatID             int64
	SuperGroupUsername string
	UserID             int
}
