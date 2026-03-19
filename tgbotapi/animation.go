package tgbotapi

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id,omitempty"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// VideoNote represents a video message (available in Telegram apps as of v.4.0).
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id,omitempty"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int        `json:"file_size,omitempty"`
}

// Dice represents an animated emoji that displays a random value.
// https://core.telegram.org/bots/api#dice
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// Story represents a message about a forwarded story in the chat.
// https://core.telegram.org/bots/api#story
type Story struct {
	Chat *Chat `json:"chat"`
	ID   int   `json:"id"`
}
