package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD4(t *testing.T) {
	t.Run("Valid MD4 hash", func(t *testing.T) {
		err := MD4("a4b2c2a19b1a3b1a10e0c5b8a47e38fc") // Example valid MD4 hash
		assert.NoError(t, err)
	})

	t.Run("Valid MD4 hash with uppercase letters", func(t *testing.T) {
		err := MD4("A4B2C2A19B1A3B1A10E0C5B8A47E38FC") // Example valid MD4 hash in uppercase
		assert.NoError(t, err)
	})

	t.Run("Invalid MD4 - too short", func(t *testing.T) {
		err := MD4("a4b2c2a19b1a3b1a10e0c5b8a47e38f") // Only 31 characters
		assert.Error(t, err)
		assert.Equal(t, "invalid MD4 value: a4b2c2a19b1a3b1a10e0c5b8a47e38f", err.Error())
	})

	t.Run("Invalid MD4 - non-hexadecimal characters", func(t *testing.T) {
		err := MD4("g4b2c2a19b1a3b1a10e0c5b8a47e38fc") // Invalid hex character 'g'
		assert.Error(t, err)
		assert.Equal(t, "invalid MD4 value: g4b2c2a19b1a3b1a10e0c5b8a47e38fc", err.Error())
	})

	t.Run("Invalid input type - integer", func(t *testing.T) {
		err := MD4(12345)
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got int", err.Error())
	})

	t.Run("Invalid input type - struct", func(t *testing.T) {
		err := MD4(struct{}{})
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got struct {}", err.Error())
	})
}
