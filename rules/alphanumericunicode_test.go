package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases for the AlphaNumericUnicode function
func TestAlphaNumericUnicode(t *testing.T) {
	t.Run("Valid: Letters and numbers", func(t *testing.T) {
		err := AlphaNumericUnicode("John123")
		assert.NoError(t, err)
	})

	t.Run("Valid: Letters and numbers", func(t *testing.T) {
		err := AlphaNumericUnicode("Producto2023")
		assert.NoError(t, err)
	})

	t.Run("Valid: Only numbers", func(t *testing.T) {
		err := AlphaNumericUnicode("123456")
		assert.NoError(t, err)
	})

	t.Run("Valid: Cyrillic letters and numbers", func(t *testing.T) {
		err := AlphaNumericUnicode("Сергей2023")
		assert.NoError(t, err)
	})

	t.Run("Valid: Chinese characters and numbers", func(t *testing.T) {
		err := AlphaNumericUnicode("中文2023")
		assert.NoError(t, err)
	})

	t.Run("Invalid: Contains special character (@)", func(t *testing.T) {
		err := AlphaNumericUnicode("User@Name")
		assert.EqualError(t, err, "invalid alpha unicode numeric: User@Name")
	})

	t.Run("Invalid: Empty string", func(t *testing.T) {
		err := AlphaNumericUnicode("")
		assert.EqualError(t, err, "invalid alpha unicode numeric: ")
	})

	t.Run("Invalid: Input is an integer", func(t *testing.T) {
		err := AlphaNumericUnicode(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("Invalid: Input is nil", func(t *testing.T) {
		err := AlphaNumericUnicode(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("Invalid: Contains space", func(t *testing.T) {
		err := AlphaNumericUnicode("John Doe")
		assert.EqualError(t, err, "invalid alpha unicode numeric: John Doe")
	})

	t.Run("Invalid: Contains hyphen", func(t *testing.T) {
		err := AlphaNumericUnicode("John-Doe")
		assert.EqualError(t, err, "invalid alpha unicode numeric: John-Doe")
	})

	t.Run("Invalid: Contains underscore", func(t *testing.T) {
		err := AlphaNumericUnicode("John_Doe")
		assert.EqualError(t, err, "invalid alpha unicode numeric: John_Doe")
	})

	t.Run("Invalid: Contains punctuation", func(t *testing.T) {
		err := AlphaNumericUnicode("John.Doe")
		assert.EqualError(t, err, "invalid alpha unicode numeric: John.Doe")
	})
}
