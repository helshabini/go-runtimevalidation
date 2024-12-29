package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsciiPrint(t *testing.T) {
	t.Run("Valid Printable ASCII", func(t *testing.T) {
		err := AsciiPrint("Hello, World!")
		assert.NoError(t, err)
	})

	t.Run("Valid Printable ASCII with symbols", func(t *testing.T) {
		err := AsciiPrint("12345!@#$%^&*()_+-=")
		assert.NoError(t, err)
	})

	t.Run("Valid ASCII spaces", func(t *testing.T) {
		err := AsciiPrint("  ")
		assert.NoError(t, err)
	})

	t.Run("Invalid ASCII with non-printable characters", func(t *testing.T) {
		err := AsciiPrint("Hello\tWorld") // contains a tab
		assert.Error(t, err)
	})

	t.Run("Invalid input with non-ASCII characters", func(t *testing.T) {
		err := AsciiPrint("こんにちは") // non-ASCII characters
		assert.Error(t, err)
	})

	t.Run("Input not a string", func(t *testing.T) {
		err := AsciiPrint(12345) // non-string input
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expected a string")
	})
}
