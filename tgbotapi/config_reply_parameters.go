package tgbotapi

import (
	"encoding/json"
	"errors"
	"fmt"
)

// ReplyParameters describes reply parameters for the message that is being sent.
type ReplyParameters struct {
	// Optional. Identifier of the message that will be replied to in the current chat, or in the chat
	// ChatIDInt/ChatIDStr if specified. Required if EphemeralMessageID isn't specified.
	MessageID int64 `json:"message_id,omitempty"`

	// Optional. If the message to be replied to is from a different chat, unique identifier for the chat or username of the channel (in the format @channelusername). Not supported for messages sent on behalf of a business account.
	ChatIDInt int64  `json:"-"`
	ChatIDStr string `json:"-"`

	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"` // Optional. Pass True if the message should be sent even if the specified message to be replied to is not found. Always False for replies in another chat or forum topic. Always True for messages sent on behalf of a business account.
	Quote                    string          `json:"quote,omitempty"`                       // Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including bold, italic, underline, strikethrough, spoiler, and custom_emoji entities. The message will fail to send if the quote isn't found in the original message.
	QuoteParseMode           string          `json:"quote_parse_mode,omitempty"`            // Optional. Mode for parsing entities in the quote. See formatting options for more details.
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`              // Optional. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of quote_parse_mode.
	QuotePosition            int64           `json:"quote_position,omitempty"`              // Optional. Position of the quote in the original message in UTF-16 code units

	// Optional. Persistent identifier of the poll option to reply to. Bot API 9.6+
	PollOptionID string `json:"poll_option_id,omitempty"`

	// Optional. Identifier of the incoming ephemeral message that will be replied to in the current
	// chat. A reply to an ephemeral message must itself be an ephemeral message. An ephemeral message
	// may only be replied to within 15 seconds of being sent. Required if MessageID isn't specified;
	// MessageID and EphemeralMessageID are mutually exclusive - set exactly one. Bot API 10.2+
	EphemeralMessageID int64 `json:"ephemeral_message_id,omitempty"`
}

// Validate checks the mutually exclusive reply target and formatting fields.
func (v ReplyParameters) Validate() error {
	if v.MessageID == 0 && v.EphemeralMessageID == 0 {
		return errors.New("message_id or ephemeral_message_id is required")
	}
	if v.MessageID != 0 && v.EphemeralMessageID != 0 {
		return errors.New("message_id and ephemeral_message_id are mutually exclusive")
	}
	if v.ChatIDInt != 0 && v.ChatIDStr != "" {
		return errors.New("ChatIDInt and ChatIDStr are mutually exclusive")
	}
	if v.QuoteParseMode != "" && len(v.QuoteEntities) > 0 {
		return errors.New("quote_parse_mode and quote_entities are mutually exclusive")
	}
	return nil
}

type replyParametersAlias ReplyParameters

// MarshalJSON emits the polymorphic chat_id field under Telegram's exact wire
// name while retaining the historical typed ChatIDInt/ChatIDStr Go API.
func (v ReplyParameters) MarshalJSON() ([]byte, error) {
	if err := v.Validate(); err != nil {
		return nil, err
	}
	raw, err := json.Marshal(replyParametersAlias(v))
	if err != nil {
		return nil, err
	}
	if v.ChatIDInt == 0 && v.ChatIDStr == "" {
		return raw, nil
	}
	var object map[string]json.RawMessage
	if err = json.Unmarshal(raw, &object); err != nil {
		return nil, err
	}
	if v.ChatIDStr != "" {
		object["chat_id"], err = json.Marshal(v.ChatIDStr)
	} else {
		object["chat_id"], err = json.Marshal(v.ChatIDInt)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to marshal reply chat_id: %w", err)
	}
	return json.Marshal(object)
}
