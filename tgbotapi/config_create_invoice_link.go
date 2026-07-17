package tgbotapi

import (
	"fmt"
	"net/url"
	"strconv"
)

var _ Sendable = (*CreateInvoiceLinkConfig)(nil)

// CreateInvoiceLinkConfig creates a link for an invoice.
// https://core.telegram.org/bots/api#createinvoicelink
//
// Unlike sendInvoice this method is NOT tied to a chat: it returns a t.me
// invoice URL the payer can open anywhere. It therefore does NOT embed BaseChat
// (there is no chat_id) and carries the invoice product fields directly. Pass an
// empty ProviderToken and Currency "XTR" for payments in Telegram Stars.
type CreateInvoiceLinkConfig struct {
	// BusinessConnectionID: unique identifier of the business connection on
	// behalf of which the link will be created.
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	Title              string         `json:"title"`                         // Product name, 1-32 characters
	Description        string         `json:"description"`                   // Product description, 1-255 characters
	Payload            string         `json:"payload"`                       // Bot-defined invoice payload, 1-128 bytes. Not shown to the user.
	ProviderToken      string         `json:"provider_token,omitempty"`      // Payment provider token; empty string for Telegram Stars.
	Currency           string         `json:"currency"`                      // Three-letter ISO 4217 currency code; "XTR" for Telegram Stars.
	Prices             []LabeledPrice `json:"prices"`                        // Price breakdown. Exactly one item for Telegram Stars.
	SubscriptionPeriod int64          `json:"subscription_period,omitempty"` // Seconds the subscription stays active before the next payment. Must be 2592000 (30 days) if used.

	MaxTipAmount        int64   `json:"max_tip_amount,omitempty"`        // Maximum accepted tip in the smallest currency units. Not supported for Telegram Stars.
	SuggestedTipAmounts []int64 `json:"suggested_tip_amounts,omitempty"` // Up to 4 suggested tip amounts, strictly increasing, not exceeding MaxTipAmount.
	ProviderData        string  `json:"provider_data,omitempty"`         // JSON-serialized data about the invoice, shared with the payment provider.
	PhotoURL            string  `json:"photo_url,omitempty"`             // URL of the product photo for the invoice.
	PhotoSize           int     `json:"photo_size,omitempty"`            // Photo size in bytes.
	PhotoWidth          int     `json:"photo_width,omitempty"`           // Photo width.
	PhotoHeight         int     `json:"photo_height,omitempty"`          // Photo height.
	NeedName            bool    `json:"need_name,omitempty"`             // Require the user's full name. Ignored for Telegram Stars.
	NeedPhoneNumber     bool    `json:"need_phone_number,omitempty"`     // Require the user's phone number. Ignored for Telegram Stars.
	NeedEmail           bool    `json:"need_email,omitempty"`            // Require the user's email. Ignored for Telegram Stars.
	NeedShippingAddress bool    `json:"need_shipping_address,omitempty"` // Require the user's shipping address. Ignored for Telegram Stars.

	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"` // Send the phone number to the provider. Ignored for Telegram Stars.
	SendEmailToProvider       bool `json:"send_email_to_provider,omitempty"`        // Send the email to the provider. Ignored for Telegram Stars.
	IsFlexible                bool `json:"is_flexible,omitempty"`                   // Final price depends on the shipping method. Ignored for Telegram Stars.
}

func (*CreateInvoiceLinkConfig) TelegramMethod() string {
	return "createInvoiceLink"
}

// Values returns the url.Values representation of CreateInvoiceLinkConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v *CreateInvoiceLinkConfig) Values() (url.Values, error) {
	values := url.Values{}
	values.Add("title", v.Title)
	values.Add("description", v.Description)
	values.Add("payload", v.Payload)
	values.Add("currency", v.Currency)
	// provider_token is required by the API but MUST be sent as an empty string
	// for Telegram Stars payments, so it is always included.
	values.Add("provider_token", v.ProviderToken)

	if len(v.Prices) > 0 {
		b, err := encodeToJson(v.Prices)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal invoice prices as JSON: %w", err)
		}
		values.Add("prices", string(b))
	}
	if v.BusinessConnectionID != "" {
		values.Add("business_connection_id", v.BusinessConnectionID)
	}
	if v.SubscriptionPeriod != 0 {
		values.Add("subscription_period", strconv.FormatInt(v.SubscriptionPeriod, 10))
	}
	if v.MaxTipAmount != 0 {
		values.Add("max_tip_amount", strconv.FormatInt(v.MaxTipAmount, 10))
	}
	if len(v.SuggestedTipAmounts) > 0 {
		b, err := encodeToJson(v.SuggestedTipAmounts)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal suggested_tip_amounts as JSON: %w", err)
		}
		values.Add("suggested_tip_amounts", string(b))
	}
	if v.ProviderData != "" {
		values.Add("provider_data", v.ProviderData)
	}
	if v.PhotoURL != "" {
		values.Add("photo_url", v.PhotoURL)
	}
	if v.PhotoSize != 0 {
		values.Add("photo_size", strconv.Itoa(v.PhotoSize))
	}
	if v.PhotoWidth != 0 {
		values.Add("photo_width", strconv.Itoa(v.PhotoWidth))
	}
	if v.PhotoHeight != 0 {
		values.Add("photo_height", strconv.Itoa(v.PhotoHeight))
	}
	if v.NeedName {
		values.Add("need_name", "true")
	}
	if v.NeedPhoneNumber {
		values.Add("need_phone_number", "true")
	}
	if v.NeedEmail {
		values.Add("need_email", "true")
	}
	if v.NeedShippingAddress {
		values.Add("need_shipping_address", "true")
	}
	if v.SendPhoneNumberToProvider {
		values.Add("send_phone_number_to_provider", "true")
	}
	if v.SendEmailToProvider {
		values.Add("send_email_to_provider", "true")
	}
	if v.IsFlexible {
		values.Add("is_flexible", "true")
	}
	return values, nil
}
