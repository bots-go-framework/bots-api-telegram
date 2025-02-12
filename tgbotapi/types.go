package tgbotapi

//go:generate ffjson $GOFILE

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	botsgocore "github.com/bots-go-framework/bots-go-core"
	"net/url"
	"strconv"
	"strings"
)

// APIResponse is a response from the Telegram API with the result
// stored raw.
type APIResponse struct {
	Ok          bool            `json:"ok"`
	Result      json.RawMessage `json:"result"`
	ErrorCode   int             `json:"error_code"`
	Description string          `json:"description"`
}

func (r APIResponse) Error() string {
	if r.Ok || r.ErrorCode != 0 {
		return fmt.Sprintf("APIResponse{Ok: %v, ErrorCode: %d, Description: %v}", r.Ok, r.ErrorCode, r.Description)
	}
	return ""
}

// User is a user on Telegram.
type User struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name,omitempty"`
	LastName     string `json:"last_name,omitempty"`     // optional
	UserName     string `json:"username,omitempty"`      // optional
	LanguageCode string `json:"language_code,omitempty"` // optional
}

// ChatMember holds information about chat member
type ChatMember struct {
	User
	IsBot bool `json:"is_bot,omitempty"`
}

// IsBotUser indicates if chat member is a bot
func (chatMember ChatMember) IsBotUser() bool {
	return chatMember.IsBot
}

// Platform returns 'Telegram'
func (u User) Platform() string {
	return "telegram"
}

// GetID returns Telegram user ID
func (u User) GetID() interface{} {
	return u.ID
}

// GetFirstName returns first name of the user
func (u User) GetFirstName() string {
	return u.FirstName
}

// GetLastName returns last name of the user
func (u User) GetLastName() string {
	return u.LastName
}

// GetUserName returns user name of the user
func (u User) GetUserName() string {
	return u.UserName
}

// GetFullName returns full name of the user
func (u User) GetFullName() string {
	var buffer bytes.Buffer
	if u.FirstName != "" {
		buffer.WriteString(u.FirstName)
		if u.LastName != "" {
			buffer.WriteString(" ")
			buffer.WriteString(u.LastName)
			return buffer.String()
		}
		if u.UserName != "" {
			_, _ = fmt.Fprintf(&buffer, " (@%s)", u.UserName)
		}
		return buffer.String()
	}
	if u.LastName != "" {
		buffer.WriteString(u.LastName)
		if u.UserName != "" {
			_, _ = fmt.Fprintf(&buffer, " (@%s)", u.UserName)
		}
		return buffer.String()
	}

	if u.UserName != "" {
		return "@" + u.UserName
	}

	return "#" + strconv.Itoa(u.ID)
}

// GetLanguage returns preferred language of the user
func (u User) GetLanguage() string {
	return u.LanguageCode
}

// String displays a simple text version of a user.
//
// It is normally a user's username, but falls back to a first/last
// name as available.
func (u *User) String() string {
	if u.UserName != "" {
		return u.UserName
	}

	name := u.FirstName
	if u.LastName != "" {
		name += " " + u.LastName
	}

	return name
}

// GroupChat is a group chat.
type GroupChat struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// Chat contains information about the place a message was sent.
type Chat struct {
	ID        int64  `json:"id"`
	Type      string `json:"type"`
	Title     string `json:"title"`      // optional
	UserName  string `json:"username"`   // optional
	FirstName string `json:"first_name"` // optional
	LastName  string `json:"last_name"`  // optional
}

// IsPrivate returns if the Chat is a private conversation.
func (c *Chat) IsPrivate() bool {
	return c.Type == "private"
}

// IsGroup returns if the Chat is a group.
func (c *Chat) IsGroup() bool {
	return c.Type == "group"
}

// IsSuperGroup returns if the Chat is a supergroup.
func (c *Chat) IsSuperGroup() bool {
	return c.Type == "supergroup"
}

// IsChannel returns if the Chat is a channel.
func (c *Chat) IsChannel() bool {
	return c.Type == "channel"
}

// MessageEntity contains information about data in a Message.
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"` // optional
}

func (entity *MessageEntity) Validate() error {
	if entity.Type == "" {
		return errors.New("MessageEntity.Type is empty")
	}
	if entity.Offset < 0 {
		return errors.New("MessageEntity.Offset is negative")
	}
	if entity.Length <= 0 {
		return errors.New("MessageEntity.Length is not positive")
	}
	return nil
}

// ParseURL attempts to parse a URL contained within a MessageEntity.
func (entity *MessageEntity) ParseURL() (*url.URL, error) {
	if entity.URL == "" {
		return nil, ErrBadURL
	}

	return url.Parse(entity.URL)
}

// Audio contains information about audio.
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"` // optional
	Title     string `json:"title"`     // optional
	MimeType  string `json:"mime_type"` // optional
	FileSize  int    `json:"file_size"` // optional
}

// Document contains information about a document.
type Document struct {
	FileID    string     `json:"file_id"`
	Thumbnail *PhotoSize `json:"thumb,omitempty"`     // optional
	FileName  string     `json:"file_name,omitempty"` // optional
	MimeType  string     `json:"mime_type,omitempty"` // optional
	FileSize  int        `json:"file_size,omitempty"` // optional
}

// Sticker contains information about a sticker.
type Sticker struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Thumbnail *PhotoSize `json:"thumb,omitempty"`     // optional
	FileSize  int        `json:"file_size,omitempty"` // optional
}

// Video contains information about a video.
type Video struct {
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb,omitempty"`     // optional
	MimeType  string     `json:"mime_type,omitempty"` // optional
	FileSize  int        `json:"file_size,omitempty"` // optional
}

// Voice contains information about a voice.
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type,omitempty"` // optional
	FileSize int    `json:"file_size,omitempty"` // optional
}

// Contact contains information about a contact.
// Note that LastName, UserID, VCard may be empty.
type Contact struct {

	// PhoneNumber must always be presented
	PhoneNumber string `json:"phone_number"`

	// FirstName must always be presented
	FirstName string `json:"first_name"`

	// Optional
	LastName string `json:"last_name,omitempty"` // optional

	// UserID (optional) is a Contact's user identifier in Telegram.
	// It has at most 52 significant bits,
	// so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserID int64 `json:"user_id,omitempty"` // optional

	// VCard (optional) additional data about the contact in the form of https://en.wikipedia.org/wiki/VCard
	VCard string `json:"vcard,omitempty"`
}

// Location contains information about a place.
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Venue contains information about a venue, including its Location.
type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID string   `json:"foursquare_id,omitempty"` // optional
}

// UserProfilePhotos contains a set of user profile photos.
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// File contains information about a file to download from Telegram.
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size,omitempty"` // optional
	FilePath string `json:"file_path,omitempty"` // optional
}

// Link returns a full path to the download URL for a File.
//
// It requires the Bot Token to create the link.
func (f *File) Link(token string) string {
	return fmt.Sprintf(FileEndpoint, token, f.FilePath)
}

// ReplyKeyboardMarkup allows the Bot to set a custom keyboard.
type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard,omitempty"`   // optional
	OneTimeKeyboard bool               `json:"one_time_keyboard,omitempty"` // optional
	Selective       bool               `json:"selective,omitempty"`         // optional
}

// KeyboardType returns KeyboardTypeBottom
func (*ReplyKeyboardMarkup) KeyboardType() botsgocore.KeyboardType {
	return botsgocore.KeyboardTypeBottom
}

var _ botsgocore.Keyboard = (*ReplyKeyboardMarkup)(nil)

// KeyboardButtonRequestUsers represents a request from the bot to send users
// https://core.telegram.org/bots/api#keyboardbuttonrequestusers
type KeyboardButtonRequestUsers struct {
	// Signed 32-bit identifier of the request, which will be received back in the ChatShared object.
	// Must be unique within the message
	RequestID int `json:"request_id"`

	// Optional.
	// Pass True to request bots, pass False to request regular users.
	// If not specified, no additional restrictions are applied.
	UserIsBot bool `json:"user_is_bot,omitempty"`

	// Optional.
	// Pass True to request premium users, pass False to request non-premium users.
	// If not specified, no additional restrictions are applied.
	UserIsPremium bool `json:"user_is_premium,omitempty"`

	// Optional. The maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity int `json:"max_quantity,omitempty"`

	// Optional. Pass True to request the users' first and last names
	RequestName bool `json:"request_name,omitempty"`

	// Optional. Pass True to request the users' usernames
	RequestUsername bool `json:"request_username,omitempty"`

	// Optional. Pass True to request the users' photos
	RequestPhoto bool `json:"request_photo,omitempty"`
}

// KeyboardButtonRequestChat represents a request from the bot to send a chat
// https://core.telegram.org/bots/api#keyboardbuttonrequestchat
type KeyboardButtonRequestChat struct {
	// Signed 32-bit identifier of the request, which will be received back in the ChatShared object.
	// Must be unique within the message
	RequestID int `json:"request_id"`

	// Pass True to request a channel chat, pass False to request a group or a supergroup chat.
	ChatIsChannel bool `json:"chat_is_channel"`

	// Pass True to request a forum supergroup, pass False to request a non-forum chat.
	// If not specified, no additional restrictions are applied.
	ChatIsForum bool `json:"chat_is_forum,omitempty"`

	// Pass True to request a supergroup or a channel with a username,
	// pass False to request a chat without a username.
	// If not specified, no additional restrictions are applied.
	ChatHasUsername bool `json:"chat_has_username,omitempty"`

	// Pass True to request a chat owned by the user. Otherwise, no additional restrictions are applied.
	ChatIsCreated bool `json:"chat_is_created,omitempty"`

	// A JSON-serialized object listing the required administrator rights of the user in the chat.
	// The rights must be a superset of bot_administrator_rights.
	// If not specified, no additional restrictions are applied.
	UserAdministratorRights *ChatAdministratorRights `json:"user_administrator_rights,omitempty"`

	// A JSON-serialized object listing the required administrator rights of the bot in the chat.
	// The rights must be a subset of user_administrator_rights.
	// If not specified, no additional restrictions are applied.
	BotAdministratorRights *ChatAdministratorRights `json:"bot_administrator_rights,omitempty"`

	// Pass True to request a chat with the bot as a member. Otherwise, no additional restrictions are applied.
	BotIsMember bool `json:"bot_is_member,omitempty"`

	// Pass True to request the chat's title
	RequestTitle bool `json:"request_title,omitempty"`

	// Pass True to request the chat's username
	RequestUsername bool `json:"request_username,omitempty"`

	// Pass True to request the chat's photo
	RequestPhoto bool `json:"request_photo,omitempty"`
}

// KeyboardButtonPollType represents the type of poll to be created
// https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	// Optional.
	// If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	// If regular is passed, only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	Type string `json:"type,omitempty"`
}

type ChatAdministratorRights struct {
	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`
}

// KeyboardButton is a button within a custom keyboard.
type KeyboardButton struct {
	Text            string                      `json:"text"`
	RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
	RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
	RequestContact  bool                        `json:"request_contact"`
	RequestLocation bool                        `json:"request_location"`
	RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`
	Webapp          *WebAppInfo                 `json:"web_app,omitempty"`
}

// Validate checks if the keyboard button is valid
func (j *KeyboardButton) Validate() error {
	if j.Text == "" {
		return errors.New("keyboard button requires 'text' field")
	}
	return nil
}

// WebAppInfo represents a web app to be opened with the button
// https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {

	// An HTTPS URL of a Web App to be opened with additional data
	// as specified in https://core.telegram.org/bots/webapps#initializing-mini-apps
	Url string `json:"url"`
}

func (v WebAppInfo) Validate() error {
	if v.Url == "" {
		return errors.New("url is empty")
	}
	return nil
}

// ReplyKeyboardHide allows the Bot to hide a custom keyboard.
type ReplyKeyboardHide struct {
	HideKeyboard bool `json:"hide_keyboard"`
	Selective    bool `json:"selective,omitempty"` // optional
}

// KeyboardType returns KeyboardTypeHide
func (_ *ReplyKeyboardHide) KeyboardType() botsgocore.KeyboardType {
	return botsgocore.KeyboardTypeHide
}

var _ botsgocore.Keyboard = (*ReplyKeyboardHide)(nil)

// InlineKeyboardMarkup is a custom keyboard presented for an inline bot.
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (v *InlineKeyboardMarkup) Validate() error {
	for _, row := range v.InlineKeyboard {
		for _, button := range row {
			if err := button.Validate(); err != nil {
				return err
			}
		}
	}
	return nil
}

// KeyboardType returns KeyboardTypeInline
func (*InlineKeyboardMarkup) KeyboardType() botsgocore.KeyboardType {
	return botsgocore.KeyboardTypeInline
}

var _ botsgocore.Keyboard = (*InlineKeyboardMarkup)(nil)

// LoginUrl represents a parameter of the inline keyboard button used to automatically authorize a user.
// https://core.telegram.org/bots/api#loginurl
type LoginUrl struct {
	// An HTTPS URL to be opened with user authorization data added to the query string when the button is pressed.
	// If the user refuses to provide authorization data,
	// the original URL without information about the user will be opened.
	// The data added is the same as described in Receiving authorization data.
	Url string `json:"url"`

	// Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`

	// Optional.
	// Username of a bot, which will be used for user authorization.
	// See Setting up a bot for more details.
	// If not specified, the current bot's username will be assumed.
	// The url's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot for more details.
	BotUsername string `json:"bot_username,omitempty"`

	// Optional. Pass True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

func (v LoginUrl) Validate() error {
	if v.Url == "" {
		return errors.New("url is empty")
	}
	if strings.TrimSpace(v.BotUsername) != v.BotUsername {
		return errors.New("bot_username must not have leading or trailing spaces")
	}
	return nil
}

// InlineKeyboardButton  represents one button of an inline keyboard.
// !!Exactly one of the optional fields must be used to specify type of the button.
// Note that some values are references as even an empty string will change behavior.
// Documentation: https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {

	// Label text on the button
	Text string `json:"text"`

	// Optional.
	// HTTP or tg:// URL to be opened when the button is pressed.
	// Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username,
	// if this is allowed by their privacy settings.
	URL string `json:"url,omitempty"`

	// Optional. Data to be sent in a callback query to the bot when the button is pressed, 1-64 bytes
	CallbackData string `json:"callback_data,omitempty"`

	// Optional. Description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	// Available only in private chats between a user and the bot. Not supported for messages sent on behalf of a Telegram Business account.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// Optional. An HTTPS URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	LoginUrl *LoginUrl `json:"login_url,omitempty"`

	// Optional. If set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// May be empty, in which case just the bot's username will be inserted.
	// Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"` // we use pointer as empty string is non zero value in this case

	// Optional. If set, pressing the button will insert the bot's username
	// and the specified inline query in the current chat's input field.
	// May be empty, in which case only the bot's username will be inserted.
	//
	// This offers a quick way for the user to open your bot
	// in inline mode in the same chat - good for selecting something from multiple options.
	// Not supported in channels and for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"` // we use pointer as empty string is non zero value in this case

	// Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// Not supported for messages sent on behalf of a Telegram Business account.
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`

	CopyText *CopyTextButton `json:"copy_text,omitempty"`

	// Optional. Description of the game that will be launched when the user presses the button.
	//
	//NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Optional. Specify True, to send a Pay button.
	//  Substrings “⭐” and “XTR” in the buttons's text will be replaced with a Telegram Star icon.
	//
	// NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
	Pay bool `json:"pay,omitempty"`
}

func (v InlineKeyboardButton) Validate() error {
	var populatedFields []string
	if v.URL != "" {
		populatedFields = append(populatedFields, "url")
	}
	if v.CallbackData != "" {
		populatedFields = append(populatedFields, "callback_data")
		if l := len(v.CallbackData); l > 64 {
			return fmt.Errorf("callback_data is too long, must be less than 64 bytes, got %d", l)
		}
	}
	if v.WebApp != nil {
		populatedFields = append(populatedFields, "web_app")
		if err := v.WebApp.Validate(); err != nil {
			return fmt.Errorf("WebAppInfo is invalid: %w", err)
		}
	}
	if v.LoginUrl != nil {
		populatedFields = append(populatedFields, "login_url")
		if err := v.LoginUrl.Validate(); err != nil {
			return fmt.Errorf("LoginUrl is invalid: %w", err)
		}
	}
	if v.SwitchInlineQuery != nil {
		populatedFields = append(populatedFields, "switch_inline_query")
	}
	if v.SwitchInlineQueryCurrentChat != nil {
		populatedFields = append(populatedFields, "switch_inline_query_current_chat")
	}
	if v.SwitchInlineQueryChosenChat != nil {
		populatedFields = append(populatedFields, "switch_inline_query_chosen_chat")
		if err := v.SwitchInlineQueryChosenChat.Validate(); err != nil {
			return fmt.Errorf("invalid field SwitchInlineQueryChosenChat: %w", err)
		}
	}
	if v.CallbackGame != nil {
		populatedFields = append(populatedFields, "callback_game")
		if err := v.CallbackGame.Validate(); err != nil {
			return fmt.Errorf("invaldi field CallbackGame: %w", err)
		}
	}
	if v.CopyText != nil {
		populatedFields = append(populatedFields, "copy_text")
		if err := v.CopyText.Validate(); err != nil {
			return fmt.Errorf("invalid field CopyText: %w", err)
		}
	}
	if len(populatedFields) != 1 {
		return fmt.Errorf("exactly one of the optional fields must be used to specify type of the button, got: %s", strings.Join(populatedFields, ", "))
	}
	return nil
}

// CopyTextButton represents an inline keyboard button that copies specified text to the clipboard.
type CopyTextButton struct {
	Text string `json:"text"`
}

func (v CopyTextButton) Validate() error {
	if v.Text == "" {
		return errors.New("text is empty")
	}
	return nil
}

// SwitchInlineQueryChosenChat represents an inline button that switches the current user to inline mode
// in a chosen chat, with an optional default inline query.
// Documentation: https://core.telegram.org/bots/api#switchinlinequerychosenchat
type SwitchInlineQueryChosenChat struct {
	// Optional.
	//The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	Query string `json:"query,omitempty"`

	// Optional. True, if private chats with users can be chosen
	AllowUserChats bool `json:"allow_user_chats"` // DO not omit this field, as it defaults to TRUE

	// Optional. True, if private chats with bots can be chosen
	AllowBotChats bool `json:"allow_bot_chats"` // DO not omit this field, as it defaults to TRUE

	// Optional. True, if group and supergroup chats can be chosen
	AllowGroupChats bool `json:"allow_group_chats"` // DO not omit this field, as it defaults to TRUE

	// Optional. True, if channel chats can be chosen
	AllowChannelChats bool `json:"allow_channel_chats"` // DO not omit this field, as it defaults to TRUE
}

func (v SwitchInlineQueryChosenChat) Validate() error {
	return nil
}

// CallbackGame is a placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct {
	// A placeholder, currently holds no information. Use BotFather to set up your game.
}

func (v CallbackGame) Validate() error {
	return nil
}

// ForceReply allows the Bot to have users directly reply to it without
// additional interaction.
type ForceReply struct {
	ForceReply bool `json:"force_reply"`
	Selective  bool `json:"selective,omitempty"` // optional
}

// KeyboardType returns KeyboardTypeForceReply
func (ForceReply) KeyboardType() botsgocore.KeyboardType {
	return botsgocore.KeyboardTypeForceReply
}

var _ botsgocore.Keyboard = (*ForceReply)(nil)

// ChosenInlineResult is an inline query result chosen by a User
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}
