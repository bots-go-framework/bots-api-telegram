package tgbotapi

// ForumTopic represents a forum topic.
// https://core.telegram.org/bots/api#forumtopic
type ForumTopic struct {
	MessageThreadID   int    `json:"message_thread_id"`
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
}

// ForumTopicCreated represents a service message about a new forum topic.
// https://core.telegram.org/bots/api#forumtopiccreated
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
}

// ForumTopicEdited represents a service message about an edited forum topic.
// https://core.telegram.org/bots/api#forumtopicedited
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicClosed represents a service message about a closed forum topic.
// https://core.telegram.org/bots/api#forumtopicclosed
type ForumTopicClosed struct{}

// ForumTopicReopened represents a service message about a reopened forum topic.
// https://core.telegram.org/bots/api#forumtopicreopened
type ForumTopicReopened struct{}

// GeneralForumTopicHidden represents a service message about the General forum topic hidden in a chat.
// https://core.telegram.org/bots/api#generalforumtopichidden
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden represents a service message about the General forum topic unhidden in a chat.
// https://core.telegram.org/bots/api#generalforumtopicunhidden
type GeneralForumTopicUnhidden struct{}
