package tgbotapi

import "github.com/bots-go-framework/bots-go-core/botkb"

type keyboardType string

type Keyboard interface {
	botkb.Keyboard
	telegramKeyboardType() keyboardType
}
