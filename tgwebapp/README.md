# tgwebapp — Telegram Web App (Mini App) init-data validation

The `tgwebapp` package validates the `initData` received from a
[Telegram Web App (Mini App)](https://core.telegram.org/bots/webapps)
and parses it into a typed `InitData` struct.

## Installation

```sh
go get github.com/bots-go-framework/bots-api-telegram/tgwebapp
```

## Usage

Call `AuthenticateTelegramWebApp` from your HTTP handler. It reads the
POST body, validates the HMAC signature, and invokes `complete` with the
parsed `InitData`. On any error it writes the appropriate HTTP status and
calls `complete` with an empty `InitData`.

```go
import "github.com/bots-go-framework/bots-api-telegram/tgwebapp"

http.HandleFunc("/webapp", func(w http.ResponseWriter, r *http.Request) {
    tgwebapp.AuthenticateTelegramWebApp(w, r,
        func(bot string) string {
            return os.Getenv("TELEGRAM_BOT_TOKEN")
        },
        func(initData *tgwebapp.InitData) {
            if initData.QueryID == "" {
                return // authentication failed; HTTP error already written
            }
            // initData is verified — handle the request
            fmt.Fprintf(w, "Hello, auth_date=%d", initData.AuthDate)
        },
    )
})
```

## Exported API

### `AuthenticateTelegramWebApp`

```go
func AuthenticateTelegramWebApp(
    w          http.ResponseWriter,
    r          *http.Request,
    getToken   func(bot string) string,
    complete   func(initData *InitData),
)
```

Validates the request and calls `complete` with the parsed init data.
The `bot` parameter passed to `getToken` comes from the `?bot=` query
string of the request URL, allowing a single handler to serve multiple bots.

### `InitData`

Parsed representation of the
[`WebAppInitData`](https://core.telegram.org/bots/webapps#webappinitdata)
object:

```go
type InitData struct {
    QueryID      string `json:"query_id"`
    ChatType     string `json:"chat_type,omitempty"`
    ChatInstance string `json:"chat_instance,omitempty"`
    StartParam   string `json:"start_param,omitempty"`
    CanSendAfter int    `json:"can_send_after,omitempty"`
    AuthDate     int    `json:"auth_date"`
    Hash         string `json:"hash"`
}
```

### `NewInitDataFromUrlValues`

```go
func NewInitDataFromUrlValues(values url.Values) InitData
```

Parses a `url.Values` map (e.g. from a POST body) into an `InitData` struct.

## Telegram documentation

- [Web Apps overview](https://core.telegram.org/bots/webapps)
- [Validating init data](https://core.telegram.org/bots/webapps#validating-data-received-via-the-web-app)
- [`WebAppInitData` reference](https://core.telegram.org/bots/webapps#webappinitdata)
