package tgbotapi

// PaidMedia describes paid media. It can be one of:
// - PaidMediaPreview
// - PaidMediaPhoto
// - PaidMediaVideo
// https://core.telegram.org/bots/api#paidmedia
type PaidMedia struct {
	// Type of the paid media — "preview", "photo", or "video"
	Type string `json:"type"`

	// For "preview": media width if known
	Width int `json:"width,omitempty"`

	// For "preview": media height if known
	Height int `json:"height,omitempty"`

	// For "preview": duration of the media in seconds if known
	Duration int `json:"duration,omitempty"`

	// For "photo": the photo
	Photo []PhotoSize `json:"photo,omitempty"`

	// For "video": the video
	Video *Video `json:"video,omitempty"`
}

// PaidMediaInfo describes the paid media added to a message.
// https://core.telegram.org/bots/api#paidmediainfo
type PaidMediaInfo struct {
	StarCount int         `json:"star_count"`
	PaidMedia []PaidMedia `json:"paid_media"`
}

// PaidMediaPurchased contains information about a paid media purchase.
// https://core.telegram.org/bots/api#paidmediapurchased
type PaidMediaPurchased struct {
	From             User   `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}
