package tgbotapi

// ReactionType describes a reaction type. Currently, it can be one of:
// - ReactionTypeEmoji
// - ReactionTypeCustomEmoji
// - ReactionTypePaid
// https://core.telegram.org/bots/api#reactiontype
type ReactionType struct {
	// Type of the reaction — "emoji", "custom_emoji", or "paid"
	Type string `json:"type"`

	// For "emoji": the emoji itself
	Emoji string `json:"emoji,omitempty"`

	// For "custom_emoji": custom emoji identifier
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// MessageReactionUpdated represents a change of a reaction on a message performed by a user.
// https://core.telegram.org/bots/api#messagereactionupdated
type MessageReactionUpdated struct {
	Chat        Chat           `json:"chat"`
	MessageID   int            `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

// ReactionCount represents a reaction added to a message along with the number of times it was added.
// https://core.telegram.org/bots/api#reactioncount
type ReactionCount struct {
	Type       ReactionType `json:"type"`
	TotalCount int          `json:"total_count"`
}

// MessageReactionCountUpdated represents reaction changes on a message with anonymous reactions.
// https://core.telegram.org/bots/api#messagereactioncountupdated
type MessageReactionCountUpdated struct {
	Chat      Chat            `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int             `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}
