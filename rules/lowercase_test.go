package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowercase(t *testing.T) {
	t.Run("Valid lowercase string", func(t *testing.T) {
		err := Lowercase("hello")
		assert.NoError(t, err)
	})

	t.Run("Valid lowercase string with numbers and symbols", func(t *testing.T) {
		err := Lowercase("hello123!")
		assert.NoError(t, err)
	})

	t.Run("Valid empty string", func(t *testing.T) {
		err := Lowercase("")
		assert.NoError(t, err)
	})

	t.Run("Valid string with only numbers and symbols", func(t *testing.T) {
		err := Lowercase("123!@#")
		assert.NoError(t, err)
	})

	t.Run("Invalid string with uppercase letters", func(t *testing.T) {
		err := Lowercase("Hello")
		assert.Error(t, err)
	})

	t.Run("Invalid string with mixed lowercase and uppercase", func(t *testing.T) {
		err := Lowercase("heLLo")
		assert.Error(t, err)
	})

	t.Run("Invalid string with uppercase and non-letters", func(t *testing.T) {
		err := Lowercase("Hello123!")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid lowercase: Hello123!")
	})

	t.Run("Non-string input (integer)", func(t *testing.T) {
		err := Lowercase(12345)
		assert.Error(t, err)
	})

	t.Run("Non-string input (boolean)", func(t *testing.T) {
		err := Lowercase(true)
		assert.Error(t, err)
	})
}
