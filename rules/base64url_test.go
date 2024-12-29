package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Url(t *testing.T) {
	t.Run("valid Base64Url string", func(t *testing.T) {
		err := Base64Url("SGVsbG8gd29ybGQ=") // "Hello world" in Base64 URL
		assert.NoError(t, err)
	})

	t.Run("invalid Base64Url string", func(t *testing.T) {
		err := Base64Url("InvalidBase64UrlString!")
		assert.EqualError(t, err, "invalid Base64Url string: InvalidBase64UrlString!")
	})

	t.Run("non-string input", func(t *testing.T) {
		err := Base64Url(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
