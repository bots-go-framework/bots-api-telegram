package tgbotapi

// ChatMemberStatus represents the status of a chat member.
type ChatMemberStatus string

const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

// ChatMemberUpdated represents changes in the status of a chat member.
// https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	Chat                    Chat        `json:"chat"`
	From                    User        `json:"from"`
	Date                    int         `json:"date"`
	OldChatMember           interface{} `json:"old_chat_member"`
	NewChatMember           interface{} `json:"new_chat_member"`
	InviteLink              interface{} `json:"invite_link,omitempty"`
	ViaJoinRequest          bool        `json:"via_join_request,omitempty"`
	ViaChatFolderInviteLink bool        `json:"via_chat_folder_invite_link,omitempty"`
}

// ChatJoinRequest represents a join request sent to a chat.
// https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	Chat       Chat        `json:"chat"`
	From       User        `json:"from"`
	UserChatID int64       `json:"user_chat_id"`
	Date       int         `json:"date"`
	Bio        string      `json:"bio,omitempty"`
	InviteLink interface{} `json:"invite_link,omitempty"`
}
