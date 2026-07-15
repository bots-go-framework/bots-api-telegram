package tgbotapi

import (
	"fmt"
	"net/url"
	"strconv"
)

// baseEphemeralMessageEdit carries the addressing fields shared by every editEphemeralMessage* method
// and deleteEphemeralMessage (Bot API 10.2 Ephemeral Messages): the group/supergroup chat, the user who
// received the ephemeral message, and the ephemeral message's identifier within that chat.
type baseEphemeralMessageEdit struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format
	// @username)
	ChatID          int64  `json:"-"`
	ChannelUsername string `json:"-"`

	// Identifier of the user who received the message
	ReceiverUserID int64 `json:"receiver_user_id"`

	// Identifier of the ephemeral message to edit/delete
	EphemeralMessageID int64 `json:"ephemeral_message_id"`

	// Optional. A JSON-serialized object for an inline keyboard
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Values returns url.Values representation of the fields shared by every ephemeral message edit/delete
// method.
//
//goland:noinspection GoMixedReceiverTypes
func (v baseEphemeralMessageEdit) Values() (url.Values, error) {
	values := url.Values{}
	if v.ChannelUsername != "" {
		values.Add("chat_id", v.ChannelUsername)
	} else {
		values.Add("chat_id", strconv.FormatInt(v.ChatID, 10))
	}
	values.Add("receiver_user_id", strconv.FormatInt(v.ReceiverUserID, 10))
	values.Add("ephemeral_message_id", strconv.FormatInt(v.EphemeralMessageID, 10))

	if v.ReplyMarkup != nil {
		data, err := encodeToJson(v.ReplyMarkup)
		if err != nil {
			return values, err
		}
		values.Add("reply_markup", string(data))
	}

	return values, nil
}

// EditEphemeralMessageTextConfig allows you to modify the text of an ephemeral message (Bot API 10.2
// Ephemeral Messages).
//
// https://core.telegram.org/bots/api#editephemeralmessagetext
type EditEphemeralMessageTextConfig struct {
	baseEphemeralMessageEdit

	// New text of the message, 1-4096 characters after entity parsing
	Text string `json:"text"`

	// Optional. Mode for parsing entities in the message text
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. A JSON-serialized list of special entities that appear in message text, which can be
	// specified instead of ParseMode
	Entities []MessageEntity `json:"entities,omitempty"`

	// Optional. Link preview generation options for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

// Values returns URL values representation of EditEphemeralMessageTextConfig
//
//goland:noinspection GoMixedReceiverTypes
func (v EditEphemeralMessageTextConfig) Values() (url.Values, error) {
	values, err := v.baseEphemeralMessageEdit.Values()
	if err != nil {
		return values, err
	}

	values.Add("text", v.Text)
	if v.ParseMode != "" {
		values.Add("parse_mode", v.ParseMode)
	}
	if len(v.Entities) > 0 {
		data, err := encodeToJson(v.Entities)
		if err != nil {
			return values, fmt.Errorf("failed to marshal entities as JSON: %w", err)
		}
		values.Add("entities", string(data))
	}
	if v.LinkPreviewOptions != nil {
		data, err := encodeToJson(v.LinkPreviewOptions)
		if err != nil {
			return values, fmt.Errorf("failed to marshal link_preview_options as JSON: %w", err)
		}
		values.Add("link_preview_options", string(data))
	}

	return values, nil
}

//goland:noinspection GoMixedReceiverTypes
func (EditEphemeralMessageTextConfig) TelegramMethod() string {
	return "editEphemeralMessageText"
}

var _ Sendable = EditEphemeralMessageTextConfig{}

// EditEphemeralMessageMediaConfig allows you to modify the media of an ephemeral message (Bot API 10.2
// Ephemeral Messages). A new file can't be uploaded; use a previously uploaded file via its file_id or
// specify a URL.
//
// https://core.telegram.org/bots/api#editephemeralmessagemedia
type EditEphemeralMessageMediaConfig struct {
	baseEphemeralMessageEdit

	// A JSON-serialized object for the new media content of the message. Must be one of
	// *InputMediaAnimation, *InputMediaAudio, *InputMediaPhoto, or *InputMediaVideo.
	Media any `json:"media"`
}

// Values returns URL values representation of EditEphemeralMessageMediaConfig
//
//goland:noinspection GoMixedReceiverTypes
func (v EditEphemeralMessageMediaConfig) Values() (url.Values, error) {
	values, err := v.baseEphemeralMessageEdit.Values()
	if err != nil {
		return values, err
	}

	data, err := encodeToJson(v.Media)
	if err != nil {
		return values, fmt.Errorf("failed to marshal media as JSON: %w", err)
	}
	values.Add("media", string(data))

	return values, nil
}

//goland:noinspection GoMixedReceiverTypes
func (EditEphemeralMessageMediaConfig) TelegramMethod() string {
	return "editEphemeralMessageMedia"
}

var _ Sendable = EditEphemeralMessageMediaConfig{}

// EditEphemeralMessageCaptionConfig allows you to modify the caption of an ephemeral message (Bot API
// 10.2 Ephemeral Messages).
//
// https://core.telegram.org/bots/api#editephemeralmessagecaption
type EditEphemeralMessageCaptionConfig struct {
	baseEphemeralMessageEdit

	// Optional. New caption of the message, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the message caption
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. A JSON-serialized list of special entities that appear in the caption, which can be
	// specified instead of ParseMode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}

// Values returns URL values representation of EditEphemeralMessageCaptionConfig
//
//goland:noinspection GoMixedReceiverTypes
func (v EditEphemeralMessageCaptionConfig) Values() (url.Values, error) {
	values, err := v.baseEphemeralMessageEdit.Values()
	if err != nil {
		return values, err
	}

	values.Add("caption", v.Caption)
	if v.ParseMode != "" {
		values.Add("parse_mode", v.ParseMode)
	}
	if len(v.CaptionEntities) > 0 {
		data, err := encodeToJson(v.CaptionEntities)
		if err != nil {
			return values, fmt.Errorf("failed to marshal caption_entities as JSON: %w", err)
		}
		values.Add("caption_entities", string(data))
	}

	return values, nil
}

//goland:noinspection GoMixedReceiverTypes
func (EditEphemeralMessageCaptionConfig) TelegramMethod() string {
	return "editEphemeralMessageCaption"
}

var _ Sendable = EditEphemeralMessageCaptionConfig{}

// EditEphemeralMessageReplyMarkupConfig allows you to modify the reply markup of an ephemeral message
// (Bot API 10.2 Ephemeral Messages).
//
// https://core.telegram.org/bots/api#editephemeralmessagereplymarkup
type EditEphemeralMessageReplyMarkupConfig struct {
	baseEphemeralMessageEdit
}

//goland:noinspection GoMixedReceiverTypes
func (v EditEphemeralMessageReplyMarkupConfig) Values() (url.Values, error) {
	return v.baseEphemeralMessageEdit.Values()
}

//goland:noinspection GoMixedReceiverTypes
func (EditEphemeralMessageReplyMarkupConfig) TelegramMethod() string {
	return "editEphemeralMessageReplyMarkup"
}

var _ Sendable = EditEphemeralMessageReplyMarkupConfig{}

// DeleteEphemeralMessageConfig deletes an ephemeral message (Bot API 10.2 Ephemeral Messages).
//
// https://core.telegram.org/bots/api#deleteephemeralmessage
type DeleteEphemeralMessageConfig struct {
	baseEphemeralMessageEdit
}

//goland:noinspection GoMixedReceiverTypes
func (v DeleteEphemeralMessageConfig) Values() (url.Values, error) {
	return v.baseEphemeralMessageEdit.Values()
}

//goland:noinspection GoMixedReceiverTypes
func (DeleteEphemeralMessageConfig) TelegramMethod() string {
	return "deleteEphemeralMessage"
}

var _ Sendable = DeleteEphemeralMessageConfig{}
