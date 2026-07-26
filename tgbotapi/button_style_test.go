package tgbotapi

import (
	"strings"
	"testing"
)

func TestCurrentButtonStyles(t *testing.T) {
	for _, style := range []string{
		"",
		ButtonStylePrimary,
		ButtonStyleSuccess,
		ButtonStyleDanger,
	} {
		t.Run(style, func(t *testing.T) {
			inline := NewInlineKeyboardButtonData("Play", "play")
			inline.Style = style
			if err := inline.Validate(); err != nil {
				t.Fatalf("InlineKeyboardButton.Validate() error = %v", err)
			}

			keyboard := KeyboardButton{Text: "Play", Style: style}
			if err := keyboard.Validate(); err != nil {
				t.Fatalf("KeyboardButton.Validate() error = %v", err)
			}
		})
	}
}

func TestRejectsObsoleteButtonStyles(t *testing.T) {
	for _, style := range []string{"default", "positive", "destructive"} {
		t.Run(style, func(t *testing.T) {
			inline := NewInlineKeyboardButtonData("Play", "play")
			inline.Style = style
			if err := inline.Validate(); err == nil || !strings.Contains(err.Error(), "invalid button style") {
				t.Fatalf("InlineKeyboardButton.Validate() error = %v, want invalid button style", err)
			}

			keyboard := KeyboardButton{Text: "Play", Style: style}
			if err := keyboard.Validate(); err == nil || !strings.Contains(err.Error(), "invalid button style") {
				t.Fatalf("KeyboardButton.Validate() error = %v, want invalid button style", err)
			}
		})
	}
}
