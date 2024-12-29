package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsbn10(t *testing.T) {
	// Test with a valid ISBN-10
	t.Run("Valid ISBN-10", func(t *testing.T) {
		err := Isbn10("0471958697")
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	// Test with an invalid ISBN-10
	t.Run("Invalid ISBN-10", func(t *testing.T) {
		err := Isbn10("invalid123")
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "invalid ISBN10:")
		}
	})

	// Test with a non-string input
	t.Run("Non-string input", func(t *testing.T) {
		err := Isbn10(12345)
		if assert.Error(t, err) {
			assert.Equal(t, "expected a string, got int", err.Error())
		}
	})

	// Test with an empty string (invalid ISBN-10)
	t.Run("Empty string", func(t *testing.T) {
		err := Isbn10("")
		if assert.Error(t, err) {
			assert.Contains(t, err.Error(), "invalid ISBN10:")
		}
	})

	// Test with a valid ISBN-10 that includes the check digit 'X'
	t.Run("Valid ISBN-10 with check digit X", func(t *testing.T) {
		err := Isbn10("0306406152")
		assert.NoError(t, err)
	})
}
