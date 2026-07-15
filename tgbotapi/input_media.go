package tgbotapi

// This file adds the InputMedia* variants needed as dependencies of the InputRichBlock* family and
// InputRichMessageMedia (Bot API 10.2 Rich Messages). InputMediaAnimation, InputMediaAudio,
// InputMediaPhoto, and InputMediaVideo are pre-existing Bot API types not previously bound in this
// (partial) package; they are added here purely as typed homes for the "animation"/"audio"/"photo"/
// "video" fields of InputRichBlockAnimation/InputRichBlockAudio/InputRichBlockPhoto/InputRichBlockVideo
// and the "media" field of InputRichMessageMedia. InputMediaVoiceNote is new in Bot API 10.2.

// InputMediaAnimation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be
// sent.
//
// https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	// Type of the media, must be "animation"
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers, pass an HTTP URL
	// for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new
	// one using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side.
	Thumbnail string `json:"thumbnail,omitempty"`

	// Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing. Ignored
	// when this InputMediaAnimation is used as InputRichBlockAnimation.Animation; use
	// InputRichBlockAnimation.Caption instead.
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the animation caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified instead of
	// ParseMode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Pass True if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// Optional. Animation width.
	Width int `json:"width,omitempty"`

	// Optional. Animation height.
	Height int `json:"height,omitempty"`

	// Optional. Animation duration in seconds.
	Duration int `json:"duration,omitempty"`

	// Optional. Pass True if the animation needs to be covered with a spoiler animation.
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// InputMediaAudio represents an audio file to be treated as music to be sent.
//
// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	// Type of the media, must be "audio"
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers, pass an HTTP URL
	// for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new
	// one using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side.
	Thumbnail string `json:"thumbnail,omitempty"`

	// Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing. Ignored when
	// this InputMediaAudio is used as InputRichBlockAudio.Audio; use InputRichBlockAudio.Caption instead.
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the audio caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified instead of
	// ParseMode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Duration of the audio in seconds.
	Duration int `json:"duration,omitempty"`

	// Optional. Performer of the audio.
	Performer string `json:"performer,omitempty"`

	// Optional. Title of the audio.
	Title string `json:"title,omitempty"`
}

// InputMediaPhoto represents a photo to be sent.
//
// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	// Type of the media, must be "photo"
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers, pass an HTTP URL
	// for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new
	// one using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing. Ignored when
	// this InputMediaPhoto is used as InputRichBlockPhoto.Photo; use InputRichBlockPhoto.Caption instead.
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the photo caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified instead of
	// ParseMode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Pass True if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// Optional. Pass True if the photo needs to be covered with a spoiler animation.
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// InputMediaVideo represents a video to be sent.
//
// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	// Type of the media, must be "video"
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers, pass an HTTP URL
	// for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new
	// one using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is
	// supported server-side.
	Thumbnail string `json:"thumbnail,omitempty"`

	// Optional. Cover for the video in the message.
	Cover string `json:"cover,omitempty"`

	// Optional. Start timestamp for the video in the message.
	StartTimestamp int `json:"start_timestamp,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing. Ignored when
	// this InputMediaVideo is used as InputRichBlockVideo.Video; use InputRichBlockVideo.Caption instead.
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the video caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified instead of
	// ParseMode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Pass True if the caption must be shown above the message media.
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// Optional. Video width.
	Width int `json:"width,omitempty"`

	// Optional. Video height.
	Height int `json:"height,omitempty"`

	// Optional. Video duration in seconds.
	Duration int `json:"duration,omitempty"`

	// Optional. Pass True if the uploaded video is suitable for streaming.
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// Optional. Pass True if the video needs to be covered with a spoiler animation.
	HasSpoiler bool `json:"has_spoiler,omitempty"`
}

// InputMediaVoiceNote represents a voice message file to be sent (Bot API 10.2).
//
// https://core.telegram.org/bots/api#inputmediavoicenote
type InputMediaVoiceNote struct {
	// Type of the media, must be "voice_note"
	Type string `json:"type"`

	// File to send. Pass a file_id to send a file that exists on the Telegram servers, pass an HTTP URL
	// for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new
	// one using multipart/form-data under <file_attach_name> name.
	Media string `json:"media"`

	// Optional. Caption of the voice message to be sent, 0-1024 characters after entities parsing.
	// Ignored when this InputMediaVoiceNote is used as InputRichBlockVoiceNote.VoiceNote; use
	// InputRichBlockVoiceNote.Caption instead.
	Caption string `json:"caption,omitempty"`

	// Optional. Mode for parsing entities in the voice message caption.
	ParseMode string `json:"parse_mode,omitempty"`

	// Optional. List of special entities that appear in the caption, which can be specified instead of
	// ParseMode.
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Duration of the voice message in seconds.
	Duration int `json:"duration,omitempty"`
}
