package tgbotapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserSupportsJoinRequestQueries verifies the new supports_join_request_queries field on User
// (Bot API 10.1 Join Request Queries).
func TestUserSupportsJoinRequestQueries(t *testing.T) {
	data := `{"id": 123, "is_bot": true, "first_name": "Gatekeeper", "supports_join_request_queries": true}`
	var u User
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	assert.True(t, u.SupportsJoinRequestQueries)
}

// TestChatJoinRequestQueryID verifies the new query_id field on ChatJoinRequest
// (Bot API 10.1 Join Request Queries).
func TestChatJoinRequestQueryID(t *testing.T) {
	data := `{
		"chat": {"id": -100111, "type": "supergroup"},
		"from": {"id": 5, "is_bot": false, "first_name": "Joiner"},
		"user_chat_id": 5,
		"date": 1700000000,
		"query_id": "join-query-1"
	}`
	var r ChatJoinRequest
	require.NoError(t, json.Unmarshal([]byte(data), &r))
	assert.Equal(t, "join-query-1", r.QueryID)
}

// TestRichTextPlainString verifies that a bare JSON string round-trips through RichText.PlainText.
func TestRichTextPlainString(t *testing.T) {
	var rt RichText
	require.NoError(t, json.Unmarshal([]byte(`"hello world"`), &rt))
	assert.Equal(t, "hello world", rt.PlainText)
	assert.Empty(t, rt.Type)

	data, err := json.Marshal(rt)
	require.NoError(t, err)
	assert.JSONEq(t, `"hello world"`, string(data))
}

// TestRichTextArray verifies that a JSON array round-trips through RichText.Items.
func TestRichTextArray(t *testing.T) {
	data := `["a", {"type": "bold", "text": "b"}]`
	var rt RichText
	require.NoError(t, json.Unmarshal([]byte(data), &rt))
	require.Len(t, rt.Items, 2)
	assert.Equal(t, "a", rt.Items[0].PlainText)
	assert.Equal(t, RichTextTypeBold, rt.Items[1].Type)
	require.NotNil(t, rt.Items[1].Text)
	assert.Equal(t, "b", rt.Items[1].Text.PlainText)

	out, err := json.Marshal(rt)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

// TestRichTextBold verifies RichTextBold JSON round-tripping via the flattened RichText struct.
func TestRichTextBold(t *testing.T) {
	data := `{"type": "bold", "text": "important"}`
	var rt RichText
	require.NoError(t, json.Unmarshal([]byte(data), &rt))
	assert.Equal(t, RichTextTypeBold, rt.Type)
	require.NotNil(t, rt.Text)
	assert.Equal(t, "important", rt.Text.PlainText)

	out, err := json.Marshal(rt)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

// TestRichTextUrl verifies RichTextUrl JSON round-tripping (url + nested text).
func TestRichTextUrl(t *testing.T) {
	data := `{"type": "url", "text": "click me", "url": "https://example.com"}`
	var rt RichText
	require.NoError(t, json.Unmarshal([]byte(data), &rt))
	assert.Equal(t, RichTextTypeUrl, rt.Type)
	assert.Equal(t, "https://example.com", rt.URL)
	require.NotNil(t, rt.Text)
	assert.Equal(t, "click me", rt.Text.PlainText)

	out, err := json.Marshal(rt)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

// TestRichTextCustomEmoji verifies RichTextCustomEmoji, a variant with no nested "text" field.
func TestRichTextCustomEmoji(t *testing.T) {
	data := `{"type": "custom_emoji", "custom_emoji_id": "emoji1", "alternative_text": "😀"}`
	var rt RichText
	require.NoError(t, json.Unmarshal([]byte(data), &rt))
	assert.Equal(t, RichTextTypeCustomEmoji, rt.Type)
	assert.Equal(t, "emoji1", rt.CustomEmojiID)
	assert.Equal(t, "😀", rt.AlternativeText)
	assert.Nil(t, rt.Text)

	out, err := json.Marshal(rt)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

// TestRichTextDateTime verifies RichTextDateTime's unix_time and date_time_format fields.
func TestRichTextDateTime(t *testing.T) {
	data := `{"type": "date_time", "text": "yesterday", "unix_time": 1700000000, "date_time_format": "r"}`
	var rt RichText
	require.NoError(t, json.Unmarshal([]byte(data), &rt))
	assert.Equal(t, RichTextTypeDateTime, rt.Type)
	assert.Equal(t, 1700000000, rt.UnixTime)
	assert.Equal(t, "r", rt.DateTimeFormat)

	out, err := json.Marshal(rt)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

func TestRichTextDateTimeFormatValidate(t *testing.T) {
	tests := []struct {
		name    string
		format  string
		wantErr bool
	}{
		{name: "relative", format: "r"},
		{name: "weekday", format: "w"},
		{name: "short date", format: "d"},
		{name: "long date", format: "D"},
		{name: "short time", format: "t"},
		{name: "long time", format: "T"},
		{name: "weekday short date", format: "wd"},
		{name: "weekday long date", format: "wD"},
		{name: "weekday short time", format: "wt"},
		{name: "weekday long time", format: "wT"},
		{name: "short date short time", format: "dt"},
		{name: "short date long time", format: "dT"},
		{name: "long date short time", format: "Dt"},
		{name: "long date long time", format: "DT"},
		{name: "weekday short date short time", format: "wdt"},
		{name: "weekday short date long time", format: "wdT"},
		{name: "weekday long date short time", format: "wDt"},
		{name: "weekday long date long time", format: "wDT"},
		{name: "empty", format: ""},
		{name: "word alias", format: "relative", wantErr: true},
		{name: "uppercase relative", format: "R", wantErr: true},
		{name: "relative with date", format: "rd", wantErr: true},
		{name: "date before weekday", format: "dw", wantErr: true},
		{name: "time before date", format: "td", wantErr: true},
		{name: "duplicate date", format: "dd", wantErr: true},
		{name: "duplicate time", format: "TT", wantErr: true},
		{name: "surrounding whitespace", format: " r ", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rt := RichText{
				Type:           RichTextTypeDateTime,
				Text:           &RichText{PlainText: "when"},
				UnixTime:       1_700_000_000,
				DateTimeFormat: tt.format,
			}
			err := rt.Validate()
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestRichTextDateTimeRequiredZeroValuesArePresentOnWire(t *testing.T) {
	rt := RichText{
		Type: RichTextTypeDateTime,
		Text: &RichText{PlainText: "when"},
	}
	require.NoError(t, rt.Validate())

	data, err := json.Marshal(rt)
	require.NoError(t, err)
	assert.JSONEq(t, `{
		"type": "date_time",
		"text": "when",
		"unix_time": 0,
		"date_time_format": ""
	}`, string(data))
}

func TestRichTextDateTimeFieldsAreOmittedForOtherVariants(t *testing.T) {
	rt := RichText{
		Type: RichTextTypeBold,
		Text: &RichText{PlainText: "bold"},
	}

	data, err := json.Marshal(rt)
	require.NoError(t, err)
	assert.NotContains(t, string(data), "date_time_format")
	assert.NotContains(t, string(data), "unix_time")
}

// TestRichBlockParagraph verifies RichBlockParagraph JSON round-tripping via the flattened RichBlock struct.
func TestRichBlockParagraph(t *testing.T) {
	data := `{"type": "paragraph", "text": "hello"}`
	var b RichBlock
	require.NoError(t, json.Unmarshal([]byte(data), &b))
	assert.Equal(t, RichBlockTypeParagraph, b.Type)
	require.NotNil(t, b.Text)
	assert.Equal(t, "hello", b.Text.PlainText)

	out, err := json.Marshal(b)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

// TestRichBlockList verifies RichBlockList and its nested RichBlockListItem.
func TestRichBlockList(t *testing.T) {
	data := `{
		"type": "list",
		"items": [
			{"label": "1", "blocks": [{"type": "paragraph", "text": "first"}], "has_checkbox": true}
		]
	}`
	var b RichBlock
	require.NoError(t, json.Unmarshal([]byte(data), &b))
	assert.Equal(t, RichBlockTypeList, b.Type)
	require.Len(t, b.Items, 1)
	assert.Equal(t, "1", b.Items[0].Label)
	assert.True(t, b.Items[0].HasCheckbox)
	require.Len(t, b.Items[0].Blocks, 1)
	assert.Equal(t, RichBlockTypeParagraph, b.Items[0].Blocks[0].Type)

	out, err := json.Marshal(b)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

// TestRichBlockPhotoWithCaption verifies a media block using the shared RichBlockCaption type.
func TestRichBlockPhotoWithCaption(t *testing.T) {
	data := `{
		"type": "photo",
		"photo": [{"file_id": "p1", "width": 10, "height": 10}],
		"has_spoiler": true,
		"caption": {"text": "a photo", "credit": "photographer"}
	}`
	var b RichBlock
	require.NoError(t, json.Unmarshal([]byte(data), &b))
	assert.Equal(t, RichBlockTypePhoto, b.Type)
	require.Len(t, b.Photo, 1)
	assert.True(t, b.HasSpoiler)
	require.NotNil(t, b.Caption)
	assert.Equal(t, "a photo", b.Caption.Text.PlainText)
	require.NotNil(t, b.Caption.Credit)
	assert.Equal(t, "photographer", b.Caption.Credit.PlainText)
	assert.Nil(t, b.TableCaption)

	out, err := json.Marshal(b)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

// TestRichBlockTableCaption verifies that RichBlockTable's "caption" field decodes as a plain RichText
// (via TableCaption) rather than a RichBlockCaption, unlike every other captioned block.
func TestRichBlockTableCaption(t *testing.T) {
	data := `{
		"type": "table",
		"cells": [[{"text": "A", "is_header": true}, {"text": "B", "is_header": true}]],
		"is_bordered": true,
		"caption": "a table"
	}`
	var b RichBlock
	require.NoError(t, json.Unmarshal([]byte(data), &b))
	assert.Equal(t, RichBlockTypeTable, b.Type)
	require.Len(t, b.Cells, 1)
	require.Len(t, b.Cells[0], 2)
	assert.True(t, b.Cells[0][0].IsHeader)
	assert.True(t, b.IsBordered)
	require.NotNil(t, b.TableCaption)
	assert.Equal(t, "a table", b.TableCaption.PlainText)
	assert.Nil(t, b.Caption)

	out, err := json.Marshal(b)
	require.NoError(t, err)
	assert.JSONEq(t, data, string(out))
}

// TestRichMessageUnmarshal verifies parsing of the top-level RichMessage class (Bot API 10.1).
func TestRichMessageUnmarshal(t *testing.T) {
	data := `{
		"blocks": [
			{"type": "paragraph", "text": "hi"},
			{"type": "divider"}
		],
		"is_rtl": true
	}`
	var rm RichMessage
	require.NoError(t, json.Unmarshal([]byte(data), &rm))
	require.Len(t, rm.Blocks, 2)
	assert.Equal(t, RichBlockTypeParagraph, rm.Blocks[0].Type)
	assert.Equal(t, RichBlockTypeDivider, rm.Blocks[1].Type)
	assert.True(t, rm.IsRTL)
}

// TestMessageRichMessageField verifies the rich_message field on Message (Bot API 10.1).
func TestMessageRichMessageField(t *testing.T) {
	data := `{
		"message_id": 70,
		"date": 1700000000,
		"rich_message": {"blocks": [{"type": "paragraph", "text": "hi there"}]}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.RichMessage)
	require.Len(t, m.RichMessage.Blocks, 1)
	assert.Equal(t, "hi there", m.RichMessage.Blocks[0].Text.PlainText)
}

// TestInputRichMessageValidate verifies that InputRichMessage requires exactly one of HTML or Markdown.
func TestInputRichMessageValidate(t *testing.T) {
	assert.Error(t, InputRichMessage{}.Validate())
	assert.Error(t, InputRichMessage{HTML: "<b>hi</b>", Markdown: "**hi**"}.Validate())
	assert.NoError(t, InputRichMessage{HTML: "<b>hi</b>"}.Validate())
	assert.NoError(t, InputRichMessage{Markdown: "**hi**"}.Validate())
}

// TestInputRichMessageContentValidate verifies InputRichMessageContent's Validate() and its wiring into
// the InputMessageContent interface used by inline/guest/Web App query results (Bot API 10.1).
func TestInputRichMessageContentValidate(t *testing.T) {
	var _ InputMessageContent = InputRichMessageContent{}

	invalid := InputRichMessageContent{}
	assert.Error(t, invalid.Validate())

	valid := InputRichMessageContent{RichMessage: InputRichMessage{Markdown: "**bold**"}}
	assert.NoError(t, valid.Validate())

	data, err := encodeToJson(valid)
	require.NoError(t, err)
	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	richMessage, ok := decoded["rich_message"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "**bold**", richMessage["markdown"])
}

// TestRichMessageConfigValues verifies that RichMessageConfig.Values() serializes the sendRichMessage
// parameters correctly (Bot API 10.1).
func TestRichMessageConfigValues(t *testing.T) {
	cfg := RichMessageConfig{
		BaseChat:    BaseChat{ChatID: 12345},
		RichMessage: InputRichMessage{HTML: "<b>hello</b>"},
	}

	assert.Equal(t, "sendRichMessage", cfg.TelegramMethod())

	values, err := cfg.Values()
	require.NoError(t, err)
	assert.Equal(t, "12345", values.Get("chat_id"))

	var rm InputRichMessage
	require.NoError(t, json.Unmarshal([]byte(values.Get("rich_message")), &rm))
	assert.Equal(t, "<b>hello</b>", rm.HTML)
}

// TestRichMessageDraftConfigValues verifies that RichMessageDraftConfig.Values() serializes the
// sendRichMessageDraft parameters correctly (Bot API 10.1).
func TestRichMessageDraftConfigValues(t *testing.T) {
	cfg := RichMessageDraftConfig{
		ChatID:      12345,
		DraftID:     7,
		RichMessage: InputRichMessage{Markdown: "**streaming**"},
	}

	assert.Equal(t, "sendRichMessageDraft", cfg.TelegramMethod())

	values, err := cfg.Values()
	require.NoError(t, err)
	assert.Equal(t, "12345", values.Get("chat_id"))
	assert.Equal(t, "7", values.Get("draft_id"))

	var rm InputRichMessage
	require.NoError(t, json.Unmarshal([]byte(values.Get("rich_message")), &rm))
	assert.Equal(t, "**streaming**", rm.Markdown)
}

// TestEditMessageTextConfigRichMessage verifies that the new rich_message parameter is serialized by
// EditMessageTextConfig.Values() (Bot API 10.1).
func TestEditMessageTextConfigRichMessage(t *testing.T) {
	cfg := EditMessageTextConfig{
		BaseEdit:    NewChatMessageEdit(12345, 99),
		Text:        "fallback text",
		RichMessage: &InputRichMessage{HTML: "<i>edited</i>"},
	}

	values, err := cfg.Values()
	require.NoError(t, err)

	var rm InputRichMessage
	require.NoError(t, json.Unmarshal([]byte(values.Get("rich_message")), &rm))
	assert.Equal(t, "<i>edited</i>", rm.HTML)
}

// TestPollMediaLink verifies the new Link class and PollMedia.Link field (Bot API 10.1 Polls).
func TestPollMediaLink(t *testing.T) {
	data := `{"link": {"url": "https://example.com/poll-option"}}`
	var pm PollMedia
	require.NoError(t, json.Unmarshal([]byte(data), &pm))
	require.NotNil(t, pm.Link)
	assert.Equal(t, "https://example.com/poll-option", pm.Link.URL)
}

// TestInputPollOptionMediaLink verifies the InputMediaLink class allowed as InputPollOptionMedia
// (Bot API 10.1 Polls), following the same flattened-fields convention as the other media kinds.
func TestInputPollOptionMediaLink(t *testing.T) {
	media := InputPollOptionMedia{Type: "link", URL: "https://example.com"}
	data, err := encodeToJson(media)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "link", decoded["type"])
	assert.Equal(t, "https://example.com", decoded["url"])
	assert.NotContains(t, decoded, "media")
}

// TestInputMediaLinkMarshal verifies marshaling of the standalone InputMediaLink type (Bot API 10.1).
func TestInputMediaLinkMarshal(t *testing.T) {
	m := InputMediaLink{Type: "link", URL: "https://example.com/x"}
	data, err := encodeToJson(m)
	require.NoError(t, err)
	assert.JSONEq(t, `{"type":"link","url":"https://example.com/x"}`, string(data))
}

// TestAnswerChatJoinRequestQueryResultConstants verifies the ChatJoinRequestQueryResult constants used
// by BotAPI.AnswerChatJoinRequestQuery (Bot API 10.1 Join Request Queries).
func TestAnswerChatJoinRequestQueryResultConstants(t *testing.T) {
	assert.Equal(t, ChatJoinRequestQueryResult("approve"), ChatJoinRequestQueryResultApprove)
	assert.Equal(t, ChatJoinRequestQueryResult("decline"), ChatJoinRequestQueryResultDecline)
	assert.Equal(t, ChatJoinRequestQueryResult("queue"), ChatJoinRequestQueryResultQueue)
}
