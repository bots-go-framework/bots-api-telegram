package tgbotapi

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestUserIDIsInt64 verifies that User.ID is int64 to handle large Telegram IDs.
func TestUserIDIsInt64(t *testing.T) {
	var u User
	u.ID = 9_999_999_999
	assert.Equal(t, int64(9_999_999_999), u.ID)
	assert.Equal(t, int64(9_999_999_999), u.GetIDInt64())
	assert.Equal(t, int64(9_999_999_999), u.GetID())
}

// TestUserNewFields verifies newly added User fields.
func TestUserNewFields(t *testing.T) {
	data := `{
		"id": 123456789,
		"is_bot": true,
		"first_name": "TestBot",
		"username": "testbot",
		"language_code": "en",
		"is_premium": false,
		"can_join_groups": true,
		"can_read_all_group_messages": false,
		"supports_inline_queries": true,
		"can_connect_to_business": false,
		"has_main_web_app": true,
		"has_topics_enabled": false,
		"allows_users_to_create_topics": true
	}`
	var u User
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	assert.Equal(t, int64(123456789), u.ID)
	assert.True(t, u.IsBot)
	assert.Equal(t, "TestBot", u.FirstName)
	assert.True(t, u.CanJoinGroups)
	assert.True(t, u.SupportsInlineQueries)
	assert.True(t, u.HasMainWebApp)
	assert.True(t, u.AllowsUsersToCreateTopics)
}

// TestUserGetFullNameFallsBackToID verifies fallback when all name fields are empty.
func TestUserGetFullNameFallsBackToID(t *testing.T) {
	u := User{ID: 42}
	assert.Equal(t, "#42", u.GetFullName())
}

// TestChatNewFields verifies newly added Chat fields.
func TestChatNewFields(t *testing.T) {
	data := `{"id": -1001234567890, "type": "supergroup", "title": "Test Group", "is_forum": true, "is_direct_messages": false}`
	var c Chat
	require.NoError(t, json.Unmarshal([]byte(data), &c))
	assert.Equal(t, int64(-1001234567890), c.ID)
	assert.True(t, c.IsForum)
	assert.False(t, c.IsDirectMessages)
}

// TestChatDirectMessages verifies is_direct_messages field.
func TestChatDirectMessages(t *testing.T) {
	data := `{"id": -100111, "type": "supergroup", "is_direct_messages": true}`
	var c Chat
	require.NoError(t, json.Unmarshal([]byte(data), &c))
	assert.True(t, c.IsDirectMessages)
}

// TestMessageEntityNewFields verifies newly added MessageEntity fields.
func TestMessageEntityNewFields(t *testing.T) {
	data := `{
		"type": "text_mention",
		"offset": 0,
		"length": 5,
		"user": {"id": 123, "is_bot": false, "first_name": "Alice"}
	}`
	var e MessageEntity
	require.NoError(t, json.Unmarshal([]byte(data), &e))
	assert.Equal(t, "text_mention", e.Type)
	require.NotNil(t, e.User)
	assert.Equal(t, int64(123), e.User.ID)
	assert.Equal(t, "Alice", e.User.FirstName)
}

// TestMessageEntityCustomEmoji tests custom emoji entity fields.
func TestMessageEntityCustomEmoji(t *testing.T) {
	data := `{"type": "custom_emoji", "offset": 0, "length": 1, "custom_emoji_id": "5368324170671202286"}`
	var e MessageEntity
	require.NoError(t, json.Unmarshal([]byte(data), &e))
	assert.Equal(t, "custom_emoji", e.Type)
	assert.Equal(t, "5368324170671202286", e.CustomEmojiID)
}

// TestMessageEntityDateTime tests date_time entity fields (Bot API 9.5).
func TestMessageEntityDateTime(t *testing.T) {
	data := `{"type": "date_time", "offset": 0, "length": 10, "unix_time": 1700000000, "date_time_format": "date"}`
	var e MessageEntity
	require.NoError(t, json.Unmarshal([]byte(data), &e))
	assert.Equal(t, "date_time", e.Type)
	assert.Equal(t, 1700000000, e.UnixTime)
	assert.Equal(t, "date", e.DateTimeFormat)
}

// TestMediaThumbnailJSONTag verifies the thumbnail JSON tag was updated from "thumb" to "thumbnail".
func TestMediaThumbnailJSONTag(t *testing.T) {
	// Document
	docData := `{"file_id": "doc1", "file_unique_id": "uniq1", "thumbnail": {"file_id": "thumb1", "width": 100, "height": 100}}`
	var doc Document
	require.NoError(t, json.Unmarshal([]byte(docData), &doc))
	assert.Equal(t, "doc1", doc.FileID)
	assert.Equal(t, "uniq1", doc.FileUniqueID)
	require.NotNil(t, doc.Thumbnail)
	assert.Equal(t, "thumb1", doc.Thumbnail.FileID)

	// Video
	videoData := `{"file_id": "vid1", "file_unique_id": "uniq2", "width": 1920, "height": 1080, "duration": 60, "thumbnail": {"file_id": "vthumb1", "width": 320, "height": 180}}`
	var video Video
	require.NoError(t, json.Unmarshal([]byte(videoData), &video))
	assert.Equal(t, "vid1", video.FileID)
	assert.Equal(t, "uniq2", video.FileUniqueID)
	require.NotNil(t, video.Thumbnail)
	assert.Equal(t, "vthumb1", video.Thumbnail.FileID)

	// Sticker
	stickerData := `{"file_id": "stk1", "file_unique_id": "uniq3", "width": 512, "height": 512, "thumbnail": {"file_id": "sthumb1", "width": 100, "height": 100}}`
	var sticker Sticker
	require.NoError(t, json.Unmarshal([]byte(stickerData), &sticker))
	assert.Equal(t, "stk1", sticker.FileID)
	require.NotNil(t, sticker.Thumbnail)
	assert.Equal(t, "sthumb1", sticker.Thumbnail.FileID)
}

// TestAudioFileUniqueID verifies FileUniqueID is parsed in Audio.
func TestAudioFileUniqueID(t *testing.T) {
	data := `{"file_id": "audio1", "file_unique_id": "uaudio1", "duration": 120}`
	var a Audio
	require.NoError(t, json.Unmarshal([]byte(data), &a))
	assert.Equal(t, "audio1", a.FileID)
	assert.Equal(t, "uaudio1", a.FileUniqueID)
}

// TestUpdateNewFields verifies that Update parses new update fields.
func TestUpdateNewFields(t *testing.T) {
	data := `{
		"update_id": 123456,
		"poll": {
			"id": "poll1",
			"question": "Test question",
			"options": [
				{"text": "Yes", "voter_count": 5},
				{"text": "No", "voter_count": 3}
			],
			"total_voter_count": 8,
			"is_closed": false,
			"is_anonymous": true,
			"type": "regular",
			"allows_multiple_answers": false
		}
	}`
	var u Update
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	assert.Equal(t, 123456, u.UpdateID)
	require.NotNil(t, u.Poll)
	assert.Equal(t, "poll1", u.Poll.ID)
	assert.Equal(t, "Test question", u.Poll.Question)
	assert.Len(t, u.Poll.Options, 2)
	assert.Equal(t, 8, u.Poll.TotalVoterCount)
}

// TestUpdatePollAnswer verifies PollAnswer in Update.
func TestUpdatePollAnswer(t *testing.T) {
	data := `{
		"update_id": 789,
		"poll_answer": {
			"poll_id": "poll1",
			"user": {"id": 42, "is_bot": false, "first_name": "Bob"},
			"option_ids": [0]
		}
	}`
	var u Update
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	require.NotNil(t, u.PollAnswer)
	assert.Equal(t, "poll1", u.PollAnswer.PollID)
	require.NotNil(t, u.PollAnswer.User)
	assert.Equal(t, int64(42), u.PollAnswer.User.ID)
	assert.Equal(t, []int{0}, u.PollAnswer.OptionIDs)
}

// TestUpdateBusinessConnection verifies BusinessConnection in Update.
func TestUpdateBusinessConnection(t *testing.T) {
	data := `{
		"update_id": 999,
		"business_connection": {
			"id": "bc1",
			"user": {"id": 100, "is_bot": false, "first_name": "Alice"},
			"user_chat_id": 100,
			"date": 1700000000,
			"can_reply": true,
			"is_enabled": true
		}
	}`
	var u Update
	require.NoError(t, json.Unmarshal([]byte(data), &u))
	require.NotNil(t, u.BusinessConnection)
	assert.Equal(t, "bc1", u.BusinessConnection.ID)
	assert.True(t, u.BusinessConnection.CanReply)
	assert.True(t, u.BusinessConnection.IsEnabled)
}

// TestUpdateChatFromBusinessMessage checks Chat() method on update with business message.
func TestUpdateChatFromBusinessMessage(t *testing.T) {
	chat := &Chat{ID: 555, Type: "private"}
	u := Update{
		BusinessMessage: &Message{Chat: chat},
	}
	assert.Equal(t, chat, u.Chat())
}

// TestMessageNewFields verifies new Message fields are parsed.
func TestMessageNewFields(t *testing.T) {
	data := `{
		"message_id": 1,
		"date": 1700000000,
		"chat": {"id": 100, "type": "private"},
		"from": {"id": 42, "is_bot": false, "first_name": "Alice"},
		"edit_date": 1700001000,
		"has_protected_content": true,
		"is_paid_post": false,
		"media_group_id": "mg1",
		"author_signature": "Alice",
		"paid_star_count": 5,
		"effect_id": "effect123",
		"sender_boost_count": 3,
		"sender_tag": "VIP"
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	assert.Equal(t, 1, m.MessageID)
	assert.Equal(t, 1700001000, m.EditDate)
	assert.True(t, m.HasProtectedContent)
	assert.Equal(t, "mg1", m.MediaGroupID)
	assert.Equal(t, "Alice", m.AuthorSignature)
	assert.Equal(t, 5, m.PaidStarCount)
	assert.Equal(t, "effect123", m.EffectID)
	assert.Equal(t, 3, m.SenderBoostCount)
	assert.Equal(t, "VIP", m.SenderTag)
}

// TestMessageAnimation verifies Animation field parsing.
func TestMessageAnimation(t *testing.T) {
	data := `{
		"message_id": 2, "date": 1, "chat": {"id": 1, "type": "private"},
		"animation": {
			"file_id": "anim1", "file_unique_id": "uanim1",
			"width": 320, "height": 240, "duration": 5,
			"file_name": "test.gif", "mime_type": "image/gif"
		}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.Animation)
	assert.Equal(t, "anim1", m.Animation.FileID)
	assert.Equal(t, "uanim1", m.Animation.FileUniqueID)
	assert.Equal(t, 320, m.Animation.Width)
	assert.Equal(t, 5, m.Animation.Duration)
}

// TestMessageVideoNote verifies VideoNote field parsing.
func TestMessageVideoNote(t *testing.T) {
	data := `{
		"message_id": 3, "date": 1, "chat": {"id": 1, "type": "private"},
		"video_note": {"file_id": "vn1", "file_unique_id": "uvn1", "length": 240, "duration": 10}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.VideoNote)
	assert.Equal(t, "vn1", m.VideoNote.FileID)
	assert.Equal(t, 240, m.VideoNote.Length)
	assert.Equal(t, 10, m.VideoNote.Duration)
}

// TestMessageDice verifies Dice field parsing.
func TestMessageDice(t *testing.T) {
	data := `{
		"message_id": 4, "date": 1, "chat": {"id": 1, "type": "private"},
		"dice": {"emoji": "🎲", "value": 6}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.Dice)
	assert.Equal(t, "🎲", m.Dice.Emoji)
	assert.Equal(t, 6, m.Dice.Value)
}

// TestMessagePoll verifies Poll field parsing.
func TestMessagePoll(t *testing.T) {
	data := `{
		"message_id": 5, "date": 1, "chat": {"id": 1, "type": "private"},
		"poll": {
			"id": "p1", "question": "Q?",
			"options": [{"text": "A", "voter_count": 0}],
			"total_voter_count": 0, "is_closed": false,
			"is_anonymous": true, "type": "regular",
			"allows_multiple_answers": false
		}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.Poll)
	assert.Equal(t, "p1", m.Poll.ID)
	assert.Equal(t, "Q?", m.Poll.Question)
}

// TestMessageForumTopicCreated verifies ForumTopicCreated field.
func TestMessageForumTopicCreated(t *testing.T) {
	data := `{
		"message_id": 6, "date": 1, "chat": {"id": 1, "type": "supergroup"},
		"forum_topic_created": {"name": "New Topic", "icon_color": 16478047}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.ForumTopicCreated)
	assert.Equal(t, "New Topic", m.ForumTopicCreated.Name)
	assert.Equal(t, 16478047, m.ForumTopicCreated.IconColor)
}

// TestMessageGiveaway verifies Giveaway field.
func TestMessageGiveaway(t *testing.T) {
	data := `{
		"message_id": 7, "date": 1, "chat": {"id": 1, "type": "channel"},
		"giveaway": {
			"chats": [{"id": -1001, "type": "channel"}],
			"winners_selection_date": 1700100000,
			"winner_count": 5
		}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.Giveaway)
	assert.Equal(t, 5, m.Giveaway.WinnerCount)
	assert.Len(t, m.Giveaway.Chats, 1)
}

// TestMessageWebAppData verifies WebAppData field.
func TestMessageWebAppData(t *testing.T) {
	data := `{
		"message_id": 8, "date": 1, "chat": {"id": 1, "type": "private"},
		"web_app_data": {"data": "some_data", "button_text": "Open App"}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.WebAppData)
	assert.Equal(t, "some_data", m.WebAppData.Data)
	assert.Equal(t, "Open App", m.WebAppData.ButtonText)
}

// TestMessageNewChatMembersAsUsers verifies new_chat_members is []User.
func TestMessageNewChatMembersAsUsers(t *testing.T) {
	data := `{
		"message_id": 9, "date": 1, "chat": {"id": 1, "type": "group"},
		"new_chat_members": [
			{"id": 100, "is_bot": false, "first_name": "Alice"},
			{"id": 200, "is_bot": true, "first_name": "BotName"}
		]
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.Len(t, m.NewChatMembers, 2)
	assert.Equal(t, int64(100), m.NewChatMembers[0].ID)
	assert.Equal(t, "Alice", m.NewChatMembers[0].FirstName)
	assert.Equal(t, int64(200), m.NewChatMembers[1].ID)
	assert.True(t, m.NewChatMembers[1].IsBot)
}

// TestMessageLeftChatMemberAsUser verifies left_chat_member is *User.
func TestMessageLeftChatMemberAsUser(t *testing.T) {
	data := `{
		"message_id": 10, "date": 1, "chat": {"id": 1, "type": "group"},
		"left_chat_member": {"id": 300, "is_bot": false, "first_name": "Bob"}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.LeftChatMember)
	assert.Equal(t, int64(300), m.LeftChatMember.ID)
	assert.Equal(t, "Bob", m.LeftChatMember.FirstName)
}

// TestMessageReplyMarkup verifies ReplyMarkup field.
func TestMessageReplyMarkup(t *testing.T) {
	data := `{
		"message_id": 11, "date": 1, "chat": {"id": 1, "type": "private"},
		"reply_markup": {
			"inline_keyboard": [[{"text": "Click me", "callback_data": "btn1"}]]
		}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.ReplyMarkup)
	require.Len(t, m.ReplyMarkup.InlineKeyboard, 1)
	assert.Equal(t, "Click me", m.ReplyMarkup.InlineKeyboard[0][0].Text)
	assert.Equal(t, "btn1", m.ReplyMarkup.InlineKeyboard[0][0].CallbackData)
}

// TestInlineKeyboardButtonNewFields verifies IconCustomEmojiID and Style (Bot API 9.4).
func TestInlineKeyboardButtonNewFields(t *testing.T) {
	data := `{"text": "Star", "callback_data": "s1", "icon_custom_emoji_id": "emoji123", "style": "positive"}`
	var btn InlineKeyboardButton
	require.NoError(t, json.Unmarshal([]byte(data), &btn))
	assert.Equal(t, "Star", btn.Text)
	assert.Equal(t, "emoji123", btn.IconCustomEmojiID)
	assert.Equal(t, "positive", btn.Style)
}

// TestKeyboardButtonNewFields verifies IconCustomEmojiID and Style on regular keyboard button.
func TestKeyboardButtonNewFields(t *testing.T) {
	data := `{"text": "Go", "icon_custom_emoji_id": "emoji456", "style": "destructive"}`
	var btn KeyboardButton
	require.NoError(t, json.Unmarshal([]byte(data), &btn))
	assert.Equal(t, "Go", btn.Text)
	assert.Equal(t, "emoji456", btn.IconCustomEmojiID)
	assert.Equal(t, "destructive", btn.Style)
}

// TestChatAdministratorRightsFields verifies all fields of ChatAdministratorRights.
func TestChatAdministratorRightsFields(t *testing.T) {
	data := `{
		"is_anonymous": true,
		"can_manage_chat": true,
		"can_delete_messages": true,
		"can_manage_video_chats": true,
		"can_restrict_members": false,
		"can_promote_members": false,
		"can_change_info": true,
		"can_invite_users": true,
		"can_post_messages": true,
		"can_edit_messages": false,
		"can_pin_messages": true,
		"can_post_stories": true,
		"can_edit_stories": false,
		"can_delete_stories": false,
		"can_manage_topics": true,
		"can_manage_direct_messages": true,
		"can_manage_tags": false
	}`
	var r ChatAdministratorRights
	require.NoError(t, json.Unmarshal([]byte(data), &r))
	assert.True(t, r.IsAnonymous)
	assert.True(t, r.CanManageChat)
	assert.True(t, r.CanDeleteMessages)
	assert.True(t, r.CanManageVideoChats)
	assert.False(t, r.CanRestrictMembers)
	assert.True(t, r.CanChangeInfo)
	assert.True(t, r.CanInviteUsers)
	assert.True(t, r.CanPostMessages)
	assert.False(t, r.CanEditMessages)
	assert.True(t, r.CanPinMessages)
	assert.True(t, r.CanPostStories)
	assert.True(t, r.CanManageTopics)
	assert.True(t, r.CanManageDirectMessages)
	assert.False(t, r.CanManageTags)
}

// TestReactionType verifies ReactionType parsing.
func TestReactionType(t *testing.T) {
	t.Run("emoji", func(t *testing.T) {
		data := `{"type": "emoji", "emoji": "👍"}`
		var r ReactionType
		require.NoError(t, json.Unmarshal([]byte(data), &r))
		assert.Equal(t, "emoji", r.Type)
		assert.Equal(t, "👍", r.Emoji)
	})
	t.Run("custom_emoji", func(t *testing.T) {
		data := `{"type": "custom_emoji", "custom_emoji_id": "ce123"}`
		var r ReactionType
		require.NoError(t, json.Unmarshal([]byte(data), &r))
		assert.Equal(t, "custom_emoji", r.Type)
		assert.Equal(t, "ce123", r.CustomEmojiID)
	})
}

// TestMessageReactionUpdated verifies MessageReactionUpdated parsing.
func TestMessageReactionUpdated(t *testing.T) {
	data := `{
		"chat": {"id": 1, "type": "group"},
		"message_id": 100,
		"user": {"id": 42, "is_bot": false, "first_name": "Alice"},
		"date": 1700000000,
		"old_reaction": [],
		"new_reaction": [{"type": "emoji", "emoji": "❤️"}]
	}`
	var r MessageReactionUpdated
	require.NoError(t, json.Unmarshal([]byte(data), &r))
	assert.Equal(t, int64(1), r.Chat.ID)
	assert.Equal(t, 100, r.MessageID)
	require.NotNil(t, r.User)
	assert.Equal(t, int64(42), r.User.ID)
	assert.Empty(t, r.OldReaction)
	require.Len(t, r.NewReaction, 1)
	assert.Equal(t, "❤️", r.NewReaction[0].Emoji)
}

// TestUpdateChatMethod verifies the Chat() method covers all cases.
func TestUpdateChatMethod(t *testing.T) {
	chat := &Chat{ID: 1, Type: "private"}

	tests := []struct {
		name   string
		update Update
	}{
		{"message", Update{Message: &Message{Chat: chat}}},
		{"edited_message", Update{EditedMessage: &Message{Chat: chat}}},
		{"channel_post", Update{ChannelPost: &Message{Chat: chat}}},
		{"edited_channel_post", Update{EditedChannelPost: &Message{Chat: chat}}},
		{"business_message", Update{BusinessMessage: &Message{Chat: chat}}},
		{"edited_business_message", Update{EditedBusinessMessage: &Message{Chat: chat}}},
		{"callback_query", Update{CallbackQuery: &CallbackQuery{Message: &Message{Chat: chat}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, chat, tt.update.Chat())
		})
	}
}

// TestUpdateChatNil verifies Chat() returns nil for update with no chat.
func TestUpdateChatNil(t *testing.T) {
	u := Update{UpdateID: 1}
	assert.Nil(t, u.Chat())
}

// TestPollParsing tests Poll and PollOption parsing.
func TestPollParsing(t *testing.T) {
	data := `{
		"id": "12345",
		"question": "What is your favorite color?",
		"options": [
			{"text": "Red", "voter_count": 10},
			{"text": "Blue", "voter_count": 15},
			{"text": "Green", "voter_count": 5}
		],
		"total_voter_count": 30,
		"is_closed": false,
		"is_anonymous": true,
		"type": "regular",
		"allows_multiple_answers": false
	}`
	var p Poll
	require.NoError(t, json.Unmarshal([]byte(data), &p))
	assert.Equal(t, "12345", p.ID)
	assert.Equal(t, "What is your favorite color?", p.Question)
	assert.Len(t, p.Options, 3)
	assert.Equal(t, "Red", p.Options[0].Text)
	assert.Equal(t, 10, p.Options[0].VoterCount)
	assert.Equal(t, 30, p.TotalVoterCount)
	assert.True(t, p.IsAnonymous)
}

// TestBusinessConnectionParsing verifies BusinessConnection parsing.
func TestBusinessConnectionParsing(t *testing.T) {
	data := `{
		"id": "bc_test",
		"user": {"id": 777, "is_bot": false, "first_name": "Merchant"},
		"user_chat_id": 777,
		"date": 1700000000,
		"can_reply": true,
		"is_enabled": true
	}`
	var bc BusinessConnection
	require.NoError(t, json.Unmarshal([]byte(data), &bc))
	assert.Equal(t, "bc_test", bc.ID)
	assert.Equal(t, int64(777), bc.User.ID)
	assert.True(t, bc.CanReply)
	assert.True(t, bc.IsEnabled)
}

// TestChatBoostParsing verifies ChatBoostUpdated and ChatBoostRemoved.
func TestChatBoostParsing(t *testing.T) {
	updData := `{
		"chat": {"id": 1, "type": "channel"},
		"boost": {
			"boost_id": "boost1",
			"add_date": 1700000000,
			"expiration_date": 1800000000,
			"source": {"source": "premium", "user": {"id": 42, "is_bot": false, "first_name": "Alice"}}
		}
	}`
	var bu ChatBoostUpdated
	require.NoError(t, json.Unmarshal([]byte(updData), &bu))
	assert.Equal(t, "boost1", bu.Boost.BoostID)
	assert.Equal(t, "premium", bu.Boost.Source.Source)
	require.NotNil(t, bu.Boost.Source.User)
	assert.Equal(t, int64(42), bu.Boost.Source.User.ID)
}

// TestGiveawayParsing verifies Giveaway type parsing.
func TestGiveawayParsing(t *testing.T) {
	data := `{
		"chats": [{"id": -1001, "type": "channel", "title": "Test"}],
		"winners_selection_date": 1800000000,
		"winner_count": 10,
		"only_new_members": true,
		"prize_description": "Premium subscription",
		"premium_subscription_month_count": 3
	}`
	var g Giveaway
	require.NoError(t, json.Unmarshal([]byte(data), &g))
	assert.Len(t, g.Chats, 1)
	assert.Equal(t, 10, g.WinnerCount)
	assert.True(t, g.OnlyNewMembers)
	assert.Equal(t, "Premium subscription", g.PrizeDescription)
	assert.Equal(t, 3, g.PremiumSubscriptionMonthCount)
}

// TestAnimationParsing verifies Animation type.
func TestAnimationParsing(t *testing.T) {
	data := `{
		"file_id": "anim1",
		"file_unique_id": "uanim1",
		"width": 640,
		"height": 480,
		"duration": 3,
		"file_name": "funny.gif",
		"mime_type": "image/gif",
		"file_size": 51200
	}`
	var a Animation
	require.NoError(t, json.Unmarshal([]byte(data), &a))
	assert.Equal(t, "anim1", a.FileID)
	assert.Equal(t, "uanim1", a.FileUniqueID)
	assert.Equal(t, 640, a.Width)
	assert.Equal(t, "funny.gif", a.FileName)
	assert.Equal(t, "image/gif", a.MimeType)
}

// TestSuggestedPostTypes verifies suggested post service message types (Bot API 9.2).
func TestSuggestedPostTypes(t *testing.T) {
	data := `{
		"message_id": 50, "date": 1, "chat": {"id": 1, "type": "supergroup"},
		"suggested_post_approved": {"post_date": 1700050000},
		"suggested_post_info": {"suggested_post_date": 1700000000}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.SuggestedPostApproved)
	assert.Equal(t, 1700050000, m.SuggestedPostApproved.PostDate)
	require.NotNil(t, m.SuggestedPostInfo)
	assert.Equal(t, 1700000000, m.SuggestedPostInfo.SuggestedPostDate)
}

// TestChatOwnerChangedTypes verifies ChatOwnerLeft and ChatOwnerChanged (Bot API 9.4).
func TestChatOwnerChangedTypes(t *testing.T) {
	data := `{
		"message_id": 60, "date": 1, "chat": {"id": 1, "type": "group"},
		"chat_owner_changed": {
			"old_owner": {"id": 100, "is_bot": false, "first_name": "OldOwner"},
			"new_owner": {"id": 200, "is_bot": false, "first_name": "NewOwner"}
		}
	}`
	var m Message
	require.NoError(t, json.Unmarshal([]byte(data), &m))
	require.NotNil(t, m.ChatOwnerChanged)
	assert.Equal(t, int64(100), m.ChatOwnerChanged.OldOwner.ID)
	assert.Equal(t, int64(200), m.ChatOwnerChanged.NewOwner.ID)

	leftData := `{
		"message_id": 61, "date": 1, "chat": {"id": 1, "type": "group"},
		"chat_owner_left": {}
	}`
	var m2 Message
	require.NoError(t, json.Unmarshal([]byte(leftData), &m2))
	require.NotNil(t, m2.ChatOwnerLeft)
}
