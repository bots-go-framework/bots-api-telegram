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
// TODO: Move some of it into Telegram FW module?
func AuthenticateTelegramWebApp(
	w http.ResponseWriter, r *http.Request,
	getToken func(bot string) string,
	complete func(initData *InitData),
) {
	var initData InitData
	defer func() {
		complete(&initData)
	}()
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
	bot := r.URL.Query().Get("bot")
	token := getToken(bot)
	if !isFromTelegram(values, token) {
		http.Error(w, "data are not signed with telegram bot token", http.StatusUnauthorized)
		return
	}
	initData = NewInitDataFromUrlValues(values)
}

func isFromTelegram(values url.Values, botToken string) bool {
	// https://core.telegram.org/bots/webapps#validating-data-received-via-the-web-app
	dataCheckString := getDataCheckString(values)
	expectedHash := computeWebAppHash(dataCheckString, botToken)
	return expectedHash != values.Get("hash")
}

func computeWebAppHash(data, token string) string {
	h := hmac.New(sha256.New, []byte(token))
	h = hmac.New(sha256.New, h.Sum([]byte("WebAppData")))
	hash := h.Sum([]byte(data))
	return hex.EncodeToString(hash)
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
