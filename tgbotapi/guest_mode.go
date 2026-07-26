package tgbotapi

// SentGuestMessage describes a message sent by the bot in reply to a guest query, i.e. a message
// sent within a chat the bot is not a member of via Guest Mode.
//
// https://core.telegram.org/bots/api#sentguestmessage
type SentGuestMessage struct {
	// Unique identifier of the inline message sent by the guest bot.
	InlineMessageID string `json:"inline_message_id"`
}
