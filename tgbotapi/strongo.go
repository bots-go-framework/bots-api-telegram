package tgbotapi

import (
	"fmt"
	"net/http"
)

// ReplyToResponse replies to response
func ReplyToResponse(chattable Sendable, w http.ResponseWriter) (string, error) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	values, err := chattable.Values()

	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("TelegramMethod=%v&%v", chattable.TelegramMethod(), values.Encode())

	_, err = w.Write([]byte(s))

	return s, err
}

// NewReplyKeyboardUsingStrings creates reply keyboard from strings arrays
func NewReplyKeyboardUsingStrings(buttons [][]string) *ReplyKeyboardMarkup {
	kb := make([][]KeyboardButton, len(buttons))

	for i, row := range buttons {
		kbRow := make([]KeyboardButton, len(row))
		for j, text := range row {
			kbRow[j] = KeyboardButton{Text: text}
		}
		kb[i] = kbRow
	}

	return &ReplyKeyboardMarkup{Keyboard: kb, ResizeKeyboard: true}
}
