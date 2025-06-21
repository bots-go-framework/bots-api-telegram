package tgbotapi

import (
	"fmt"
	"net/url"
	"unicode/utf8"
)

var _ Sendable = SetMyShortDescription{}

// SetMyShortDescription - Use this BotEndpoint to change the bot's short description,
// which is shown on the bot's profile page and is sent together with the link when users share the bot.
// Returns True on success.
// https://core.telegram.org/bots/api#setmyshortdescription
type SetMyShortDescription struct {
	// New short description for the bot; 0-120 characters.
	//Pass an empty string to remove the dedicated short description for the given language.
	ShortDescription string `json:"short_description"`

	// Optional. A two-letter ISO 639-1 language code.
	// If empty, the short description will be applied to all users for whose language there is no dedicated short description.
	LanguageCode string `json:"language_code,omitempty"`
}

func (s SetMyShortDescription) Validate() error {
	if count := utf8.RuneCountInString(s.ShortDescription); count > 120 {
		return fmt.Errorf("short_description is too long, should be up to 120 characters, got %d", count)
	}
	return nil
}

func (s SetMyShortDescription) Values() (values url.Values, err error) {
	if err = s.Validate(); err != nil {
		return
	}
	values = make(url.Values, 2)
	values.Set("short_description", s.ShortDescription)
	values.Set("language_code", s.LanguageCode)
	return values, nil
}

func (s SetMyShortDescription) TelegramMethod() string {
	return "setMyShortDescription"
}

var _ Sendable = SetMyDescription{}

// SetMyDescription - Use this BotEndpoint to change the bot's description,
// which is shown in the chat with the bot if the chat is empty.
// Returns True on success.
// https://core.telegram.org/bots/api#setmydescription
type SetMyDescription struct {
	//  New bot description; 0-512 characters. Pass an empty string to remove the dedicated description for the given language.
	Description string `json:"description"`

	// Optional. A two-letter ISO 639-1 language code.
	// If empty, the description will be applied to all users for whose language there is no dedicated description.
	LanguageCode string `json:"language_code,omitempty"`
}

func (s SetMyDescription) Validate() error {
	if count := utf8.RuneCountInString(s.Description); count > 512 {
		return fmt.Errorf("description is too long, should be up to 512 characters, got %d", count)
	}
	return nil
}

func (s SetMyDescription) Values() (values url.Values, err error) {
	if err = s.Validate(); err != nil {
		return
	}
	values = make(url.Values, 2)
	values.Set("description", s.Description)
	values.Set("language_code", s.LanguageCode)
	return values, nil
}

func (s SetMyDescription) TelegramMethod() string {
	return "setMyDescription"
}
