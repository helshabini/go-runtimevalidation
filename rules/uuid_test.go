package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUUID(t *testing.T) {
	// Test valid UUID
	t.Run("Valid UUID", func(t *testing.T) {
		err := UUID("123e4567-e89b-12d3-a456-426614174000")
		assert.NoError(t, err)
	})

	// Test invalid UUID with incorrect format
	t.Run("Invalid UUID with incorrect format", func(t *testing.T) {
		err := UUID("123e4567-e89b-12d3-a456")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID value: 123e4567-e89b-12d3-a456")
	})

	// Test invalid UUID with non-hex characters
	t.Run("Invalid UUID with non-hex characters", func(t *testing.T) {
		err := UUID("123e4567-e89b-12d3-a456-42661417g000")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID value: 123e4567-e89b-12d3-a456-42661417g000")
	})

	// Test UUID input is not a string
	t.Run("Input is not a string", func(t *testing.T) {
		err := UUID(1234567890)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})

	// Test valid UUID with uppercase characters
	t.Run("Valid UUID with uppercase", func(t *testing.T) {
		err := UUID("123E4567-E89B-12D3-A456-426614174000")
		assert.NoError(t, err)
	})

	// Test empty string input
	t.Run("Empty string", func(t *testing.T) {
		err := UUID("")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid UUID value: ")
	})
}
