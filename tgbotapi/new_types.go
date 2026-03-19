package tgbotapi

// SuggestedPostInfo contains information about suggested post parameters.
// https://core.telegram.org/bots/api#suggestedpostinfo
type SuggestedPostInfo struct {
	Price             *SuggestedPostPrice `json:"price,omitempty"`
	SuggestedPostDate int                 `json:"suggested_post_date,omitempty"`
}

// SuggestedPostPrice describes the price of a suggested post.
// https://core.telegram.org/bots/api#suggestedpostprice
type SuggestedPostPrice struct {
	StarCount int `json:"star_count"`
}

// SuggestedPostApproved represents a service message about the approval of a suggested post.
// https://core.telegram.org/bots/api#suggestedpostapproved
type SuggestedPostApproved struct {
	PostDate int `json:"post_date"`
}

// SuggestedPostDeclined represents a service message about the rejection of a suggested post.
// https://core.telegram.org/bots/api#suggestedpostdeclined
type SuggestedPostDeclined struct{}

// SuggestedPostApprovalFailed represents a service message about the failed approval of a suggested post.
// https://core.telegram.org/bots/api#suggestedpostapprovalfailed
type SuggestedPostApprovalFailed struct{}

// SuggestedPostPaid represents a service message about a successful payment for a suggested post.
// https://core.telegram.org/bots/api#suggestedpostpaid
type SuggestedPostPaid struct {
	StarCount int `json:"star_count"`
}

// SuggestedPostRefunded represents a service message about a payment refund for a suggested post.
// https://core.telegram.org/bots/api#suggestedpostrefunded
type SuggestedPostRefunded struct {
	StarCount int `json:"star_count"`
}

// DirectMessagesTopic represents information about a direct messages chat topic.
// https://core.telegram.org/bots/api#directmessagestopic
type DirectMessagesTopic struct {
	MessageThreadID int    `json:"message_thread_id"`
	Name            string `json:"name,omitempty"`
}

// ChatOwnerLeft represents a service message: chat owner has left the chat.
// https://core.telegram.org/bots/api#chatownerleft
type ChatOwnerLeft struct{}

// ChatOwnerChanged represents a service message: chat ownership has been transferred.
// https://core.telegram.org/bots/api#chatownerchanged
type ChatOwnerChanged struct {
	OldOwner User `json:"old_owner"`
	NewOwner User `json:"new_owner"`
}

// UserRating contains information about a user's rating.
// https://core.telegram.org/bots/api#userrating
type UserRating struct {
	Rating int `json:"rating"`
}

// Checklist represents a checklist message.
// https://core.telegram.org/bots/api#checklist
type Checklist struct {
	Title         string          `json:"title"`
	TitleEntities []MessageEntity `json:"title_entities,omitempty"`
	Tasks         []ChecklistTask `json:"tasks"`
	OthersCanAdd  bool            `json:"others_can_add,omitempty"`
	OthersCanMark bool            `json:"others_can_mark,omitempty"`
}

// ChecklistTask represents a task in a checklist.
// https://core.telegram.org/bots/api#checklisttask
type ChecklistTask struct {
	ID              int             `json:"id"`
	Text            string          `json:"text"`
	TextEntities    []MessageEntity `json:"text_entities,omitempty"`
	CompletedByUser *User           `json:"completed_by_user,omitempty"`
	CompletedByChat *Chat           `json:"completed_by_chat,omitempty"`
	CompletionDate  int             `json:"completion_date,omitempty"`
}

// ChecklistTasksDone represents a service message about tasks in a checklist marked as done or not done.
// https://core.telegram.org/bots/api#checklisttasksdone
type ChecklistTasksDone struct {
	ChecklistMessageID int   `json:"checklist_message_id"`
	MarkedAsDone       []int `json:"marked_as_done,omitempty"`
	MarkedAsNotDone    []int `json:"marked_as_not_done,omitempty"`
}

// ChecklistTasksAdded represents a service message about tasks added to a checklist.
// https://core.telegram.org/bots/api#checklisttasksadded
type ChecklistTasksAdded struct {
	ChecklistMessageID int             `json:"checklist_message_id"`
	Tasks              []ChecklistTask `json:"tasks"`
}

// DirectMessagePriceChanged represents a service message about the price change for paid messages.
// https://core.telegram.org/bots/api#directmessagepricechanged
type DirectMessagePriceChanged struct {
	DirectMessageStarCount int `json:"direct_message_star_count"`
}

// PaidMessagePriceChanged represents a service message about a change in the price of paid messages.
// https://core.telegram.org/bots/api#paidmessagepricechanged
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}

// VideoQuality represents the quality of a video.
// https://core.telegram.org/bots/api#videoquality
type VideoQuality struct {
	Type     string `json:"type"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size,omitempty"`
}
