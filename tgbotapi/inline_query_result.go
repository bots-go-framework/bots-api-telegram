package tgbotapi

//go:generate ffjson $GOFILE

import (
	"errors"
	"fmt"
	"strings"
)

type InlineQueryResultType string

const (
	InlineQueryResultTypeArticle  InlineQueryResultType = "article"
	InlineQueryResultTypeAudio    InlineQueryResultType = "audio"
	InlineQueryResultTypeContact  InlineQueryResultType = "contact"
	InlineQueryResultTypeGame     InlineQueryResultType = "game"
	InlineQueryResultTypeDocument InlineQueryResultType = "document"
	InlineQueryResultTypeGIF      InlineQueryResultType = "gif"
	InlineQueryResultTypeLocation InlineQueryResultType = "location"
	InlineQueryResultTypeMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultTypePhoto    InlineQueryResultType = "photo"
	InlineQueryResultTypeVenue    InlineQueryResultType = "venue"
	InlineQueryResultTypeSticker  InlineQueryResultType = "sticker"
	InlineQueryResultTypeVideo    InlineQueryResultType = "video"
	InlineQueryResultTypeVoice    InlineQueryResultType = "voice"
)

type InlineQueryResult interface {
	GetType() InlineQueryResultType
	GetID() string
	Validate() error
}

var _ InlineQueryResult = (*InlineQueryResultBase)(nil)

type InlineQueryResultBase struct {
	Type InlineQueryResultType `json:"type"`

	// Unique identifier for this result, 1-64 bytes
	ID string `json:"id"`

	// Title for the result. Few results do not support title.
	Title string `json:"title,omitempty"`

	// Optional. Inline keyboard attached to the message.
	// https://core.telegram.org/bots/features#inline-keyboards
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (r InlineQueryResultBase) Validate() error {
	return errors.New("should be implemented in child struct")
}

func (r InlineQueryResultBase) validate(expectedType InlineQueryResultType) error {
	if r.Type != expectedType {
		return fmt.Errorf("expected to have type=%s, got: %s", expectedType, r.Type)
	}
	if r.ID == "" {
		return fmt.Errorf("ID is empty")
	} else if l := len(r.ID); l > 64 {
		return fmt.Errorf("ID length is %d, should be less than 64 bytes", l)
	}
	if strings.TrimSpace(r.Title) != r.Title {
		return fmt.Errorf("title has leading or trailing white spaces")
	}
	if r.ReplyMarkup != nil {
		if err := r.ReplyMarkup.Validate(); err != nil {
			return fmt.Errorf("invalid InlineQueryResultBase.ReplyMarkup: %w", err)
		}
	}
	return nil
}

func (r InlineQueryResultBase) GetType() InlineQueryResultType {
	return r.Type
}

func (r InlineQueryResultBase) GetID() string {
	return r.ID
}

func (r InlineQueryResultBase) GetTitle() string {
	return r.Title
}

// InlineQueryResultArticle is an inline query response article.
type InlineQueryResultArticle struct {
	InlineQueryResultBase
	InputMessageContent interface{} `json:"input_message_content"` // required
	URL                 string      `json:"url,omitempty"`
	HideURL             bool        `json:"hide_url,omitempty"`
	Description         string      `json:"description,omitempty"`
	ThumbURL            string      `json:"thumb_url,omitempty"`
	ThumbWidth          int         `json:"thumb_width,omitempty"`
	ThumbHeight         int         `json:"thumb_height,omitempty"`
}

func (r InlineQueryResultArticle) Validate() error {
	return r.validate(InlineQueryResultTypeArticle)
}

// InlineQueryResultPhoto is an inline query response photo.
type InlineQueryResultPhoto struct {
	InlineQueryResultBase
	URL                 string      `json:"photo_url"` // required
	MimeType            string      `json:"mime_type,omitempty"`
	Width               int         `json:"photo_width,omitempty"`
	Height              int         `json:"photo_height,omitempty"`
	ThumbURL            string      `json:"thumb_url,omitempty"`
	Description         string      `json:"description,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultPhoto) Validate() error {
	return r.validate(InlineQueryResultTypePhoto)
}

// InlineQueryResultGIF is an inline query response GIF.
type InlineQueryResultGIF struct {
	InlineQueryResultBase
	URL                 string      `json:"gif_url"` // required
	Width               int         `json:"gif_width,omitempty"`
	Height              int         `json:"gif_height,omitempty"`
	ThumbURL            string      `json:"thumb_url,omitempty"`
	Caption             string      `json:"caption,omitempty"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultGIF) Validate() error {
	return r.validate(InlineQueryResultTypeGIF)
}

// InlineQueryResultMPEG4GIF is an inline query response MPEG4 GIF.
type InlineQueryResultMPEG4GIF struct {
	InlineQueryResultBase
	URL                 string      `json:"mpeg4_url"` // required
	Width               int         `json:"mpeg4_width"`
	Height              int         `json:"mpeg4_height"`
	ThumbURL            string      `json:"thumb_url"`
	Caption             string      `json:"caption"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultMPEG4GIF) Validate() error {
	return r.validate(InlineQueryResultTypeMpeg4Gif)
}

// InlineQueryResultVideo is an inline query response video.
type InlineQueryResultVideo struct {
	InlineQueryResultBase
	URL                 string      `json:"video_url"` // required
	MimeType            string      `json:"mime_type"` // required
	ThumbURL            string      `json:"thumb_url"`
	Caption             string      `json:"caption"`
	Width               int         `json:"video_width"`
	Height              int         `json:"video_height"`
	Duration            int         `json:"video_duration"`
	Description         string      `json:"description"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultVideo) Validate() error {
	return r.validate(InlineQueryResultTypeVideo)
}

// InlineQueryResultAudio is an inline query response audio.
type InlineQueryResultAudio struct {
	InlineQueryResultBase
	URL                 string      `json:"audio_url"` // required
	Performer           string      `json:"performer"`
	Duration            int         `json:"audio_duration"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultAudio) Validate() error {
	return r.validate(InlineQueryResultTypeAudio)
}

// InlineQueryResultVoice is an inline query response voice.
type InlineQueryResultVoice struct {
	InlineQueryResultBase
	URL                 string      `json:"voice_url"` // required
	Duration            int         `json:"voice_duration"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultVoice) Validate() error {
	return r.validate(InlineQueryResultTypeVoice)
}

// InlineQueryResultDocument is an inline query response document.
type InlineQueryResultDocument struct {
	InlineQueryResultBase
	Caption             string      `json:"caption"`
	URL                 string      `json:"document_url"` // required
	MimeType            string      `json:"mime_type"`    // required
	Description         string      `json:"description"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	ThumbURL            string      `json:"thumb_url"`
	ThumbWidth          int         `json:"thumb_width"`
	ThumbHeight         int         `json:"thumb_height"`
}

func (r InlineQueryResultDocument) Validate() error {
	return r.validate(InlineQueryResultTypeDocument)
}

// InlineQueryResultLocation is an inline query response location.
type InlineQueryResultLocation struct {
	InlineQueryResultBase
	Latitude            float64     `json:"latitude"`  // required
	Longitude           float64     `json:"longitude"` // required
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	ThumbURL            string      `json:"thumb_url,omitempty"`
	ThumbWidth          int         `json:"thumb_width,omitempty"`
	ThumbHeight         int         `json:"thumb_height,omitempty"`
}

func (r InlineQueryResultLocation) Validate() error {
	return r.validate(InlineQueryResultTypeLocation)
}

type FoursquareFields struct {

	// Optional. Foursquare identifier of the venue if known
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Optional. Foursquare type of the venue, if known.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
}

type GooglePlaceFields struct {
	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// Optional. Google Places type of the venue.
	// https://developers.google.com/maps/documentation/places/web-service/supported_types
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

type InlineQueryResultVenue struct {
	InlineQueryResultBase
	Latitude  float64 `json:"latitude"`  // required
	Longitude float64 `json:"longitude"` // required
	Address   string  `json:"address"`   // required
	FoursquareFields
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	ThumbURL            string      `json:"thumb_url,omitempty"`
	ThumbWidth          int         `json:"thumb_width,omitempty"`
	ThumbHeight         int         `json:"thumb_height,omitempty"`
}

func (r InlineQueryResultVenue) Validate() error {
	return r.validate(InlineQueryResultTypeVenue)
}

type InlineQueryResultCachedSticker struct {
	InlineQueryResultBase

	// A valid file identifier of the sticker
	StickerFileID       string              `json:"sticker_file_id"` // required
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

func (r InlineQueryResultCachedSticker) Validate() error {
	if r.Title != "" {
		return errors.New("InlineQueryResultCachedSticker.Title should be empty")
	}
	if r.InputMessageContent != nil {
		if err := r.InputMessageContent.Validate(); err != nil {
			return fmt.Errorf("invalid InlineQueryResultCachedSticker.InputMessageContent: %w", err)
		}
	}
	return r.validate(InlineQueryResultTypeSticker)
}

type InlineQueryResultContact struct {
	InlineQueryResultBase
	PhoneNumber         string              `json:"phone_number"`
	FirstName           string              `json:"first_name"`
	LastName            string              `json:"last_name,omitempty"`
	Vcard               string              `json:"vcard,omitempty"`
	InputMessageContent InputMessageContent `json:"input_message_content,omitempty"`
}

type InlineQueryResultGame struct {
	InlineQueryResultBase
	// Short name of the game
	GameShortName string `json:"game_short_name"`
}

func (v InlineQueryResultGame) Validate() error {
	if v.Title != "" {
		return errors.New("InlineQueryResultGame.Title should be empty")
	}
	if v.ReplyMarkup != nil {
		if err := v.ReplyMarkup.Validate(); err != nil {
			return fmt.Errorf("invalid InlineQueryResultGame.ReplyMarkup: %w", err)
		}
	}
	return v.validate(InlineQueryResultTypeGame)
}

func (r InlineQueryResultContact) Validate() error {
	return r.validate(InlineQueryResultTypeContact)
}

// NewInlineQueryResultArticle creates a new inline query article.
func NewInlineQueryResultArticle(id, title, messageText string) InlineQueryResultArticle {
	return InlineQueryResultArticle{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypeArticle,
			ID:    id,
			Title: title,
		},
		InputMessageContent: InputTextMessageContent{
			MessageText: messageText,
		},
	}
}

// NewInlineQueryResultGIF creates a new inline query GIF.
func NewInlineQueryResultGIF(id, url, title string) InlineQueryResultGIF {
	return InlineQueryResultGIF{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypeGIF,
			ID:    id,
			Title: title,
		},
		URL: url,
	}
}

// NewInlineQueryResultMPEG4GIF creates a new inline query MPEG4 GIF.
func NewInlineQueryResultMPEG4GIF(id, url, title string) *InlineQueryResultMPEG4GIF {
	return &InlineQueryResultMPEG4GIF{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypeMpeg4Gif,
			ID:    id,
			Title: title,
		},
		URL: url,
	}
}

// NewInlineQueryResultPhoto creates a new inline query photo.
func NewInlineQueryResultPhoto(id, url, title string) *InlineQueryResultPhoto {
	return &InlineQueryResultPhoto{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypePhoto,
			ID:    id,
			Title: title,
		},
		URL: url,
	}
}

// NewInlineQueryResultVideo creates a new inline query video.
func NewInlineQueryResultVideo(id, url, title string) *InlineQueryResultVideo {
	return &InlineQueryResultVideo{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypeVideo,
			ID:    id,
			Title: title,
		},
		URL: url,
	}
}

// NewInlineQueryResultAudio creates a new inline query audio.
func NewInlineQueryResultAudio(id, url, title string) *InlineQueryResultAudio {
	return &InlineQueryResultAudio{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypeAudio,
			ID:    id,
			Title: title,
		},
		URL: url,
	}
}

// NewInlineQueryResultVoice creates a new inline query voice.
func NewInlineQueryResultVoice(id, url, title string) *InlineQueryResultVoice {
	return &InlineQueryResultVoice{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypeVoice,
			ID:    id,
			Title: title,
		},
		URL: url,
	}
}

// NewInlineQueryResultDocument creates a new inline query document.
func NewInlineQueryResultDocument(id, url, title, mimeType string) *InlineQueryResultDocument {
	return &InlineQueryResultDocument{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypeDocument,
			ID:    id,
			Title: title,
		},
		URL:      url,
		MimeType: mimeType,
	}
}

// NewInlineQueryResultLocation creates a new inline query location.
func NewInlineQueryResultLocation(id, title string, latitude, longitude float64) *InlineQueryResultLocation {
	return &InlineQueryResultLocation{
		InlineQueryResultBase: InlineQueryResultBase{
			Type:  InlineQueryResultTypeLocation,
			ID:    id,
			Title: title,
		},
		Latitude:  latitude,
		Longitude: longitude,
	}
}
