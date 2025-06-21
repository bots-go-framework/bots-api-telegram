package tgbotapi

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

var _ Sendable = AnswerPreCheckoutQueryConfig{}

type AnswerPreCheckoutQueryConfig struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`

	// Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	OK bool `json:"ok"`

	// ErrorMessage is optional. Required if ok is False. Error message in human readable form
	// that explains the reason for failure to proceed with the checkout
	// (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your Payment details. Please choose a different color or garment!").
	// Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

func (c AnswerPreCheckoutQueryConfig) Values() (values url.Values, err error) {
	if c.OK && c.ErrorMessage != "" {
		err = fmt.Errorf("has OK=true and error message: %s", c.ErrorMessage)
		return
	}
	if !c.OK && c.ErrorMessage == "" {
		err = errors.New("has OK=false and no error message")
		return
	}
	values = make(url.Values)
	values.Set("pre_checkout_query_id", c.PreCheckoutQueryID)
	values.Set("ok", strconv.FormatBool(c.OK))
	if c.ErrorMessage != "" {
		values.Set("error_message", c.ErrorMessage)
	}
	return
}

func (AnswerPreCheckoutQueryConfig) TelegramMethod() string {
	return "answerPreCheckoutQuery"
}

func AnswerPreCheckoutQueryWithOK(preCheckoutQueryID string) AnswerPreCheckoutQueryConfig {
	return AnswerPreCheckoutQueryConfig{
		PreCheckoutQueryID: preCheckoutQueryID,
		OK:                 true,
	}
}

func AnswerPreCheckoutQueryWithNotOK(preCheckoutQueryID, errorMessage string) AnswerPreCheckoutQueryConfig {
	return AnswerPreCheckoutQueryConfig{
		PreCheckoutQueryID: preCheckoutQueryID,
		OK:                 false,
		ErrorMessage:       errorMessage,
	}
}
