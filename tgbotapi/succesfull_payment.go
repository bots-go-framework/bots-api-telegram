package tgbotapi

type SuccessfulPayment struct {
	payment

	// Optional. Expiration date of the subscription, in Unix time; for recurring payments only
	SubscriptionExpirationDate int `json:"subscription_expiration_date"`

	// Optional. True, if the payment is a recurring payment for a subscription
	IsRecurring bool `json:"is_recurring,omitempty"`

	// Optional. True, if the payment is the first payment for a subscription
	IsFirstRecurring bool `json:"is_first_recurring,omitempty"`

	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// Optional. Order information provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}
