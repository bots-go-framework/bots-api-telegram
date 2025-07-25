package tgbotapi

import "net/url"

var _ Sendable = (*CreateInvoiceLinkConfig)(nil)

type CreateInvoiceLinkConfig struct {
	BaseChat

	// Unique identifier of the business connection on behalf of which the message will be sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
}

func (*CreateInvoiceLinkConfig) TelegramMethod() string {
	return "createInvoiceLink"
}

// Values returns url.Values representation of InvoiceConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v *CreateInvoiceLinkConfig) Values() (url.Values, error) {
	values, _ := v.BaseChat.Values()
	if v.BusinessConnectionID != "" {
		values.Add("business_connection_id", v.BusinessConnectionID)
	}
	return values, nil
}
