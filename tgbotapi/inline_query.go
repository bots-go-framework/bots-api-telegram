package tgbotapi

//go:generate ffjson $GOFILE

import (
	"errors"
	"github.com/pquerna/ffjson/ffjson"
	"net/url"
	"strconv"
)

// InlineQueryResultsButton represents a button to be shown above inline query results.
// You must use exactly one of the optional fields.
type InlineQueryResultsButton struct {

	// Label text on the button
	Text string `json:"text,omitempty"`

	// Optional. Description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to switch back to the inline mode using the method switchInlineQuery inside the Web App.
	WebApp *WebAppInfo `json:"url,omitempty"`

	// Optional. Deep-linking parameter for the /start message sent to the bot when a user presses the button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
	//
	// Example: An inline bot that sends YouTube videos can ask the user to connect the bot
	// to their YouTube account to adapt search results accordingly.
	// To do this, it displays a 'Connect your YouTube account' button above the results, or even before showing any.
	// The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an OAuth link.
	// Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
	StartParameter string `json:"start_parameter,omitempty"`
}

func (b InlineQueryResultsButton) Validate() error {
	if b.Text == "" {
		return errors.New("InlineQueryResultsButton.Text is empty")
	}
	if b.WebApp != nil && b.StartParameter != "" {
		return errors.New("InlineQueryResultsButton has more than one field populated")
	}
	return nil
}

// InlineConfig contains information on making an InlineQuery response.
type InlineConfig struct {
	InlineQueryID string `json:"inline_query_id"`

	Results []InlineQueryResult `json:"results,omitempty"`

	// Optional.
	// The maximum amount of time in seconds that the result of the inline query may be cached on the server.
	// Defaults to 300.
	CacheTime int `json:"cache_time"`

	// Optional	Pass True if results may be cached on the server side only for the user that sent the query.
	// By default, results may be returned to any user who sends the same query.
	IsPersonal bool `json:"is_personal,omitempty"`

	// Optional.
	// Pass the offset that a client should send in the next query with the same text to receive more results.
	// Pass an empty string if there are no more results or if you don't support pagination.
	// Offset length can't exceed 64 bytes.
	NextOffset string `json:"next_offset,omitempty"`

	Button *InlineQueryResultsButton `json:"button,omitempty"`
}

//goland:noinspection GoMixedReceiverTypes
func (config InlineConfig) method() string {
	return "answerInlineQuery"
}

// Values returns URL values representation of InlineConfig
//
//goland:noinspection GoMixedReceiverTypes
func (config InlineConfig) Values() (url.Values, error) {
	if len(config.NextOffset) > 64 {
		return nil, errors.New("NextOffset length can't exceed 64 bytes")
	}
	if len(config.Results) == 0 {
		return nil, errors.New("InlineConfig.Results is empty")
	}
	v := url.Values{}

	v.Add("inline_query_id", config.InlineQueryID)
	if config.CacheTime != 0 {
		v.Add("cache_time", strconv.Itoa(config.CacheTime))
	}
	if config.IsPersonal {
		v.Add("is_personal", strconv.FormatBool(config.IsPersonal))
	}
	if config.NextOffset != "" {
		v.Add("next_offset", config.NextOffset)
	}
	//if config.SwitchPMText != "" {
	//	v.Add("switch_pm_text", config.SwitchPMText)
	//}
	//if config.SwitchPMParameter != "" {
	//	v.Add("switch_pm_parameter", config.SwitchPMParameter)
	//}

	data, err := ffjson.Marshal(config.Results)
	if err != nil {
		ffjson.Pool(data)
		return v, err
	}
	v.Add("results", string(data))
	ffjson.Pool(data)

	return v, nil
}

// InlineQuery is a Query from Telegram for an inline request.
type InlineQuery struct {
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location,omitempty"` // optional
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}
