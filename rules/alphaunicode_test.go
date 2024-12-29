package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases for the AlphaUnicode function
func TestAlphaUnicode(t *testing.T) {
	t.Run("Valid English letters", func(t *testing.T) {
		err := AlphaUnicode("JohnDoe")
		assert.NoError(t, err)
	})

	t.Run("Valid accented characters", func(t *testing.T) {
		err := AlphaUnicode("José")
		assert.NoError(t, err)
	})

	t.Run("Valid Cyrillic characters", func(t *testing.T) {
		err := AlphaUnicode("Сергей")
		assert.NoError(t, err)
	})

	t.Run("Valid Chinese characters", func(t *testing.T) {
		err := AlphaUnicode("中文")
		assert.NoError(t, err)
	})

	t.Run("Invalid: Contains numbers", func(t *testing.T) {
		err := AlphaUnicode("123ABC")
		assert.EqualError(t, err, "invalid alpha unicode: 123ABC")
	})

	t.Run("Invalid: Contains underscore", func(t *testing.T) {
		err := AlphaUnicode("John_Doe")
		assert.EqualError(t, err, "invalid alpha unicode: John_Doe")
	})

	t.Run("Invalid: Empty string", func(t *testing.T) {
		err := AlphaUnicode("")
		assert.EqualError(t, err, "invalid alpha unicode: ")
	})

	t.Run("Invalid: Input is an integer", func(t *testing.T) {
		err := AlphaUnicode(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("Invalid: Input is nil", func(t *testing.T) {
		err := AlphaUnicode(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("Invalid: Contains special characters", func(t *testing.T) {
		err := AlphaUnicode("John@Doe")
		assert.EqualError(t, err, "invalid alpha unicode: John@Doe")
	})

	t.Run("Invalid: Contains spaces", func(t *testing.T) {
		err := AlphaUnicode("John Doe")
		assert.EqualError(t, err, "invalid alpha unicode: John Doe")
	})

	t.Run("Invalid: Contains punctuation", func(t *testing.T) {
		err := AlphaUnicode("John.Doe")
		assert.EqualError(t, err, "invalid alpha unicode: John.Doe")
	})

	t.Run("Valid: Mixed case letters", func(t *testing.T) {
		err := AlphaUnicode("JohnDoe")
		assert.NoError(t, err)
	})

	t.Run("Valid: Single letter", func(t *testing.T) {
		err := AlphaUnicode("A")
		assert.NoError(t, err)
	})
}
