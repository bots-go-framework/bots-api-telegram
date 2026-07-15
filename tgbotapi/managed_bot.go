package tgbotapi

// KeyboardButtonRequestManagedBot defines the parameters for the creation of a managed bot.
// Information about the created bot will be shared with the bot using the update managed_bot
// and a Message with the field managed_bot_created.
//
// https://core.telegram.org/bots/api#keyboardbuttonrequestmanagedbot
type KeyboardButtonRequestManagedBot struct {
	// Signed 32-bit identifier of the request. Must be unique within the message.
	RequestID int `json:"request_id"`

	// Optional. Suggested name for the bot
	SuggestedName string `json:"suggested_name,omitempty"`

	// Optional. Suggested username for the bot
	SuggestedUsername string `json:"suggested_username,omitempty"`
}

// ManagedBotCreated contains information about the bot that was created to be managed by the current bot.
//
// https://core.telegram.org/bots/api#managedbotcreated
type ManagedBotCreated struct {
	// Information about the bot. The bot's token can be fetched using the method getManagedBotToken.
	Bot User `json:"bot"`
}

// ManagedBotUpdated contains information about the creation, token update, or owner update of a bot
// that is managed by the current bot.
//
// https://core.telegram.org/bots/api#managedbotupdated
type ManagedBotUpdated struct {
	// User that created the bot
	User User `json:"user"`

	// Information about the bot. Token of the bot can be fetched using the method getManagedBotToken.
	Bot User `json:"bot"`
}

// PreparedKeyboardButton describes a keyboard button to be used by a user of a Mini App.
//
// https://core.telegram.org/bots/api#preparedkeyboardbutton
type PreparedKeyboardButton struct {
	// Unique identifier of the keyboard button
	ID string `json:"id"`
}

// BotAccessSettings describes the access settings granted to a bot for managing a bot it created. Used
// by getManagedBotAccessSettings and setManagedBotAccessSettings.
//
// https://core.telegram.org/bots/api#botaccesssettings
type BotAccessSettings struct {
	// Optional. True, if the bot can read messages sent to the managed bot
	CanReadMessages bool `json:"can_read_messages,omitempty"`

	// Optional. True, if the bot can send messages on behalf of the managed bot
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// Optional. True, if the bot can change the managed bot's settings, such as its name,
	// description and commands
	CanManageSettings bool `json:"can_manage_settings,omitempty"`
}

// PollOptionAdded describes a service message about a new option being added to a poll.
//
// https://core.telegram.org/bots/api#polloptionadded
type PollOptionAdded struct {
	// Optional. Message containing the poll
	PollMessage *Message `json:"poll_message,omitempty"`

	// Unique identifier of the added option
	OptionPersistentID string `json:"option_persistent_id"`

	// Option text
	OptionText string `json:"option_text"`

	// Optional. Special entities that appear in the option text
	OptionTextEntities []MessageEntity `json:"option_text_entities,omitempty"`
}

// PollOptionDeleted describes a service message about an option being deleted from a poll.
//
// https://core.telegram.org/bots/api#polloptiondeleted
type PollOptionDeleted struct {
	// Optional. Message containing the poll
	PollMessage *Message `json:"poll_message,omitempty"`

	// Unique identifier of the deleted option
	OptionPersistentID string `json:"option_persistent_id"`

	// Option text
	OptionText string `json:"option_text"`

	// Optional. Special entities that appear in the option text
	OptionTextEntities []MessageEntity `json:"option_text_entities,omitempty"`
}
