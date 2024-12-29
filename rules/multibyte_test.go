package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiByte(t *testing.T) {
	t.Run("Valid MultiByte characters (Japanese)", func(t *testing.T) {
		err := MultiByte("こんにちは") // Japanese characters
		assert.NoError(t, err)
	})

	t.Run("Valid MultiByte characters (Chinese)", func(t *testing.T) {
		err := MultiByte("你好") // Chinese characters
		assert.NoError(t, err)
	})

	t.Run("Valid MultiByte characters with spaces", func(t *testing.T) {
		err := MultiByte("こんにちは 你好") // Japanese and Chinese characters with space
		assert.NoError(t, err)
	})

	t.Run("Invalid input with ASCII characters", func(t *testing.T) {
		err := MultiByte("Hello") // ASCII characters
		assert.Error(t, err)
	})

	t.Run("Invalid input with mixed ASCII and MultiByte", func(t *testing.T) {
		err := MultiByte("こんにちはHello") // Mixed ASCII and Japanese
		assert.Error(t, err)
	})

	t.Run("Input not a string", func(t *testing.T) {
		err := MultiByte(12345) // non-string input
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "expected a string")
	})
}
