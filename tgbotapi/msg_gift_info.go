package tgbotapi

type Gift struct {
	ID               string  `json:"id"`                           // Unique identifier of the gift
	Sticker          Sticker `json:"sticker"`                      // The sticker that represents the gift
	StarCount        int     `json:"star_count"`                   // The number of Telegram Stars that must be paid to send the sticker
	UpgradeStarCount int     `json:"upgrade_star_count,omitempty"` // Optional. The number of Telegram Stars that must be paid to upgrade the gift to a unique one
	TotalCount       int     `json:"total_count,omitempty"`        // Optional. The total number of the gifts of this type that can be sent; for limited gifts only
	RemainingCount   int     `json:"remaining_count,omitempty"`    // Optional. The number of remaining gifts of this type that can be sent; for limited gifts only
}

type GiftInfo struct {
	Gift Gift `json:"gift"` // Information about the gift

	OwnedGiftID string `json:"owned_gift_id,omitempty"` // Optional. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts

	ConvertStarCount        int `json:"convert_star_count,omitempty"`         // Optional. Number of Telegram Stars that can be claimed by the receiver by converting the gift; omitted if conversion to Telegram Stars is impossible
	PrepaidUpgradeStarCount int `json:"prepaid_upgrade_star_count,omitempty"` // Optional. Number of Telegram Stars that were prepaid by the sender for the ability to upgrade the gift

	CanBeUpgraded bool `json:"can_be_upgraded,omitempty"` // Optional. True, if the gift can be upgraded to a unique gift

	Text string `json:"text,omitempty"` // Optional. Text of the message that was added to the gift

	Entities []MessageEntity `json:"entities,omitempty"` // Optional. Special entities that appear in the text

	IsPrivate bool `json:"is_private,omitempty"` // Optional. True, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them
}

// UniqueGift
// https://core.telegram.org/bots/api#uniquegift
type UniqueGift struct {
	BaseName string             `json:"base_name"` // Human-readable name of the regular gift from which this unique gift was upgraded
	Name     string             `json:"name"`      // Unique name of the gift. This name can be used in https://t.me/nft/... links and story areas
	Number   int                `json:"number"`    // Unique number of the upgraded gift among gifts upgraded from the same regular gift
	Model    UniqueGiftModel    `json:"model"`     // Model of the gift
	Symbol   UniqueGiftSymbol   `json:"symbol"`    // Symbol of the gift
	Backdrop UniqueGiftBackdrop `json:"backdrop"`  // Backdrop of the gift
}

// UniqueGiftModel
// https://core.telegram.org/bots/api#uniquegiftmodel
type UniqueGiftModel struct {
	Name           string  `json:"name"`             // Name of the model
	Sticker        Sticker `json:"sticker"`          // The sticker that represents the unique gift
	RarityPerMille int     `json:"rarity_per_mille"` // The number of unique gifts that receive this model for every 1000 gifts upgraded
}

// UniqueGiftSymbol
// https://core.telegram.org/bots/api#uniquegiftsymbol
type UniqueGiftSymbol struct {
	Name           string  `json:"name"`             // Name of the symbol
	Sticker        Sticker `json:"sticker"`          // The sticker that represents the unique gift
	RarityPerMille int     `json:"rarity_per_mille"` // The number of unique gifts that receive this model for every 1000 gifts upgraded
}

type GiftOrigin string

const (
	GiftOriginUpgrade  GiftOrigin = "upgrade"
	GiftOriginTransfer GiftOrigin = "transfer"
)

type UniqueGiftInfo struct {
	Gift UniqueGift `json:"gift"` // Information about the gift

	// Origin of the gift. Currently, either “upgrade” or “transfer”
	Origin GiftOrigin `json:"origin,omitempty"`

	// Optional. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts
	OwnedGiftID string `json:"owned_gift_id,omitempty"`

	// Optional. Number of Telegram Stars that must be paid to transfer the gift; omitted if the bot cannot transfer the gift
	TransferStarCount int `json:"transfer_star_count,omitempty"`
}

// UniqueGiftBackdrop
// https://core.telegram.org/bots/api#uniquegiftbackdrop
type UniqueGiftBackdrop struct {
	Name   string                   `json:"name"` // Name of the backdrop
	Colors UniqueGiftBackdropColors // Colors of the backdrop

	// The number of unique gifts that receive this backdrop for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille,omitempty"`
}

// UniqueGiftBackdropColors
// https://core.telegram.org/bots/api#uniquegiftbackdropcolors
type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"` // The color in the center of the backdrop in RGB format
	EdgeColor   int `json:"edge_color"`   // The color on the edges of the backdrop in RGB format
	SymbolColor int `json:"symbol_color"` // The color to be applied to the symbol in RGB format
	TextColor   int `json:"text_color"`   // The color for the text on the backdrop in RGB format
}

type OwnedGiftType string

const (
	OwnedGiftTypeRegular = OwnedGiftType("regular")
	OwnedGiftTypeUnique  = OwnedGiftType("unique")
)

type OwnedGift interface {
	GetType() OwnedGiftType
	GetOwnedGiftID() string
	GetSenderUser() *User
	GetSendDate() int
	GetIsSaved() bool
}

// ownedGift describes a gift received and owned by a user or a chat. Currently, it can be one of
// - OwnedGiftRegular
// - OwnedGiftUnique
type ownedGift struct {
	Type        OwnedGiftType `json:"type"`                    // Type of the gift, always “regular” for OwnedGiftRegular or "unique" for OwnedGiftUnique
	OwnedGiftID string        `json:"owned_gift_id,omitempty"` // Optional. Unique identifier of the gift for the bot; for gifts received on behalf of business accounts only
	SenderUser  *User         `json:"sender_user,omitempty"`   // Optional. Sender of the gift if it is a known user
	SendDate    int           `json:"send_date,omitempty"`     // Date the gift was sent in Unix time
	IsSaved     bool          `json:"is_saved,omitempty"`      // Optional. True, if the gift is displayed on the account's profile page; for gifts received on behalf of business accounts only
}

func (v *ownedGift) GetOwnedGiftID() string {
	return v.OwnedGiftID
}

func (v *ownedGift) GetSenderUser() *User {
	return v.SenderUser
}

func (v *ownedGift) GetSendDate() int {
	return v.SendDate
}

func (v *ownedGift) GetIsSaved() bool {
	return v.IsSaved
}

var _ OwnedGift = (*OwnedGiftRegular)(nil)

// OwnedGiftRegular
// https://core.telegram.org/bots/api#ownedgiftregular
type OwnedGiftRegular struct {
	ownedGift
	Text                    string          `json:"text,omitempty"`                       // Optional. Text of the message that was added to the gift
	Entities                []MessageEntity `json:"entities,omitempty"`                   // Optional. Special entities that appear in the text
	IsPrivate               bool            `json:"is_private,omitempty"`                 // Optional. True, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them
	CanBeUpgraded           bool            `json:"can_be_upgraded,omitempty"`            // Optional. True, if the gift can be upgraded to a unique gift; for gifts received on behalf of business accounts only
	WasRefunded             bool            `json:"was_refunded,omitempty"`               // Optional. True, if the gift was refunded and isn't available anymore
	ConvertStarCount        int             `json:"convert_star_count,omitempty"`         // Optional. Number of Telegram Stars that can be claimed by the receiver instead of the gift; omitted if the gift cannot be converted to Telegram Stars
	PrepaidUpgradeStarCount int             `json:"prepaid_upgrade_star_count,omitempty"` // Optional. Number of Telegram Stars that were paid by the sender for the ability to upgrade the gift
}

func (*OwnedGiftRegular) GetType() OwnedGiftType {
	return OwnedGiftTypeRegular
}

var _ OwnedGift = (*OwnedGiftUnique)(nil)

// OwnedGiftUnique
// https://core.telegram.org/bots/api#ownedgiftunique
type OwnedGiftUnique struct {
	ownedGift
}

func (*OwnedGiftUnique) GetType() OwnedGiftType {
	return OwnedGiftTypeUnique
}

type OwnedGifts struct {
	TotalCount int         `json:"total_count"`           // The total number of gifts owned by the user or the chat
	Gifts      []OwnedGift `json:"gifts"`                 // The list of gifts
	NextOffset string      `json:"next_offset,omitempty"` // Optional. Offset for the next request. If empty, then there are no more results
}
