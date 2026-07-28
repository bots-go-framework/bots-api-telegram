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

// AuthenticateTelegramWebApp validates the signature of Telegram web app init
// data. Callers must reject stale AuthDate values according to their session
// policy.
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
	if botToken == "" {
		return false
	}
	suppliedHash, err := hex.DecodeString(values.Get("hash"))
	if err != nil || len(suppliedHash) != sha256.Size {
		return false
	}
	expectedHash := computeWebAppHash(getDataCheckString(values), botToken)
	return hmac.Equal(expectedHash, suppliedHash)
}

func computeWebAppHash(data, token string) []byte {
	secretKeyMAC := hmac.New(sha256.New, []byte("WebAppData"))
	_, _ = secretKeyMAC.Write([]byte(token))

	dataMAC := hmac.New(sha256.New, secretKeyMAC.Sum(nil))
	_, _ = dataMAC.Write([]byte(data))
	return dataMAC.Sum(nil)
}

func getDataCheckString(values url.Values) string {
	keys := make([]string, 0, len(values))
	for k := range values {
		if k == "hash" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var s strings.Builder
	for i, k := range keys {
		if i > 0 {
			s.WriteByte('\n')
		}
		fmt.Fprintf(&s, "%s=%s", k, values.Get(k))
	}
	return s.String()
}
