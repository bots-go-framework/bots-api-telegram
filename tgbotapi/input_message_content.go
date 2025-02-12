package tgbotapi

//go:generate ffjson $GOFILE

import (
	"errors"
	"fmt"
)

type inputMessageContentBase struct {
}

func (v inputMessageContentBase) base() inputMessageContentBase {
	return v
}

type InputMessageContent interface {
	Validate() error
	base() inputMessageContentBase
}

var _ InputMessageContent = (*InputTextMessageContent)(nil)

// InputTextMessageContent contains text for displaying
// as an inline query result.
type InputTextMessageContent struct {
	inputMessageContentBase
	MessageText           string          `json:"message_text"`
	ParseMode             string          `json:"parse_mode"`
	Entities              []MessageEntity `json:"entities"`
	DisableWebPagePreview bool            `json:"disable_web_page_preview"`
}

func (v InputTextMessageContent) Validate() error {
	if v.MessageText == "" {
		return errors.New("MessageText is empty")
	}
	switch v.ParseMode {
	case "", "MarkdownV2", "HTML":
	// OK
	default:
		return errors.New("ParseMode is not empty, MarkdownV2 or HTML")
	}
	for i, entity := range v.Entities {
		if err := entity.Validate(); err != nil {
			return fmt.Errorf("invalid entities[%d]: %w", i, err)
		}
	}
	return nil
}

// InputLocationMessageContent contains a location for displaying
// as an inline query result.
type InputLocationMessageContent struct {
	inputMessageContentBase
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (v InputLocationMessageContent) Validate() error {
	return nil
}

// InputVenueMessageContent contains a venue for displaying
// as an inline query result.
type InputVenueMessageContent struct {
	inputMessageContentBase
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareID string  `json:"foursquare_id"`
}

func (v InputVenueMessageContent) Validate() error {
	if v.Title == "" {
		return errors.New("title is empty")
	}
	if v.Address == "" {
		return errors.New("address is empty")
	}
	return nil
}

// InputContactMessageContent contains a contact for displaying
// as an inline query result.
type InputContactMessageContent struct {
	inputMessageContentBase

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name"`

	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	VCard string `json:"vcard"`
}

func (v InputContactMessageContent) Validate() error {
	if v.PhoneNumber == "" {
		return errors.New("phone_number is empty")
	}
	if v.FirstName == "" {
		return errors.New("first_name is empty")
	}
	return nil
}
