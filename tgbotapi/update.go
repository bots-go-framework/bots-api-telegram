package tgbotapi

//go:generate ffjson $GOFILE

// Update is an update response, from GetUpdates.
type Update struct {
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
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
