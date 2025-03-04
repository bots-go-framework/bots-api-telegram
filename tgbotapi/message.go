package tgbotapi

//go:generate ffjson $GOFILE

import (
	"strings"
	"time"
)

// Message is returned by almost every request, and contains data about almost anything.
type Message struct {
	MessageID             int              `json:"message_id"`
	From                  *User            `json:"from,omitempty"` // optional
	Date                  int              `json:"date"`
	Chat                  *Chat            `json:"chat,omitempty"`
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
