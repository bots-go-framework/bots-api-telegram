package tgbotapi

import (
	"errors"
	"fmt"
	"github.com/pquerna/ffjson/ffjson"
	"net/url"
)

var _ Sendable = SetMyCommandsConfig{}

type MyCommandsBase struct {
	// Optional. A JSON-serialized object, describing scope of users for which the commands are relevant. Defaults to BotCommandScopeDefault.
	Scope *BotCommandScope `json:"scope,omitempty"`

	// Optional. A two-letter ISO 639-1 language code. If empty, commands will be applied to all users from the given scope, for whose language there are no dedicated commands
	LanguageCode string `json:"language_code,omitempty"`
}

func (s MyCommandsBase) Validate() error {
	if s.Scope != nil {
		if err := s.Scope.Validate(); err != nil {
			return err
		}
	}
	switch l := len(s.LanguageCode); l {
	case 0, 2: //OK
		break
	default:
		return errors.New("language code length must be 0 or 2")
	}
	return nil
}

func (s MyCommandsBase) Values() (values url.Values, err error) {
	if err = s.Validate(); err != nil {
		return
	}
	values = make(url.Values)
	if s.LanguageCode != "" {
		values.Set("language_code", s.LanguageCode)
	}
	return
}

var _ Sendable = SetMyCommandsConfig{}

type GetMyCommandsConfig = MyCommandsBase

func (v GetMyCommandsConfig) TelegramMethod() string {
	return "getMyCommands"
}

type SetMyCommandsConfig struct {
	MyCommandsBase

	// A JSON-serialized list of bot commands to be set as the list of the bot's commands. At most 100 commands can be specified.
	Commands []TelegramBotCommand `json:"commands"`
}

func (s SetMyCommandsConfig) Validate() error {
	if err := s.MyCommandsBase.Validate(); err != nil {
		return err
	}
	if count := len(s.Commands); count == 0 {
		return errors.New("commands must have at least one command")
	} else if count > 100 {
		return errors.New("at most 100 commands can be specified")
	}
	for i, cmd := range s.Commands {
		if err := cmd.Validate(); err != nil {
			return fmt.Errorf("commands[%d]: %s", i, err)
		}
	}
	return nil
}

func (s SetMyCommandsConfig) Values() (values url.Values, err error) {
	if err = s.Validate(); err != nil {
		return
	}
	if values, err = s.MyCommandsBase.Values(); err != nil {
		return
	}
	if len(s.Commands) > 0 {
		var b []byte
		if b, err = ffjson.Marshal(s.Commands); err != nil {
			err = fmt.Errorf("failed to serialize commands to JSON: %w", err)
			return
		}
		values.Set("commands", string(b))
	}
	return
}

func (s SetMyCommandsConfig) TelegramMethod() string {
	return "setMyCommands"
}

type TelegramBotCommand struct {
	Command     string `json:"command"`     // Text of the command; 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Description string `json:"description"` // Description of the command; 1-256 characters.
}

func (v TelegramBotCommand) Validate() error {
	if len(v.Command) == 0 {
		return errors.New("command is required")
	}
	if len(v.Command) > 32 {
		return errors.New("command is too long, expected to be 32 characters max")
	}
	if len(v.Description) == 0 {
		return errors.New("description is required")
	}
	if len(v.Description) > 256 {
		return errors.New("description is too long, expected to be 256 characters max")
	}
	return nil
}

// BotCommandScope represents the scope to which bot commands are applied. Currently, the following 7 scopes are supported:
//
// - BotCommandScopeDefault
// - BotCommandScopeAllPrivateChats
// - BotCommandScopeAllGroupChats
// - BotCommandScopeAllChatAdministrators
// - BotCommandScopeChat
// - BotCommandScopeChatAdministrators
// - BotCommandScopeChatMember
//
// https://core.telegram.org/bots/api#botcommandscope
type BotCommandScope struct {
	Type   BotCommandScopeType `json:"type"`
	ChatID any                 `json:"chatID"`  // Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserID int                 `json:"user_id"` // Unique identifier of the target user
}

func (v *BotCommandScope) Validate() error {
	if v == nil {
		return nil
	}
	switch v.Type {
	case // OK
		BotCommandScopeDefault,
		BotCommandScopeAllPrivateChats,
		BotCommandScopeAllGroupChats,
		BotCommandScopeAllChatAdministrators,
		BotCommandScopeChat,
		BotCommandScopeChatAdministrators,
		BotCommandScopeChatMember:
		break
	case "":
		return errors.New("scope type is required")
	default:
		return errors.New("unknown command scope type: " + string(v.Type))
	}
	if v.ChatID != nil {
		switch v.ChatID.(type) {
		case int, int64, string:
			break // OK
		default:
			return errors.New("chat_id is required")
		}
	}
	if v.UserID < 0 {
		return errors.New("user_id must be positive integer")
	}
	return nil
}

type BotCommandScopeType string

const (
	BotCommandScopeDefault               BotCommandScopeType = "default"
	BotCommandScopeAllPrivateChats       BotCommandScopeType = "all_private_chats"
	BotCommandScopeAllGroupChats         BotCommandScopeType = "all_group_chats"
	BotCommandScopeAllChatAdministrators BotCommandScopeType = "all_chat_administrators"
	BotCommandScopeChat                  BotCommandScopeType = "chat"
	BotCommandScopeChatAdministrators    BotCommandScopeType = "chat_administrators"
	BotCommandScopeChatMember            BotCommandScopeType = "chat_member"
)
