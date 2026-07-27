package tgbotapi

import (
	"strings"
	"testing"
)

func TestInputRichMessageValidate_DateTimeFallbackUTF8ByteLimit(t *testing.T) {
	for _, test := range []struct {
		name     string
		fallback string
		wantErr  string
	}{
		{name: "31 ASCII bytes are accepted", fallback: strings.Repeat("a", 31)},
		{name: "32 ASCII bytes are rejected", fallback: strings.Repeat("a", 32), wantErr: "date_time fallback must be at most 31 UTF-8 bytes; got 32"},
		{name: "30 Cyrillic UTF-8 bytes are accepted", fallback: strings.Repeat("я", 15)},
		{name: "32 Cyrillic UTF-8 bytes are rejected", fallback: strings.Repeat("я", 16), wantErr: "date_time fallback must be at most 31 UTF-8 bytes; got 32"},
	} {
		t.Run(test.name, func(t *testing.T) {
			err := nestedDateTimeFallbackMessage(test.fallback).Validate()
			if test.wantErr == "" {
				if err != nil {
					t.Fatalf("Validate() error = %v", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), test.wantErr) {
				t.Fatalf("Validate() error = %v, want %q", err, test.wantErr)
			}
		})
	}
}

func nestedDateTimeFallbackMessage(fallback string) InputRichMessage {
	return InputRichMessage{Blocks: []InputRichBlock{{
		Type:    RichBlockTypeDetails,
		Summary: &RichText{PlainText: "Details"},
		Blocks: []InputRichBlock{{
			Type: RichBlockTypeParagraph,
			Text: &RichText{Type: RichTextTypeDateTime, UnixTime: 1, DateTimeFormat: "r", Text: &RichText{
				Type: RichTextTypeBold,
				Text: &RichText{Items: []RichText{{PlainText: fallback}}},
			}},
		}},
	}}}
}
