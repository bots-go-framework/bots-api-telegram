package tgbotapi

// PreCheckoutQuery contains information about an incoming pre-checkout query.
// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	ID               string     `json:"id"`                           // Unique query identifier
	From             *User      `json:"from"`                         // User who sent the query
	Currency         string     `json:"currency"`                     // Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
	TotalAmount      int        `json:"total_amount"`                 // Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload"`              // Bot-specified invoice payload
	ShippingOptionID string     `json:"shipping_option_id,omitempty"` // Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`         // Optional. Order information provided by the user
}
