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
	"time"
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

// Update is an update response, from GetUpdates.
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
}

// Chat provides chat struct for the update
func (update Update) Chat() *Chat {
	switch {
	case update.Message != nil:
		return update.Message.Chat
	case update.EditedMessage != nil:
		return update.EditedMessage.Chat
	case update.ChannelPost != nil:
		return update.ChannelPost.Chat
	case update.EditedChannelPost != nil:
		return update.EditedChannelPost.Chat
	case update.CallbackQuery != nil:
		if update.CallbackQuery.Message != nil {
			return update.CallbackQuery.Message.Chat
		}
	}
	return nil
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

// Message is returned by almost every request, and contains data about almost anything.
type Message struct {
	MessageID             int              `json:"message_id"`
	From                  *User            `json:"from,omitempty"` // optional
	Date                  int              `json:"date"`
	Chat                  *Chat            `json:"chat,omitempty"`
	ForwardFrom           *User            `json:"forward_from,omitempty"`            // optional
	ForwardDate           int              `json:"forward_date,omitempty"`            // optional
	ReplyToMessage        *Message         `json:"reply_to_message,omitempty"`        // optional
	Text                  string           `json:"text,omitempty"`                    // optional
	Entities              *[]MessageEntity `json:"entities,omitempty"`                // optional
	Audio                 *Audio           `json:"audio,omitempty"`                   // optional
	Document              *Document        `json:"document,omitempty"`                // optional
	Photo                 *[]PhotoSize     `json:"photo,omitempty"`                   // optional
	Sticker               *Sticker         `json:"sticker,omitempty"`                 // optional
	Video                 *Video           `json:"video,omitempty"`                   // optional
	Voice                 *Voice           `json:"voice,omitempty"`                   // optional
	Caption               string           `json:"caption,omitempty"`                 // optional
	Contact               *Contact         `json:"contact,omitempty"`                 // optional
	Location              *Location        `json:"location,omitempty"`                // optional
	Venue                 *Venue           `json:"venue,omitempty"`                   // optional
	NewChatParticipant    *ChatMember      `json:"new_chat_participant,omitempty"`    // Obsolete
	NewChatMember         *ChatMember      `json:"new_chat_member,omitempty"`         // Obsolete
	NewChatMembers        []ChatMember     `json:"new_chat_members,omitempty"`        // optional
	LeftChatMember        *ChatMember      `json:"left_chat_member,omitempty"`        // optional
	NewChatTitle          string           `json:"new_chat_title,omitempty"`          // optional
	NewChatPhoto          *[]PhotoSize     `json:"new_chat_photo,omitempty"`          // optional
	DeleteChatPhoto       bool             `json:"delete_chat_photo,omitempty"`       // optional
	GroupChatCreated      bool             `json:"group_chat_created,omitempty"`      // optional
	SuperGroupChatCreated bool             `json:"supergroup_chat_created,omitempty"` // optional
	ChannelChatCreated    bool             `json:"channel_chat_created,omitempty"`    // optional
	MigrateToChatID       int64            `json:"migrate_to_chat_id,omitempty"`      // optional
	MigrateFromChatID     int64            `json:"migrate_from_chat_id,omitempty"`    // optional
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`          // optional
}

// Time converts the message timestamp into a Time.
func (m *Message) Time() time.Time {
	return time.Unix(int64(m.Date), 0)
}

// IsCommand returns true if message starts with '/'.
func (m *Message) IsCommand() bool {
	return m.Text != "" && m.Text[0] == '/'
}

// Command checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at bot syntax, it removes the bot name.
func (m *Message) Command() string {
	if !m.IsCommand() {
		return ""
	}

	command := strings.SplitN(m.Text, " ", 2)[0][1:]

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

// CommandArguments checks if the message was a command and if it was,
// returns all text after the command name. If the Message was not a
// command, it returns an empty string.
func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	split := strings.SplitN(m.Text, " ", 2)
	if len(split) != 2 {
		return ""
	}

	return strings.SplitN(m.Text, " ", 2)[1]
}

// MessageEntity contains information about data in a Message.
type MessageEntity struct {
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url,omitempty"` // optional
}

// ParseURL attempts to parse a URL contained within a MessageEntity.
func (entity MessageEntity) ParseURL() (*url.URL, error) {
	if entity.URL == "" {
		return nil, errors.New(ErrBadURL)
	}

	return url.Parse(entity.URL)
}

// PhotoSize contains information about photos.
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"` // optional
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
//
// Note that LastName and UserID may be empty.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"` // optional
	UserID      int    `json:"user_id,omitempty"`   // optional
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

	UserIsBot       bool `json:"user_is_bot,omitempty"`
	UserIsPremium   bool `json:"user_is_premium,omitempty"`
	MaxQuantity     int  `json:"max_quantity,omitempty"` // The maximum number of users to be selected; 1-10. Defaults to 1.
	RequestName     bool `json:"request_name,omitempty"`
	RequestUsername bool `json:"request_username,omitempty"`
	RequestPhoto    bool `json:"request_photo,omitempty"`
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
	Webapp          *WebappInfo                 `json:"web_app,omitempty"`
}

// Validate checks if the keyboard button is valid
func (j *KeyboardButton) Validate() error {
	if j.Text == "" {
		return errors.New("keyboard button requires 'text' field")
	}
	return nil
}

// WebappInfo represents a web app to be opened with the button
// https://core.telegram.org/bots/api#webappinfo
type WebappInfo struct {
	Url string `json:"url"`
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

// InlineKeyboardButton is a button within a custom keyboard for inline query responses.
// Note that some values are references as even an empty string will change behavior.
// Documentation: https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	URL                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *WebappInfo                  `json:"web_app,omitempty"`
	LoginUrl                     *LoginUrl                    `json:"login_url,omitempty"`
	SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`              // we use pointer as empty string is non zero value in this case
	SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"` // we use pointer as empty string is non zero value in this case
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}

// SwitchInlineQueryChosenChat represents an inline button that switches the current user to inline mode
// in a chosen chat, with an optional default inline query.
// Documentation: https://core.telegram.org/bots/api#switchinlinequerychosenchat
type SwitchInlineQueryChosenChat struct {
	// Optional.
	//The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	Query string `json:"query,omitempty"`

	// Optional. True, if private chats with users can be chosen
	AllowUserChats bool `json:"allow_user_chats,omitempty"`

	// Optional. True, if private chats with bots can be chosen
	AllowBotChats bool `json:"allow_bot_chats,omitempty"`

	// Optional. True, if group and supergroup chats can be chosen
	AllowGroupChats bool `json:"allow_group_chats,omitempty"`

	// Optional. True, if channel chats can be chosen
	AllowChannelChats bool `json:"allow_channel_chats,omitempty"`
}

// CallbackGame is a placeholder, currently holds no information. Use BotFather to set up your game.
type CallbackGame struct {
	// A placeholder, currently holds no information. Use BotFather to set up your game.
}

// CallbackQuery is data sent when a keyboard button with callback data
// is clicked.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`           // optional
	ChatInstance    string   `json:"chat_instance,omitempty"`     // optional
	InlineMessageID string   `json:"inline_message_id,omitempty"` // optional
	Data            string   `json:"data,omitempty"`              // optional
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

// InlineQuery is a Query from Telegram for an inline request.
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"` // optional
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// InlineQueryResultArticle is an inline query response article.
type InlineQueryResultArticle struct {
	Type                string                `json:"type"`                  // required
	ID                  string                `json:"id"`                    // required
	Title               string                `json:"title"`                 // required
	InputMessageContent interface{}           `json:"input_message_content"` // required
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 string                `json:"url,omitempty"`
	HideURL             bool                  `json:"hide_url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	ThumbWidth          int                   `json:"thumb_width,omitempty"`
	ThumbHeight         int                   `json:"thumb_height,omitempty"`
}

// InlineQueryResultPhoto is an inline query response photo.
type InlineQueryResultPhoto struct {
	Type                string                `json:"type"`      // required
	ID                  string                `json:"id"`        // required
	URL                 string                `json:"photo_url"` // required
	MimeType            string                `json:"mime_type,omitempty"`
	Width               int                   `json:"photo_width,omitempty"`
	Height              int                   `json:"photo_height,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// InlineQueryResultGIF is an inline query response GIF.
type InlineQueryResultGIF struct {
	Type                string                `json:"type"`    // required
	ID                  string                `json:"id"`      // required
	URL                 string                `json:"gif_url"` // required
	Width               int                   `json:"gif_width,omitempty"`
	Height              int                   `json:"gif_height,omitempty"`
	ThumbURL            string                `json:"thumb_url,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// InlineQueryResultMPEG4GIF is an inline query response MPEG4 GIF.
type InlineQueryResultMPEG4GIF struct {
	Type                string                `json:"type"`      // required
	ID                  string                `json:"id"`        // required
	URL                 string                `json:"mpeg4_url"` // required
	Width               int                   `json:"mpeg4_width"`
	Height              int                   `json:"mpeg4_height"`
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// InlineQueryResultVideo is an inline query response video.
type InlineQueryResultVideo struct {
	Type                string                `json:"type"`      // required
	ID                  string                `json:"id"`        // required
	URL                 string                `json:"video_url"` // required
	MimeType            string                `json:"mime_type"` // required
	ThumbURL            string                `json:"thumb_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption"`
	Width               int                   `json:"video_width"`
	Height              int                   `json:"video_height"`
	Duration            int                   `json:"video_duration"`
	Description         string                `json:"description"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// InlineQueryResultAudio is an inline query response audio.
type InlineQueryResultAudio struct {
	Type                string                `json:"type"`      // required
	ID                  string                `json:"id"`        // required
	URL                 string                `json:"audio_url"` // required
	Title               string                `json:"title"`     // required
	Performer           string                `json:"performer"`
	Duration            int                   `json:"audio_duration"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// InlineQueryResultVoice is an inline query response voice.
type InlineQueryResultVoice struct {
	Type                string                `json:"type"`      // required
	ID                  string                `json:"id"`        // required
	URL                 string                `json:"voice_url"` // required
	Title               string                `json:"title"`     // required
	Duration            int                   `json:"voice_duration"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// InlineQueryResultDocument is an inline query response document.
type InlineQueryResultDocument struct {
	Type                string                `json:"type"`  // required
	ID                  string                `json:"id"`    // required
	Title               string                `json:"title"` // required
	Caption             string                `json:"caption"`
	URL                 string                `json:"document_url"` // required
	MimeType            string                `json:"mime_type"`    // required
	Description         string                `json:"description"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	ThumbWidth          int                   `json:"thumb_width"`
	ThumbHeight         int                   `json:"thumb_height"`
}

// InlineQueryResultLocation is an inline query response location.
type InlineQueryResultLocation struct {
	Type                string                `json:"type"`      // required
	ID                  string                `json:"id"`        // required
	Latitude            float64               `json:"latitude"`  // required
	Longitude           float64               `json:"longitude"` // required
	Title               string                `json:"title"`     // required
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ThumbURL            string                `json:"thumb_url"`
	ThumbWidth          int                   `json:"thumb_width"`
	ThumbHeight         int                   `json:"thumb_height"`
}

// ChosenInlineResult is an inline query result chosen by a User
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

// InputTextMessageContent contains text for displaying
// as an inline query result.
type InputTextMessageContent struct {
	Text                  string `json:"message_text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

// InputLocationMessageContent contains a location for displaying
// as an inline query result.
type InputLocationMessageContent struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// InputVenueMessageContent contains a venue for displaying
// as an inline query result.
type InputVenueMessageContent struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareID string  `json:"foursquare_id"`
}

// InputContactMessageContent contains a contact for displaying
// as an inline query result.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}
