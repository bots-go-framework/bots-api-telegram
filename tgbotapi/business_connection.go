package tgbotapi

// BusinessConnection describes the connection of the bot with a business account.
// https://core.telegram.org/bots/api#businessconnection
type BusinessConnection struct {
	ID         string `json:"id"`
	User       User   `json:"user"`
	UserChatID int64  `json:"user_chat_id"`
	Date       int    `json:"date"`
	CanReply   bool   `json:"can_reply"`
	IsEnabled  bool   `json:"is_enabled"`
}

// BusinessMessagesDeleted is received when messages are deleted from a connected business account.
// https://core.telegram.org/bots/api#businessmessagesdeleted
type BusinessMessagesDeleted struct {
	BusinessConnectionID string `json:"business_connection_id"`
	Chat                 Chat   `json:"chat"`
	MessageIDs           []int  `json:"message_ids"`
}
