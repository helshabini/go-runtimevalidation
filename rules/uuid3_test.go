package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID3(t *testing.T) {
	// Test valid UUID3
	t.Run("Valid UUID3", func(t *testing.T) {
		input := "f47ac10b-58cc-3bf1-8a9a-1234567890ab"
		err := UUID3(input)
		assert.NoError(t, err)
	})

	// Test invalid UUID3 - incorrect version number
	t.Run("Invalid UUID3 with incorrect version", func(t *testing.T) {
		input := "f47ac10b-58cc-4bf1-8a9a-1234567890ab" // Should have a '3' in the version position
		err := UUID3(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID3 value: f47ac10b-58cc-4bf1-8a9a-1234567890ab")
	})

	// Test invalid UUID3 - missing segments
	t.Run("Invalid UUID3 with missing segments", func(t *testing.T) {
		input := "f47ac10b-58cc-3bf1-8a9a"
		err := UUID3(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID3 value: f47ac10b-58cc-3bf1-8a9a")
	})

	// Test valid UUID3 with uppercase characters
	t.Run("Valid UUID3 with uppercase characters", func(t *testing.T) {
		input := "F47AC10B-58CC-3BF1-8A9A-1234567890AB"
		err := UUID3(input)
		assert.NoError(t, err)
	})

	// Test invalid input type
	t.Run("Invalid input type", func(t *testing.T) {
		input := 12345 // Not a string
		err := UUID3(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})

	// Test invalid UUID3 - random string
	t.Run("Invalid UUID3 with random string", func(t *testing.T) {
		input := "not-a-uuid"
		err := UUID3(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID3 value: not-a-uuid")
	})
}
