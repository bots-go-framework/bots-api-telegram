package tgbotapi

type ChatShared struct {
	RequestID int         `json:"request_id"`      // Identifier of the request
	ChatID    int         `json:"chat_id"`         // Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot may not have access to the chat and could be unable to use this identifier, unless the chat is already known to the bot by some other means.
	Title     string      `json:"title"`           // Optional. Title of the chat, if the title was requested by the bot.
	Username  string      `json:"username"`        // Optional. Username of the chat, if the username was requested by the bot and available.
	Photo     []PhotoSize `json:"photo,omitempty"` // Optional. Available sizes of the chat photo, if the photo was requested by the bot
}
