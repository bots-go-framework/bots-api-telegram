package tgbotapi

import (
	"fmt"
	"net/url"
	"strconv"
)

// MessageDraftConfig streams a temporary text preview while a response is generated.
// Text may be empty to display Telegram's native “Thinking…” placeholder.
type MessageDraftConfig struct {
	ChatID          int64           `json:"chat_id"`
	MessageThreadID int64           `json:"message_thread_id,omitempty"`
	DraftID         int64           `json:"draft_id"`
	Text            string          `json:"text,omitempty"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	Entities        []MessageEntity `json:"entities,omitempty"`
}

func (v MessageDraftConfig) Values() (url.Values, error) {
	if v.ChatID == 0 {
		return nil, fmt.Errorf("chat_id is required")
	}
	if v.DraftID == 0 {
		return nil, fmt.Errorf("draft_id must be non-zero")
	}
	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(v.ChatID, 10))
	if v.MessageThreadID != 0 {
		values.Add("message_thread_id", strconv.FormatInt(v.MessageThreadID, 10))
	}
	values.Add("draft_id", strconv.FormatInt(v.DraftID, 10))
	// text is intentionally always included: an empty value requests Telegram's
	// native Thinking placeholder.
	values.Add("text", v.Text)
	if v.ParseMode != "" {
		values.Add("parse_mode", v.ParseMode)
	}
	if len(v.Entities) > 0 {
		data, err := encodeToJson(v.Entities)
		if err != nil {
			return nil, err
		}
		values.Add("entities", string(data))
	}
	return values, nil
}

func (MessageDraftConfig) TelegramMethod() string {
	return "sendMessageDraft"
}

var _ Sendable = MessageDraftConfig{}
