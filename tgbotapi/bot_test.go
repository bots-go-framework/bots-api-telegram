package tgbotapi

import (
	"log"
	"net/http"
	"os"
	"testing"
)

const (
	ChatID                 = 76918703
	ReplyToMessageID       = 35
	ExistingPhotoFileID    = "AgADAgADw6cxG4zHKAkr42N7RwEN3IFShCoABHQwXEtVks4EH2wBAAEC"
	ExistingDocumentFileID = "BQADAgADOQADjMcoCcioX1GrDvp3Ag"
	ExistingAudioFileID    = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	ExistingVoiceFileID    = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	ExistingVideoFileID    = "BAADAgADZgADjMcoCav432kYe0FRAg"
	ExistingStickerFileID  = "BQADAgADcwADjMcoCbdl-6eB--YPAg"
)

func getBot(t *testing.T) *BotAPI {
	token := os.Getenv("TELEGRAM_TOKEN")

	if token == "" {
		t.Skip("TODO: re-enable!")
		//t.Fatalf("TELEGRAM_TOKEN not set in os environment")
	}

	bot := NewBotAPI(token)

	if bot == nil {
		t.Fatal("bot is nil")
	}

	return bot
}

func TestNewBotAPI_notoken(t *testing.T) {
	defer func() {
		if err := recover(); err == nil {
			t.Errorf("The code did not panic")
		}
	}()
	_ = NewBotAPI("")
}

func TestGetUpdates(t *testing.T) {
	bot := getBot(t)

	u := NewUpdate(0)

	_, err := bot.GetUpdates(u)

	if err != nil {
		t.Errorf("Error getting updates: %s", err)
	}
}

func TestSendWithMessage(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending message: %s", err)
	}
}

func TestSendWithMessageReply(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ReplyToMessageID = ReplyToMessageID
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending message reply: %s", err)
	}
}

func TestSendWithMessageForward(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewForward(ChatID, ChatID, ReplyToMessageID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending message forward: %s", err)
	}
}

func TestSendWithNewPhoto(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewPhotoUpload(ChatID, "tests/image.jpg")
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending with new photo: %s", err)
	}
}

func TestSendWithNewPhotoWithFileBytes(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	data, _ := os.ReadFile("tests/image.jpg")
	b := FileBytes{Name: "image.jpg", Bytes: data}

	msg := NewPhotoUpload(ChatID, b)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending photo with file reply: %s", err)
	}
}

func TestSendWithNewPhotoWithFileReader(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	f, _ := os.Open("tests/image.jpg")
	reader := FileReader{Name: "image.jpg", Reader: f, Size: -1}

	msg := NewPhotoUpload(ChatID, reader)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending photo with file reply: %s", err)
	}
}

func TestSendWithNewPhotoReply(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewPhotoUpload(ChatID, "tests/image.jpg")
	msg.ReplyToMessageID = ReplyToMessageID

	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending photo reply: %s", err)
	}
}

func TestSendWithExistingPhoto(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewPhotoShare(ChatID, ExistingPhotoFileID)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending with existing photo: %s", err)
	}
}

func TestSendWithNewDocument(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewDocumentUpload(ChatID, "tests/image.jpg")
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending document: %s", err)
	}
}

func TestSendWithExistingDocument(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewDocumentShare(ChatID, ExistingDocumentFileID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending document: %s", err)
	}
}

func TestSendWithNewAudio(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewAudioUpload(ChatID, "tests/audio.mp3")
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"
	msg.MimeType = "audio/mpeg"
	msg.FileSize = 688
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending audio: %s", err)
	}
}

func TestSendWithExistingAudio(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewAudioShare(ChatID, ExistingAudioFileID)
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending audio: %s", err)
	}
}

func TestSendWithNewVoice(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewVoiceUpload(ChatID, "tests/voice.ogg")
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending voice: %s", err)
	}
}

func TestSendWithExistingVoice(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewVoiceShare(ChatID, ExistingVoiceFileID)
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending voice: %s", err)
	}
}

func TestSendWithLocation(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	_, err := bot.Send(NewLocation(ChatID, 40, 40))

	if err != nil {
		t.Errorf("Error sending location: %s", err)
	}
}

func TestSendWithNewVideo(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewVideoUpload(ChatID, "tests/video.mp4")
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending video: %s", err)
	}
}

func TestSendWithExistingVideo(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewVideoShare(ChatID, ExistingVideoFileID)
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending existing video: %s", err)
	}
}

func TestSendWithNewSticker(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewStickerUpload(ChatID, "tests/image.jpg")

	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending sticker: %s", err)
	}
}

func TestSendWithExistingSticker(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewStickerShare(ChatID, ExistingStickerFileID)

	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending existing sticker: %s", err)
	}
}

func TestSendWithNewStickerAndKeyboardHide(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewStickerUpload(ChatID, "tests/image.jpg")
	msg.ReplyMarkup = ReplyKeyboardHide{true, false}
	_, err := bot.Send(msg)

	if err != nil {
		t.Errorf("Error sending sticker and keyboard hide: %s", err)
	}
}

func TestSendWithExistingStickerAndKeyboardHide(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	msg := NewStickerShare(ChatID, ExistingStickerFileID)
	msg.ReplyMarkup = ReplyKeyboardHide{true, false}

	_, err := bot.Send(msg)

	if err != nil {

		t.Errorf("Error sending existing sticker and keyboard hide: %s", err)
	}
}

func TestGetFile(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	file := FileConfig{ExistingPhotoFileID}

	_, err := bot.GetFile(file)

	if err != nil {
		t.Errorf("Error getting file: %s", err)
	}
}

func TestSendChatConfig(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	_, err := bot.Send(NewChatAction(ChatID, ChatTyping))

	if err != nil {
		t.Errorf("Error sending chat action: %s", err)
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	t.Skip()
	bot := getBot(t)

	_, err := bot.GetUserProfilePhotos(*NewUserProfilePhotos(ChatID))
	if err != nil {
		t.Errorf("Error getting user profile photos: %s", err)
	}
}

//func TestListenForWebhook(t *testing.T) {
//	bot := getBot(t)
//
//	updates := bot.ListenForWebhook("/")
//
//	req, _ := http.NewRequest("GET", "", strings.NewReader("{}"))
//	w := httptest.NewRecorder()
//
//	handler.ServeHTTP(w, req)
//	if w.Code != http.StatusOK {
//		t.Errorf("Home page didn't return %v", http.StatusOK)
//	}
//}

//func TestSetWebhookWithCert(t *testing.T) {
//	bot := getBot(t)
//
//	bot.RemoveWebhook()
//
//	wh := NewWebhookWithCert("https://example.com/tgbotapi-test/"+bot.Token, "tests/cert.pem")
//	_, err := bot.SetWebhook(wh)
//	if err != nil {
//		t.Fail()
//	}
//
//	bot.RemoveWebhook()
//}
//
//func TestSetWebhookWithoutCert(t *testing.T) {
//	bot := getBot(t)
//
//	bot.RemoveWebhook()
//
//	wh := NewWebhook("https://example.com/tgbotapi-test/" + bot.Token)
//	_, err := bot.SetWebhook(wh)
//	if err != nil {
//		t.Fail()
//	}
//
//	bot.RemoveWebhook()
//}

func TestUpdatesChan(t *testing.T) {
	bot := getBot(t)

	ucfg := NewUpdate(0)
	ucfg.Timeout = 60
	_, err := bot.GetUpdatesChan(ucfg)

	if err != nil {
		t.Errorf("Error getting updates channel: %s", err)
	}
}

func ExampleNewBotAPI() {
	bot := NewBotAPI("MyAwesomeBotToken")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}

func ExampleNewWebhook() {
	bot := NewBotAPI("MyAwesomeBotToken")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err := bot.SetWebhook(*NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go func() {
		err := http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
