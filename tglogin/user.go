package tglogin

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
)

// LoginUser is a user data received from Telegram
type LoginUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"username,omitempty"`
	PhotoUrl  string `json:"photo_url,omitempty"`
	AuthDate  int    `json:"auth_date"`
	Hash      string `json:"hash"`
}

// IsFromTelegram checks if the user data are signed with given Telegram bot token
func (v LoginUser) IsFromTelegram(botToken string) bool {
	data := v.checkString()
	expectedHash := computeLoginHash(data, botToken)
	return v.Hash == expectedHash
}

func (v LoginUser) checkString() string {
	s := make([]string, 0, 6)
	s = append(s, "auth_date="+strconv.Itoa(v.AuthDate))
	s = append(s, "first_name="+v.FirstName)
	s = append(s, "id="+strconv.Itoa(v.ID))
	if v.LastName != "" {
		s = append(s, "last_name="+v.LastName)
	}
	if v.PhotoUrl != "" {
		s = append(s, "photo_url="+v.PhotoUrl)
	}
	if v.Username != "" {
		s = append(s, "username="+v.Username)
	}
	return strings.Join(s, "\n")
}

// computeLoginHash computes expected Telegram login hash
func computeLoginHash(data, token string) string {
	secretKey := sha256.Sum256([]byte(token))
	hash := hmac.New(sha256.New, secretKey[:]).Sum([]byte(data))
	return hex.EncodeToString(hash)
}
