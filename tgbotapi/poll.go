package tgbotapi

// Link represents an HTTP link.
//
// https://core.telegram.org/bots/api#link
type Link struct {
	// URL of the link
	URL string `json:"url"`
}

// InputMediaLink represents an HTTP link to be sent. It can be used as an InputPollOptionMedia.
// Bot API 10.1+
//
// https://core.telegram.org/bots/api#inputmedialink
type InputMediaLink struct {
	// Type of the media, must be "link"
	Type string `json:"type"`

	// HTTP URL of the link
	URL string `json:"url"`
}

// PollMedia describes media attached to a poll, a poll's quiz explanation, or a poll option.
// At most one of the fields is expected to be set.
//
// https://core.telegram.org/bots/api#pollmedia
type PollMedia struct {
	// Optional. Media is an animation
	Animation *Animation `json:"animation,omitempty"`

	// Optional. Media is an audio file
	Audio *Audio `json:"audio,omitempty"`

	// Optional. Media is a general file
	Document *Document `json:"document,omitempty"`

	// Optional. The HTTP link attached to the poll option. Bot API 10.1+
	Link *Link `json:"link,omitempty"`

	// Optional. Media is a live photo
	LivePhoto *LivePhoto `json:"live_photo,omitempty"`

	// Optional. Media is a shared location
	Location *Location `json:"location,omitempty"`

	// Optional. Media is a photo
	Photo []PhotoSize `json:"photo,omitempty"`

	// Optional. Media is a sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Optional. Media is a venue
	Venue *Venue `json:"venue,omitempty"`

	// Optional. Media is a video
	Video *Video `json:"video,omitempty"`
}

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

	// Optional. Media attached to the option. Bot API 10.0+
	Media *PollMedia `json:"media,omitempty"`
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

	// Optional. Media attached to the poll. Bot API 10.0+
	Media *PollMedia `json:"media,omitempty"`

	// Optional. Media attached to the quiz explanation. Bot API 10.0+
	ExplanationMedia *PollMedia `json:"explanation_media,omitempty"`

	// Optional. True, if the poll can be voted only by members of the chat it was sent to. Bot API 10.0+
	MembersOnly bool `json:"members_only,omitempty"`

	// Optional. A list of two-letter ISO 3166-1 alpha-2 country codes the poll is restricted to. Bot API 10.0+
	CountryCodes []string `json:"country_codes,omitempty"`
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
