package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID5(t *testing.T) {
	// Test with valid UUID5
	t.Run("Valid UUID5", func(t *testing.T) {
		input := "f47ac10b-58cc-5b4c-8b9a-2bbd051a3cb8"
		err := UUID5(input)
		assert.NoError(t, err)
	})

	// Test with invalid UUID5 (wrong version)
	t.Run("Invalid UUID5 version", func(t *testing.T) {
		input := "f47ac10b-58cc-4b4c-8b9a-2bbd051a3cb8"
		err := UUID5(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID5 value: f47ac10b-58cc-4b4c-8b9a-2bbd051a3cb8")
	})

	// Test with non-hyphenated UUID
	t.Run("Invalid non-hyphenated UUID", func(t *testing.T) {
		input := "f47ac10b58cc5b4c8b9a2bbd051a3cb8"
		err := UUID5(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID5 value: f47ac10b58cc5b4c8b9a2bbd051a3cb8")
	})

	// Test with string that's too short
	t.Run("UUID5 too short", func(t *testing.T) {
		input := "f47ac10b-58cc-5b4c-8b9a-2bbd051a3c"
		err := UUID5(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID5 value: f47ac10b-58cc-5b4c-8b9a-2bbd051a3c")
	})

	// Test with non-string input
	t.Run("Non-string input", func(t *testing.T) {
		input := 12345
		err := UUID5(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})

	// Test with empty string
	t.Run("Empty string", func(t *testing.T) {
		input := ""
		err := UUID5(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID5 value: ")
	})
}
