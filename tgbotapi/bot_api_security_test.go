package tgbotapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestMakeRequest_TransportErrorIsRedactedAndUnwraps(t *testing.T) {
	const (
		token       = "123456:PRIVATE-BOT-TOKEN"
		privateText = "PRIVATE-MESSAGE-TEXT"
	)
	transportErr := errors.New("transport failed against " + token)
	bot := NewBotAPIWithClient(token, &http.Client{
		Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return nil, transportErr
		}),
	})

	_, err := bot.MakeRequest("sendMessage", url.Values{"text": {privateText}})
	if err == nil {
		t.Fatal("MakeRequest() error = nil, want transport error")
	}
	if !errors.Is(err, transportErr) {
		t.Fatalf("MakeRequest() error does not unwrap transport error: %v", err)
	}
	assertDoesNotContainPrivateValues(t, err.Error(), token, privateText)
}

func TestMakeRequest_UnauthorizedErrorDoesNotExposeRequestOrResponse(t *testing.T) {
	const (
		token        = "123456:PRIVATE-BOT-TOKEN"
		privateText  = "PRIVATE-MESSAGE-TEXT"
		privateReply = "PRIVATE-TELEGRAM-RESPONSE"
	)
	bot := testBotWithResponse(token, http.StatusUnauthorized, privateReply)

	_, err := bot.MakeRequest("sendMessage", url.Values{"text": {privateText}})
	if err == nil {
		t.Fatal("MakeRequest() error = nil, want unauthorized error")
	}
	assertDoesNotContainPrivateValues(t, err.Error(), token, privateText, privateReply)
	if !strings.Contains(err.Error(), http.StatusText(http.StatusUnauthorized)) {
		t.Fatalf("MakeRequest() error = %q, want HTTP status", err)
	}
}

func TestMakeRequest_InvalidJSONErrorDoesNotExposeResponseBody(t *testing.T) {
	const (
		token        = "123456:PRIVATE-BOT-TOKEN"
		privateReply = "PRIVATE-RESPONSE-CANARY"
	)
	bot := testBotWithResponse(token, http.StatusOK, privateReply+"{")

	_, err := bot.MakeRequest("getMe", nil)
	if err == nil {
		t.Fatal("MakeRequest() error = nil, want invalid JSON error")
	}
	assertDoesNotContainPrivateValues(t, err.Error(), token, privateReply)
}

func TestMakeRequest_ProviderErrorDoesNotExposeDescription(t *testing.T) {
	const (
		token              = "123456:PRIVATE-BOT-TOKEN"
		privateDescription = "PRIVATE-PROVIDER-DESCRIPTION"
	)
	body := `{"ok":false,"error_code":400,"description":"` + privateDescription + `"}`
	bot := testBotWithResponse(token, http.StatusBadRequest, body)

	response, err := bot.MakeRequest("sendMessage", nil)
	if err == nil {
		t.Fatal("MakeRequest() error = nil, want provider error")
	}
	assertDoesNotContainPrivateValues(t, err.Error(), token, privateDescription)
	if response.Description != privateDescription {
		t.Fatalf("APIResponse.Description = %q, want structured provider description", response.Description)
	}
	var providerResponse APIResponse
	if !errors.As(err, &providerResponse) {
		t.Fatalf("MakeRequest() error does not preserve APIResponse: %v", err)
	}
	if providerResponse.ErrorCode != http.StatusBadRequest {
		t.Fatalf("provider error code = %d, want %d", providerResponse.ErrorCode, http.StatusBadRequest)
	}
	details, ok := TelegramProviderErrorDetailsFrom(fmt.Errorf("wrapped: %w", err))
	if !ok {
		t.Fatal("TelegramProviderErrorDetailsFrom() = false, want true")
	}
	if got, want := details, (TelegramProviderErrorDetails{
		Method:      "sendMessage",
		ErrorCode:   http.StatusBadRequest,
		Description: privateDescription,
	}); got != want {
		t.Fatalf("TelegramProviderErrorDetailsFrom() = %#v, want %#v", got, want)
	}
}

func TestUploadFile_ProviderErrorDoesNotExposeDescription(t *testing.T) {
	const (
		token              = "123456:PRIVATE-BOT-TOKEN"
		privateDescription = "PRIVATE-UPLOAD-DESCRIPTION"
	)
	body := `{"ok":false,"error_code":400,"description":"` + privateDescription + `"}`
	bot := testBotWithResponse(token, http.StatusBadRequest, body)

	response, err := bot.UploadFile(
		"sendPhoto",
		map[string]string{"chat_id": "123"},
		"photo",
		FileBytes{Name: "photo.jpg", Bytes: []byte("image")},
	)
	if err == nil {
		t.Fatal("UploadFile() error = nil, want provider error")
	}
	assertDoesNotContainPrivateValues(t, err.Error(), token, privateDescription)
	if response.Description != privateDescription {
		t.Fatalf("APIResponse.Description = %q, want structured provider description", response.Description)
	}
	details, ok := TelegramProviderErrorDetailsFrom(err)
	if !ok {
		t.Fatal("TelegramProviderErrorDetailsFrom() = false, want true")
	}
	if got, want := details, (TelegramProviderErrorDetails{
		Method:      "sendPhoto",
		ErrorCode:   http.StatusBadRequest,
		Description: privateDescription,
	}); got != want {
		t.Fatalf("TelegramProviderErrorDetailsFrom() = %#v, want %#v", got, want)
	}
}

func TestTelegramProviderErrorDetailsFrom_NonProviderError(t *testing.T) {
	if details, ok := TelegramProviderErrorDetailsFrom(errors.New("not a provider error")); ok || details != (TelegramProviderErrorDetails{}) {
		t.Fatalf("TelegramProviderErrorDetailsFrom() = (%#v, %t), want zero details and false", details, ok)
	}
}

func TestMakeRequest_ReadErrorIsRedactedAndUnwraps(t *testing.T) {
	const (
		token      = "123456:PRIVATE-BOT-TOKEN"
		readCanary = "PRIVATE-READ-ERROR"
	)
	readErr := errors.New(readCanary)
	body := &failingReadCloser{readErr: readErr}
	bot := testBotWithBody(token, http.StatusOK, 1, body)

	_, err := bot.MakeRequest("getMe", nil)
	if err == nil {
		t.Fatal("MakeRequest() error = nil, want read error")
	}
	if !errors.Is(err, readErr) {
		t.Fatalf("MakeRequest() error does not unwrap read error: %v", err)
	}
	assertDoesNotContainPrivateValues(t, err.Error(), token, readCanary)
	if !body.closed {
		t.Fatal("MakeRequest() did not close response body")
	}
}

func TestUploadFile_ReadErrorIsRedactedAndUnwraps(t *testing.T) {
	const readCanary = "PRIVATE-UPLOAD-READ-ERROR"
	readErr := errors.New(readCanary)
	body := &failingReadCloser{readErr: readErr}
	bot := testBotWithBody("123456:PRIVATE-BOT-TOKEN", http.StatusOK, 1, body)

	_, err := bot.UploadFile(
		"sendPhoto",
		nil,
		"photo",
		FileBytes{Name: "photo.jpg", Bytes: []byte("image")},
	)
	if err == nil {
		t.Fatal("UploadFile() error = nil, want read error")
	}
	if !errors.Is(err, readErr) {
		t.Fatalf("UploadFile() error does not unwrap read error: %v", err)
	}
	assertDoesNotContainPrivateValues(t, err.Error(), readCanary)
}

func TestMakeRequest_CloseErrorDoesNotOverrideSuccessfulResult(t *testing.T) {
	closeErr := errors.New("PRIVATE-CLOSE-ERROR")
	body := &failingReadCloser{
		reader:   strings.NewReader(`{"ok":true,"result":true}`),
		closeErr: closeErr,
	}
	bot := testBotWithBody("123456:PRIVATE-BOT-TOKEN", http.StatusOK, int64(body.reader.Len()), body)

	response, err := bot.MakeRequest("answerCallbackQuery", nil)
	if err != nil {
		t.Fatalf("MakeRequest() error = %v, want nil", err)
	}
	if !response.Ok {
		t.Fatalf("MakeRequest() response.Ok = false, want true")
	}
	if !body.closed {
		t.Fatal("MakeRequest() did not close response body")
	}
}

func testBotWithResponse(token string, status int, body string) *BotAPI {
	return testBotWithBody(
		token,
		status,
		int64(len(body)),
		io.NopCloser(strings.NewReader(body)),
	)
}

func testBotWithBody(token string, status int, contentLength int64, body io.ReadCloser) *BotAPI {
	return NewBotAPIWithClient(token, &http.Client{
		Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode:    status,
				Body:          body,
				ContentLength: contentLength,
				Header:        make(http.Header),
			}, nil
		}),
	})
}

type failingReadCloser struct {
	reader   *strings.Reader
	readErr  error
	closeErr error
	closed   bool
}

func (r *failingReadCloser) Read(p []byte) (int, error) {
	if r.reader != nil {
		return r.reader.Read(p)
	}
	return 0, r.readErr
}

func (r *failingReadCloser) Close() error {
	r.closed = true
	return r.closeErr
}

func assertDoesNotContainPrivateValues(t *testing.T, text string, privateValues ...string) {
	t.Helper()
	for _, privateValue := range privateValues {
		if strings.Contains(text, privateValue) {
			t.Fatalf("text exposes private value %q: %s", privateValue, fmt.Sprintf("%q", text))
		}
	}
}
