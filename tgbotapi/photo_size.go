package tgbotapi

//go:generate ffjson $GOFILE

import "fmt"

// PhotoSize contains information about photos.
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id,omitempty"` // optional

	Width    int `json:"width,omitempty"`     // Photo width
	Height   int `json:"height,omitempty"`    // Photo height
	FileSize int `json:"file_size,omitempty"` // Optional. File size in bytes
}

func (v *PhotoSize) String() string {
	return fmt.Sprintf("%s@%dx%d:%dbytes", v.FileID, v.Width, v.Height, v.FileSize)
}
