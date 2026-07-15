package tgbotapi

import (
	"encoding/json"
	"fmt"
)

// Rich block type discriminators for RichBlock.Type.
// https://core.telegram.org/bots/api#richblock
const (
	RichBlockTypeParagraph              = "paragraph"
	RichBlockTypeSectionHeading         = "heading"
	RichBlockTypePreformatted           = "pre"
	RichBlockTypeFooter                 = "footer"
	RichBlockTypeDivider                = "divider"
	RichBlockTypeMathematicalExpression = "mathematical_expression"
	RichBlockTypeAnchor                 = "anchor"
	RichBlockTypeList                   = "list"
	RichBlockTypeBlockQuotation         = "blockquote"
	RichBlockTypePullQuotation          = "pullquote"
	RichBlockTypeCollage                = "collage"
	RichBlockTypeSlideshow              = "slideshow"
	RichBlockTypeTable                  = "table"
	RichBlockTypeDetails                = "details"
	RichBlockTypeMap                    = "map"
	RichBlockTypeAnimation              = "animation"
	RichBlockTypeAudio                  = "audio"
	RichBlockTypePhoto                  = "photo"
	RichBlockTypeVideo                  = "video"
	RichBlockTypeVoiceNote              = "voice_note"
	RichBlockTypeThinking               = "thinking"
)

// RichBlockCaption describes the caption of a rich formatted block (Bot API 10.1 Rich Messages).
//
// https://core.telegram.org/bots/api#richblockcaption
type RichBlockCaption struct {
	// Block caption
	Text RichText `json:"text"`

	// Optional. Block credit which corresponds to the HTML tag <cite>
	Credit *RichText `json:"credit,omitempty"`
}

// RichBlockTableCell represents a cell in a RichBlockTable (Bot API 10.1 Rich Messages).
//
// https://core.telegram.org/bots/api#richblocktablecell
type RichBlockTableCell struct {
	// Optional. Text in the cell. If omitted, the cell is invisible.
	Text *RichText `json:"text,omitempty"`

	// Optional. True, if the cell is a header cell
	IsHeader bool `json:"is_header,omitempty"`

	// Optional. The number of columns the cell spans if it is bigger than 1
	Colspan int `json:"colspan,omitempty"`

	// Optional. The number of rows the cell spans if it is bigger than 1
	Rowspan int `json:"rowspan,omitempty"`

	// Horizontal cell content alignment. Currently, must be one of "left", "center", or "right".
	Align string `json:"align,omitempty"`

	// Vertical cell content alignment. Currently, must be one of "top", "middle", or "bottom".
	Valign string `json:"valign,omitempty"`
}

// RichBlockListItem represents an item of a RichBlockList (Bot API 10.1 Rich Messages).
//
// https://core.telegram.org/bots/api#richblocklistitem
type RichBlockListItem struct {
	// Label of the item
	Label string `json:"label"`

	// The content of the item
	Blocks []RichBlock `json:"blocks"`

	// Optional. True, if the item has a checkbox
	HasCheckbox bool `json:"has_checkbox,omitempty"`

	// Optional. True, if the item has a checked checkbox
	IsChecked bool `json:"is_checked,omitempty"`

	// Optional. For ordered lists, the numeric value of the item label
	Value int `json:"value,omitempty"`

	// Optional. For ordered lists, the type of the item label; must be one of "a" for lowercase letters,
	// "A" for uppercase letters, "i" for lowercase Roman numerals, "I" for uppercase Roman numerals, or
	// "1" for decimal numbers
	Type string `json:"type,omitempty"`
}

// RichBlock represents a block in a rich formatted message (Bot API 10.1 Rich Messages). It is a union
// over RichBlockParagraph, RichBlockSectionHeading, RichBlockPreformatted, RichBlockFooter,
// RichBlockDivider, RichBlockMathematicalExpression, RichBlockAnchor, RichBlockList,
// RichBlockBlockQuotation, RichBlockPullQuotation, RichBlockCollage, RichBlockSlideshow, RichBlockTable,
// RichBlockDetails, RichBlockMap, RichBlockAnimation, RichBlockAudio, RichBlockPhoto, RichBlockVideo,
// RichBlockVoiceNote, RichBlockThinking, following the same single-flattened-struct-with-Type-
// discriminator convention used elsewhere in this package for unions (see PaidMedia, PollMedia).
//
// The "caption" field is the one exception: RichBlockTable's caption is a RichText, while every other
// captioned block's caption is a RichBlockCaption. Since a single Go field can't carry both wire shapes
// under the same "caption" JSON key, this struct keeps them as two distinct fields (Caption and
// TableCaption) and a custom MarshalJSON/UnmarshalJSON pair projects whichever one applies, based on
// Type, onto the single "caption" key on the wire.
//
// https://core.telegram.org/bots/api#richblock
type RichBlock struct {
	// Type of the block, e.g. "paragraph", "heading", "list", etc.
	Type string `json:"type,omitempty"`

	// Text: for RichBlockParagraph, RichBlockSectionHeading, RichBlockPreformatted, RichBlockFooter,
	// RichBlockPullQuotation, RichBlockThinking - text of the block.
	Text *RichText `json:"text,omitempty"`

	// Size: for RichBlockSectionHeading, the relative size of the text font; 1-6, 1 is the largest,
	// 6 is the smallest.
	Size int `json:"size,omitempty"`

	// Language: for RichBlockPreformatted, the programming language of the text.
	Language string `json:"language,omitempty"`

	// Expression: for RichBlockMathematicalExpression, the mathematical expression in LaTeX format.
	Expression string `json:"expression,omitempty"`

	// Name: for RichBlockAnchor, the name of the anchor.
	Name string `json:"name,omitempty"`

	// Items: for RichBlockList, the items of the list.
	Items []RichBlockListItem `json:"items,omitempty"`

	// Blocks: for RichBlockBlockQuotation, RichBlockCollage, RichBlockSlideshow, RichBlockDetails -
	// content/elements of the block.
	Blocks []RichBlock `json:"blocks,omitempty"`

	// Credit: for RichBlockBlockQuotation, RichBlockPullQuotation - credit of the block.
	Credit *RichText `json:"credit,omitempty"`

	// Caption: for RichBlockCollage, RichBlockSlideshow, RichBlockMap, RichBlockAnimation, RichBlockAudio,
	// RichBlockPhoto, RichBlockVideo, RichBlockVoiceNote - caption of the block. Not used by
	// RichBlockTable, which uses TableCaption instead (see type doc comment).
	Caption *RichBlockCaption `json:"-"`

	// TableCaption: for RichBlockTable only, the caption of the table. See Caption for every other
	// captioned block.
	TableCaption *RichText `json:"-"`

	// Cells: for RichBlockTable, the cells of the table.
	Cells [][]RichBlockTableCell `json:"cells,omitempty"`

	// IsBordered: for RichBlockTable, true if the table has borders.
	IsBordered bool `json:"is_bordered,omitempty"`

	// IsStriped: for RichBlockTable, true if the table is striped.
	IsStriped bool `json:"is_striped,omitempty"`

	// Summary: for RichBlockDetails, the always-shown summary of the block.
	Summary *RichText `json:"summary,omitempty"`

	// IsOpen: for RichBlockDetails, true if the content of the block is visible by default.
	IsOpen bool `json:"is_open,omitempty"`

	// Location: for RichBlockMap, the location of the center of the map.
	Location *Location `json:"location,omitempty"`

	// Zoom: for RichBlockMap, the map zoom level; 13-20.
	Zoom int `json:"zoom,omitempty"`

	// Width: for RichBlockMap, the expected width of the map.
	Width int `json:"width,omitempty"`

	// Height: for RichBlockMap, the expected height of the map.
	Height int `json:"height,omitempty"`

	// Animation: for RichBlockAnimation, the animation.
	Animation *Animation `json:"animation,omitempty"`

	// HasSpoiler: for RichBlockAnimation, RichBlockPhoto, RichBlockVideo - true if the media preview is
	// covered by a spoiler animation.
	HasSpoiler bool `json:"has_spoiler,omitempty"`

	// Audio: for RichBlockAudio, the audio.
	Audio *Audio `json:"audio,omitempty"`

	// Photo: for RichBlockPhoto, the available sizes of the photo.
	Photo []PhotoSize `json:"photo,omitempty"`

	// Video: for RichBlockVideo, the video.
	Video *Video `json:"video,omitempty"`

	// VoiceNote: for RichBlockVoiceNote, the voice note.
	VoiceNote *Voice `json:"voice_note,omitempty"`
}

// richBlockAlias has the same fields as RichBlock but none of its methods, used to avoid infinite
// recursion in RichBlock's custom MarshalJSON/UnmarshalJSON.
type richBlockAlias RichBlock

// MarshalJSON projects Caption/TableCaption onto the single "caption" wire field, based on Type.
//
//goland:noinspection GoMixedReceiverTypes
func (b RichBlock) MarshalJSON() ([]byte, error) {
	raw, err := json.Marshal(richBlockAlias(b))
	if err != nil {
		return nil, err
	}

	var captionRaw json.RawMessage
	switch {
	case b.Type == RichBlockTypeTable && b.TableCaption != nil:
		if captionRaw, err = json.Marshal(b.TableCaption); err != nil {
			return nil, fmt.Errorf("failed to marshal RichBlock.TableCaption: %w", err)
		}
	case b.Type != RichBlockTypeTable && b.Caption != nil:
		if captionRaw, err = json.Marshal(b.Caption); err != nil {
			return nil, fmt.Errorf("failed to marshal RichBlock.Caption: %w", err)
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

// UnmarshalJSON extracts the "caption" wire field into Caption or TableCaption, based on Type.
//
//goland:noinspection GoMixedReceiverTypes
func (b *RichBlock) UnmarshalJSON(data []byte) error {
	var alias richBlockAlias
	if err := json.Unmarshal(data, &alias); err != nil {
		return fmt.Errorf("failed to unmarshal RichBlock: %w", err)
	}
	*b = RichBlock(alias)

	var probe struct {
		Caption json.RawMessage `json:"caption,omitempty"`
	}
	if err := json.Unmarshal(data, &probe); err != nil {
		return fmt.Errorf("failed to probe RichBlock.caption: %w", err)
	}
	if len(probe.Caption) == 0 || string(probe.Caption) == "null" {
		return nil
	}

	if b.Type == RichBlockTypeTable {
		var rt RichText
		if err := json.Unmarshal(probe.Caption, &rt); err != nil {
			return fmt.Errorf("failed to unmarshal RichBlock.TableCaption: %w", err)
		}
		b.TableCaption = &rt
	} else {
		var rc RichBlockCaption
		if err := json.Unmarshal(probe.Caption, &rc); err != nil {
			return fmt.Errorf("failed to unmarshal RichBlock.Caption: %w", err)
		}
		b.Caption = &rc
	}
	return nil
}
