package tgbotapi

import (
	"encoding/json"
	"testing"
)

func TestSharedUsersUnmarshall(t *testing.T) {
	s := `{
	"update_id": 903586833,
	"message": {
		"message_id": 1498,
		"from": {
			"id": 1234,
			"is_bot": false,
			"first_name": "Jack",
			"last_name": "Black",
			"username": "jackblack",
			"language_code": "en"
		},
		"chat": {
			"id": 1234,
			"first_name": "Jack",
			"last_name": "Black",
			"username": "jackblack",
			"type": "private"
		},
		"date": 1737789766,
		"user_shared": {
			"user_id": 12345678,
			"request_id": 1
		},
		"users_shared": {
			"user_ids": [
				12345678
			],
			"users": [
				{
					"user_id": 12345678,
					"first_name": "John",
					"last_name": "Smith",
					"photo": [
						{
							"file_id": "AgACAgQAAxUAAWeT0iEiHTM78tA1i7xnuA3KrmzmAAKutjEbKPMYUfCASmGV5ahpAQADAgADYQADNgQ",
							"file_unique_id": "AQADrrYxGyjzGFEAAQ",
							"file_size": 9499,
							"width": 160,
							"height": 160
						},
						{
							"file_id": "AgACAgQAAxUAAWeT0iEiHTM78tA1i7xnuA3KrmzmAAKutjEbKPMYUfCASmGV5ahpAQADAgADYgADNgQ",
							"file_unique_id": "AQADrrYxGyjzGFFn",
							"file_size": 30776,
							"width": 320,
							"height": 320
						},
						{
							"file_id": "AgACAgQAAxUAAWeT0iEiHTM78tA1i7xnuA3KrmzmAAKutjEbKPMYUfCASmGV5ahpAQADAgADYwADNgQ",
							"file_unique_id": "AQADrrYxGyjzGFEB",
							"file_size": 106261,
							"width": 640,
							"height": 640
						}
					]
				}
			],
			"request_id": 1
		}
	}
}`
	var update Update
	if err := json.Unmarshal([]byte(s), &update); err != nil {
		t.Fatal(err)
	}
	if update.Message == nil {
		t.Fatal("expected Message to be non-nil")
	}
	if update.Message.MessageID != 1498 {
		t.Errorf("expected 1498, got %d", update.Message.MessageID)
	}
	if update.Message.UsersShared == nil {
		t.Fatal("expected UsersShared to be non-nil")
	}
	if update.Message.UsersShared.RequestID != 1 {
		t.Errorf("expected 1, got %d", update.Message.UsersShared.RequestID)
	}
	if len(update.Message.UsersShared.Users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(update.Message.UsersShared.Users))
	}
	if len(update.Message.UsersShared.Users[0].Photo) != 3 {
		t.Fatalf("expected 3 photos, got %d", len(update.Message.UsersShared.Users))
	}
}
