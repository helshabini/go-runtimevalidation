package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID4(t *testing.T) {
	// Test with valid UUID4
	t.Run("Valid UUID4", func(t *testing.T) {
		input := "550e8400-e29b-41d4-a716-446655440000"
		err := UUID4(input)
		assert.NoError(t, err)
	})

	// Test with invalid UUID4 (wrong version)
	t.Run("Invalid UUID4 version", func(t *testing.T) {
		input := "550e8400-e29b-51d4-a716-446655440000"
		err := UUID4(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID4 value: 550e8400-e29b-51d4-a716-446655440000")
	})

	// Test with non-hyphenated UUID
	t.Run("Invalid non-hyphenated UUID", func(t *testing.T) {
		input := "550e8400e29b41d4a716446655440000"
		err := UUID4(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID4 value: 550e8400e29b41d4a716446655440000")
	})

	// Test with string that's too short
	t.Run("UUID4 too short", func(t *testing.T) {
		input := "550e8400-e29b-41d4-a716-44665544"
		err := UUID4(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID4 value: 550e8400-e29b-41d4-a716-44665544")
	})

	// Test with non-string input
	t.Run("Non-string input", func(t *testing.T) {
		input := 12345
		err := UUID4(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})

	// Test with empty string
	t.Run("Empty string", func(t *testing.T) {
		input := ""
		err := UUID4(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID4 value: ")
	})
}
