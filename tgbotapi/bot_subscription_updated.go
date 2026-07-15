package tgbotapi

// Bot subscription state constants for BotSubscriptionUpdated.State.
// https://core.telegram.org/bots/api#botsubscriptionupdated
const (
	// BotSubscriptionStateCanceled indicates the user canceled the subscription.
	BotSubscriptionStateCanceled = "canceled"

	// BotSubscriptionStateActive indicates the user re-enabled a previously canceled subscription.
	BotSubscriptionStateActive = "active"

	// BotSubscriptionStateFailed indicates payment for the subscription failed.
	BotSubscriptionStateFailed = "failed"
)

// BotSubscriptionUpdated contains information about changes to a user payment subscription toward the
// current bot (Bot API 10.2 General).
//
// https://core.telegram.org/bots/api#botsubscriptionupdated
type BotSubscriptionUpdated struct {
	// User who subscribed for payments toward the bot
	User User `json:"user"`

	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`

	// The new state of the subscription. Currently, it can be one of BotSubscriptionStateCanceled if the
	// user canceled the subscription, BotSubscriptionStateActive if the user re-enabled a previously
	// canceled subscription, or BotSubscriptionStateFailed if payment for the subscription failed.
	State string `json:"state"`
}
