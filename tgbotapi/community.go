package tgbotapi

// Community represents a community: several supergroups, channels, and bots linked together around a
// shared topic or audience (Bot API 10.2 Communities).
//
// https://core.telegram.org/bots/api#community
type Community struct {
	// Unique identifier for this community. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting it. But it has at most 52
	// significant bits, so a signed 64-bit integer or double-precision float type are safe for storing
	// this identifier.
	ID int64 `json:"id"`

	// Name of the community
	Name string `json:"name"`
}

// CommunityChatAdded describes a service message about a chat being added to a community (Bot API 10.2
// Communities).
//
// https://core.telegram.org/bots/api#communitychatadded
type CommunityChatAdded struct {
	// The new community to which the chat belongs
	Community Community `json:"community"`
}

// CommunityChatRemoved describes a service message about a chat being removed from a community.
// Currently holds no information (Bot API 10.2 Communities).
//
// https://core.telegram.org/bots/api#communitychatremoved
type CommunityChatRemoved struct {
}
