package tgbotapi

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	richMediaIDPattern        = regexp.MustCompile(`^[A-Za-z0-9_-]{1,64}$`)
	richDateTimeFormatPattern = regexp.MustCompile(`^(r|w?[dD]?[tT]?)$`)
)

func (v InputRichMessage) validate(allowThinking bool) error {
	set := 0
	if v.HTML != "" {
		set++
	}
	if v.Markdown != "" {
		set++
	}
	if len(v.Blocks) > 0 {
		set++
	}
	if set == 0 {
		return errors.New("one of HTML, Markdown, or Blocks must be set")
	}
	if set > 1 {
		return errors.New("only one of HTML, Markdown, or Blocks may be set")
	}
	if len(v.Media) > 0 && len(v.Blocks) > 0 {
		return errors.New("media can only be referenced by HTML or Markdown content")
	}
	for i := range v.Media {
		if err := v.Media[i].Validate(); err != nil {
			return fmt.Errorf("media[%d]: %w", i, err)
		}
	}
	for i := range v.Blocks {
		if err := v.Blocks[i].validate(allowThinking); err != nil {
			return fmt.Errorf("blocks[%d]: %w", i, err)
		}
	}
	return nil
}

// Validate checks the media identifier and the allowed InputMedia variant.
func (v InputRichMessageMedia) Validate() error {
	if !richMediaIDPattern.MatchString(v.ID) {
		return errors.New("id must contain 1-64 ASCII letters, digits, underscores, or hyphens")
	}
	switch media := v.Media.(type) {
	case InputMediaAnimation:
		return validateTypedMedia(media.Type, "animation", media.Media)
	case *InputMediaAnimation:
		if media == nil {
			return errors.New("media is nil")
		}
		return validateTypedMedia(media.Type, "animation", media.Media)
	case InputMediaAudio:
		return validateTypedMedia(media.Type, "audio", media.Media)
	case *InputMediaAudio:
		if media == nil {
			return errors.New("media is nil")
		}
		return validateTypedMedia(media.Type, "audio", media.Media)
	case InputMediaPhoto:
		return validateTypedMedia(media.Type, "photo", media.Media)
	case *InputMediaPhoto:
		if media == nil {
			return errors.New("media is nil")
		}
		return validateTypedMedia(media.Type, "photo", media.Media)
	case InputMediaVideo:
		return validateTypedMedia(media.Type, "video", media.Media)
	case *InputMediaVideo:
		if media == nil {
			return errors.New("media is nil")
		}
		return validateTypedMedia(media.Type, "video", media.Media)
	case InputMediaVoiceNote:
		return validateTypedMedia(media.Type, "voice_note", media.Media)
	case *InputMediaVoiceNote:
		if media == nil {
			return errors.New("media is nil")
		}
		return validateTypedMedia(media.Type, "voice_note", media.Media)
	default:
		return fmt.Errorf("unsupported media type %T", v.Media)
	}
}

func validateTypedMedia(gotType, wantType, media string) error {
	if gotType != wantType {
		return fmt.Errorf("type must be %q, got %q", wantType, gotType)
	}
	if media == "" {
		return errors.New("media is required")
	}
	return nil
}

// Validate checks the three RichText wire shapes and recursively validates
// object/array-shaped content.
func (r RichText) Validate() error {
	if r.Type == "" {
		if r.Items != nil && r.PlainText != "" {
			return errors.New("plain text and item array are mutually exclusive")
		}
		for i := range r.Items {
			if err := r.Items[i].Validate(); err != nil {
				return fmt.Errorf("items[%d]: %w", i, err)
			}
		}
		return nil
	}
	if r.Items != nil || r.PlainText != "" {
		return errors.New("object-shaped RichText can't also contain plain text or items")
	}
	requireText := func() error {
		if r.Text == nil {
			return fmt.Errorf("text is required for rich text type %q", r.Type)
		}
		return r.Text.Validate()
	}
	switch r.Type {
	case RichTextTypeBold, RichTextTypeItalic, RichTextTypeUnderline, RichTextTypeStrikethrough,
		RichTextTypeSpoiler, RichTextTypeSubscript, RichTextTypeSuperscript, RichTextTypeMarked,
		RichTextTypeCode:
		return requireText()
	case RichTextTypeDateTime:
		if err := requireText(); err != nil {
			return err
		}
		if !richDateTimeFormatPattern.MatchString(r.DateTimeFormat) {
			return fmt.Errorf(
				"date_time_format must match r|w?[dD]?[tT]?, got %q",
				r.DateTimeFormat,
			)
		}
	case RichTextTypeTextMention:
		if err := requireText(); err != nil {
			return err
		}
		if r.User == nil {
			return errors.New("user is required")
		}
	case RichTextTypeCustomEmoji:
		if r.CustomEmojiID == "" || r.AlternativeText == "" {
			return errors.New("custom_emoji_id and alternative_text are required")
		}
	case RichTextTypeMathematicalExpression:
		if r.Expression == "" {
			return errors.New("expression is required")
		}
	case RichTextTypeUrl:
		if err := requireText(); err != nil {
			return err
		}
		if r.URL == "" {
			return errors.New("url is required")
		}
	case RichTextTypeEmailAddress:
		if err := requireText(); err != nil {
			return err
		}
		if r.EmailAddress == "" {
			return errors.New("email_address is required")
		}
	case RichTextTypePhoneNumber:
		if err := requireText(); err != nil {
			return err
		}
		if r.PhoneNumber == "" {
			return errors.New("phone_number is required")
		}
	case RichTextTypeBankCardNumber:
		if err := requireText(); err != nil {
			return err
		}
		if r.BankCardNumber == "" {
			return errors.New("bank_card_number is required")
		}
	case RichTextTypeMention:
		if err := requireText(); err != nil {
			return err
		}
		if r.Username == "" {
			return errors.New("username is required")
		}
	case RichTextTypeHashtag:
		if err := requireText(); err != nil {
			return err
		}
		if r.Hashtag == "" {
			return errors.New("hashtag is required")
		}
	case RichTextTypeCashtag:
		if err := requireText(); err != nil {
			return err
		}
		if r.Cashtag == "" {
			return errors.New("cashtag is required")
		}
	case RichTextTypeBotCommand:
		if err := requireText(); err != nil {
			return err
		}
		if r.BotCommand == "" {
			return errors.New("bot_command is required")
		}
	case RichTextTypeAnchor:
		if r.Name == "" {
			return errors.New("name is required")
		}
	case RichTextTypeAnchorLink:
		return requireText()
	case RichTextTypeReference:
		if err := requireText(); err != nil {
			return err
		}
		if r.Name == "" {
			return errors.New("name is required")
		}
	case RichTextTypeReferenceLink:
		if err := requireText(); err != nil {
			return err
		}
		if r.ReferenceName == "" {
			return errors.New("reference_name is required")
		}
	default:
		return fmt.Errorf("unknown rich text type %q", r.Type)
	}
	return nil
}

func validateRichText(field string, text *RichText) error {
	if text == nil {
		return fmt.Errorf("%s is required", field)
	}
	if err := text.Validate(); err != nil {
		return fmt.Errorf("%s: %w", field, err)
	}
	return nil
}

// Validate checks a persistent InputRichBlock. Thinking blocks are only valid
// through InputRichMessage.ValidateDraft.
func (b InputRichBlock) Validate() error {
	return b.validate(false)
}

func (b InputRichBlock) validate(allowThinking bool) error {
	validateBlocks := func(blocks []InputRichBlock) error {
		if len(blocks) == 0 {
			return errors.New("blocks are required")
		}
		for i := range blocks {
			if err := blocks[i].validate(allowThinking); err != nil {
				return fmt.Errorf("blocks[%d]: %w", i, err)
			}
		}
		return nil
	}
	switch b.Type {
	case RichBlockTypeParagraph, RichBlockTypePreformatted, RichBlockTypeFooter:
		return validateRichText("text", b.Text)
	case RichBlockTypeSectionHeading:
		if err := validateRichText("text", b.Text); err != nil {
			return err
		}
		if b.Size < 1 || b.Size > 6 {
			return errors.New("heading size must be between 1 and 6")
		}
	case RichBlockTypeDivider:
		return nil
	case RichBlockTypeMathematicalExpression:
		if b.Expression == "" {
			return errors.New("expression is required")
		}
	case RichBlockTypeAnchor:
		if b.Name == "" {
			return errors.New("name is required")
		}
	case RichBlockTypeList:
		if len(b.Items) == 0 {
			return errors.New("items are required")
		}
		for i := range b.Items {
			item := b.Items[i]
			if item.IsChecked && !item.HasCheckbox {
				return fmt.Errorf("items[%d]: is_checked requires has_checkbox", i)
			}
			switch item.Type {
			case "", "a", "A", "i", "I", "1":
			default:
				return fmt.Errorf("items[%d]: invalid label type %q", i, item.Type)
			}
			if err := validateBlocks(item.Blocks); err != nil {
				return fmt.Errorf("items[%d]: %w", i, err)
			}
		}
	case RichBlockTypeBlockQuotation, RichBlockTypeCollage, RichBlockTypeSlideshow:
		return validateBlocks(b.Blocks)
	case RichBlockTypePullQuotation:
		return validateRichText("text", b.Text)
	case RichBlockTypeTable:
		if len(b.Cells) == 0 {
			return errors.New("cells are required")
		}
		for rowIndex := range b.Cells {
			if len(b.Cells[rowIndex]) == 0 {
				return fmt.Errorf("cells[%d] must not be empty", rowIndex)
			}
			for columnIndex := range b.Cells[rowIndex] {
				cell := b.Cells[rowIndex][columnIndex]
				if cell.Text != nil {
					if err := cell.Text.Validate(); err != nil {
						return fmt.Errorf("cells[%d][%d].text: %w", rowIndex, columnIndex, err)
					}
				}
				if cell.Colspan < 0 || cell.Rowspan < 0 {
					return fmt.Errorf("cells[%d][%d] spans can't be negative", rowIndex, columnIndex)
				}
				switch cell.Align {
				case "", "left", "center", "right":
				default:
					return fmt.Errorf("cells[%d][%d] has invalid align %q", rowIndex, columnIndex, cell.Align)
				}
				switch cell.Valign {
				case "", "top", "middle", "bottom":
				default:
					return fmt.Errorf("cells[%d][%d] has invalid valign %q", rowIndex, columnIndex, cell.Valign)
				}
			}
		}
	case RichBlockTypeDetails:
		if err := validateRichText("summary", b.Summary); err != nil {
			return err
		}
		return validateBlocks(b.Blocks)
	case RichBlockTypeMap:
		if b.Location == nil {
			return errors.New("location is required")
		}
		if b.Zoom < 0 || b.Zoom > 24 {
			return errors.New("zoom must be between 0 and 24")
		}
		if b.Width < 0 || b.Width > 10000 || b.Height < 0 || b.Height > 10000 || b.Width+b.Height > 10000 {
			return errors.New("map width and height must each be 0-10000 and total at most 10000")
		}
		if b.Width > 0 && b.Height > 0 && (b.Width > b.Height*20 || b.Height > b.Width*20) {
			return errors.New("map width/height ratio must be at most 20")
		}
	case RichBlockTypeAnimation:
		if b.Animation == nil {
			return errors.New("animation is required")
		}
		return validateTypedMedia(b.Animation.Type, "animation", b.Animation.Media)
	case RichBlockTypeAudio:
		if b.Audio == nil {
			return errors.New("audio is required")
		}
		return validateTypedMedia(b.Audio.Type, "audio", b.Audio.Media)
	case RichBlockTypePhoto:
		if b.Photo == nil {
			return errors.New("photo is required")
		}
		return validateTypedMedia(b.Photo.Type, "photo", b.Photo.Media)
	case RichBlockTypeVideo:
		if b.Video == nil {
			return errors.New("video is required")
		}
		return validateTypedMedia(b.Video.Type, "video", b.Video.Media)
	case RichBlockTypeVoiceNote:
		if b.VoiceNote == nil {
			return errors.New("voice_note is required")
		}
		return validateTypedMedia(b.VoiceNote.Type, "voice_note", b.VoiceNote.Media)
	case RichBlockTypeThinking:
		if !allowThinking {
			return errors.New("thinking blocks are only valid in sendRichMessageDraft")
		}
		return validateRichText("text", b.Text)
	default:
		return fmt.Errorf("unknown input rich block type %q", b.Type)
	}
	if b.Caption != nil {
		if err := b.Caption.Text.Validate(); err != nil {
			return fmt.Errorf("caption.text: %w", err)
		}
		if b.Caption.Credit != nil {
			if err := b.Caption.Credit.Validate(); err != nil {
				return fmt.Errorf("caption.credit: %w", err)
			}
		}
	}
	if b.TableCaption != nil {
		if err := b.TableCaption.Validate(); err != nil {
			return fmt.Errorf("caption: %w", err)
		}
	}
	if b.Credit != nil {
		if err := b.Credit.Validate(); err != nil {
			return fmt.Errorf("credit: %w", err)
		}
	}
	return nil
}
