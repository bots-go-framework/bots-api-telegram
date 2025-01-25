package tgbotapi

//go:generate ffjson $GOFILE

// CallbackQuery is data sent when a keyboard button with callback data
// is clicked.
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message,omitempty"`           // optional
	ChatInstance    string   `json:"chat_instance,omitempty"`     // optional
	InlineMessageID string   `json:"inline_message_id,omitempty"` // optional
	Data            string   `json:"data,omitempty"`              // optional
}
