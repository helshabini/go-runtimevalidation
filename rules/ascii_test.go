package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAscii(t *testing.T) {
	t.Run("valid ASCII string", func(t *testing.T) {
		err := Ascii("Hello, World!")
		assert.NoError(t, err)
	})

	t.Run("valid ASCII string with numbers and symbols", func(t *testing.T) {
		err := Ascii("12345!@#$%^&*()")
		assert.NoError(t, err)
	})

	t.Run("invalid string with non-ASCII character", func(t *testing.T) {
		err := Ascii("Hello, 世界!")
		assert.Error(t, err)
		assert.Equal(t, "invalid ASCII: Hello, 世界!", err.Error())
	})

	t.Run("empty string is valid ASCII", func(t *testing.T) {
		err := Ascii("")
		assert.NoError(t, err)
	})

	t.Run("input is not a string", func(t *testing.T) {
		err := Ascii(12345)
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got int", err.Error())
	})
}
