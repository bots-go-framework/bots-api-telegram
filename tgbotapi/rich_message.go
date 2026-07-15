package tgbotapi

import "errors"

// RichMessage represents a rich formatted message (Bot API 10.1 Rich Messages).
//
// https://core.telegram.org/bots/api#richmessage
type RichMessage struct {
	// Content of the message
	Blocks []RichBlock `json:"blocks"`

	// Optional. True, if the rich message must be shown right-to-left
	IsRTL bool `json:"is_rtl,omitempty"`
}

// InputRichMessage describes a rich message to be sent (Bot API 10.1 Rich Messages). Exactly one of
// HTML or Markdown must be used.
//
// Note: the live Bot API docs additionally list a "blocks" field (Array of InputRichBlock) and a
// "media" field (Array of InputRichMessageMedia) on this class, added together with the InputRichBlock*
// family in Bot API 10.2. Those are intentionally not modeled here; this type only implements the
// Bot API 10.1 shape (html/markdown/is_rtl/skip_entity_detection).
//
// https://core.telegram.org/bots/api#inputrichmessage
type InputRichMessage struct {
	// Optional. Content of the rich message to send described using HTML formatting. See rich message
	// formatting options for more details.
	HTML string `json:"html,omitempty"`

	// Optional. Content of the rich message to send described using Markdown formatting. See rich
	// message formatting options for more details.
	Markdown string `json:"markdown,omitempty"`

	// Optional. Pass True if the rich message must be shown right-to-left
	IsRTL bool `json:"is_rtl,omitempty"`

	// Optional. Pass True to skip automatic detection of entities (e.g., URLs, email addresses,
	// username mentions, hashtags, cashtags, bot commands, or phone numbers) in the text
	SkipEntityDetection bool `json:"skip_entity_detection,omitempty"`
}

// Validate checks that exactly one of HTML or Markdown is set.
func (v InputRichMessage) Validate() error {
	if v.HTML == "" && v.Markdown == "" {
		return errors.New("one of HTML or Markdown must be set")
	}
	if v.HTML != "" && v.Markdown != "" {
		return errors.New("only one of HTML or Markdown may be set")
	}
	return nil
}

var _ InputMessageContent = (*InputRichMessageContent)(nil)

// InputRichMessageContent represents the content of a rich message to be sent as the result of an
// inline, guest, or Web App query (Bot API 10.1 Rich Messages).
//
// https://core.telegram.org/bots/api#inputrichmessagecontent
type InputRichMessageContent struct {
	inputMessageContentBase

	// The message to be sent
	RichMessage InputRichMessage `json:"rich_message"`
}

func (v InputRichMessageContent) Validate() error {
	return v.RichMessage.Validate()
}
