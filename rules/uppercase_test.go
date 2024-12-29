package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUppercase(t *testing.T) {
	t.Run("Valid uppercase string", func(t *testing.T) {
		err := Uppercase("HELLO WORLD")
		assert.NoError(t, err)
	})

	t.Run("Invalid uppercase string", func(t *testing.T) {
		err := Uppercase("hello WORLD")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid uppercase: hello WORLD")
	})

	t.Run("Valid uppercase with numbers and symbols", func(t *testing.T) {
		err := Uppercase("HELLO123!@#")
		assert.NoError(t, err)
	})

	t.Run("Valid empty string", func(t *testing.T) {
		err := Uppercase("")
		assert.NoError(t, err)
	})

	t.Run("Invalid string with lowercase letters", func(t *testing.T) {
		err := Uppercase("Hello WORLD")
		assert.Error(t, err)
	})

	t.Run("Invalid string with mixed case", func(t *testing.T) {
		err := Uppercase("HELLO world")
		assert.Error(t, err)
	})

	t.Run("Invalid string with only one lowercase letter", func(t *testing.T) {
		err := Uppercase("HELLO wORLD")
		assert.Error(t, err)
	})

	t.Run("Valid string with non-letter characters", func(t *testing.T) {
		err := Uppercase("!!! @@@ ### 123")
		assert.NoError(t, err)
	})
}
