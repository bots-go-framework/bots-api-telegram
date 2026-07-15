package tgbotapi

import (
	"encoding/json"
	"fmt"
)

// InputRichBlockListItem represents an item of an InputRichBlockList to be sent (Bot API 10.2 Rich
// Messages).
//
// https://core.telegram.org/bots/api#inputrichblocklistitem
type InputRichBlockListItem struct {
	// The content of the item
	Blocks []InputRichBlock `json:"blocks"`

	// Optional. Pass True if the item has a checkbox
	HasCheckbox bool `json:"has_checkbox,omitempty"`

	// Optional. Pass True if the item has a checked checkbox
	IsChecked bool `json:"is_checked,omitempty"`

	// Optional. For ordered lists, the numeric value of the item label
	Value int `json:"value,omitempty"`

	// Optional. For ordered lists, the type of the item label; must be one of "a" for lowercase letters,
	// "A" for uppercase letters, "i" for lowercase Roman numerals, "I" for uppercase Roman numerals, or
	// "1" for decimal numbers
	Type string `json:"type,omitempty"`
}

// InputRichBlock represents a block in a rich formatted message to be sent (Bot API 10.2 Rich
// Messages). It is the send-side mirror of RichBlock and is a union over InputRichBlockParagraph,
// InputRichBlockSectionHeading, InputRichBlockPreformatted, InputRichBlockFooter, InputRichBlockDivider,
// InputRichBlockMathematicalExpression, InputRichBlockAnchor, InputRichBlockList,
// InputRichBlockBlockQuotation, InputRichBlockPullQuotation, InputRichBlockCollage,
// InputRichBlockSlideshow, InputRichBlockTable, InputRichBlockDetails, InputRichBlockMap,
// InputRichBlockAnimation, InputRichBlockAudio, InputRichBlockPhoto, InputRichBlockVideo,
// InputRichBlockVoiceNote, InputRichBlockThinking, following the same single-flattened-struct-with-
// Type-discriminator convention used by RichBlock (see rich_block.go). The Type discriminator values
// are shared with RichBlock; reuse the RichBlockType* constants (RichBlockTypeParagraph,
// RichBlockTypeSectionHeading, etc.) when constructing an InputRichBlock.
//
// Being outgoing-only, InputRichBlock does not need a custom UnmarshalJSON. It does need a custom
// MarshalJSON for the same reason as RichBlock: the "caption" field is a RichBlockCaption for every
// captioned block except InputRichBlockTable, whose "caption" is a plain RichText. See Caption and
// TableCaption below.
//
// https://core.telegram.org/bots/api#inputrichblock
type InputRichBlock struct {
	// Type of the block, e.g. "paragraph", "heading", "list", etc. Reuse the RichBlockType* constants.
	Type string `json:"type,omitempty"`

	// Text: for InputRichBlockParagraph, InputRichBlockSectionHeading, InputRichBlockPreformatted,
	// InputRichBlockFooter, InputRichBlockPullQuotation, InputRichBlockThinking - text of the block.
	Text *RichText `json:"text,omitempty"`

	// Size: for InputRichBlockSectionHeading, the relative size of the text font; 1-6, 1 is the largest,
	// 6 is the smallest.
	Size int `json:"size,omitempty"`

	// Language: for InputRichBlockPreformatted, the programming language of the text.
	Language string `json:"language,omitempty"`

	// Expression: for InputRichBlockMathematicalExpression, the mathematical expression in LaTeX format.
	Expression string `json:"expression,omitempty"`

	// Name: for InputRichBlockAnchor, the name of the anchor.
	Name string `json:"name,omitempty"`

	// Items: for InputRichBlockList, the items of the list.
	Items []InputRichBlockListItem `json:"items,omitempty"`

	// Blocks: for InputRichBlockBlockQuotation, InputRichBlockCollage, InputRichBlockSlideshow,
	// InputRichBlockDetails - content/elements of the block.
	Blocks []InputRichBlock `json:"blocks,omitempty"`

	// Credit: for InputRichBlockBlockQuotation, InputRichBlockPullQuotation - credit of the block.
	Credit *RichText `json:"credit,omitempty"`

	// Caption: for InputRichBlockCollage, InputRichBlockSlideshow, InputRichBlockMap,
	// InputRichBlockAnimation, InputRichBlockAudio, InputRichBlockPhoto, InputRichBlockVideo,
	// InputRichBlockVoiceNote - caption of the block. Not used by InputRichBlockTable, which uses
	// TableCaption instead (see type doc comment).
	Caption *RichBlockCaption `json:"-"`

	// TableCaption: for InputRichBlockTable only, the caption of the table. See Caption for every other
	// captioned block.
	TableCaption *RichText `json:"-"`

	// Cells: for InputRichBlockTable, the cells of the table.
	Cells [][]RichBlockTableCell `json:"cells,omitempty"`

	// IsBordered: for InputRichBlockTable, pass True if the table has borders.
	IsBordered bool `json:"is_bordered,omitempty"`

	// IsStriped: for InputRichBlockTable, pass True if the table is striped.
	IsStriped bool `json:"is_striped,omitempty"`

	// Summary: for InputRichBlockDetails, the always-shown summary of the block.
	Summary *RichText `json:"summary,omitempty"`

	// IsOpen: for InputRichBlockDetails, pass True if the content of the block is visible by default.
	IsOpen bool `json:"is_open,omitempty"`

	// Location: for InputRichBlockMap, the location of the center of the map.
	Location *Location `json:"location,omitempty"`

	// Zoom: for InputRichBlockMap, the map zoom level; 0-24.
	Zoom int `json:"zoom,omitempty"`

	// Width: for InputRichBlockMap, the map width; 0-10000.
	Width int `json:"width,omitempty"`

	// Height: for InputRichBlockMap, the map height; 0-10000.
	Height int `json:"height,omitempty"`

	// Animation: for InputRichBlockAnimation, the animation. Its Caption/ParseMode/CaptionEntities/
	// ShowCaptionAboveMedia fields are ignored; use Caption on this struct instead.
	Animation *InputMediaAnimation `json:"animation,omitempty"`

	// Audio: for InputRichBlockAudio, the audio. Its Caption/ParseMode/CaptionEntities fields are
	// ignored; use Caption on this struct instead.
	Audio *InputMediaAudio `json:"audio,omitempty"`

	// Photo: for InputRichBlockPhoto, the photo. Its Caption/ParseMode/CaptionEntities/
	// ShowCaptionAboveMedia fields are ignored; use Caption on this struct instead.
	Photo *InputMediaPhoto `json:"photo,omitempty"`

	// Video: for InputRichBlockVideo, the video. Its Caption/ParseMode/CaptionEntities/
	// ShowCaptionAboveMedia fields are ignored; use Caption on this struct instead.
	Video *InputMediaVideo `json:"video,omitempty"`

	// VoiceNote: for InputRichBlockVoiceNote, the voice note. Its Caption/ParseMode/CaptionEntities
	// fields are ignored; use Caption on this struct instead.
	VoiceNote *InputMediaVoiceNote `json:"voice_note,omitempty"`
}

// inputRichBlockAlias has the same fields as InputRichBlock but none of its methods, used to avoid
// infinite recursion in InputRichBlock's custom MarshalJSON.
type inputRichBlockAlias InputRichBlock

// MarshalJSON projects Caption/TableCaption onto the single "caption" wire field, based on Type. See
// RichBlock.MarshalJSON for the receive-side counterpart of this projection.
//
//goland:noinspection GoMixedReceiverTypes
func (b InputRichBlock) MarshalJSON() ([]byte, error) {
	raw, err := json.Marshal(inputRichBlockAlias(b))
	if err != nil {
		return nil, err
	}

	var captionRaw json.RawMessage
	switch {
	case b.Type == RichBlockTypeTable && b.TableCaption != nil:
		if captionRaw, err = json.Marshal(b.TableCaption); err != nil {
			return nil, fmt.Errorf("failed to marshal InputRichBlock.TableCaption: %w", err)
		}
	case b.Type != RichBlockTypeTable && b.Caption != nil:
		if captionRaw, err = json.Marshal(b.Caption); err != nil {
			return nil, fmt.Errorf("failed to marshal InputRichBlock.Caption: %w", err)
		}
	}
	if captionRaw == nil {
		return raw, nil
	}

	var m map[string]json.RawMessage
	if err = json.Unmarshal(raw, &m); err != nil {
		return nil, err
	}
	m["caption"] = captionRaw
	return json.Marshal(m)
}
