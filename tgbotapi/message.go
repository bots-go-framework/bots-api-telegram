package tgbotapi

//go:generate ffjson $GOFILE

import (
	"strings"
	"time"
)

// Message is returned by almost every request and contains data about almost anything.
type Message struct {

	// Unique message identifier inside this chat. In specific instances (e.g., message containing a video sent to a big chat), the server might automatically schedule a message instead of sending it immediately. In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageID int `json:"message_id"`

	// Optional. Unique identifier of a message thread to which the message belongs; for supergroups only
	MessageThreadID int `json:"message_thread_id"`

	// Optional. Sender of the message; may be empty for messages sent to channels. For backward compatibility, if the message was sent on behalf of a chat, the field contains a fake sender user in non-channel chats
	From *User `json:"from,omitempty"`

	// Optional. Sender of the message when sent on behalf of a chat. For example, the supergroup itself for messages sent by its anonymous administrators or a linked channel for messages automatically forwarded to the channel's discussion group. For backward compatibility, if the message was sent on behalf of a chat, the field from contains a fake sender user in non-channel chats.
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// Optional. If the sender of the message boosted the chat, the number of boosts added by the user
	SenderBootCount int `json:"sender_boot_count,omitempty"`

	// Optional. The bot that actually sent the message on behalf of the business account. Available only for outgoing messages sent on behalf of the connected business account.
	SenderBusinessBot *User `json:"sender_business_bot,omitempty"`

	// Date the message was sent in Unix time. It is always a positive number, representing a valid date.
	Date int `json:"date"`

	// Optional. Unique identifier of the business connection from which the message was received.
	// If non-empty, the message belongs to a chat of the corresponding business account
	// that is independent from any potential bot chat which might share the same identifier.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Chat the message belongs to
	Chat *Chat `json:"chat,omitempty"`

	// Optional. Information about the original message for forwarded messages
	ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`

	// Optional. True, if the message is sent to a forum topic
	IsTopicMessage bool `json:"is_topic_message,omitempty"`

	// Optional. True, if the message is a channel post that was automatically forwarded to the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`

	UserShared            *UserShared      `json:"user_shared,omitempty"`             // optional NON-DOCUMENTED FIELD
	UsersShared           *UsersShared     `json:"users_shared,omitempty"`            // optional
	ForwardFrom           *User            `json:"forward_from,omitempty"`            // optional
	ForwardDate           int              `json:"forward_date,omitempty"`            // optional
	ReplyToMessage        *Message         `json:"reply_to_message,omitempty"`        // optional
	Text                  string           `json:"text,omitempty"`                    // optional
	Entities              *[]MessageEntity `json:"entities,omitempty"`                // optional
	Audio                 *Audio           `json:"audio,omitempty"`                   // optional
	Document              *Document        `json:"document,omitempty"`                // optional
	Photo                 *[]PhotoSize     `json:"photo,omitempty"`                   // optional
	Sticker               *Sticker         `json:"sticker,omitempty"`                 // optional
	Video                 *Video           `json:"video,omitempty"`                   // optional
	Voice                 *Voice           `json:"voice,omitempty"`                   // optional
	Caption               string           `json:"caption,omitempty"`                 // optional
	Contact               *Contact         `json:"contact,omitempty"`                 // optional
	Location              *Location        `json:"location,omitempty"`                // optional
	Venue                 *Venue           `json:"venue,omitempty"`                   // optional
	NewChatParticipant    *ChatMember      `json:"new_chat_participant,omitempty"`    // Obsolete
	NewChatMember         *ChatMember      `json:"new_chat_member,omitempty"`         // Obsolete
	NewChatMembers        []ChatMember     `json:"new_chat_members,omitempty"`        // optional
	LeftChatMember        *ChatMember      `json:"left_chat_member,omitempty"`        // optional
	NewChatTitle          string           `json:"new_chat_title,omitempty"`          // optional
	NewChatPhoto          *[]PhotoSize     `json:"new_chat_photo,omitempty"`          // optional
	DeleteChatPhoto       bool             `json:"delete_chat_photo,omitempty"`       // optional
	GroupChatCreated      bool             `json:"group_chat_created,omitempty"`      // optional
	SuperGroupChatCreated bool             `json:"supergroup_chat_created,omitempty"` // optional
	ChannelChatCreated    bool             `json:"channel_chat_created,omitempty"`    // optional
	MigrateToChatID       int64            `json:"migrate_to_chat_id,omitempty"`      // optional
	MigrateFromChatID     int64            `json:"migrate_from_chat_id,omitempty"`    // optional
	PinnedMessage         *Message         `json:"pinned_message,omitempty"`          // optional

	// Optional. Message is an invoice for a Payment, information about the invoice.
	// https://core.telegram.org/bots/api#payments
	Invoice *InvoiceConfig `json:"invoice"`

	// Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot"`

	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"` // optional
	RefundedPayment   *RefundedPayment   `json:"refunded_payment,omitempty"`   //	optional

	ChatShared *ChatShared `json:"chat_shared,omitempty"`

	Gift       *GiftInfo       `json:"gift,omitempty"`
	UniqueGift *UniqueGiftInfo `json:"unique_gift,omitempty"`
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
