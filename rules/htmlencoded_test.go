package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTMLEncoded(t *testing.T) {
	t.Run("Valid HTML-encoded string", func(t *testing.T) {
		err := HTMLEncoded("&amp;")
		assert.NoError(t, err)
	})

	t.Run("Invalid HTML-encoded string (numbers)", func(t *testing.T) {
		err := HTMLEncoded("123")
		assert.Error(t, err)
		assert.Equal(t, "invalid html encoded: 123", err.Error())
	})

	t.Run("Valid numeric HTML entity", func(t *testing.T) {
		err := HTMLEncoded("&#123;")
		assert.NoError(t, err)
	})

	t.Run("Invalid HTML-encoded string (missing semicolon)", func(t *testing.T) {
		err := HTMLEncoded("&amp")
		assert.Error(t, err)
		assert.Equal(t, "invalid html encoded: &amp", err.Error())
	})

	t.Run("Invalid HTML-encoded string (non-entity)", func(t *testing.T) {
		err := HTMLEncoded("hello")
		assert.Error(t, err)
		assert.Equal(t, "invalid html encoded: hello", err.Error())
	})

	t.Run("Empty string input", func(t *testing.T) {
		err := HTMLEncoded("")
		assert.Error(t, err)
		assert.Equal(t, "invalid html encoded: ", err.Error())
	})

	t.Run("Invalid input type (int instead of string)", func(t *testing.T) {
		err := HTMLEncoded(123)
		assert.Error(t, err)
		assert.Equal(t, "invalid html encoded: 123", err.Error())
	})

	t.Run("Mixed content with valid HTML encoding", func(t *testing.T) {
		err := HTMLEncoded("Hello &amp; welcome!")
		assert.NoError(t, err)
	})
}
