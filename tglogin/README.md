# tglogin — Telegram Login Widget authentication

The `tglogin` package validates user data received from the
[Telegram Login Widget](https://core.telegram.org/widgets/login).

## Installation

```sh
go get github.com/bots-go-framework/bots-api-telegram/tglogin
```

## Usage

```go
import "github.com/bots-go-framework/bots-api-telegram/tglogin"

// Populate from the widget callback query parameters
user := tglogin.LoginUser{
    ID:        123456789,
    FirstName: "Alice",
    AuthDate:  1700000000,
    Hash:      "<hash from Telegram>",
}

if user.IsFromTelegram(botToken) {
    // Data is authentic — safe to create a session for this user
} else {
    // Reject: data was not signed by Telegram
}
```

## Exported API

### `LoginUser`

```go
type LoginUser struct {
    ID        int    `json:"id"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name,omitempty"`
    Username  string `json:"username,omitempty"`
    PhotoUrl  string `json:"photo_url,omitempty"`
    AuthDate  int    `json:"auth_date"`
    Hash      string `json:"hash"`
}
```

| Method | Description |
|---|---|
| `IsFromTelegram(botToken string) bool` | Returns `true` if the struct fields are authentically signed by Telegram using the given bot token |

## Telegram documentation

- [Telegram Login blog post](https://telegram.org/blog/login)
- [Login widget reference](https://core.telegram.org/widgets/login)