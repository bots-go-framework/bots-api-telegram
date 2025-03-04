package tgbotapi

//go:generate ffjson $GOFILE

// UserShared is a NON-DOCUMENTED TELEGRAM MESSAGE FIELD
type UserShared struct {
	UserID    int `json:"userID"`
	RequestID int `json:"request_id"`
}

// UsersShared contains information about the users whose identifiers were shared with the bot using a KeyboardButtonRequestUsers button.
// https://core.telegram.org/bots/api#usersshared
type UsersShared struct {
	RequestID int          `json:"request_id"`
	UserIDs   []int        `json:"user_ids"`
	Users     []SharedUser `json:"users"`
}

// SharedUser contains information about a user that was shared with the bot using a KeyboardButtonRequestUsers button.
// https://core.telegram.org/bots/api#shareduser
type SharedUser struct {
	UserID int `json:"user_id"`

	// Optional. First name of the user, if the name was requested by the bot
	FirstName string `json:"first_name,omitempty"`

	// Optional. Last name of the user, if the name was requested by the bot
	LastName string `json:"last_name,omitempty"`

	// Optional. Username of the user, if the name was requested by the bot
	Username string `json:"username,omitempty"`

	Photo []PhotoSize `json:"photo,omitempty"`
}
