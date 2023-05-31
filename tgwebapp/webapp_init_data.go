package tgwebapp

import (
	"net/url"
	"strconv"
)

// InitData describes Telegram web app init data
// https://core.telegram.org/bots/webapps#webappinitdata
type InitData struct {
	// QueryID - A unique identifier for the Web App session, required for sending messages via the answerWebAppQuery method.
	QueryID string `json:"query_id"`

	// ChatType - Type of the chat from which the Web App was opened.
	// Can be either “sender” for a private chat with the user opening the link,
	// “private”, “group”, “supergroup”, or “channel”.
	//Returned only for Web Apps launched from direct links.
	ChatType string `json:"chat_type,omitempty"`

	// ChatInstance - Global identifier, uniquely corresponding to the chat from which the Web App was opened.
	// Returned only for Web Apps launched from a direct link.
	ChatInstance string `json:"chat_instance,omitempty"`

	// StartParam - The value of the startattach parameter, passed via link. Only returned for Web Apps when launched from the attachment menu via link.
	StartParam string `json:"start_param,omitempty"`

	// CanSendAfter - Time in seconds, after which a message can be sent via the answerWebAppQuery method.
	CanSendAfter int `json:"can_send_after,omitempty"`

	// AuthDate - Unix time when the form was opened.
	AuthDate int `json:"auth_date"`

	// Hash of all passed parameters, which the bot server can use to check their validity.
	Hash string `json:"hash"`
}

func NewInitDataFromUrlValues(values url.Values) InitData {
	initData := InitData{
		QueryID:      values.Get("query_id"),
		ChatType:     values.Get("chat_type"),
		ChatInstance: values.Get("chat_instance"),
		StartParam:   values.Get("start_param"),
		Hash:         values.Get("hash"),
	}
	initData.CanSendAfter, _ = strconv.Atoi(values.Get("can_send_after"))
	initData.AuthDate, _ = strconv.Atoi(values.Get("auth_date"))
	return initData
}
