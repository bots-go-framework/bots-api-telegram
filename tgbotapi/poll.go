package tgbotapi

// PollOption contains information about one answer option in a poll.
// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	// Unique persistent identifier of the option. Bot API 9.6+
	PersistentID string          `json:"persistent_id"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int             `json:"voter_count"`

	// Optional. User that added the option, if it was added after the poll was created. Bot API 9.6+
	AddedByUser *User `json:"added_by_user,omitempty"`

	// Optional. Chat that added the option, if it was added after the poll was created. Bot API 9.6+
	AddedByChat *Chat `json:"added_by_chat,omitempty"`

	// Optional. Date when the option was added, in Unix time. Bot API 9.6+
	AdditionDate int `json:"addition_date,omitempty"`
}

// Poll contains information about a poll.
// https://core.telegram.org/bots/api#poll
type Poll struct {
	ID                    string          `json:"id"`
	Question              string          `json:"question"`
	QuestionEntities      []MessageEntity `json:"question_entities,omitempty"`
	Options               []PollOption    `json:"options"`
	TotalVoterCount       int             `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`

	// 0-based identifiers of correct answer options. Optional, returned only for polls in quiz mode
	// and only to the chat member who created the poll if the poll is not closed.
	// Replaces the deprecated single-value CorrectOptionID. Bot API 9.6+
	CorrectOptionIDs []int `json:"correct_option_ids,omitempty"`

	Explanation         string          `json:"explanation,omitempty"`
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod          int             `json:"open_period,omitempty"`
	CloseDate           int             `json:"close_date,omitempty"`

	// Optional. True, if the poll allows to change the chosen answer options. Bot API 9.6+
	AllowsRevoting bool `json:"allows_revoting,omitempty"`

	// Optional. Description of the poll. Bot API 9.6+
	Description string `json:"description,omitempty"`

	// Optional. Special entities that appear in the poll description. Bot API 9.6+
	DescriptionEntities []MessageEntity `json:"description_entities,omitempty"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	PollID    string `json:"poll_id"`
	VoterChat *Chat  `json:"voter_chat,omitempty"`
	User      *User  `json:"user,omitempty"`
	OptionIDs []int  `json:"option_ids"`

	// Optional. Persistent identifiers of the chosen options. Bot API 9.6+
	OptionPersistentIDs []string `json:"option_persistent_ids,omitempty"`
}
