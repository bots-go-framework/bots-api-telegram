package tgbotapi

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestReplyToResponse_UsesTelegramMethodParameter(t *testing.T) {
	w := httptest.NewRecorder()
	_, err := ReplyToResponse(AnswerCallbackQueryConfig{
		CallbackQueryID: "callback-id",
		Text:            "Done",
	}, w)
	if err != nil {
		t.Fatalf("ReplyToResponse() error = %v", err)
	}

	values, err := url.ParseQuery(w.Body.String())
	if err != nil {
		t.Fatalf("ParseQuery() error = %v", err)
	}
	if got, want := values.Get("method"), "answerCallbackQuery"; got != want {
		t.Errorf("method = %q, want %q", got, want)
	}
	if got := values.Get("TelegramMethod"); got != "" {
		t.Errorf("unexpected legacy TelegramMethod = %q", got)
	}
}
