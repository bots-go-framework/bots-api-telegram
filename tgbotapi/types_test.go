package tgbotapi

import (
	"testing"
	"time"
)

func TestUserStringWith(t *testing.T) {
	user := User{0, "Test", "Test", "", ""}

	if user.String() != "Test Test" {
		t.Fail()
	}
}

func TestUserStringWithUserName(t *testing.T) {
	user := User{0, "Test", "Test", "@test", ""}

	if user.String() != "@test" {
		t.Fail()
	}
}

func TestMessageTime(t *testing.T) {
	message := Message{Date: 0}

	date := time.Unix(0, 0)
	if message.Time() != date {
		t.Fail()
	}
}

func TestMessageIsCommandWithCommand(t *testing.T) {
	message := Message{Text: "/command"}

	if message.IsCommand() != true {
		t.Fail()
	}
}

func TestIsCommandWithText(t *testing.T) {
	message := Message{Text: "some text"}

	if message.IsCommand() != false {
		t.Fail()
	}
}

func TestIsCommandWithEmptyText(t *testing.T) {
	message := Message{Text: ""}

	if message.IsCommand() != false {
		t.Fail()
	}
}

func TestCommandWithCommand(t *testing.T) {
	message := Message{Text: "/testcommand"}

	if message.Command() != "testcommand" {
		t.Fatal("Expected `/testcommand`, got: " + message.Command())
	}
}

func TestCommandWithEmptyText(t *testing.T) {
	message := Message{Text: ""}

	if message.Command() != "" {
		t.Fail()
	}
}

func TestCommandWithNonCommand(t *testing.T) {
	message := Message{Text: "test text"}

	if message.Command() != "" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsWithArguments(t *testing.T) {
	message := Message{Text: "/command with arguments"}
	if message.CommandArguments() != "with arguments" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsWithoutArguments(t *testing.T) {
	message := Message{Text: "/command"}
	if message.CommandArguments() != "" {
		t.Fail()
	}
}

func TestMessageCommandArgumentsForNonCommand(t *testing.T) {
	message := Message{Text: "test text"}
	if message.CommandArguments() != "" {
		t.Fail()
	}
}

func TestChatIsPrivate(t *testing.T) {
	chat := Chat{ID: 10, Type: "private"}

	if chat.IsPrivate() != true {
		t.Fail()
	}
}

func TestChatIsGroup(t *testing.T) {
	chat := Chat{ID: 10, Type: "group"}

	if chat.IsGroup() != true {
		t.Fail()
	}
}

func TestChatIsChannel(t *testing.T) {
	chat := Chat{ID: 10, Type: "channel"}

	if chat.IsChannel() != true {
		t.Fail()
	}
}

func TestChatIsSuperGroup(t *testing.T) {
	chat := Chat{ID: 10, Type: "supergroup"}

	if !chat.IsSuperGroup() {
		t.Fail()
	}
}

func TestFileLink(t *testing.T) {
	file := File{FilePath: "test/test.txt"}

	if file.Link("token") != "https://api.telegram.org/file/bottoken/test/test.txt" {
		t.Fail()
	}
}
