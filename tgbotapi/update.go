package tgbotapi

//go:generate ffjson $GOFILE

// Update is an update response, from GetUpdates.
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditedChannelPost  *Message            `json:"edited_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query,omitempty"`
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
	case update.CallbackQuery != nil:
		if update.CallbackQuery.Message != nil {
			return update.CallbackQuery.Message.Chat
		}
	}
	return nil
}
