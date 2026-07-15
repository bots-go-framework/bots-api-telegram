package tgbotapi

import (
	"fmt"
	"net/url"
)

// LivePhotoConfig contains information about a sendLivePhoto request.
//
// https://core.telegram.org/bots/api#sendlivephoto
type LivePhotoConfig struct {
	BaseChat

	// LivePhoto should be PhotoUrl, FileID or InputFile
	Photo                 Photo  `json:"live_photo"`
	Caption               string `json:"caption,omitempty"`
	ParseMode             string `json:"parse_mode,omitempty"`
	ShowCaptionAboveMedia bool   `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool   `json:"has_spoiler,omitempty"`
}

// Values returns url.Values representation of LivePhotoConfig.
func (v LivePhotoConfig) Values() (url.Values, error) {
	const livePhotoParamName = "live_photo"

	values, _ := v.BaseChat.Values()

	switch photoType := v.Photo.PhotoType(); photoType {
	case PhotoTypeUrl:
		values.Add(livePhotoParamName, string(v.Photo.(PhotoUrl)))
	case PhotoTypeFileID:
		values.Add(livePhotoParamName, string(v.Photo.(FileID)))
	case PhotoTypeInputFile:
		return values, fmt.Errorf("not implemented yet for photo type: %v", photoType)
	}
	if v.Caption != "" {
		values.Add("caption", v.Caption)
	}
	if v.ShowCaptionAboveMedia {
		values.Add("show_caption_above_media", "true")
	}
	if v.HasSpoiler {
		values.Add("has_spoiler", "true")
	}
	return values, nil
}

// TelegramMethod returns Telegram API method name for sending a LivePhoto.
func (LivePhotoConfig) TelegramMethod() string {
	return "sendLivePhoto"
}
