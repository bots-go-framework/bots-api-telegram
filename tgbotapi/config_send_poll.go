package tgbotapi

import (
	"fmt"
	"net/url"
)

// InputPollOption contains information about one answer option in a poll to be sent.
// https://core.telegram.org/bots/api#inputpolloption
type InputPollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`

	// Optional. Mode for parsing entities in the text. See formatting options for more details.
	// Currently, only custom emoji entities are allowed.
	TextParseMode string `json:"text_parse_mode,omitempty"`

	// Optional. A JSON-serialized list of special entities that appear in the option text.
	// It can be specified instead of text_parse_mode.
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
}

var _ Sendable = (*PollConfig)(nil)

// PollConfig contains information about a sendPoll request.
// https://core.telegram.org/bots/api#sendpoll
type PollConfig struct {
	BaseChat

	// Poll question, 1-300 characters
	Question string `json:"question"`

	// Optional. Mode for parsing entities in the question. Currently, only custom emoji entities are allowed.
	QuestionParseMode string `json:"question_parse_mode,omitempty"`

	// Optional. A JSON-serialized list of special entities that appear in the poll question.
	// It can be specified instead of question_parse_mode.
	QuestionEntities []MessageEntity `json:"question_entities,omitempty"`

	// A JSON-serialized list of 1-12 answer options
	Options []InputPollOption `json:"options"`

	// Optional. True, if the poll needs to be anonymous, defaults to True.
	// NOTE: because the zero value of bool is false, this config can't distinguish
	// "not set" from "explicitly false"; leaving it unset sends nothing to Telegram
	// and the anonymous default applies. To send a non-anonymous poll, this field
	// must be explicitly set to true when marshalling is not desired - see Values().
	IsAnonymous bool `json:"is_anonymous,omitempty"`

	// Optional. Poll type, "quiz" or "regular", defaults to "regular"
	Type string `json:"type,omitempty"`

	// Optional. Pass True if the poll allows multiple answers, defaults to False.
	// Bot API 9.6 allows this to be combined with quiz-mode polls that have multiple correct answers.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`

	// Optional. Pass True if the poll allows to change the chosen answer options,
	// defaults to False for quizzes and to True for regular polls. Bot API 9.6+
	AllowsRevoting bool `json:"allows_revoting,omitempty"`

	// Optional. Pass True if the poll options must be shown in random order. Bot API 9.6+
	ShuffleOptions bool `json:"shuffle_options,omitempty"`

	// Optional. Pass True if answer options can be added to the poll after creation;
	// not supported for anonymous polls and quizzes. Bot API 9.6+
	AllowAddingOptions bool `json:"allow_adding_options,omitempty"`

	// Optional. Pass True if poll results must be shown only after the poll closes. Bot API 9.6+
	HideResultsUntilCloses bool `json:"hide_results_until_closes,omitempty"`

	// Optional. 0-based identifiers of the correct answer options, required for polls in quiz mode.
	// Supports multiple correct answers as of Bot API 9.6.
	CorrectOptionIDs []int `json:"correct_option_ids,omitempty"`

	// Optional. Text shown when a user chooses an incorrect answer or taps the lamp icon in a quiz
	Explanation string `json:"explanation,omitempty"`

	// Optional. Mode for parsing entities in the explanation.
	ExplanationParseMode string `json:"explanation_parse_mode,omitempty"`

	// Optional. A JSON-serialized list of special entities that appear in the poll explanation.
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`

	// Optional. Amount of time in seconds the poll will be active after creation, 5-2628000.
	// Can't be used together with CloseDate.
	OpenPeriod int `json:"open_period,omitempty"`

	// Optional. Point in time (Unix timestamp) when the poll will be automatically closed.
	// Can't be used together with OpenPeriod.
	CloseDate int `json:"close_date,omitempty"`

	// Optional. Pass True if the poll needs to be immediately closed.
	IsClosed bool `json:"is_closed,omitempty"`

	// Optional. Description of the poll to be sent, 0-1024 characters after entities parsing. Bot API 9.6+
	Description string `json:"description,omitempty"`

	// Optional. Mode for parsing entities in the poll description. Bot API 9.6+
	DescriptionParseMode string `json:"description_parse_mode,omitempty"`

	// Optional. A JSON-serialized list of special entities that appear in the poll description,
	// which can be specified instead of DescriptionParseMode. Bot API 9.6+
	DescriptionEntities []MessageEntity `json:"description_entities,omitempty"`
}

// TelegramMethod returns Telegram API method name for sending a Poll.
func (*PollConfig) TelegramMethod() string {
	return "sendPoll"
}

// Values returns url.Values representation of PollConfig.
//
//goland:noinspection GoMixedReceiverTypes
func (v *PollConfig) Values() (url.Values, error) {
	values, err := v.BaseChat.Values()
	if err != nil {
		return values, err
	}

	values.Add("question", v.Question)
	if v.QuestionParseMode != "" {
		values.Add("question_parse_mode", v.QuestionParseMode)
	}
	if len(v.QuestionEntities) > 0 {
		if b, err := encodeToJson(v.QuestionEntities); err != nil {
			return values, fmt.Errorf("failed to marshal question entities as JSON: %w", err)
		} else {
			values.Add("question_entities", string(b))
		}
	}

	if len(v.Options) > 0 {
		if b, err := encodeToJson(v.Options); err != nil {
			return values, fmt.Errorf("failed to marshal poll options as JSON: %w", err)
		} else {
			values.Add("options", string(b))
		}
	}

	if v.IsAnonymous {
		values.Add("is_anonymous", "true")
	}
	if v.Type != "" {
		values.Add("type", v.Type)
	}
	if v.AllowsMultipleAnswers {
		values.Add("allows_multiple_answers", "true")
	}
	if v.AllowsRevoting {
		values.Add("allows_revoting", "true")
	}
	if v.ShuffleOptions {
		values.Add("shuffle_options", "true")
	}
	if v.AllowAddingOptions {
		values.Add("allow_adding_options", "true")
	}
	if v.HideResultsUntilCloses {
		values.Add("hide_results_until_closes", "true")
	}

	if len(v.CorrectOptionIDs) > 0 {
		if b, err := encodeToJson(v.CorrectOptionIDs); err != nil {
			return values, fmt.Errorf("failed to marshal correct option ids as JSON: %w", err)
		} else {
			values.Add("correct_option_ids", string(b))
		}
	}

	if v.Explanation != "" {
		values.Add("explanation", v.Explanation)
	}
	if v.ExplanationParseMode != "" {
		values.Add("explanation_parse_mode", v.ExplanationParseMode)
	}
	if len(v.ExplanationEntities) > 0 {
		if b, err := encodeToJson(v.ExplanationEntities); err != nil {
			return values, fmt.Errorf("failed to marshal explanation entities as JSON: %w", err)
		} else {
			values.Add("explanation_entities", string(b))
		}
	}

	if v.OpenPeriod != 0 {
		values.Add("open_period", fmt.Sprintf("%d", v.OpenPeriod))
	}
	if v.CloseDate != 0 {
		values.Add("close_date", fmt.Sprintf("%d", v.CloseDate))
	}
	if v.IsClosed {
		values.Add("is_closed", "true")
	}

	if v.Description != "" {
		values.Add("description", v.Description)
	}
	if v.DescriptionParseMode != "" {
		values.Add("description_parse_mode", v.DescriptionParseMode)
	}
	if len(v.DescriptionEntities) > 0 {
		if b, err := encodeToJson(v.DescriptionEntities); err != nil {
			return values, fmt.Errorf("failed to marshal description entities as JSON: %w", err)
		} else {
			values.Add("description_entities", string(b))
		}
	}

	return values, nil
}
