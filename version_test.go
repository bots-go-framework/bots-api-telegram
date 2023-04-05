package bots_api_telegram

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersion(t *testing.T) {
	assert.NotEmptyf(t, Version, "Version is empty")
}
