package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsbn13(t *testing.T) {
	// Test valid ISBN-13
	t.Run("Valid ISBN-13", func(t *testing.T) {
		err := Isbn13("9783161484100")
		assert.NoError(t, err)
	})

	// Test valid ISBN-13 with dashes
	t.Run("Valid ISBN-13 with dashes", func(t *testing.T) {
		err := Isbn13("978-3-16-148410-0")
		assert.NoError(t, err)
	})

	// Test invalid ISBN-13
	t.Run("Invalid ISBN-13", func(t *testing.T) {
		err := Isbn13("978316148410X")
		assert.Error(t, err)
	})

	// Test input not a string
	t.Run("Non-string input", func(t *testing.T) {
		err := Isbn13(1234567890123)
		assert.Error(t, err)
	})

	// Test short ISBN-13
	t.Run("Short ISBN-13", func(t *testing.T) {
		err := Isbn13("97831614841")
		assert.Error(t, err)
	})

	// Test empty input
	t.Run("Empty input", func(t *testing.T) {
		err := Isbn13("")
		assert.Error(t, err)
	})

	// Test valid ISBN-13 with spaces
	t.Run("Valid ISBN-13 with spaces", func(t *testing.T) {
		err := Isbn13("978 3 16 148410 0")
		assert.NoError(t, err)
	})

	// Test valid ISBN-13 with mixed spaces and dashes
	t.Run("Valid ISBN-13 with mixed spaces and dashes", func(t *testing.T) {
		err := Isbn13("978-3 16-148410-0")
		assert.NoError(t, err)
	})

	// Test ISBN-13 with invalid characters
	t.Run("ISBN-13 with invalid characters", func(t *testing.T) {
		err := Isbn13("97831614841@0")
		assert.Error(t, err)
	})

	// Test ISBN-13 with correct length but invalid checksum
	t.Run("ISBN-13 with invalid checksum", func(t *testing.T) {
		err := Isbn13("9783161484101")
		assert.Error(t, err)
	})
}
