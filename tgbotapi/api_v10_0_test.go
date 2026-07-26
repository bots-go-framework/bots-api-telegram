package tgbotapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserSupportsGuestQueries verifies the new supports_guest_queries field on User (Bot API 10.0).
func TestUserSupportsGuestQueries(t *testing.T) {
	data := `{"id": 123, "is_bot": true, "first_name": "Guesty", "supports_guest_queries": true}`
	var u User
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	assert.True(t, u.SupportsGuestQueries)
}

// TestMessageGuestFields verifies the guest_bot_caller_user, guest_bot_caller_chat and guest_query_id
// fields on Message (Bot API 10.0 Guest Mode).
func TestMessageGuestFields(t *testing.T) {
	data := `{
		"message_id": 40,
		"date": 1700000000,
		"guest_bot_caller_user": {"id": 7, "is_bot": false, "first_name": "Caller"},
		"guest_bot_caller_chat": {"id": -100777, "type": "supergroup"},
		"guest_query_id": "guest-query-1"
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.GuestBotCallerUser)
	assert.Equal(t, int64(7), m.GuestBotCallerUser.ID)
	require.NotNil(t, m.GuestBotCallerChat)
	assert.Equal(t, int64(-100777), m.GuestBotCallerChat.ID)
	assert.Equal(t, "guest-query-1", m.GuestQueryID)
}

// TestUpdateGuestMessage verifies the guest_message field on Update (Bot API 10.0 Guest Mode).
func TestUpdateGuestMessage(t *testing.T) {
	data := `{
		"update_id": 1,
		"guest_message": {
			"message_id": 41,
			"date": 1700000000,
			"guest_query_id": "guest-query-2",
			"text": "hi from a guest"
		}
	}`
	var u Update
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	require.NotNil(t, u.GuestMessage)
	assert.Equal(t, "guest-query-2", u.GuestMessage.GuestQueryID)
	assert.Equal(t, "hi from a guest", u.GuestMessage.Text)
}

// TestSentGuestMessage verifies round-tripping of the SentGuestMessage class returned by
// answerGuestQuery (Bot API 10.0).
func TestSentGuestMessage(t *testing.T) {
	data := `{"inline_message_id": "guest-inline-55"}`
	var sent SentGuestMessage
	require.NoError(t, json.Unmarshal([]byte(data), &sent))
	assert.Equal(t, "guest-inline-55", sent.InlineMessageID)
}

// TestBotAccessSettings verifies round-tripping of BotAccessSettings, used by
// getManagedBotAccessSettings/setManagedBotAccessSettings (Bot API 10.0).
func TestBotAccessSettings(t *testing.T) {
	settings := BotAccessSettings{
		IsAccessRestricted: true,
		AddedUsers: []User{
			{ID: 42, FirstName: "Player"},
		},
	}

	data, err := encodeToJson(settings)
	require.NoError(t, err)

	var decoded BotAccessSettings
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.True(t, decoded.IsAccessRestricted)
	require.Len(t, decoded.AddedUsers, 1)
	assert.EqualValues(t, 42, decoded.AddedUsers[0].ID)
	assert.Equal(t, "Player", decoded.AddedUsers[0].FirstName)
}

// TestLivePhotoUnmarshal verifies parsing of the LivePhoto class (Bot API 10.0 Live Photos).
func TestLivePhotoUnmarshal(t *testing.T) {
	data := `{
		"photo": [{"file_id": "photo1", "width": 100, "height": 100}],
		"file_id": "video1",
		"file_unique_id": "video1-unique",
		"width": 720,
		"height": 1280,
		"duration": 3,
		"mime_type": "video/mp4"
	}`
	var lp LivePhoto
	require.NoError(t, json.Unmarshal([]byte(data), &lp))
	require.Len(t, lp.Photo, 1)
	assert.Equal(t, "photo1", lp.Photo[0].FileID)
	assert.Equal(t, "video1", lp.FileID)
	assert.Equal(t, 3, lp.Duration)
}

// TestMessageLivePhoto verifies the live_photo field on Message (Bot API 10.0).
func TestMessageLivePhoto(t *testing.T) {
	data := `{
		"message_id": 60,
		"date": 1700000000,
		"live_photo": {"file_id": "video2", "width": 720, "height": 1280, "duration": 4}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.LivePhoto)
	assert.Equal(t, "video2", m.LivePhoto.FileID)
}

// TestExternalReplyInfoLivePhoto verifies the live_photo field on ExternalReplyInfo (Bot API 10.0).
func TestExternalReplyInfoLivePhoto(t *testing.T) {
	data := `{
		"live_photo": {"file_id": "video3", "width": 720, "height": 1280, "duration": 2}
	}`
	var eri ExternalReplyInfo
	require.NoError(t, json.Unmarshal([]byte(data), &eri))
	require.NotNil(t, eri.LivePhoto)
	assert.Equal(t, "video3", eri.LivePhoto.FileID)
}

// TestPaidMediaLivePhotoField verifies the flattened PaidMedia struct can decode a "live_photo" entry
// (Bot API 10.0 PaidMediaLivePhoto).
func TestPaidMediaLivePhotoField(t *testing.T) {
	data := `{"type": "live_photo", "live_photo": {"file_id": "video4", "width": 720, "height": 1280, "duration": 5}}`
	var pm PaidMedia
	require.NoError(t, json.Unmarshal([]byte(data), &pm))
	assert.Equal(t, "live_photo", pm.Type)
	require.NotNil(t, pm.LivePhoto)
	assert.Equal(t, "video4", pm.LivePhoto.FileID)
}

// TestInputMediaLivePhotoMarshal verifies marshaling of InputMediaLivePhoto (Bot API 10.0).
func TestInputMediaLivePhotoMarshal(t *testing.T) {
	m := InputMediaLivePhoto{
		Type:    "live_photo",
		Media:   "attach://live1",
		Caption: "a live photo",
	}
	data, err := encodeToJson(m)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "live_photo", decoded["type"])
	assert.Equal(t, "attach://live1", decoded["media"])
	assert.Equal(t, "a live photo", decoded["caption"])
}

// TestInputPaidMediaLivePhotoMarshal verifies marshaling of InputPaidMediaLivePhoto (Bot API 10.0).
func TestInputPaidMediaLivePhotoMarshal(t *testing.T) {
	m := InputPaidMediaLivePhoto{Type: "live_photo", Media: "file-id-1"}
	data, err := encodeToJson(m)
	require.NoError(t, err)
	assert.JSONEq(t, `{"type":"live_photo","media":"file-id-1"}`, string(data))
}

// TestLivePhotoConfigValues verifies that LivePhotoConfig.Values() serializes the sendLivePhoto
// parameters correctly (Bot API 10.0).
func TestLivePhotoConfigValues(t *testing.T) {
	cfg := LivePhotoConfig{
		BaseChat:              BaseChat{ChatID: 12345},
		Photo:                 FileID("live-file-id"),
		Caption:               "look at this",
		ShowCaptionAboveMedia: true,
		HasSpoiler:            true,
	}

	assert.Equal(t, "sendLivePhoto", cfg.TelegramMethod())

	values, err := cfg.Values()
	require.NoError(t, err)

	assert.Equal(t, "12345", values.Get("chat_id"))
	assert.Equal(t, "live-file-id", values.Get("live_photo"))
	assert.Equal(t, "look at this", values.Get("caption"))
	assert.Equal(t, "true", values.Get("show_caption_above_media"))
	assert.Equal(t, "true", values.Get("has_spoiler"))
}

// TestPollMediaAndOptionMedia verifies the new media/explanation_media fields on Poll and the
// media field on PollOption (Bot API 10.0 Poll media).
func TestPollMediaAndOptionMedia(t *testing.T) {
	data := `{
		"id": "poll2",
		"question": "Pick a photo",
		"options": [
			{"persistent_id": "opt1", "text": "A", "voter_count": 0, "media": {"photo": [{"file_id": "p1", "width": 10, "height": 10}]}}
		],
		"total_voter_count": 0,
		"is_closed": false,
		"is_anonymous": true,
		"type": "regular",
		"allows_multiple_answers": false,
		"media": {"video": {"file_id": "v1", "width": 10, "height": 10, "duration": 5}},
		"explanation_media": {"sticker": {"file_id": "s1", "width": 10, "height": 10}},
		"members_only": true,
		"country_codes": ["US", "GB"]
	}`
	var p Poll
	require.NoError(t, json.Unmarshal([]byte(data), &p))
	require.NotNil(t, p.Media)
	require.NotNil(t, p.Media.Video)
	assert.Equal(t, "v1", p.Media.Video.FileID)
	require.NotNil(t, p.ExplanationMedia)
	require.NotNil(t, p.ExplanationMedia.Sticker)
	assert.Equal(t, "s1", p.ExplanationMedia.Sticker.FileID)
	assert.True(t, p.MembersOnly)
	assert.Equal(t, []string{"US", "GB"}, p.CountryCodes)
	require.Len(t, p.Options, 1)
	require.NotNil(t, p.Options[0].Media)
	require.Len(t, p.Options[0].Media.Photo, 1)
	assert.Equal(t, "p1", p.Options[0].Media.Photo[0].FileID)
}

// TestPollConfigMediaValues verifies that PollConfig.Values() serializes the Bot API 10.0 sendPoll
// media parameters (media, explanation_media, members_only, country_codes).
func TestPollConfigMediaValues(t *testing.T) {
	cfg := &PollConfig{
		BaseChat: BaseChat{ChatID: 12345},
		Question: "Pick a photo",
		Options: []InputPollOption{
			{Text: "Option A", Media: &InputPollOptionMedia{Type: "photo", Media: "attach://opt-a"}},
		},
		Media:            &InputPollMedia{Type: "photo", Media: "attach://poll-photo"},
		ExplanationMedia: &InputPollMedia{Type: "video", Media: "attach://poll-video"},
		MembersOnly:      true,
		CountryCodes:     []string{"US", "CA"},
	}

	values, err := cfg.Values()
	require.NoError(t, err)

	assert.Equal(t, "true", values.Get("members_only"))

	var media InputPollMedia
	require.NoError(t, json.Unmarshal([]byte(values.Get("media")), &media))
	assert.Equal(t, "photo", media.Type)
	assert.Equal(t, "attach://poll-photo", media.Media)

	var explanationMedia InputPollMedia
	require.NoError(t, json.Unmarshal([]byte(values.Get("explanation_media")), &explanationMedia))
	assert.Equal(t, "video", explanationMedia.Type)
	assert.Equal(t, "attach://poll-video", explanationMedia.Media)

	var countryCodes []string
	require.NoError(t, json.Unmarshal([]byte(values.Get("country_codes")), &countryCodes))
	assert.Equal(t, []string{"US", "CA"}, countryCodes)

	var options []InputPollOption
	require.NoError(t, json.Unmarshal([]byte(values.Get("options")), &options))
	require.Len(t, options, 1)
	require.NotNil(t, options[0].Media)
	assert.Equal(t, "photo", options[0].Media.Type)
	assert.Equal(t, "attach://opt-a", options[0].Media.Media)
}
