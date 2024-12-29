package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases for the Alpha function
func TestAlpha(t *testing.T) {
	t.Run("Valid alphabetic string", func(t *testing.T) {
		err := Alpha("HelloWorld")
		assert.NoError(t, err)
	})

	t.Run("Valid all uppercase", func(t *testing.T) {
		err := Alpha("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
		assert.NoError(t, err)
	})

	t.Run("Valid all lowercase", func(t *testing.T) {
		err := Alpha("abcdefghijklmnopqrstuvwxyz")
		assert.NoError(t, err)
	})

	t.Run("Invalid, contains numbers", func(t *testing.T) {
		err := Alpha("Hello123")
		assert.EqualError(t, err, "invalid alpha: Hello123")
	})

	t.Run("Invalid, all numbers", func(t *testing.T) {
		err := Alpha("12345")
		assert.EqualError(t, err, "invalid alpha: 12345")
	})

	t.Run("Invalid, contains hyphen", func(t *testing.T) {
		err := Alpha("Hello-World")
		assert.EqualError(t, err, "invalid alpha: Hello-World")
	})

	t.Run("Invalid, contains space", func(t *testing.T) {
		err := Alpha("Hello World")
		assert.EqualError(t, err, "invalid alpha: Hello World")
	})

	t.Run("Invalid, empty string", func(t *testing.T) {
		err := Alpha("")
		assert.EqualError(t, err, "invalid alpha: ")
	})

	t.Run("Invalid, input is not a string", func(t *testing.T) {
		err := Alpha(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("Invalid, input is nil", func(t *testing.T) {
		err := Alpha(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})
}
