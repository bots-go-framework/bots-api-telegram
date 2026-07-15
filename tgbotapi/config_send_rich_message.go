package tgbotapi

import (
	"fmt"
	"net/url"
	"strconv"
)

var _ Sendable = (*RichMessageConfig)(nil)

// RichMessageConfig contains information about a sendRichMessage request.
//
// https://core.telegram.org/bots/api#sendrichmessage
type RichMessageConfig struct {
	BaseChat

	// The message to be sent
	RichMessage InputRichMessage `json:"rich_message"`
}

// Values returns url.Values representation of RichMessageConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v RichMessageConfig) Values() (url.Values, error) {
	values, err := v.BaseChat.Values()
	if err != nil {
		return values, err
	}

	if b, err := encodeToJson(v.RichMessage); err != nil {
		return values, fmt.Errorf("failed to marshal rich message as JSON: %w", err)
	} else {
		values.Add("rich_message", string(b))
	}

	return values, nil
}

// TelegramMethod returns Telegram API method name for sending a RichMessage.
//
//goland:noinspection GoMixedReceiverTypes
func (RichMessageConfig) TelegramMethod() string {
	return "sendRichMessage"
}

var _ Sendable = (*RichMessageDraftConfig)(nil)

// RichMessageDraftConfig contains information about a sendRichMessageDraft request, used to stream a
// partial rich message to a user while it is being generated. The streamed draft is ephemeral and acts
// as a temporary 30-second preview - once the output is finalized, RichMessageConfig must be sent with
// the complete message to persist it in the user's chat. Returns True on success, so this should be sent
// via BotAPI.SendCustomMessage rather than BotAPI.Send.
//
// https://core.telegram.org/bots/api#sendrichmessagedraft
type RichMessageDraftConfig struct {
	// Unique identifier for the target private chat
	ChatID int64 `json:"chat_id"`

	// Optional. Unique identifier for the target message thread
	MessageThreadID int64 `json:"message_thread_id,omitempty"`

	// Unique identifier of the message draft; must be non-zero. Changes to drafts with the same
	// identifier are animated.
	DraftID int64 `json:"draft_id"`

	// The partial message to be streamed. Direct upload of new files isn't supported.
	RichMessage InputRichMessage `json:"rich_message"`
}

// Values returns url.Values representation of RichMessageDraftConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v RichMessageDraftConfig) Values() (url.Values, error) {
	values := url.Values{}
	values.Add("chat_id", strconv.FormatInt(v.ChatID, 10))
	if v.MessageThreadID != 0 {
		values.Add("message_thread_id", strconv.FormatInt(v.MessageThreadID, 10))
	}
	values.Add("draft_id", strconv.FormatInt(v.DraftID, 10))

	if b, err := encodeToJson(v.RichMessage); err != nil {
		return values, fmt.Errorf("failed to marshal rich message draft as JSON: %w", err)
	} else {
		values.Add("rich_message", string(b))
	}

	return values, nil
}

// TelegramMethod returns Telegram API method name for streaming a RichMessage draft.
//
//goland:noinspection GoMixedReceiverTypes
func (RichMessageDraftConfig) TelegramMethod() string {
	return "sendRichMessageDraft"
}
