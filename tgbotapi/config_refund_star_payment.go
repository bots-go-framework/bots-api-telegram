package tgbotapi

import (
	"net/url"
	"strconv"
)

var _ Sendable = (*RefundStarPaymentConfig)(nil)

// RefundStarPaymentConfig refunds a successful payment in Telegram Stars.
// https://core.telegram.org/bots/api#refundstarpayment
//
// Refunding needs BOTH the payer's user id and the telegram_payment_charge_id
// from the successful_payment update (there is no retrieve-charge API to recover
// one from the other). Returns True on success.
type RefundStarPaymentConfig struct {
	UserID                  int64  `json:"user_id"`                    // Identifier of the user whose payment will be refunded.
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"` // Telegram payment identifier of the charge to refund.
}

func (*RefundStarPaymentConfig) TelegramMethod() string {
	return "refundStarPayment"
}

// Values returns the url.Values representation of RefundStarPaymentConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v *RefundStarPaymentConfig) Values() (url.Values, error) {
	values := url.Values{}
	values.Add("user_id", strconv.FormatInt(v.UserID, 10))
	values.Add("telegram_payment_charge_id", v.TelegramPaymentChargeID)
	return values, nil
}
