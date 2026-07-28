package tgwebapp

import (
	"encoding/hex"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testBotToken = "123456789:TEST_TOKEN"
	testHash     = "63d3bdf25d0ff88a990f3be8e24f22e02c40993054c25afbb43c82db11fb576d"
)

func TestAuthenticateTelegramWebApp(t *testing.T) {
	t.Run("accepts valid Telegram init data", func(t *testing.T) {
		values := validInitDataValues()
		request := httptest.NewRequest(
			http.MethodPost,
			"/webapp?bot=chessraiders",
			strings.NewReader(values.Encode()),
		)
		response := httptest.NewRecorder()
		var completedWith *InitData

		AuthenticateTelegramWebApp(
			response,
			request,
			func(bot string) string {
				assert.Equal(t, "chessraiders", bot)
				return testBotToken
			},
			func(initData *InitData) {
				completedWith = initData
			},
		)

		require.NotNil(t, completedWith)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "AAEAAAE", completedWith.QueryID)
		assert.Equal(t, 1710000000, completedWith.AuthDate)
		assert.Equal(t, testHash, completedWith.Hash)
	})

	for _, tc := range []struct {
		name   string
		mutate func(values url.Values) string
		token  string
	}{
		{
			name: "rejects a tampered field",
			mutate: func(values url.Values) string {
				values.Set("query_id", "tampered")
				return values.Encode()
			},
			token: testBotToken,
		},
		{
			name: "rejects a missing hash",
			mutate: func(values url.Values) string {
				values.Del("hash")
				return values.Encode()
			},
			token: testBotToken,
		},
		{
			name: "rejects a malformed hash",
			mutate: func(values url.Values) string {
				values.Set("hash", "not-hex")
				return values.Encode()
			},
			token: testBotToken,
		},
		{
			name: "rejects a hash of the wrong length",
			mutate: func(values url.Values) string {
				values.Set("hash", "00")
				return values.Encode()
			},
			token: testBotToken,
		},
		{
			name: "rejects an unavailable bot token",
			mutate: func(values url.Values) string {
				return values.Encode()
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			request := httptest.NewRequest(
				http.MethodPost,
				"/webapp",
				strings.NewReader(tc.mutate(validInitDataValues())),
			)
			response := httptest.NewRecorder()
			completedWith := &InitData{QueryID: "not empty"}

			AuthenticateTelegramWebApp(
				response,
				request,
				func(string) string { return tc.token },
				func(initData *InitData) {
					completedWith = initData
				},
			)

			assert.Equal(t, http.StatusUnauthorized, response.Code)
			assert.Equal(t, &InitData{}, completedWith)
		})
	}

	t.Run("rejects non-POST requests before looking up a token", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/webapp", nil)
		response := httptest.NewRecorder()
		tokenRequested := false
		var completedWith *InitData

		AuthenticateTelegramWebApp(
			response,
			request,
			func(string) string {
				tokenRequested = true
				return testBotToken
			},
			func(initData *InitData) {
				completedWith = initData
			},
		)

		assert.Equal(t, http.StatusMethodNotAllowed, response.Code)
		assert.False(t, tokenRequested)
		assert.Equal(t, &InitData{}, completedWith)
	})

	t.Run("rejects a malformed form body", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/webapp", strings.NewReader("%"))
		response := httptest.NewRecorder()
		var completedWith *InitData

		AuthenticateTelegramWebApp(
			response,
			request,
			func(string) string { return testBotToken },
			func(initData *InitData) {
				completedWith = initData
			},
		)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, &InitData{}, completedWith)
	})

	t.Run("reports a request body read failure", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/webapp", nil)
		request.Body = io.NopCloser(failingReader{})
		response := httptest.NewRecorder()
		var completedWith *InitData

		AuthenticateTelegramWebApp(
			response,
			request,
			func(string) string { return testBotToken },
			func(initData *InitData) {
				completedWith = initData
			},
		)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, &InitData{}, completedWith)
	})
}

func TestGetDataCheckString(t *testing.T) {
	values := validInitDataValues()

	assert.Equal(
		t,
		"auth_date=1710000000\n"+
			"query_id=AAEAAAE\n"+
			"signature=test-signature\n"+
			`user={"id":42,"first_name":"Ada"}`,
		getDataCheckString(values),
	)
}

func TestComputeWebAppHash(t *testing.T) {
	hash := computeWebAppHash(getDataCheckString(validInitDataValues()), testBotToken)

	assert.Equal(t, testHash, hex.EncodeToString(hash))
}

func validInitDataValues() url.Values {
	return url.Values{
		"auth_date": {"1710000000"},
		"hash":      {testHash},
		"query_id":  {"AAEAAAE"},
		"signature": {"test-signature"},
		"user":      {`{"id":42,"first_name":"Ada"}`},
	}
}

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) {
	return 0, errors.New("read failed")
}
