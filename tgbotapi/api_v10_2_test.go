package tgbotapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInputRichBlockParagraphMarshal verifies InputRichBlockParagraph marshaling via the flattened
// InputRichBlock struct (Bot API 10.2 Rich Messages).
func TestInputRichBlockParagraphMarshal(t *testing.T) {
	b := InputRichBlock{Type: RichBlockTypeParagraph, Text: &RichText{PlainText: "hello"}}
	data, err := encodeToJson(b)
	require.NoError(t, err)
	assert.JSONEq(t, `{"type":"paragraph","text":"hello"}`, string(data))
}

// TestInputRichBlockListMarshal verifies InputRichBlockList and its nested InputRichBlockListItem.
func TestInputRichBlockListMarshal(t *testing.T) {
	b := InputRichBlock{
		Type: RichBlockTypeList,
		Items: []InputRichBlockListItem{
			{
				Blocks:      []InputRichBlock{{Type: RichBlockTypeParagraph, Text: &RichText{PlainText: "first"}}},
				HasCheckbox: true,
			},
		},
	}
	data, err := encodeToJson(b)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "list", decoded["type"])
	items, ok := decoded["items"].([]any)
	require.True(t, ok)
	require.Len(t, items, 1)
	item, ok := items[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, true, item["has_checkbox"])
}

// TestInputRichBlockPhotoCaption verifies that InputRichBlockPhoto's caption is projected as a
// RichBlockCaption object onto the "caption" wire field via InputRichBlock.MarshalJSON, and that the
// nested InputMediaPhoto is embedded under "photo".
func TestInputRichBlockPhotoCaption(t *testing.T) {
	b := InputRichBlock{
		Type:    RichBlockTypePhoto,
		Photo:   &InputMediaPhoto{Type: "photo", Media: "file123"},
		Caption: &RichBlockCaption{Text: RichText{PlainText: "a photo"}},
	}
	data, err := encodeToJson(b)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "photo", decoded["type"])
	photo, ok := decoded["photo"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "file123", photo["media"])
	caption, ok := decoded["caption"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "a photo", caption["text"])
}

// TestInputRichBlockTableCaption verifies that InputRichBlockTable's caption is projected as a plain
// RichText onto the "caption" wire field, unlike every other captioned InputRichBlock.
func TestInputRichBlockTableCaption(t *testing.T) {
	b := InputRichBlock{
		Type:         RichBlockTypeTable,
		Cells:        [][]RichBlockTableCell{{{Text: &RichText{PlainText: "A"}, IsHeader: true}}},
		TableCaption: &RichText{PlainText: "a table"},
	}
	data, err := encodeToJson(b)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "a table", decoded["caption"])
}

// TestInputRichBlockMapFields verifies InputRichBlockMap's location/zoom/width/height fields.
func TestInputRichBlockMapFields(t *testing.T) {
	b := InputRichBlock{
		Type:     RichBlockTypeMap,
		Location: &Location{Latitude: 1.5, Longitude: 2.5},
		Zoom:     15,
		Width:    300,
		Height:   200,
	}
	data, err := encodeToJson(b)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "map", decoded["type"])
	assert.Equal(t, float64(15), decoded["zoom"])
	loc, ok := decoded["location"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, 1.5, loc["latitude"])
}

// TestInputRichBlockDetails verifies InputRichBlockDetails' summary/blocks/is_open fields.
func TestInputRichBlockDetails(t *testing.T) {
	b := InputRichBlock{
		Type:    RichBlockTypeDetails,
		Summary: &RichText{PlainText: "click to expand"},
		Blocks:  []InputRichBlock{{Type: RichBlockTypeParagraph, Text: &RichText{PlainText: "hidden"}}},
		IsOpen:  true,
	}
	data, err := encodeToJson(b)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "details", decoded["type"])
	assert.Equal(t, true, decoded["is_open"])
	assert.Equal(t, "click to expand", decoded["summary"])
}

// TestInputRichBlockDivider verifies InputRichBlockDivider, a block with only a "type" field.
func TestInputRichBlockDivider(t *testing.T) {
	b := InputRichBlock{Type: RichBlockTypeDivider}
	data, err := encodeToJson(b)
	require.NoError(t, err)
	assert.JSONEq(t, `{"type":"divider"}`, string(data))
}

// TestInputRichMessageMediaMarshal verifies InputRichMessageMedia marshaling with each supported
// InputMedia* variant assigned to Media (Bot API 10.2 Rich Messages).
func TestInputRichMessageMediaMarshal(t *testing.T) {
	m := InputRichMessageMedia{
		ID:    "photo1",
		Media: &InputMediaPhoto{Type: "photo", Media: "file_id_1"},
	}
	data, err := encodeToJson(m)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "photo1", decoded["id"])
	media, ok := decoded["media"].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "photo", media["type"])
	assert.Equal(t, "file_id_1", media["media"])
}

// TestInputRichMessageValidateWithBlocks verifies that InputRichMessage.Validate() accepts Blocks as a
// third alternative to HTML/Markdown, and rejects combinations thereof (Bot API 10.2).
func TestInputRichMessageValidateWithBlocks(t *testing.T) {
	assert.Error(t, InputRichMessage{}.Validate())
	assert.NoError(t, InputRichMessage{Blocks: []InputRichBlock{{Type: RichBlockTypeDivider}}}.Validate())
	assert.Error(t, InputRichMessage{
		HTML:   "<b>hi</b>",
		Blocks: []InputRichBlock{{Type: RichBlockTypeDivider}},
	}.Validate())
	assert.Error(t, InputRichMessage{HTML: "<b>hi</b>", Markdown: "**hi**"}.Validate())
}

// TestInputRichMessageMediaField verifies that InputRichMessage.Media round-trips through JSON encoding.
func TestInputRichMessageMediaField(t *testing.T) {
	msg := InputRichMessage{
		HTML: `<img src="tg://photo?id=photo1"/>`,
		Media: []InputRichMessageMedia{
			{ID: "photo1", Media: &InputMediaPhoto{Type: "photo", Media: "file_id_1"}},
		},
	}
	data, err := encodeToJson(msg)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	mediaList, ok := decoded["media"].([]any)
	require.True(t, ok)
	require.Len(t, mediaList, 1)
}

// TestBotCommandIsEphemeral verifies the new is_ephemeral field on TelegramBotCommand (Bot API 10.2
// Ephemeral Messages).
func TestBotCommandIsEphemeral(t *testing.T) {
	cmd := TelegramBotCommand{Command: "secret", Description: "an ephemeral command", IsEphemeral: true}
	data, err := encodeToJson(cmd)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, true, decoded["is_ephemeral"])
}

// TestMessageReceiverUserAndEphemeralMessageID verifies the new receiver_user and ephemeral_message_id
// fields on Message (Bot API 10.2 Ephemeral Messages).
func TestMessageReceiverUserAndEphemeralMessageID(t *testing.T) {
	data := `{
		"message_id": 0,
		"date": 1700000000,
		"receiver_user": {"id": 42, "is_bot": false, "first_name": "Receiver"},
		"ephemeral_message_id": 7
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.ReceiverUser)
	assert.EqualValues(t, 42, m.ReceiverUser.ID)
	assert.Equal(t, 7, m.EphemeralMessageID)
}

// TestReplyParametersEphemeralMessageID verifies the new ephemeral_message_id field on ReplyParameters,
// and that message_id is omitted from the wire encoding when unset (Bot API 10.2 Ephemeral Messages).
func TestReplyParametersEphemeralMessageID(t *testing.T) {
	rp := ReplyParameters{EphemeralMessageID: 99}
	data, err := encodeToJson(rp)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, float64(99), decoded["ephemeral_message_id"])
	assert.NotContains(t, decoded, "message_id")
}

// TestEditEphemeralMessageTextConfigValues verifies that EditEphemeralMessageTextConfig.Values()
// serializes the editEphemeralMessageText parameters correctly (Bot API 10.2 Ephemeral Messages).
func TestEditEphemeralMessageTextConfigValues(t *testing.T) {
	cfg := EditEphemeralMessageTextConfig{
		baseEphemeralMessageEdit: baseEphemeralMessageEdit{
			ChatID:             -100111,
			ReceiverUserID:     42,
			EphemeralMessageID: 7,
		},
		Text: "updated text",
	}

	assert.Equal(t, "editEphemeralMessageText", cfg.TelegramMethod())

	values, err := cfg.Values()
	require.NoError(t, err)
	assert.Equal(t, "-100111", values.Get("chat_id"))
	assert.Equal(t, "42", values.Get("receiver_user_id"))
	assert.Equal(t, "7", values.Get("ephemeral_message_id"))
	assert.Equal(t, "updated text", values.Get("text"))
}

// TestDeleteEphemeralMessageConfigValues verifies that DeleteEphemeralMessageConfig.Values() serializes
// the deleteEphemeralMessage parameters correctly (Bot API 10.2 Ephemeral Messages).
func TestDeleteEphemeralMessageConfigValues(t *testing.T) {
	cfg := DeleteEphemeralMessageConfig{
		baseEphemeralMessageEdit{
			ChatID:             -100111,
			ReceiverUserID:     42,
			EphemeralMessageID: 7,
		},
	}

	assert.Equal(t, "deleteEphemeralMessage", cfg.TelegramMethod())

	values, err := cfg.Values()
	require.NoError(t, err)
	assert.Equal(t, "-100111", values.Get("chat_id"))
	assert.Equal(t, "42", values.Get("receiver_user_id"))
	assert.Equal(t, "7", values.Get("ephemeral_message_id"))
}

// TestMessageConfigReceiverUserID verifies that the new receiver_user_id/callback_query_id parameters
// on BaseChat are serialized by MessageConfig.Values() (Bot API 10.2 Ephemeral Messages).
func TestMessageConfigReceiverUserID(t *testing.T) {
	cfg := MessageConfig{
		BaseChat: BaseChat{ChatID: -100111, ReceiverUserID: 42, CallbackQueryID: "cbq1"},
		Text:     "psst",
	}

	values, err := cfg.Values()
	require.NoError(t, err)
	assert.Equal(t, "42", values.Get("receiver_user_id"))
	assert.Equal(t, "cbq1", values.Get("callback_query_id"))
}

// TestCommunityChatAddedUnmarshal verifies the Community, CommunityChatAdded, and CommunityChatRemoved
// types, and their wiring into Message (Bot API 10.2 Communities).
func TestCommunityChatAddedUnmarshal(t *testing.T) {
	data := `{
		"message_id": 5,
		"date": 1700000000,
		"community_chat_added": {"community": {"id": 555, "name": "Book Club"}}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.CommunityChatAdded)
	assert.EqualValues(t, 555, m.CommunityChatAdded.Community.ID)
	assert.Equal(t, "Book Club", m.CommunityChatAdded.Community.Name)
}

// TestCommunityChatRemovedUnmarshal verifies the community_chat_removed field on Message.
func TestCommunityChatRemovedUnmarshal(t *testing.T) {
	data := `{
		"message_id": 6,
		"date": 1700000000,
		"community_chat_removed": {}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	assert.NotNil(t, m.CommunityChatRemoved)
}

// TestBotSubscriptionUpdatedUnmarshal verifies the BotSubscriptionUpdated type and its wiring into
// Update.Subscription (Bot API 10.2 General).
func TestBotSubscriptionUpdatedUnmarshal(t *testing.T) {
	data := `{
		"update_id": 1,
		"subscription": {
			"user": {"id": 1, "is_bot": false, "first_name": "Subscriber"},
			"invoice_payload": "sub-payload",
			"state": "canceled"
		}
	}`
	var u Update
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	require.NotNil(t, u.Subscription)
	assert.EqualValues(t, 1, u.Subscription.User.ID)
	assert.Equal(t, "sub-payload", u.Subscription.InvoicePayload)
	assert.Equal(t, BotSubscriptionStateCanceled, u.Subscription.State)
}
