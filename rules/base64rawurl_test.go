package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64RawUrl(t *testing.T) {
	t.Run("valid Base64RawUrl string", func(t *testing.T) {
		err := Base64RawUrl("SGVsbG8gd29ybGQ") // "Hello world" in Base64 URL without padding
		assert.NoError(t, err)
	})

	t.Run("invalid Base64RawUrl string", func(t *testing.T) {
		err := Base64RawUrl("InvalidBase64RawUrlString!")
		assert.EqualError(t, err, "invalid Base64RawUrl string: InvalidBase64RawUrlString!")
	})

	t.Run("non-string input", func(t *testing.T) {
		err := Base64RawUrl(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
