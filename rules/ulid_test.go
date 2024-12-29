package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestULID(t *testing.T) {
	t.Run("Valid ULID", func(t *testing.T) {
		input := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
		err := ULID(input)
		assert.NoError(t, err)
	})

	t.Run("Invalid ULID format", func(t *testing.T) {
		input := "invalid-ulid"
		err := ULID(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid ULID value: invalid-ulid")
	})

	t.Run("Non-string input", func(t *testing.T) {
		input := 12345
		err := ULID(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("Empty string", func(t *testing.T) {
		input := ""
		err := ULID(input)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid ULID value: ")
	})
}
