package tgbotapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserCanManageBots verifies the new can_manage_bots field on User (Bot API 9.6).
func TestUserCanManageBots(t *testing.T) {
	data := `{"id": 123, "is_bot": true, "first_name": "Manager", "can_manage_bots": true}`
	var u User
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	assert.True(t, u.CanManageBots)
}

// TestKeyboardButtonRequestManagedBot verifies the new request_managed_bot field on KeyboardButton (Bot API 9.6).
func TestKeyboardButtonRequestManagedBot(t *testing.T) {
	btn := KeyboardButton{
		Text: "Create a bot",
		RequestManagedBot: &KeyboardButtonRequestManagedBot{
			RequestID:         1,
			SuggestedName:     "My Bot",
			SuggestedUsername: "my_suggested_bot",
		},
	}

	data, err := json.Marshal(btn)
	require.NoError(t, err)

	var round KeyboardButton
	require.NoError(t, json.Unmarshal(data, &round))
	require.NotNil(t, round.RequestManagedBot)
	assert.Equal(t, 1, round.RequestManagedBot.RequestID)
	assert.Equal(t, "My Bot", round.RequestManagedBot.SuggestedName)
	assert.Equal(t, "my_suggested_bot", round.RequestManagedBot.SuggestedUsername)
}

// TestManagedBotCreated verifies parsing of the ManagedBotCreated service message (Bot API 9.6).
func TestManagedBotCreated(t *testing.T) {
	data := `{"bot": {"id": 999, "is_bot": true, "first_name": "NewBot", "username": "new_bot"}}`
	var mbc ManagedBotCreated
	require.NoError(t, json.Unmarshal([]byte(data), &mbc))
	assert.Equal(t, int64(999), mbc.Bot.ID)
	assert.Equal(t, "new_bot", mbc.Bot.UserName)
}

// TestManagedBotUpdated verifies parsing of the ManagedBotUpdated update (Bot API 9.6).
func TestManagedBotUpdated(t *testing.T) {
	data := `{
		"update_id": 1,
		"managed_bot": {
			"user": {"id": 42, "is_bot": false, "first_name": "Owner"},
			"bot": {"id": 999, "is_bot": true, "first_name": "NewBot"}
		}
	}`
	var u Update
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	require.NotNil(t, u.ManagedBot)
	assert.Equal(t, int64(42), u.ManagedBot.User.ID)
	assert.Equal(t, int64(999), u.ManagedBot.Bot.ID)
}

// TestMessageManagedBotCreated verifies the managed_bot_created field on Message (Bot API 9.6).
func TestMessageManagedBotCreated(t *testing.T) {
	data := `{
		"message_id": 10,
		"date": 1700000000,
		"managed_bot_created": {"bot": {"id": 999, "is_bot": true, "first_name": "NewBot"}}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.ManagedBotCreated)
	assert.Equal(t, int64(999), m.ManagedBotCreated.Bot.ID)
}

// TestPreparedKeyboardButton verifies round-tripping of PreparedKeyboardButton (Bot API 9.6).
func TestPreparedKeyboardButton(t *testing.T) {
	pkb := PreparedKeyboardButton{ID: "abc123"}
	data, err := json.Marshal(pkb)
	require.NoError(t, err)
	assert.JSONEq(t, `{"id":"abc123"}`, string(data))
}

// TestPollCorrectOptionIDsAndNewFields verifies the Poll fields added/renamed in Bot API 9.6:
// correct_option_ids (renamed from correct_option_id), allows_revoting, description, description_entities.
func TestPollCorrectOptionIDsAndNewFields(t *testing.T) {
	data := `{
		"id": "poll1",
		"question": "Pick correct answers",
		"options": [
			{"persistent_id": "opt1", "text": "A", "voter_count": 1},
			{"persistent_id": "opt2", "text": "B", "voter_count": 2}
		],
		"total_voter_count": 3,
		"is_closed": false,
		"is_anonymous": true,
		"type": "quiz",
		"allows_multiple_answers": true,
		"correct_option_ids": [0, 1],
		"allows_revoting": true,
		"description": "Choose wisely",
		"description_entities": [{"type": "bold", "offset": 0, "length": 6}]
	}`
	var p Poll
	require.NoError(t, json.Unmarshal([]byte(data), &p))
	assert.Equal(t, []int{0, 1}, p.CorrectOptionIDs)
	assert.True(t, p.AllowsRevoting)
	assert.Equal(t, "Choose wisely", p.Description)
	require.Len(t, p.DescriptionEntities, 1)
	assert.Equal(t, "bold", p.DescriptionEntities[0].Type)
	require.Len(t, p.Options, 2)
	assert.Equal(t, "opt1", p.Options[0].PersistentID)
}

// TestPollOptionAddedByUserAndChat verifies PollOption's added_by_user/added_by_chat/addition_date (Bot API 9.6).
func TestPollOptionAddedByUserAndChat(t *testing.T) {
	data := `{
		"persistent_id": "opt3",
		"text": "C",
		"voter_count": 0,
		"added_by_user": {"id": 7, "is_bot": false, "first_name": "Voter"},
		"added_by_chat": {"id": -100123, "type": "supergroup"},
		"addition_date": 1700000100
	}`
	var opt PollOption
	require.NoError(t, json.Unmarshal([]byte(data), &opt))
	require.NotNil(t, opt.AddedByUser)
	assert.Equal(t, int64(7), opt.AddedByUser.ID)
	require.NotNil(t, opt.AddedByChat)
	assert.Equal(t, int64(-100123), opt.AddedByChat.ID)
	assert.Equal(t, 1700000100, opt.AdditionDate)
}

// TestPollAnswerOptionPersistentIDs verifies option_persistent_ids on PollAnswer (Bot API 9.6).
func TestPollAnswerOptionPersistentIDs(t *testing.T) {
	data := `{"poll_id": "poll1", "option_ids": [0, 2], "option_persistent_ids": ["opt1", "opt3"]}`
	var pa PollAnswer
	require.NoError(t, json.Unmarshal([]byte(data), &pa))
	assert.Equal(t, []string{"opt1", "opt3"}, pa.OptionPersistentIDs)
}

// TestMessagePollOptionAddedAndDeleted verifies the poll_option_added/poll_option_deleted service
// messages on Message (Bot API 9.6).
func TestMessagePollOptionAddedAndDeleted(t *testing.T) {
	data := `{
		"message_id": 20,
		"date": 1700000000,
		"poll_option_added": {
			"option_persistent_id": "opt4",
			"option_text": "New option"
		}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.PollOptionAdded)
	assert.Equal(t, "opt4", m.PollOptionAdded.OptionPersistentID)
	assert.Equal(t, "New option", m.PollOptionAdded.OptionText)

	data2 := `{
		"message_id": 21,
		"date": 1700000001,
		"poll_option_deleted": {
			"option_persistent_id": "opt5",
			"option_text": "Removed option"
		}
	}`
	var m2 Message
	require.NoError(t, json.Unmarshal([]byte(data2), &m2))
	require.NotNil(t, m2.PollOptionDeleted)
	assert.Equal(t, "opt5", m2.PollOptionDeleted.OptionPersistentID)
}

// TestReplyParametersPollOptionID verifies the poll_option_id field on ReplyParameters (Bot API 9.6).
func TestReplyParametersPollOptionID(t *testing.T) {
	rp := ReplyParameters{MessageID: 5, PollOptionID: "opt1"}
	data, err := json.Marshal(rp)
	require.NoError(t, err)

	var round ReplyParameters
	require.NoError(t, json.Unmarshal(data, &round))
	assert.Equal(t, "opt1", round.PollOptionID)
}

// TestMessageReplyToPollOptionID verifies reply_to_poll_option_id on Message (Bot API 9.6).
func TestMessageReplyToPollOptionID(t *testing.T) {
	data := `{"message_id": 30, "date": 1700000000, "reply_to_poll_option_id": "opt9"}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	assert.Equal(t, "opt9", m.ReplyToPollOptionID)
}

// TestPollConfigValues verifies that PollConfig.Values() serializes the Bot API 9.6 sendPoll
// parameters correctly and round-trips the options/correct_option_ids through JSON.
func TestPollConfigValues(t *testing.T) {
	cfg := &PollConfig{
		BaseChat: BaseChat{ChatID: 12345},
		Question: "Pick one or more",
		Options: []InputPollOption{
			{Text: "Option A"},
			{Text: "Option B"},
		},
		Type:                   "quiz",
		AllowsMultipleAnswers:  true,
		AllowsRevoting:         true,
		ShuffleOptions:         true,
		AllowAddingOptions:     true,
		HideResultsUntilCloses: true,
		CorrectOptionIDs:       []int{0, 1},
		Description:            "poll description",
	}

	assert.Equal(t, "sendPoll", cfg.TelegramMethod())

	values, err := cfg.Values()
	require.NoError(t, err)

	assert.Equal(t, "12345", values.Get("chat_id"))
	assert.Equal(t, "Pick one or more", values.Get("question"))
	assert.Equal(t, "quiz", values.Get("type"))
	assert.Equal(t, "true", values.Get("allows_multiple_answers"))
	assert.Equal(t, "true", values.Get("allows_revoting"))
	assert.Equal(t, "true", values.Get("shuffle_options"))
	assert.Equal(t, "true", values.Get("allow_adding_options"))
	assert.Equal(t, "true", values.Get("hide_results_until_closes"))
	assert.Equal(t, "poll description", values.Get("description"))

	var options []InputPollOption
	require.NoError(t, json.Unmarshal([]byte(values.Get("options")), &options))
	require.Len(t, options, 2)
	assert.Equal(t, "Option A", options[0].Text)

	var correctIDs []int
	require.NoError(t, json.Unmarshal([]byte(values.Get("correct_option_ids")), &correctIDs))
	assert.Equal(t, []int{0, 1}, correctIDs)
}

// TestSavePreparedKeyboardButtonEncoding verifies that a KeyboardButton using the new
// request_managed_bot field (the only type currently allowed for savePreparedKeyboardButton
// besides request_users/request_chat) marshals to the JSON payload Telegram expects for the
// "button" parameter of savePreparedKeyboardButton.
func TestSavePreparedKeyboardButtonEncoding(t *testing.T) {
	button := KeyboardButton{
		Text: "Create managed bot",
		RequestManagedBot: &KeyboardButtonRequestManagedBot{
			RequestID:     7,
			SuggestedName: "Helper Bot",
		},
	}

	data, err := encodeToJson(button)
	require.NoError(t, err)

	var decoded map[string]any
	require.NoError(t, json.Unmarshal(data, &decoded))
	assert.Equal(t, "Create managed bot", decoded["text"])
	require.Contains(t, decoded, "request_managed_bot")
	rmb := decoded["request_managed_bot"].(map[string]any)
	assert.Equal(t, float64(7), rmb["request_id"])
	assert.Equal(t, "Helper Bot", rmb["suggested_name"])
	assert.NotContains(t, rmb, "suggested_username")
}
