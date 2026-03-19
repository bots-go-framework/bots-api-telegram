package tgbotapi

// ShippingAddress represents a shipping address.
// Note: shipping_address.go already has ShippingAddress, this file has ShippingQuery.

// ShippingQuery contains information about an incoming shipping query.
// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	ID              string          `json:"id"`
	From            User            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}
