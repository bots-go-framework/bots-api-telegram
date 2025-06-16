package tgbotapi

type ShippingAddress struct {
	CountryCode string `json:"country_code"`    // Two-letter ISO 3166-1 alpha-2 country code
	State       string `json:"state,omitempty"` // State, if applicable
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"` // First line for the address
	StreetLine2 string `json:"street_line2"` // Second line for the address
	PostCode    string `json:"post_code"`    // Address post code
}
