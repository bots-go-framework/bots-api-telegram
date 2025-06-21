package tgbotapi

type SuccessfulPayment struct {
	Payment

	// Optional. Expiration date of the subscription, in Unix time; for recurring payments only
	SubscriptionExpirationDate int64 `json:"subscription_expiration_date"`

	// Optional. True, if the Payment is a recurring Payment for a subscription
	IsRecurring bool `json:"is_recurring,omitempty"`

	// Optional. True, if the Payment is the first Payment for a subscription
	IsFirstRecurring bool `json:"is_first_recurring,omitempty"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}
