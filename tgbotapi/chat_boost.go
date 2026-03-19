package tgbotapi

// ChatBoostSource describes the source of a chat boost.
// https://core.telegram.org/bots/api#chatboostsource
type ChatBoostSource struct {
	// Source of the boost: "premium", "gift_code", or "giveaway"
	Source string `json:"source"`

	// For "premium" and "gift_code": the user that boosted the chat
	User *User `json:"user,omitempty"`

	// For "giveaway": identifier of a message in the chat with the giveaway; 0 if the message is not yet available
	GiveawayMessageID int `json:"giveaway_message_id,omitempty"`

	// For "giveaway": True, if the giveaway was completed but no user won the boost
	IsUnclaimed bool `json:"is_unclaimed,omitempty"`
}

// ChatBoost contains information about a chat boost.
// https://core.telegram.org/bots/api#chatboost
type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int             `json:"add_date"`
	ExpirationDate int             `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}

// ChatBoostUpdated represents a boost added to a chat or changed.
// https://core.telegram.org/bots/api#chatboostupdated
type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

// ChatBoostRemoved represents a boost removed from a chat.
// https://core.telegram.org/bots/api#chatboostremoved
type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int             `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

// ChatBoostAdded represents a service message about a user boosting a chat.
// https://core.telegram.org/bots/api#chatboostadded
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}
