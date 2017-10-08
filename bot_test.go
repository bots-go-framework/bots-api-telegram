package tgbotapi

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

const (
	TestToken              = "153667468:AAHlSHlMqSt1f_uFmVRJbm5gntu2HI4WW8I"
	ChatID                 = 76918703
	ReplyToMessageID       = 35
	ExistingPhotoFileID    = "AgADAgADw6cxG4zHKAkr42N7RwEN3IFShCoABHQwXEtVks4EH2wBAAEC"
	ExistingDocumentFileID = "BQADAgADOQADjMcoCcioX1GrDvp3Ag"
	ExistingAudioFileID    = "BQADAgADRgADjMcoCdXg3lSIN49lAg"
	ExistingVoiceFileID    = "AwADAgADWQADjMcoCeul6r_q52IyAg"
	ExistingVideoFileID    = "BAADAgADZgADjMcoCav432kYe0FRAg"
	ExistingStickerFileID  = "BQADAgADcwADjMcoCbdl-6eB--YPAg"
)

func getBot(t *testing.T) (*BotAPI) {
	bot := NewBotAPI(TestToken)

	if bot == nil {
		t.Fail()
	}

	return bot
}

func TestNewBotAPI_notoken(t *testing.T) {
	botApi := NewBotAPI("")

	if botApi == nil {
		t.Fail()
	}
}

func TestGetUpdates(t *testing.T) {
	bot := getBot(t)

	u := NewUpdate(0)

	_, err := bot.GetUpdates(u)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithMessage(t *testing.T) {
	bot := getBot(t)

	msg := NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ParseMode = "markdown"
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithMessageReply(t *testing.T) {
	bot := getBot(t)

	msg := NewMessage(ChatID, "A test message from the test library in telegram-bot-api")
	msg.ReplyToMessageID = ReplyToMessageID
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithMessageForward(t *testing.T) {
	bot := getBot(t)

	msg := NewForward(ChatID, ChatID, ReplyToMessageID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewPhoto(t *testing.T) {
	bot := getBot(t)

	msg := NewPhotoUpload(ChatID, "tests/image.jpg")
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewPhotoWithFileBytes(t *testing.T) {
	bot := getBot(t)

	data, _ := ioutil.ReadFile("tests/image.jpg")
	b := FileBytes{Name: "image.jpg", Bytes: data}

	msg := NewPhotoUpload(ChatID, b)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewPhotoWithFileReader(t *testing.T) {
	bot := getBot(t)

	f, _ := os.Open("tests/image.jpg")
	reader := FileReader{Name: "image.jpg", Reader: f, Size: -1}

	msg := NewPhotoUpload(ChatID, reader)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewPhotoReply(t *testing.T) {
	bot := getBot(t)

	msg := NewPhotoUpload(ChatID, "tests/image.jpg")
	msg.ReplyToMessageID = ReplyToMessageID

	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingPhoto(t *testing.T) {
	bot := getBot(t)

	msg := NewPhotoShare(ChatID, ExistingPhotoFileID)
	msg.Caption = "Test"
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewDocument(t *testing.T) {
	bot := getBot(t)

	msg := NewDocumentUpload(ChatID, "tests/image.jpg")
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingDocument(t *testing.T) {
	bot := getBot(t)

	msg := NewDocumentShare(ChatID, ExistingDocumentFileID)
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewAudio(t *testing.T) {
	bot := getBot(t)

	msg := NewAudioUpload(ChatID, "tests/audio.mp3")
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"
	msg.MimeType = "audio/mpeg"
	msg.FileSize = 688
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingAudio(t *testing.T) {
	bot := getBot(t)

	msg := NewAudioShare(ChatID, ExistingAudioFileID)
	msg.Title = "TEST"
	msg.Duration = 10
	msg.Performer = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewVoice(t *testing.T) {
	bot := getBot(t)

	msg := NewVoiceUpload(ChatID, "tests/voice.ogg")
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingVoice(t *testing.T) {
	bot := getBot(t)

	msg := NewVoiceShare(ChatID, ExistingVoiceFileID)
	msg.Duration = 10
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithLocation(t *testing.T) {
	bot := getBot(t)

	_, err := bot.Send(NewLocation(ChatID, 40, 40))

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewVideo(t *testing.T) {
	bot := getBot(t)

	msg := NewVideoUpload(ChatID, "tests/video.mp4")
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingVideo(t *testing.T) {
	bot := getBot(t)

	msg := NewVideoShare(ChatID, ExistingVideoFileID)
	msg.Duration = 10
	msg.Caption = "TEST"

	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewSticker(t *testing.T) {
	bot := getBot(t)

	msg := NewStickerUpload(ChatID, "tests/image.jpg")

	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingSticker(t *testing.T) {
	bot := getBot(t)

	msg := NewStickerShare(ChatID, ExistingStickerFileID)

	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithNewStickerAndKeyboardHide(t *testing.T) {
	bot := getBot(t)

	msg := NewStickerUpload(ChatID, "tests/image.jpg")
	msg.ReplyMarkup = ReplyKeyboardHide{true, false}
	_, err := bot.Send(msg)

	if err != nil {
		t.Fail()
	}
}

func TestSendWithExistingStickerAndKeyboardHide(t *testing.T) {
	bot := getBot(t)

	msg := NewStickerShare(ChatID, ExistingStickerFileID)
	msg.ReplyMarkup = ReplyKeyboardHide{true, false}

	_, err := bot.Send(msg)

	if err != nil {

		t.Fail()
	}
}

func TestGetFile(t *testing.T) {
	bot := getBot(t)

	file := FileConfig{ExistingPhotoFileID}

	_, err := bot.GetFile(file)

	if err != nil {
		t.Fail()
	}
}

func TestSendChatConfig(t *testing.T) {
	bot := getBot(t)

	_, err := bot.Send(NewChatAction(ChatID, ChatTyping))

	if err != nil {
		t.Fail()
	}
}

func TestGetUserProfilePhotos(t *testing.T) {
	bot := getBot(t)

	_, err := bot.GetUserProfilePhotos(NewUserProfilePhotos(ChatID))
	if err != nil {
		t.Fail()
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

	var ucfg UpdateConfig = NewUpdate(0)
	ucfg.Timeout = 60
	_, err := bot.GetUpdatesChan(ucfg)

	if err != nil {
		t.Fail()
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

		bot.Send(msg)
	}
}

func ExampleNewWebhook() {
	bot := NewBotAPI("MyAwesomeBotToken")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err := bot.SetWebhook(NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
