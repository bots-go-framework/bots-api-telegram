package tgbotapi

import (
	"strings"
	"testing"
)

func TestCreateInvoiceLinkConfig_TelegramMethod(t *testing.T) {
	if (&CreateInvoiceLinkConfig{}).TelegramMethod() != "createInvoiceLink" {
		t.Fatal("wrong telegram method")
	}
}

func TestCreateInvoiceLinkConfig_Values_stars(t *testing.T) {
	cfg := &CreateInvoiceLinkConfig{
		Title:       "Unlock",
		Description: "Sourcer results",
		Payload:     "pch_1",
		Currency:    "XTR",
		Prices:      []LabeledPrice{{Label: "Unlock", Amount: 50}},
	}
	v, err := cfg.Values()
	if err != nil {
		t.Fatalf("Values: %v", err)
	}
	// No chat_id must be emitted — createInvoiceLink is chatless.
	if v.Has("chat_id") {
		t.Fatalf("createInvoiceLink must not carry chat_id, got %v", v)
	}
	if v.Get("title") != "Unlock" || v.Get("description") != "Sourcer results" || v.Get("payload") != "pch_1" {
		t.Fatalf("missing required fields: %v", v)
	}
	if v.Get("currency") != "XTR" {
		t.Fatalf("wrong currency: %v", v)
	}
	// provider_token MUST be present as an empty string for Stars.
	if !v.Has("provider_token") || v.Get("provider_token") != "" {
		t.Fatalf("provider_token must be present and empty for Stars, got %q present=%v", v.Get("provider_token"), v.Has("provider_token"))
	}
	if !strings.Contains(v.Get("prices"), `"amount":50`) {
		t.Fatalf("prices missing amount: %q", v.Get("prices"))
	}
}

func TestCreateInvoiceLinkConfig_Values_optionalFields(t *testing.T) {
	cfg := &CreateInvoiceLinkConfig{
		Title:                     "T",
		Description:               "D",
		Payload:                   "P",
		Currency:                  "EUR",
		ProviderToken:             "prov_123",
		Prices:                    []LabeledPrice{{Label: "x", Amount: 100}},
		BusinessConnectionID:      "biz_1",
		SubscriptionPeriod:        2592000,
		MaxTipAmount:              500,
		SuggestedTipAmounts:       []int64{100, 200},
		ProviderData:              `{"k":"v"}`,
		PhotoURL:                  "https://img",
		PhotoSize:                 1,
		PhotoWidth:                2,
		PhotoHeight:               3,
		NeedName:                  true,
		NeedPhoneNumber:           true,
		NeedEmail:                 true,
		NeedShippingAddress:       true,
		SendPhoneNumberToProvider: true,
		SendEmailToProvider:       true,
		IsFlexible:                true,
	}
	v, err := cfg.Values()
	if err != nil {
		t.Fatalf("Values: %v", err)
	}
	checks := map[string]string{
		"provider_token":                "prov_123",
		"business_connection_id":        "biz_1",
		"subscription_period":           "2592000",
		"max_tip_amount":                "500",
		"provider_data":                 `{"k":"v"}`,
		"photo_url":                     "https://img",
		"photo_size":                    "1",
		"photo_width":                   "2",
		"photo_height":                  "3",
		"need_name":                     "true",
		"need_phone_number":             "true",
		"need_email":                    "true",
		"need_shipping_address":         "true",
		"send_phone_number_to_provider": "true",
		"send_email_to_provider":        "true",
		"is_flexible":                   "true",
	}
	for k, want := range checks {
		if v.Get(k) != want {
			t.Errorf("%s = %q, want %q", k, v.Get(k), want)
		}
	}
	if !strings.Contains(v.Get("suggested_tip_amounts"), "100") || !strings.Contains(v.Get("suggested_tip_amounts"), "200") {
		t.Errorf("suggested_tip_amounts = %q", v.Get("suggested_tip_amounts"))
	}
}

func TestRefundStarPaymentConfig(t *testing.T) {
	cfg := &RefundStarPaymentConfig{UserID: 42, TelegramPaymentChargeID: "tgcharge_9"}
	if cfg.TelegramMethod() != "refundStarPayment" {
		t.Fatal("wrong telegram method")
	}
	v, err := cfg.Values()
	if err != nil {
		t.Fatalf("Values: %v", err)
	}
	if v.Get("user_id") != "42" || v.Get("telegram_payment_charge_id") != "tgcharge_9" {
		t.Fatalf("unexpected values: %v", v)
	}
}
