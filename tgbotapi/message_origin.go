package tgbotapi

type MessageOriginType string

const (
	MessageOriginTypeUser       MessageOriginType = "user"
	MessageOriginTypeHiddenUser MessageOriginType = "hidden_user"
)

type messageOrigin struct {
	// Type of the message origin, either "user", "hidden_user", "chat", "channel"
	Type MessageOriginType `json:"type"`
	Date int               `json:"date"`
}
type MessageOrigin interface {
	MessageOriginType() MessageOriginType
}

var _ MessageOrigin = (*MessageOriginUser)(nil)

type MessageOriginUser struct {
	messageOrigin
	// User that sent the message originally
	SenderUser *User `json:"sender_user,omitempty"`
}

func (MessageOriginUser) MessageOriginType() MessageOriginType {
	return MessageOriginTypeUser
}

var _ MessageOrigin = (*MessageOriginHiddenUser)(nil)

type MessageOriginHiddenUser struct {
	messageOrigin
	// Name of the user that sent the message originally
	SenderUserName string `json:"sender_user_name"`
}

func (MessageOriginHiddenUser) MessageOriginType() MessageOriginType {
	return MessageOriginTypeHiddenUser
}

type MessageOriginChat struct {
	messageOrigin
	// Chat that sent the message originally
	SenderChat Chat `json:"sender_chat"`
	// Optional. For messages originally sent by an anonymous chat administrator, original message author signature
	AuthorSignature string `json:"author_signature,omitempty"`
}

type MessageOriginChannel struct {
	messageOrigin

	// Channel chat to which the message was originally sent
	Chat Chat `json:"chat"`

	// Unique message identifier inside the chat
	MessageID int `json:"message_id"`

	// Optional. For messages originally sent by an anonymous chat administrator, original message author signature
	AuthorSignature string `json:"author_signature,omitempty"`
}
