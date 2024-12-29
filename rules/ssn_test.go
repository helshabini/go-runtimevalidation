package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSSN(t *testing.T) {
	// Test valid SSN
	t.Run("Valid SSN", func(t *testing.T) {
		input := "123-45-6789"
		err := SSN(input)
		assert.NoError(t, err)
	})

	// Test obscured SSN
	t.Run("Obscured SSN", func(t *testing.T) {
		input := "123-45-XXXX"
		err := SSN(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid SSN value: 123-45-XXXX")
	})

	// Test invalid SSN with wrong format
	t.Run("Invalid SSN - wrong format", func(t *testing.T) {
		input := "123-456-789"
		err := SSN(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid SSN value: 123-456-789")
	})

	// Test invalid SSN with letters
	t.Run("Invalid SSN - contains letters", func(t *testing.T) {
		input := "123-45-ABCD"
		err := SSN(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid SSN value: 123-45-ABCD")
	})

	// Test invalid SSN - too short
	t.Run("Invalid SSN - too short", func(t *testing.T) {
		input := "123-45-67"
		err := SSN(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid SSN value: 123-45-67")
	})

	// Test invalid SSN - too long
	t.Run("Invalid SSN - too long", func(t *testing.T) {
		input := "123-45-67890"
		err := SSN(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid SSN value: 123-45-67890")
	})

	// Test invalid obscured SSN - too long
	t.Run("Invalid SSN - too long", func(t *testing.T) {
		input := "123-45-XXXXX"
		err := SSN(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid SSN value: 123-45-XXXXX")
	})

	// Test input that is not a string
	t.Run("Non-string input", func(t *testing.T) {
		input := 123456789
		err := SSN(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
