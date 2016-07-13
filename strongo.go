package tgbotapi

import (
	"fmt"
	"net/http"
)

func ReplyToResponse(chattable Chattable, w http.ResponseWriter) (string, error) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

	values, err := chattable.values()

	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("method=%v&%v", chattable.method(), values.Encode())

	_, err = w.Write([]byte(s))

	return s, err
}

func NewReplyKeyboardUsingStrings(buttons [][]string) ReplyKeyboardMarkup {
	kb := make([][]KeyboardButton, len(buttons))

	for i, row := range buttons {
		kbRow := make([]KeyboardButton, len(row))
		for j, text := range row {
			kbRow[j] = KeyboardButton{Text: text}
		}
		kb[i] = kbRow
	}

	return ReplyKeyboardMarkup{Keyboard: kb, ResizeKeyboard: true}
}
