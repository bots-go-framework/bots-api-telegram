package tgbotapi

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Rich text type discriminators for RichText.Type.
// https://core.telegram.org/bots/api#richtext
const (
	RichTextTypeBold                   = "bold"
	RichTextTypeItalic                 = "italic"
	RichTextTypeUnderline              = "underline"
	RichTextTypeStrikethrough          = "strikethrough"
	RichTextTypeSpoiler                = "spoiler"
	RichTextTypeDateTime               = "date_time"
	RichTextTypeTextMention            = "text_mention"
	RichTextTypeSubscript              = "subscript"
	RichTextTypeSuperscript            = "superscript"
	RichTextTypeMarked                 = "marked"
	RichTextTypeCode                   = "code"
	RichTextTypeCustomEmoji            = "custom_emoji"
	RichTextTypeMathematicalExpression = "mathematical_expression"
	RichTextTypeUrl                    = "url"
	RichTextTypeEmailAddress           = "email_address"
	RichTextTypePhoneNumber            = "phone_number"
	RichTextTypeBankCardNumber         = "bank_card_number"
	RichTextTypeMention                = "mention"
	RichTextTypeHashtag                = "hashtag"
	RichTextTypeCashtag                = "cashtag"
	RichTextTypeBotCommand             = "bot_command"
	RichTextTypeAnchor                 = "anchor"
	RichTextTypeAnchorLink             = "anchor_link"
	RichTextTypeReference              = "reference"
	RichTextTypeReferenceLink          = "reference_link"
)

// RichText represents a rich formatted text (Bot API 10.1 Rich Messages). On the wire, RichText can be
// encoded as a bare JSON string (plain text), a JSON array of RichText, or a JSON object with a "type"
// discriminator identifying one of the RichText* variants (RichTextBold, RichTextItalic, RichTextUnderline,
// RichTextStrikethrough, RichTextSpoiler, RichTextDateTime, RichTextTextMention, RichTextSubscript,
// RichTextSuperscript, RichTextMarked, RichTextCode, RichTextCustomEmoji, RichTextMathematicalExpression,
// RichTextUrl, RichTextEmailAddress, RichTextPhoneNumber, RichTextBankCardNumber, RichTextMention,
// RichTextHashtag, RichTextCashtag, RichTextBotCommand, RichTextAnchor, RichTextAnchorLink, RichTextReference,
// RichTextReferenceLink).
//
// This follows the same single-flattened-struct-with-Type-discriminator convention used elsewhere in this
// package for unions (see PaidMedia, PollMedia), extended with a custom MarshalJSON/UnmarshalJSON pair to
// additionally support the plain-string and array wire encodings that RichText, uniquely among this
// package's unions, also allows.
//
// https://core.telegram.org/bots/api#richtext
type RichText struct {
	// Type of the rich text, e.g. "bold", "italic", "url", etc. Empty when this RichText is plain text
	// (see PlainText) or a list of nested rich texts (see Items).
	Type string `json:"type,omitempty"`

	// PlainText holds the value when this RichText is encoded on the wire as a bare JSON string rather
	// than an object or array. Not itself a Bot API JSON field; see MarshalJSON/UnmarshalJSON.
	PlainText string `json:"-"`

	// Items holds the value when this RichText is encoded on the wire as a JSON array of RichText.
	// Not itself a Bot API JSON field; see MarshalJSON/UnmarshalJSON.
	Items []RichText `json:"-"`

	// Text is the nested rich text wrapped by most format variants: RichTextBold, RichTextItalic,
	// RichTextUnderline, RichTextStrikethrough, RichTextSpoiler, RichTextDateTime, RichTextTextMention,
	// RichTextSubscript, RichTextSuperscript, RichTextMarked, RichTextCode, RichTextUrl,
	// RichTextEmailAddress, RichTextPhoneNumber, RichTextBankCardNumber, RichTextMention, RichTextHashtag,
	// RichTextCashtag, RichTextBotCommand, RichTextAnchorLink, RichTextReference, RichTextReferenceLink.
	Text *RichText `json:"text,omitempty"`

	// UnixTime: for RichTextDateTime, the Unix time associated with the entity.
	// https://core.telegram.org/bots/api#richtextdatetime
	UnixTime int `json:"unix_time,omitempty"`

	// DateTimeFormat: for RichTextDateTime, the string that defines the formatting of the date and time.
	DateTimeFormat string `json:"date_time_format,omitempty"`

	// User: for RichTextTextMention, the mentioned user.
	// https://core.telegram.org/bots/api#richtexttextmention
	User *User `json:"user,omitempty"`

	// CustomEmojiID: for RichTextCustomEmoji, the unique identifier of the custom emoji. Use
	// GetCustomEmojiStickers to get full information about the sticker.
	// https://core.telegram.org/bots/api#richtextcustomemoji
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`

	// AlternativeText: for RichTextCustomEmoji, the alternative emoji for the custom emoji.
	AlternativeText string `json:"alternative_text,omitempty"`

	// Expression: for RichTextMathematicalExpression, the expression in LaTeX format.
	// https://core.telegram.org/bots/api#richtextmathematicalexpression
	Expression string `json:"expression,omitempty"`

	// URL: for RichTextUrl, the URL of the link.
	// https://core.telegram.org/bots/api#richtexturl
	URL string `json:"url,omitempty"`

	// EmailAddress: for RichTextEmailAddress, the email address.
	EmailAddress string `json:"email_address,omitempty"`

	// PhoneNumber: for RichTextPhoneNumber, the phone number.
	PhoneNumber string `json:"phone_number,omitempty"`

	// BankCardNumber: for RichTextBankCardNumber, the bank card number.
	BankCardNumber string `json:"bank_card_number,omitempty"`

	// Username: for RichTextMention, the mentioned username.
	Username string `json:"username,omitempty"`

	// Hashtag: for RichTextHashtag, the hashtag.
	Hashtag string `json:"hashtag,omitempty"`

	// Cashtag: for RichTextCashtag, the cashtag.
	Cashtag string `json:"cashtag,omitempty"`

	// BotCommand: for RichTextBotCommand, the bot command.
	BotCommand string `json:"bot_command,omitempty"`

	// Name: for RichTextAnchor, the name of the anchor. For RichTextReference, the name of the reference.
	Name string `json:"name,omitempty"`

	// AnchorName: for RichTextAnchorLink, the name of the anchor being linked to. If empty, the link
	// brings back to the top of the message.
	AnchorName string `json:"anchor_name,omitempty"`

	// ReferenceName: for RichTextReferenceLink, the name of the reference being linked to.
	ReferenceName string `json:"reference_name,omitempty"`
}

// richTextAlias has the same fields as RichText but none of its methods, used to avoid infinite
// recursion in RichText's custom MarshalJSON/UnmarshalJSON when the object-shaped wire encoding applies.
type richTextAlias RichText

// MarshalJSON implements the three-shape RichText wire encoding: a bare JSON string for plain text,
// a JSON array for a list of nested rich texts, or a JSON object for a named format variant.
//
//goland:noinspection GoMixedReceiverTypes
func (r RichText) MarshalJSON() ([]byte, error) {
	if r.Type == "" {
		if r.Items != nil {
			return json.Marshal(r.Items)
		}
		return json.Marshal(r.PlainText)
	}
	return json.Marshal(richTextAlias(r))
}

// UnmarshalJSON implements the three-shape RichText wire encoding: a bare JSON string for plain text,
// a JSON array for a list of nested rich texts, or a JSON object for a named format variant.
//
//goland:noinspection GoMixedReceiverTypes
func (r *RichText) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || string(trimmed) == "null" {
		*r = RichText{}
		return nil
	}
	switch trimmed[0] {
	case '"':
		var s string
		if err := json.Unmarshal(trimmed, &s); err != nil {
			return fmt.Errorf("failed to unmarshal RichText plain text: %w", err)
		}
		*r = RichText{PlainText: s}
		return nil
	case '[':
		var items []RichText
		if err := json.Unmarshal(trimmed, &items); err != nil {
			return fmt.Errorf("failed to unmarshal RichText array: %w", err)
		}
		*r = RichText{Items: items}
		return nil
	case '{':
		var alias richTextAlias
		if err := json.Unmarshal(trimmed, &alias); err != nil {
			return fmt.Errorf("failed to unmarshal RichText object: %w", err)
		}
		*r = RichText(alias)
		return nil
	default:
		return fmt.Errorf("unexpected RichText JSON token: %s", string(trimmed))
	}
}
