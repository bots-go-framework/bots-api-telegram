package tgbotapi

// SuggestedPostParameters contains parameters of a post suggested by the bot.
type SuggestedPostParameters struct {
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int64               `json:"send_date,omitempty"`
}
