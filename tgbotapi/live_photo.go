package tgbotapi

// LivePhoto represents a photo with an accompanying short video, played back automatically alongside the photo.
//
// https://core.telegram.org/bots/api#livephoto
type LivePhoto struct {
	// Photo, in different sizes, that is shown as the still frame of the live photo
	Photo []PhotoSize `json:"photo,omitempty"`

	// Identifier for the video part of the live photo, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for the video part of the live photo, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"`

	// Video width as defined by the sender
	Width int `json:"width,omitempty"`

	// Video height as defined by the sender
	Height int `json:"height,omitempty"`

	// Duration of the video in seconds as defined by the sender
	Duration int `json:"duration,omitempty"`

	// Optional. MIME type of the video, as defined by the sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size in bytes
	FileSize int `json:"file_size,omitempty"`
}

// InputMediaLivePhoto represents a live photo to be sent as part of a media group or used to edit
// the media of a message.
//
// https://core.telegram.org/bots/api#inputmedialivephoto
type InputMediaLivePhoto struct {
	// Type of the result, must be "live_photo"
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers, pass an HTTP URL
	// for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new
	// one using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`

	// Optional. Caption of the live photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the live photo caption
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. A JSON-serialized list of special entities that appear in the caption,
	// which can be specified instead of ParseMode
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Pass True, if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// Optional. Pass True if the live photo needs to be covered with a spoiler animation
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// PaidMediaLivePhoto describes a paid media that is a live photo.
//
// https://core.telegram.org/bots/api#paidmedialivephoto
type PaidMediaLivePhoto struct {
	// Type of the paid media, always "live_photo"
	Type string `json:"type"`

	// The live photo
	LivePhoto LivePhoto `json:"live_photo"`
}

// InputPaidMediaLivePhoto describes a live photo to be sent as paid media.
//
// https://core.telegram.org/bots/api#inputpaidmedialivephoto
type InputPaidMediaLivePhoto struct {
	// Type of the media, must be "live_photo"
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers, pass an HTTP URL
	// for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new
	// one using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`
}
