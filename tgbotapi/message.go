package tgbotapi

//go:generate ffjson $GOFILE

import (
	"strconv"
	"strings"
	"time"
)

// Message is returned by almost every request and contains data about almost anything.
// https://core.telegram.org/bots/api#message
type Message struct {

	// Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageID int `json:"message_id"`

	// Optional. Unique identifier of a message thread or forum topic to which the message belongs; for supergroups and private chats only
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Optional. Information about the direct messages chat topic that contains the message
	DirectMessagesTopic *DirectMessagesTopic `json:"direct_messages_topic,omitempty"`

	// Optional. Sender of the message; may be empty for messages sent to channels
	From *User `json:"from,omitempty"`

	// Optional. Sender of the message when sent on behalf of a chat
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBoostCount int `json:"sender_boost_count,omitempty"`

	// Optional. The bot that actually sent the message on behalf of the business account
	SenderBusinessBot *User `json:"sender_business_bot,omitempty"`

	// Optional. Tag or custom title of the sender of the message; for supergroups only
	SenderTag string `json:"sender_tag,omitempty"`

	// Date the message was sent in Unix time
	Date int `json:"date"`

	// Optional. Unique identifier of the business connection from which the message was received
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Chat the message belongs to
	Chat *Chat `json:"chat,omitempty"`

	// Optional. Information about the original message for forwarded messages
	ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`

	// Optional. True, if the message is sent to a forum topic or a private chat with the bot
	IsTopicMessage bool `json:"is_topic_message,omitempty"`

	// Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`

	// Optional. For replies in the same chat and message thread, the original message
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// Optional. Information about the message being replied to, from another chat or forum topic
	ExternalReply *ExternalReplyInfo `json:"external_reply,omitempty"`

	// Optional. For replies that quote part of the original message, the quoted part of the message
	Quote *TextQuote `json:"quote,omitempty"`

	// Optional. For replies to a story, the original story
	ReplyToStory *Story `json:"reply_to_story,omitempty"`

	// Optional. Identifier of the specific checklist task that is being replied to
	ReplyToChecklistTaskID int `json:"reply_to_checklist_task_id,omitempty"`

	// Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`

	// Optional. Date the message was last edited in Unix time
	EditDate int `json:"edit_date,omitempty"`

	// Optional. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content,omitempty"`

	// Optional. True, if the message was sent by an implicit action (away or greeting business message, or scheduled)
	IsFromOffline bool `json:"is_from_offline,omitempty"`

	// Optional. True, if the message is a paid post
	IsPaidPost bool `json:"is_paid_post,omitempty"`

	// Optional. The unique identifier of a media message group this message belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`

	// Optional. Signature of the post author for messages in channels, or the custom title of an anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`

	// Optional. The number of Telegram Stars paid by the sender to send the message
	PaidStarCount int `json:"paid_star_count,omitempty"`

	// Optional. For text messages, the actual UTF-8 text of the message
	Text string `json:"text,omitempty"`

	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc.
	Entities *[]MessageEntity `json:"entities,omitempty"`

	// Optional. Options used for link preview generation for the message
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`

	// Optional. Information about suggested post parameters if the message is a suggested post
	SuggestedPostInfo *SuggestedPostInfo `json:"suggested_post_info,omitempty"`

	// Optional. Unique identifier of the message effect added to the message
	EffectID string `json:"effect_id,omitempty"`

	// Optional. Message is an animation, information about the animation
	Animation *Animation `json:"animation,omitempty"`

	// Optional. Message is an audio file
	Audio *Audio `json:"audio,omitempty"`

	// Optional. Message is a general file
	Document *Document `json:"document,omitempty"`

	// Optional. Message contains paid media
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`

	// Optional. Message is a photo
	Photo *[]PhotoSize `json:"photo,omitempty"`

	// Optional. Message is a sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Optional. Message is a forwarded story
	Story *Story `json:"story,omitempty"`

	// Optional. Message is a video
	Video *Video `json:"video,omitempty"`

	// Optional. Message is a video note
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Optional. Message is a voice message
	Voice *Voice `json:"voice,omitempty"`

	// Optional. Caption for the animation, audio, document, paid media, photo, video or voice
	Caption string `json:"caption,omitempty"`

	// Optional. For messages with a caption, special entities in the caption
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Optional. Message is a checklist
	Checklist *Checklist `json:"checklist,omitempty"`

	// Optional. Message is a shared contact
	Contact *Contact `json:"contact,omitempty"`

	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`

	// Optional. Message is a game
	Game interface{} `json:"game,omitempty"`

	// Optional. Message is a native poll
	Poll *Poll `json:"poll,omitempty"`

	// Optional. Message is a venue
	Venue *Venue `json:"venue,omitempty"`

	// Optional. Message is a shared location
	Location *Location `json:"location,omitempty"`

	// New members added to the group or supergroup
	NewChatMembers []User `json:"new_chat_members,omitempty"`

	// Optional. A member was removed from the group
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// Optional. Service message: chat owner has left
	ChatOwnerLeft *ChatOwnerLeft `json:"chat_owner_left,omitempty"`

	// Optional. Service message: chat owner has changed
	ChatOwnerChanged *ChatOwnerChanged `json:"chat_owner_changed,omitempty"`

	NewChatTitle          string       `json:"new_chat_title,omitempty"`          // optional
	NewChatPhoto          *[]PhotoSize `json:"new_chat_photo,omitempty"`          // optional
	DeleteChatPhoto       bool         `json:"delete_chat_photo,omitempty"`       // optional
	GroupChatCreated      bool         `json:"group_chat_created,omitempty"`      // optional
	SuperGroupChatCreated bool         `json:"supergroup_chat_created,omitempty"` // optional
	ChannelChatCreated    bool         `json:"channel_chat_created,omitempty"`    // optional

	// Optional. Service message: auto-delete timer settings changed
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`

	MigrateToChatID   int64    `json:"migrate_to_chat_id,omitempty"`   // optional
	MigrateFromChatID int64    `json:"migrate_from_chat_id,omitempty"` // optional
	PinnedMessage     *Message `json:"pinned_message,omitempty"`       // optional

	// Optional. Message is an invoice for a Payment
	// https://core.telegram.org/bots/api#payments
	Invoice *InvoiceConfig `json:"invoice,omitempty"`

	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"` // optional
	RefundedPayment   *RefundedPayment   `json:"refunded_payment,omitempty"`   // optional

	UserShared  *UserShared  `json:"user_shared,omitempty"`  // deprecated NON-DOCUMENTED FIELD
	UsersShared *UsersShared `json:"users_shared,omitempty"` // optional
	ChatShared  *ChatShared  `json:"chat_shared,omitempty"`

	// Optional. The domain name of the website on which the user has logged in
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// Optional. Service message: the user allowed the bot to write messages
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`

	// Optional. Service message. A user triggered another user's proximity alert
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`

	// Optional. Service message: user boosted the chat
	BoostAdded *ChatBoostAdded `json:"boost_added,omitempty"`

	// Optional. Service message: chat background set
	ChatBackgroundSet *ChatBackground `json:"chat_background_set,omitempty"`

	// Optional. Service message: some tasks in a checklist were marked as done or not done
	ChecklistTasksDone *ChecklistTasksDone `json:"checklist_tasks_done,omitempty"`

	// Optional. Service message: tasks were added to a checklist
	ChecklistTasksAdded *ChecklistTasksAdded `json:"checklist_tasks_added,omitempty"`

	// Optional. Service message: the price for paid messages in the direct messages chat has changed
	DirectMessagePriceChanged *DirectMessagePriceChanged `json:"direct_message_price_changed,omitempty"`

	// Optional. Service message: forum topic created
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`

	// Optional. Service message: forum topic edited
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`

	// Optional. Service message: forum topic closed
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed,omitempty"`

	// Optional. Service message: forum topic reopened
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened,omitempty"`

	// Optional. Service message: the 'General' forum topic hidden
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`

	// Optional. Service message: the 'General' forum topic unhidden
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`

	// Optional. Service message: a scheduled giveaway was created
	GiveawayCreated *GiveawayCreated `json:"giveaway_created,omitempty"`

	// Optional. The message is a scheduled giveaway message
	Giveaway *Giveaway `json:"giveaway,omitempty"`

	// Optional. A giveaway with public winners was completed
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`

	// Optional. Service message: a giveaway without public winners was completed
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed,omitempty"`

	// Optional. Service message: the price for paid messages has changed in the chat
	PaidMessagePriceChanged *PaidMessagePriceChanged `json:"paid_message_price_changed,omitempty"`

	// Optional. Service message: a suggested post was approved
	SuggestedPostApproved *SuggestedPostApproved `json:"suggested_post_approved,omitempty"`

	// Optional. Service message: approval of a suggested post has failed
	SuggestedPostApprovalFailed *SuggestedPostApprovalFailed `json:"suggested_post_approval_failed,omitempty"`

	// Optional. Service message: a suggested post was declined
	SuggestedPostDeclined *SuggestedPostDeclined `json:"suggested_post_declined,omitempty"`

	// Optional. Service message: payment for a suggested post was received
	SuggestedPostPaid *SuggestedPostPaid `json:"suggested_post_paid,omitempty"`

	// Optional. Service message: payment for a suggested post was refunded
	SuggestedPostRefunded *SuggestedPostRefunded `json:"suggested_post_refunded,omitempty"`

	// Optional. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`

	// Optional. Service message: video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`

	// Optional. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`

	// Optional. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`

	// Optional. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data,omitempty"`

	// Optional. Service message: upgrade of a gift was purchased
	GiftUpgradeSent *GiftInfo `json:"gift_upgrade_sent,omitempty"`

	Gift       *GiftInfo       `json:"gift,omitempty"`
	UniqueGift *UniqueGiftInfo `json:"unique_gift,omitempty"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// Deprecated: use ForwardOrigin instead
	ForwardFrom *User `json:"forward_from,omitempty"`
	// Deprecated: use ForwardOrigin instead
	ForwardDate int `json:"forward_date,omitempty"`
}

func (m *Message) GetMessageID() string {
	return strconv.Itoa(m.MessageID)
}

// Time converts the message timestamp into a Time.
func (m *Message) Time() time.Time {
	return time.Unix(int64(m.Date), 0)
}

// IsCommand returns true if message starts with '/'.
func (m *Message) IsCommand() bool {
	return m.Text != "" && m.Text[0] == '/'
}

// Command checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at bot syntax, it removes the bot name.
func (m *Message) Command() string {
	if !m.IsCommand() {
		return ""
	}

	command := strings.SplitN(m.Text, " ", 2)[0][1:]

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

// CommandArguments checks if the message was a command and if it was,
// returns all text after the command name. If the Message was not a
// command, it returns an empty string.
func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	split := strings.SplitN(m.Text, " ", 2)
	if len(split) != 2 {
		return ""
	}

	return strings.SplitN(m.Text, " ", 2)[1]
}
