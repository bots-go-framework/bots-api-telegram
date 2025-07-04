package tgbotapi

import (
	"fmt"
	"net/url"
)

type LabeledPrice struct {
	Label  string `json:"label"`  // Portion label
	Amount int    `json:"amount"` // Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

var _ Sendable = (*InvoiceConfig)(nil)

type InvoiceConfig struct {
	BaseChat
	Title               string         `json:"title"`                           // Product name, 1-32 characters
	Description         string         `json:"description"`                     // Product description, 1-255 characters
	Payload             string         `json:"payload"`                         // Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
	ProviderToken       string         `json:"provider_token,omitempty"`        // 	Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
	Currency            string         `json:"currency"`                        // Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
	Prices              []LabeledPrice `json:"prices"`                          // Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
	MaxTipAmount        int64          `json:"max_tip_amount,omitempty"`        // The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
	SuggestedTipAmounts []int64        `json:"suggested_tip_amounts,omitempty"` //A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	StartParameter      string         `json:"start_parameter,omitempty"`       // Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
	ProviderData        string         `json:"provider_data,omitempty"`         // JSON-serialized data about the invoice, which will be shared with the Payment provider. A detailed description of required fields should be provided by the Payment provider.
	PhotoURL            string         `json:"photo_url,omitempty"`             // URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize           int            `json:"photo_size,omitempty"`            // Photo size in bytes
	PhotoWidth          int            `json:"photo_width,omitempty"`           // Photo width
	PhotoHeight         int            `json:"photo_height,omitempty"`          // Photo height
	NeedName            bool           `json:"need_name,omitempty"`             // Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
	NeedPhoneNumber     bool           `json:"need_phone_number,omitempty"`     // Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
	NeedEmail           bool           `json:"need_email,omitempty"`            // Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
	NeedShippingAddress bool           `json:"need_shipping_address,omitempty"` // Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.

	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"` // Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
	SendEmailToProvider       bool `json:"send_email_to_provider,omitempty"`        // Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
	IsFlexible                bool `json:"is_flexible,omitempty"`                   // Pass True if the final price depends on the shipping BotEndpoint. Ignored for payments in Telegram Stars.
}

func (*InvoiceConfig) TelegramMethod() string {
	return "sendInvoice"
}

// Values returns url.Values representation of InvoiceConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v *InvoiceConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()
	values.Add("title", v.Title)
	values.Add("description", v.Description)
	values.Add("payload", v.Payload)
	values.Add("currency", v.Currency)
	if v.ProviderToken != "" {
		values.Add("provider_token", v.ProviderToken)
	}
	if len(v.Prices) > 0 {
		if b, err := encodeToJson(v.Prices); err != nil {
			return nil, fmt.Errorf("failed to marshal invoice prices as JSON: %v", err)
		} else {
			values.Add("prices", string(b))
		}
	}
	return values, nil
}
