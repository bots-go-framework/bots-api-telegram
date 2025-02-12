package tgbotapi

//go:generate ffjson $GOFILE

// LinkPreviewOptions Describes the options used for link preview generation.
type LinkPreviewOptions struct {

	// Optional. True, if the link preview is disabled.
	IsDisabled bool `json:"is_disabled,omitempty"`

	// Optional. URL to use for the link preview.
	// If empty, then the first URL found in the message text will be used.
	Url string `json:"url,omitempty"`

	// Optional. True, if the media in the link preview is supposed to be shrunk;
	// ignored if the URL isn't explicitly specified or media size change isn't supported for the preview.
	PreferSmallMedia bool `json:"prefer_small_media,omitempty"`

	// Optional. True, if the media in the link preview is supposed to be enlarged;
	// ignored if the URL isn't explicitly specified or media size change isn't supported for the preview.
	PreferLargeMedia bool `json:"prefer_large_media,omitempty"`

	// Optional. True, if the link preview must be shown above the message text;
	// otherwise, the link preview will be shown below the message text
	ShowAboveText bool `json:"show_above_text,omitempty"`
}
