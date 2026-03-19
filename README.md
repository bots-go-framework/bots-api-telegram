# Go bindings for the Telegram Bot API

[![Go CI](https://github.com/bots-go-framework/bots-api-telegram/actions/workflows/ci.yml/badge.svg)](https://github.com/bots-go-framework/bots-api-telegram/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/bots-go-framework/bots-api-telegram)](https://goreportcard.com/report/github.com/bots-go-framework/bots-api-telegram)
[![GoDoc](https://pkg.go.dev/badge/github.com/bots-go-framework/bots-api-telegram)](https://pkg.go.dev/github.com/bots-go-framework/bots-api-telegram)

Go bindings for the [Telegram Bot API](https://core.telegram.org/bots/api), tracking **Bot API 9.5**.

This module also includes sub-packages for [Telegram Login Widget](#tglogin) and [Telegram Web Apps](#tgwebapp).

## Packages

| Package | Import path | Description |
|---|---|---|
| `tgbotapi` | `github.com/bots-go-framework/bots-api-telegram/tgbotapi` | Core Bot API types and HTTP client |
| `tglogin` | `github.com/bots-go-framework/bots-api-telegram/tglogin` | Telegram Login Widget authentication |
| `tgwebapp` | `github.com/bots-go-framework/bots-api-telegram/tgwebapp` | Telegram Web App (Mini App) init-data validation |

## Installation

```sh
go get github.com/bots-go-framework/bots-api-telegram
```

## Examples

### Long-polling

A minimal bot that echoes every message back to the sender:

```go
package main

import (
	"log"

	"github.com/bots-go-framework/bots-api-telegram/tgbotapi"
)

func main() {
	bot := tgbotapi.NewBotAPI("MyAwesomeBotToken")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatal(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}
}
```

### Webhook

To receive updates via webhook your server must be reachable over HTTPS.
Telegram supports ports **443, 80, 88, and 8443**.

```go
package main

import (
	"log"
	"net/http"

	"github.com/bots-go-framework/bots-api-telegram/tgbotapi"
)

func main() {
	bot := tgbotapi.NewBotAPI("MyAwesomeBotToken")

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err := bot.SetWebhook(tgbotapi.NewWebhookWithCert(
		"https://example.com:8443/"+bot.Token, "cert.pem",
	))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)
	go func() {
		if err := http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil); err != nil {
			log.Fatal(err)
		}
	}()

	for update := range updates {
		log.Printf("%+v\n", update)
	}
}
```

#### Self-signed TLS certificate

If you don't have a certificate from a trusted CA (e.g. [Let's Encrypt](https://letsencrypt.org)),
you can generate a self-signed one. Pass the public certificate to `NewWebhookWithCert` so
Telegram knows to trust it.

```sh
openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 3650 \
  -subj "/O=Org/CN=example.com" -nodes
```

## tglogin

The `tglogin` sub-package validates data received from the
[Telegram Login Widget](https://core.telegram.org/widgets/login).

```go
import "github.com/bots-go-framework/bots-api-telegram/tglogin"

user := tglogin.LoginUser{ /* fields from the widget callback */ }
if user.IsFromTelegram(botToken) {
    // user data is authentic
}
```

See [`tglogin/README.md`](tglogin/README.md) for full details.

## tgwebapp

The `tgwebapp` sub-package validates init data received from a
[Telegram Web App (Mini App)](https://core.telegram.org/bots/webapps).

```go
import "github.com/bots-go-framework/bots-api-telegram/tgwebapp"

tgwebapp.AuthenticateTelegramWebApp(w, r,
    func(bot string) string { return myBotToken },
    func(initData *tgwebapp.InitData) {
        // initData is verified and ready to use
    },
)
```

See [`tgwebapp/README.md`](tgwebapp/README.md) for full details.

## Used by

- [debtstracker.io](https://debtstracker.io/) — personal debt tracking app with a [Telegram bot](https://t.me/DebtsTrackerBot)

## Related

- [`bots-go-framework`](https://github.com/bots-go-framework) — framework for building Telegram bots in Go
