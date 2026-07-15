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

// InputRichMessage describes a rich message to be sent (Bot API 10.1 Rich Messages, extended in Bot API
// 10.2 with Blocks and Media). Exactly one of HTML, Markdown, or Blocks must be used.
//
// https://core.telegram.org/bots/api#inputrichmessage
type InputRichMessage struct {
	// Optional. Content of the rich message to send described as a list of blocks. Bot API 10.2+
	Blocks []InputRichBlock `json:"blocks,omitempty"`

	// Optional. Content of the rich message to send described using HTML formatting. See rich message
	// formatting options for more details. Use Media to specify the media used in the message.
	HTML string `json:"html,omitempty"`

	// Optional. Content of the rich message to send described using Markdown formatting. See rich
	// message formatting options for more details. Use Media to specify the media used in the message.
	Markdown string `json:"markdown,omitempty"`

	// Optional. List of media referenced in HTML or Markdown using tg://photo?id=, tg://video?id=, and
	// tg://audio?id= links. Bot API 10.2+
	Media []InputRichMessageMedia `json:"media,omitempty"`

	// Optional. Pass True if the rich message must be shown right-to-left
	IsRTL bool `json:"is_rtl,omitempty"`

	// Optional. Pass True to skip automatic detection of entities (e.g., URLs, email addresses,
	// username mentions, hashtags, cashtags, bot commands, or phone numbers) in the text
	SkipEntityDetection bool `json:"skip_entity_detection,omitempty"`
}

// Validate checks that exactly one of HTML, Markdown, or Blocks is set.
func (v InputRichMessage) Validate() error {
	set := 0
	if v.HTML != "" {
		set++
	}
	if v.Markdown != "" {
		set++
	}
	if len(v.Blocks) > 0 {
		set++
	}
	switch set {
	case 0:
		return errors.New("one of HTML, Markdown, or Blocks must be set")
	case 1:
		return nil
	default:
		return errors.New("only one of HTML, Markdown, or Blocks may be set")
	}
}

// InputRichMessageMedia describes a media element embedded in an outgoing rich message, referenced from
// InputRichMessage's HTML or Markdown content via a tg://photo?id=, tg://video?id=, or tg://audio?id=
// link (Bot API 10.2 Rich Messages).
//
// https://core.telegram.org/bots/api#inputrichmessagemedia
type InputRichMessageMedia struct {
	// Unique identifier of the media used in a tg://photo?id=, tg://video?id=, or tg://audio?id= link.
	// 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
	ID string `json:"id"`

	// The media to be sent. Must be one of *InputMediaAnimation, *InputMediaAudio, *InputMediaPhoto,
	// *InputMediaVideo, or *InputMediaVoiceNote. Everything except the media file itself and its
	// type-specific properties (e.g. width/height/duration) is ignored - in particular, any caption set
	// on the referenced InputMedia* value is ignored.
	Media any `json:"media"`
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
