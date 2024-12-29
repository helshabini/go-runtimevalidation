package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases for the AlphaNumeric function
func TestAlphaNumeric(t *testing.T) {
	t.Run("ValidAlphanumericString", func(t *testing.T) {
		err := AlphaNumeric("User123")
		assert.NoError(t, err)
	})

	t.Run("ValidMixedCaseAlphanumericString", func(t *testing.T) {
		err := AlphaNumeric("ABCdef123")
		assert.NoError(t, err)
	})

	t.Run("ValidAllNumericString", func(t *testing.T) {
		err := AlphaNumeric("1234567890")
		assert.NoError(t, err)
	})

	t.Run("InvalidSpecialCharacters", func(t *testing.T) {
		err := AlphaNumeric("Hello!@#")
		assert.EqualError(t, err, "invalid alphanumeric: Hello!@#")
	})

	t.Run("InvalidContainsSpace", func(t *testing.T) {
		err := AlphaNumeric("Hello World")
		assert.EqualError(t, err, "invalid alphanumeric: Hello World")
	})

	t.Run("InvalidEmptyString", func(t *testing.T) {
		err := AlphaNumeric("")
		assert.EqualError(t, err, "invalid alphanumeric: ")
	})

	t.Run("InvalidNonStringInput", func(t *testing.T) {
		err := AlphaNumeric(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("InvalidNilInput", func(t *testing.T) {
		err := AlphaNumeric(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("ValidSingleCharacter", func(t *testing.T) {
		err := AlphaNumeric("A")
		assert.NoError(t, err)
	})

	t.Run("ValidSingleDigit", func(t *testing.T) {
		err := AlphaNumeric("1")
		assert.NoError(t, err)
	})

	t.Run("InvalidUnicodeCharacters", func(t *testing.T) {
		err := AlphaNumeric("こんにちは")
		assert.EqualError(t, err, "invalid alphanumeric: こんにちは")
	})

	t.Run("InvalidWhitespaceCharacters", func(t *testing.T) {
		err := AlphaNumeric("User 123")
		assert.EqualError(t, err, "invalid alphanumeric: User 123")
	})

	t.Run("InvalidSpecialCharactersOnly", func(t *testing.T) {
		err := AlphaNumeric("!@#$%^&*()")
		assert.EqualError(t, err, "invalid alphanumeric: !@#$%^&*()")
	})

	t.Run("ValidLongAlphanumericString", func(t *testing.T) {
		err := AlphaNumeric("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
		assert.NoError(t, err)
	})
}
