package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexadecimal(t *testing.T) {
	t.Run("valid hexadecimal with prefix", func(t *testing.T) {
		err := Hexadecimal("0x1A3F")
		assert.NoError(t, err)
	})

	t.Run("valid hexadecimal without prefix", func(t *testing.T) {
		err := Hexadecimal("1A3F")
		assert.NoError(t, err)
	})

	t.Run("valid hexadecimal with uppercase prefix", func(t *testing.T) {
		err := Hexadecimal("0X4B2C")
		assert.NoError(t, err)
	})

	t.Run("valid hexadecimal without prefix", func(t *testing.T) {
		err := Hexadecimal("4B2C")
		assert.NoError(t, err)
	})

	t.Run("invalid: empty after prefix", func(t *testing.T) {
		err := Hexadecimal("0x")
		assert.EqualError(t, err, "invalid hexadecimal: 0x")
	})

	t.Run("invalid: empty string", func(t *testing.T) {
		err := Hexadecimal("")
		assert.EqualError(t, err, "invalid hexadecimal: ")
	})

	t.Run("invalid: contains invalid characters", func(t *testing.T) {
		err := Hexadecimal("GHIJK")
		assert.EqualError(t, err, "invalid hexadecimal: GHIJK")
	})

	t.Run("valid hexadecimal", func(t *testing.T) {
		err := Hexadecimal("123")
		assert.NoError(t, err)
	})

	t.Run("valid hexadecimal with uppercase letters", func(t *testing.T) {
		err := Hexadecimal("0xABCDEF")
		assert.NoError(t, err)
	})

	t.Run("invalid: nil input", func(t *testing.T) {
		err := Hexadecimal(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("invalid: non-string type", func(t *testing.T) {
		err := Hexadecimal(123)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
