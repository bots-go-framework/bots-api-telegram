package tgbotapi

import (
	"fmt"
	"net/url"
)

// PhotoConfig contains information about a SendPhoto request.
type PhotoConfig struct {
	BaseChat

	// Photo should be PhotoUrl, FileID or InputFile
	Photo                 Photo  `json:"photo"`
	Caption               string `json:"caption,omitempty"`
	ParseMode             string `json:"parse_mode,omitempty"`
	ShowCaptionAboveMedia bool   `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool   `json:"has_spoiler,omitempty"`
}

type Photo interface {
	PhotoType() PhotoType
}

type PhotoUrl string

func (PhotoUrl) PhotoType() PhotoType {
	return PhotoTypeUrl
}

type FileID string

func (FileID) PhotoType() PhotoType {
	return PhotoTypeFileID
}

type PhotoType int

const (
	PhotoTypeInputFile PhotoType = 1
	PhotoTypeFileID    PhotoType = 2
	PhotoTypeUrl       PhotoType = 3
)

// Values returns url.Values representation of PhotoConfig.
func (v PhotoConfig) Values() (url.Values, error) {
	const photoParamName = "photo"

	values, _ := v.BaseChat.Values()

	switch photoType := v.Photo.PhotoType(); photoType {
	case PhotoTypeUrl:
		values.Add(photoParamName, string(v.Photo.(PhotoUrl)))
	case PhotoTypeFileID:
		values.Add(photoParamName, string(v.Photo.(FileID)))
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

// TelegramMethod returns Telegram API method name for sending Photo.
func (PhotoConfig) TelegramMethod() string {
	return "sendPhoto"
}
