package tgbotapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestBaseChatCurrentFieldsValues(t *testing.T) {
	values, err := (BaseChat{
		ChatID:                -1001,
		MessageThreadID:       2,
		DirectMessagesTopicID: 3,
		BusinessConnectionID:  "business",
		AllowPaidBroadcast:    true,
		SuggestedPostParameters: &SuggestedPostParameters{
			Price: &SuggestedPostPrice{Currency: "XTR", Amount: 25},
		},
	}).Values()
	if err != nil {
		t.Fatal(err)
	}
	if got := values.Get("direct_messages_topic_id"); got != "3" {
		t.Errorf("direct_messages_topic_id = %q", got)
	}
	if got := values.Get("business_connection_id"); got != "business" {
		t.Errorf("business_connection_id = %q", got)
	}
	if got := values["allow_paid_broadcast"]; len(got) != 1 || got[0] != "true" {
		t.Errorf("allow_paid_broadcast = %#v", got)
	}
	var suggested SuggestedPostParameters
	if err = json.Unmarshal([]byte(values.Get("suggested_post_parameters")), &suggested); err != nil {
		t.Fatal(err)
	}
	if suggested.Price == nil || suggested.Price.Amount != 25 {
		t.Errorf("suggested post = %#v", suggested)
	}
}

func TestRichMessageRecursiveValidation(t *testing.T) {
	cell := RichText{PlainText: "Score"}
	valid := InputRichMessage{Blocks: []InputRichBlock{{
		Type:       RichBlockTypeTable,
		IsBordered: true,
		Cells:      [][]RichBlockTableCell{{{Text: &cell, IsHeader: true}}},
	}}}
	if err := valid.Validate(); err != nil {
		t.Fatalf("valid table: %v", err)
	}
	if _, err := (RichMessageConfig{
		BaseChat:    BaseChat{ChatID: 1},
		RichMessage: valid,
	}).Values(); err != nil {
		t.Fatalf("valid rich config: %v", err)
	}

	thinking := InputRichMessage{Blocks: []InputRichBlock{{
		Type: RichBlockTypeThinking,
		Text: &RichText{PlainText: "Thinking…"},
	}}}
	if err := thinking.Validate(); err == nil {
		t.Fatal("persistent thinking block should be rejected")
	}
	if err := thinking.ValidateDraft(); err != nil {
		t.Fatalf("draft thinking block: %v", err)
	}

	invalidMedia := InputRichMessage{
		HTML: "<img src=\"tg://photo?id=bad id\"/>",
		Media: []InputRichMessageMedia{{
			ID:    "bad id",
			Media: InputMediaPhoto{Type: "photo", Media: "file"},
		}},
	}
	if err := invalidMedia.Validate(); err == nil {
		t.Fatal("invalid rich media identifier should be rejected")
	}
}

func TestPollMediaTypedUnionAndLegacyLocation(t *testing.T) {
	typed := NewInputPollMedia(InputMediaVideo{
		Type:     "video",
		Media:    "video-file",
		Width:    640,
		Height:   360,
		Duration: 5,
	})
	data, err := json.Marshal(typed)
	if err != nil {
		t.Fatal(err)
	}
	var video map[string]any
	if err = json.Unmarshal(data, &video); err != nil {
		t.Fatal(err)
	}
	if video["width"] != float64(640) || video["duration"] != float64(5) {
		t.Errorf("typed video lost properties: %s", data)
	}

	legacy := InputPollOptionMedia{
		Type: "location",
		Location: &Location{
			Latitude:           53.3,
			Longitude:          -6.2,
			HorizontalAccuracy: 4.5,
		},
	}
	data, err = json.Marshal(legacy)
	if err != nil {
		t.Fatal(err)
	}
	var location map[string]any
	if err = json.Unmarshal(data, &location); err != nil {
		t.Fatal(err)
	}
	if _, nested := location["location"]; nested {
		t.Fatalf("location must use flattened InputMediaLocation schema: %s", data)
	}
	if location["latitude"] != 53.3 || location["horizontal_accuracy"] != 4.5 {
		t.Errorf("location = %s", data)
	}
}

func TestEphemeralConstructorsAndValidation(t *testing.T) {
	textEdit := NewEditEphemeralMessageText(-1001, 42, 7, "Your private hand")
	values, err := textEdit.Values()
	if err != nil {
		t.Fatal(err)
	}
	if values.Get("receiver_user_id") != "42" || values.Get("ephemeral_message_id") != "7" {
		t.Errorf("ephemeral address = %v", values)
	}
	if _, err = NewDeleteEphemeralMessage(0, 42, 7).Values(); err == nil {
		t.Fatal("missing chat_id should be rejected")
	}
	if _, err = NewEditEphemeralMessageText(-1001, 42, 7, "").Values(); err == nil {
		t.Fatal("empty edit text should be rejected")
	}
}

func TestMessageDraftAllowsEmptyText(t *testing.T) {
	values, err := (MessageDraftConfig{ChatID: 1, DraftID: 9}).Values()
	if err != nil {
		t.Fatal(err)
	}
	if _, exists := values["text"]; !exists {
		t.Fatal("empty text must still be sent for Telegram's Thinking placeholder")
	}
}

type parityRoundTripFunc func(*http.Request) (*http.Response, error)

func (f parityRoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return f(request)
}

func parityBot(t *testing.T, result string, inspect func(string, url.Values)) *BotAPI {
	t.Helper()
	envelope := `{"ok":true,"result":` + result + `}`
	return NewBotAPIWithClient("1:test", &http.Client{
		Transport: parityRoundTripFunc(func(request *http.Request) (*http.Response, error) {
			if err := request.ParseForm(); err != nil {
				t.Fatal(err)
			}
			inspect(request.URL.Path, request.PostForm)
			return &http.Response{
				StatusCode:    http.StatusOK,
				ContentLength: int64(len(envelope)),
				Body:          io.NopCloser(strings.NewReader(envelope)),
				Header:        make(http.Header),
			}, nil
		}),
	})
}

func TestManagedAccessAndChatManagementWireParameters(t *testing.T) {
	t.Run("managed access", func(t *testing.T) {
		bot := parityBot(t, "true", func(path string, values url.Values) {
			if !strings.HasSuffix(path, "/setManagedBotAccessSettings") {
				t.Errorf("path = %q", path)
			}
			if values.Get("is_access_restricted") != "true" || strings.TrimSpace(values.Get("added_user_ids")) != "[7,8]" {
				t.Errorf("values = %v", values)
			}
		})
		if _, err := bot.SetManagedBotAccessSettings(5, true, []int64{7, 8}); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("return bots", func(t *testing.T) {
		bot := parityBot(t, "[]", func(path string, values url.Values) {
			if values.Get("return_bots") != "true" {
				t.Errorf("values = %v", values)
			}
		})
		if _, err := bot.GetChatAdministrators("-1001", true); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("personal messages limit", func(t *testing.T) {
		bot := parityBot(t, "[]", func(path string, values url.Values) {
			if values.Get("limit") != "12" {
				t.Errorf("values = %v", values)
			}
		})
		if _, err := bot.GetUserPersonalChatMessages(5, 12); err != nil {
			t.Fatal(err)
		}
	})
}
