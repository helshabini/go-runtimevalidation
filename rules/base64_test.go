package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64(t *testing.T) {
	t.Run("valid Base64 string", func(t *testing.T) {
		err := Base64("SGVsbG8gd29ybGQ=") // "Hello world" in Base64
		assert.NoError(t, err)
	})

	t.Run("invalid Base64 string", func(t *testing.T) {
		err := Base64("InvalidBase64String!")
		assert.EqualError(t, err, "invalid Base64 string: InvalidBase64String!")
	})

	t.Run("non-string input", func(t *testing.T) {
		err := Base64(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
