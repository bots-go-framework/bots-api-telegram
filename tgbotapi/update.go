package tgbotapi

//go:generate ffjson $GOFILE

// Update is an update response, from GetUpdates.
// https://core.telegram.org/bots/api#update
type Update struct {
	UpdateID int `json:"update_id"`

	// Optional. New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// Optional. New version of a message that is known to the bot and was edited
	EditedMessage *Message `json:"edited_message,omitempty"`

	// Optional. New incoming channel post of any kind
	ChannelPost *Message `json:"channel_post,omitempty"`

	// Optional. New version of a channel post that is known to the bot and was edited
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// Optional. The bot was connected to or disconnected from a business account
	BusinessConnection *BusinessConnection `json:"business_connection,omitempty"`

	// Optional. New message from a connected business account
	BusinessMessage *Message `json:"business_message,omitempty"`

	// Optional. New version of a message from a connected business account
	EditedBusinessMessage *Message `json:"edited_business_message,omitempty"`

	// Optional. Messages were deleted from a connected business account
	DeletedBusinessMessages *BusinessMessagesDeleted `json:"deleted_business_messages,omitempty"`

	// Optional. A reaction to a message was changed by a user
	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`

	// Optional. Reactions to a message with anonymous reactions were changed
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`

	// Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	// Optional. The result of an inline query that was chosen by a user
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// Optional. New incoming pre-checkout query
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// Optional. A user purchased paid media with a non-empty payload
	PurchasedPaidMedia *PaidMediaPurchased `json:"purchased_paid_media,omitempty"`

	// Optional. New poll state
	Poll *Poll `json:"poll,omitempty"`

	// Optional. A user changed their answer in a non-anonymous poll
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// Optional. The bot's chat member status was updated in a chat
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`

	// Optional. A chat member's status was updated in a chat
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`

	// Optional. A request to join the chat has been sent
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`

	// Optional. A chat boost was added or changed
	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`

	// Optional. A boost was removed from a chat
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
}

// Chat provides chat struct for the update
func (update Update) Chat() *Chat {
	switch {
	case update.Message != nil:
		return update.Message.Chat
	case update.EditedMessage != nil:
		return update.EditedMessage.Chat
	case update.ChannelPost != nil:
		return update.ChannelPost.Chat
	case update.EditedChannelPost != nil:
		return update.EditedChannelPost.Chat
	case update.BusinessMessage != nil:
		return update.BusinessMessage.Chat
	case update.EditedBusinessMessage != nil:
		return update.EditedBusinessMessage.Chat
	case update.CallbackQuery != nil:
		if update.CallbackQuery.Message != nil {
			return update.CallbackQuery.Message.Chat
		}
	}
	return nil
}
