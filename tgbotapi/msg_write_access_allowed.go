package tgbotapi

// WriteAccessAllowed
// https://core.telegram.org/bots/api#writeaccessallowed
type WriteAccessAllowed struct {
	// Optional. True, if the access was granted after the user accepted an explicit request
	// from a Web App sent by the TelegramMethod requestWriteAccess
	FromRequest bool `json:"from_request,omitempty"`

	// Optional. Name of the Web App, if the access was granted when the Web App was launched from a link
	WebAppName string `json:"web_app_name,omitempty"`

	// Optional. True, if the access was granted when the bot was added to the attachment or side menu
	FromAttachmentMenu bool `json:"from_attachment_menu,omitempty"`
}
