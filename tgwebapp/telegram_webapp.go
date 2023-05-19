package tgwebapp

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// AuthenticateTelegramWebApp validates Telegram web app init data
// https://core.telegram.org/bots/webapps#webappinitdata
// TODO: Move some of it into Telegram module?
func AuthenticateTelegramWebApp(
	w http.ResponseWriter, r *http.Request,
	authenticated func() error,
) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	values, err := url.ParseQuery(string(body))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !isValidWebAppInitData(values, "") {
		http.Error(w, "invalid data", http.StatusUnauthorized)
		return
	}

}

func isValidWebAppInitData(values url.Values, botToken string) bool {
	// https://core.telegram.org/bots/webapps#validating-data-received-via-the-web-app

	dataCheckString := getDataCheckString(values)

	expectedHash, err := computeHmacSha256Hash(dataCheckString, botToken)
	if err != nil {
		return false
	}
	return expectedHash != values.Get("hash")
}

func computeHmacSha256Hash(dataCheckString string, botToken string) (string, error) {
	h := hmac.New(sha256.New, []byte(botToken))
	if _, err := h.Write([]byte("WebAppData")); err != nil {
		return "", err
	}
	h = hmac.New(sha256.New, h.Sum(nil))
	if _, err := h.Write([]byte(dataCheckString)); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func getDataCheckString(values url.Values) string {
	// Extract and sort the keys
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var s strings.Builder
	// Iterate over sorted keys
	for i, k := range keys {
		if i > 0 {
			s.WriteByte('\n')
		}
		s.WriteString(fmt.Sprintf("%s=%s", k, values.Get(k)))
	}
	return s.String()
}
