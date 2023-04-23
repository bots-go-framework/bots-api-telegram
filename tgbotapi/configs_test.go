package tgbotapi

import "testing"

func TestConstants(t *testing.T) {
	_ = ModeMarkdown
	_ = ModeHTML

	_ = ChatTyping
	_ = ChatUploadPhoto
	_ = ChatRecordVideo
	_ = ChatUploadVideo
	_ = ChatRecordAudio
	_ = ChatUploadAudio
	_ = ChatUploadDocument
	_ = ChatFindLocation
}
