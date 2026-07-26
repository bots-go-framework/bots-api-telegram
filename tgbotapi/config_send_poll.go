package tgbotapi

import (
	"encoding/json"
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

	// Optional. Media to attach to the option. Bot API 10.0+
	Media *InputPollOptionMedia `json:"media,omitempty"`
}

// InputPollMedia describes media to attach to a poll or its quiz explanation.
//
// https://core.telegram.org/bots/api#inputpollmedia
type InputPollMedia struct {
	// Input is an optional fully typed InputMedia* union member. When set, it is
	// serialized directly and the legacy flattened fields below are ignored.
	Input any `json:"-"`

	// Type of the media, one of "animation", "audio", "document", "live_photo", "location", "photo",
	// "venue", "video"
	Type string `json:"type"`

	// File to send, required for file-based types ("animation", "audio", "document", "live_photo",
	// "photo", "sticker", "video"). Pass a file_id to send a file that exists on the Telegram servers,
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass "attach://<file_attach_name>"
	// to upload a new one using multipart/form-data.
	Media string `json:"media,omitempty"`

	// Location to attach, required when Type is "location"
	Location *Location `json:"location,omitempty"`

	// Venue to attach, required when Type is "venue"
	Venue *Venue `json:"venue,omitempty"`
}

// InputPollOptionMedia describes media to attach to a poll option.
//
// https://core.telegram.org/bots/api#inputpolloptionmedia
type InputPollOptionMedia struct {
	// Input is an optional fully typed InputMedia* union member. When set, it is
	// serialized directly and the legacy flattened fields below are ignored.
	Input any `json:"-"`

	// Type of the media, one of "animation", "link", "live_photo", "location", "photo", "sticker",
	// "venue", "video". "link" added in Bot API 10.1
	Type string `json:"type"`

	// File to send, required for file-based types ("animation", "live_photo", "photo", "sticker", "video").
	// Pass a file_id to send a file that exists on the Telegram servers, pass an HTTP URL for Telegram to
	// get a file from the Internet, or pass "attach://<file_attach_name>" to upload a new one using
	// multipart/form-data.
	Media string `json:"media,omitempty"`

	// HTTP URL of the link, required when Type is "link". Bot API 10.1+
	URL string `json:"url,omitempty"`

	// Location to attach, required when Type is "location"
	Location *Location `json:"location,omitempty"`

	// Venue to attach, required when Type is "venue"
	Venue *Venue `json:"venue,omitempty"`
}

// NewInputPollMedia constructs a full-fidelity InputPollMedia union.
func NewInputPollMedia(input any) InputPollMedia {
	return InputPollMedia{Input: input}
}

// NewInputPollOptionMedia constructs a full-fidelity InputPollOptionMedia union.
func NewInputPollOptionMedia(input any) InputPollOptionMedia {
	return InputPollOptionMedia{Input: input}
}

func (v InputPollMedia) MarshalJSON() ([]byte, error) {
	if v.Input != nil {
		if err := validatePollMediaInput(v.Input, false); err != nil {
			return nil, err
		}
		return json.Marshal(v.Input)
	}
	return marshalLegacyPollMedia(v.Type, v.Media, "", v.Location, v.Venue, false)
}

func (v InputPollOptionMedia) MarshalJSON() ([]byte, error) {
	if v.Input != nil {
		if err := validatePollMediaInput(v.Input, true); err != nil {
			return nil, err
		}
		return json.Marshal(v.Input)
	}
	return marshalLegacyPollMedia(v.Type, v.Media, v.URL, v.Location, v.Venue, true)
}

func marshalLegacyPollMedia(mediaType, media, linkURL string, location *Location, venue *Venue, option bool) ([]byte, error) {
	switch mediaType {
	case "location":
		if location == nil {
			return nil, fmt.Errorf("location is required for poll media type %q", mediaType)
		}
		return json.Marshal(InputMediaLocation{
			Type:               mediaType,
			Latitude:           location.Latitude,
			Longitude:          location.Longitude,
			HorizontalAccuracy: location.HorizontalAccuracy,
		})
	case "venue":
		if venue == nil {
			return nil, fmt.Errorf("venue is required for poll media type %q", mediaType)
		}
		return json.Marshal(InputMediaVenue{
			Type:            mediaType,
			Latitude:        venue.Location.Latitude,
			Longitude:       venue.Location.Longitude,
			Title:           venue.Title,
			Address:         venue.Address,
			FoursquareID:    venue.FoursquareID,
			FoursquareType:  venue.FoursquareType,
			GooglePlaceID:   venue.GooglePlaceID,
			GooglePlaceType: venue.GooglePlaceType,
		})
	case "link":
		if !option {
			return nil, fmt.Errorf("link is only valid for poll-option media")
		}
		if linkURL == "" {
			return nil, fmt.Errorf("url is required for poll media type %q", mediaType)
		}
		return json.Marshal(InputMediaLink{Type: mediaType, URL: linkURL})
	default:
		if mediaType == "" {
			return nil, fmt.Errorf("poll media type is required")
		}
		allowed := mediaType == "animation" || mediaType == "live_photo" || mediaType == "photo" || mediaType == "video"
		if option {
			allowed = allowed || mediaType == "sticker"
		} else {
			allowed = allowed || mediaType == "audio" || mediaType == "document"
		}
		if !allowed {
			return nil, fmt.Errorf("media type %q is not valid for this poll media union", mediaType)
		}
		if media == "" {
			return nil, fmt.Errorf("media is required for poll media type %q", mediaType)
		}
		return json.Marshal(struct {
			Type  string `json:"type"`
			Media string `json:"media"`
		}{Type: mediaType, Media: media})
	}
}

func validatePollMediaInput(input any, option bool) error {
	switch media := input.(type) {
	case InputMediaAnimation:
		return validateTypedMedia(media.Type, "animation", media.Media)
	case *InputMediaAnimation:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validateTypedMedia(media.Type, "animation", media.Media)
	case InputMediaLivePhoto:
		return validateTypedMedia(media.Type, "live_photo", media.Media)
	case *InputMediaLivePhoto:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validateTypedMedia(media.Type, "live_photo", media.Media)
	case InputMediaLocation:
		if media.Type != "location" {
			return fmt.Errorf("location type must be %q", "location")
		}
		return nil
	case *InputMediaLocation:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validatePollMediaInput(*media, option)
	case InputMediaPhoto:
		return validateTypedMedia(media.Type, "photo", media.Media)
	case *InputMediaPhoto:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validateTypedMedia(media.Type, "photo", media.Media)
	case InputMediaVenue:
		if media.Type != "venue" || media.Title == "" || media.Address == "" {
			return fmt.Errorf("venue type, title, and address are required")
		}
		return nil
	case *InputMediaVenue:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validatePollMediaInput(*media, option)
	case InputMediaVideo:
		return validateTypedMedia(media.Type, "video", media.Media)
	case *InputMediaVideo:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validateTypedMedia(media.Type, "video", media.Media)
	case InputMediaAudio:
		if option {
			return fmt.Errorf("audio is not valid poll-option media")
		}
		return validateTypedMedia(media.Type, "audio", media.Media)
	case *InputMediaAudio:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validatePollMediaInput(*media, option)
	case InputMediaDocument:
		if option {
			return fmt.Errorf("document is not valid poll-option media")
		}
		return validateTypedMedia(media.Type, "document", media.Media)
	case *InputMediaDocument:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validatePollMediaInput(*media, option)
	case InputMediaLink:
		if !option || media.Type != "link" || media.URL == "" {
			return fmt.Errorf("link is only valid for poll options and requires type and url")
		}
		return nil
	case *InputMediaLink:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validatePollMediaInput(*media, option)
	case InputMediaSticker:
		if !option {
			return fmt.Errorf("sticker is only valid for poll options")
		}
		return validateTypedMedia(media.Type, "sticker", media.Media)
	case *InputMediaSticker:
		if media == nil {
			return fmt.Errorf("poll media is nil")
		}
		return validatePollMediaInput(*media, option)
	default:
		return fmt.Errorf("unsupported poll media type %T", input)
	}
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

	// Optional. Media to attach to the poll. Bot API 10.0+
	Media *InputPollMedia `json:"media,omitempty"`

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

	// Optional. Media to attach to the quiz explanation. Bot API 10.0+
	ExplanationMedia *InputPollMedia `json:"explanation_media,omitempty"`

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

	// Optional. Pass True if the poll can be voted only by members of the chat it is sent to. Bot API 10.0+
	MembersOnly bool `json:"members_only,omitempty"`

	// Optional. A JSON-serialized list of two-letter ISO 3166-1 alpha-2 country codes the poll is
	// restricted to. Bot API 10.0+
	CountryCodes []string `json:"country_codes,omitempty"`
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

	if v.Media != nil {
		if b, err := encodeToJson(v.Media); err != nil {
			return values, fmt.Errorf("failed to marshal poll media as JSON: %w", err)
		} else {
			values.Add("media", string(b))
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

	if v.ExplanationMedia != nil {
		if b, err := encodeToJson(v.ExplanationMedia); err != nil {
			return values, fmt.Errorf("failed to marshal poll explanation media as JSON: %w", err)
		} else {
			values.Add("explanation_media", string(b))
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

	if v.MembersOnly {
		values.Add("members_only", "true")
	}

	if len(v.CountryCodes) > 0 {
		if b, err := encodeToJson(v.CountryCodes); err != nil {
			return values, fmt.Errorf("failed to marshal poll country codes as JSON: %w", err)
		} else {
			values.Add("country_codes", string(b))
		}
	}

	return values, nil
}
