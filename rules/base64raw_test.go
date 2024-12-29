package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase64Raw(t *testing.T) {
	t.Run("valid Base64Raw string", func(t *testing.T) {
		err := Base64Raw("SGVsbG8gd29ybGQ") // "Hello world" in Base64 without padding
		assert.NoError(t, err)
	})

	t.Run("invalid Base64Raw string", func(t *testing.T) {
		err := Base64Raw("InvalidBase64RawString!")
		assert.EqualError(t, err, "invalid Base64Raw string: InvalidBase64RawString!")
	})

	t.Run("non-string input", func(t *testing.T) {
		err := Base64Raw(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
