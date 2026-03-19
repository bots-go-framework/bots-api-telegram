package tgbotapi

// VideoChatScheduled represents a service message about a video chat scheduled in the chat.
// https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	StartDate int `json:"start_date"`
}

// VideoChatStarted represents a service message about a video chat started in the chat.
// https://core.telegram.org/bots/api#videochatstarted
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a video chat ended in the chat.
// https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited represents a service message about new members invited to a video chat.
// https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	Users []User `json:"users"`
}

// WebAppData contains data sent from a Web App to the bot.
// https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

// ProximityAlertTriggered represents the content of a service message, sent whenever a user in the chat
// triggers a proximity alert set by another user.
// https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	Traveler User `json:"traveler"`
	Watcher  User `json:"watcher"`
	Distance int  `json:"distance"`
}

// MessageAutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings.
// https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// ChatBackground represents a chat background.
// https://core.telegram.org/bots/api#chatbackground
type ChatBackground struct {
	Type interface{} `json:"type"`
}
