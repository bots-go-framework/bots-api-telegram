package tgbotapi

import (
	"encoding/json"
	"testing"
)

func TestSuccessfulPayment_UnmarshalJSON(t *testing.T) {

	jsonStr := `{
            "currency": "XTR",
            "total_amount": 2,
            "invoice_payload": "topped_up",
            "telegram_payment_charge_id": "some_charge_id",
            "provider_payment_charge_id": "1234567890_12",
			"is_recurring": true
        }`

	v := SuccessfulPayment{}
	if err := json.Unmarshal([]byte(jsonStr), &v); err != nil {
		t.Fatal(err)
	}
	if !v.IsRecurring {
		t.Errorf("IsRecurring expected to be true")
	}
	if v.Currency != "XTR" {
		t.Errorf("expected currency=XTR, got %s", v.Currency)
	}
	if v.TotalAmount != 2 {
		t.Errorf("expected TotalAmount=2, got %d", v.TotalAmount)
	}
	if v.InvoicePayload != "topped_up" {
		t.Errorf("expected InvoicePayload=topped_up, got %s", v.InvoicePayload)
	}
	if v.TelegramPaymentChargeID != "some_charge_id" {
		t.Errorf("expected TelegramPaymentChargeID=some_charge_id, got %s", v.InvoicePayload)
	}
	if v.ProviderPaymentChargeID != "1234567890_12" {
		t.Errorf("expected ProviderPaymentChargeID=1234567890_12, got %s", v.ProviderPaymentChargeID)
	}
}
