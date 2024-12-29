package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMD5(t *testing.T) {
	t.Run("Valid MD5 hash", func(t *testing.T) {
		err := MD5("d41d8cd98f00b204e9800998ecf8427e") // Empty string MD5
		assert.NoError(t, err)
	})

	t.Run("Valid MD5 hash with uppercase letters", func(t *testing.T) {
		err := MD5("D41D8CD98F00B204E9800998ECF8427E") // Empty string MD5 in uppercase
		assert.NoError(t, err)
	})

	t.Run("Invalid MD5 - too short", func(t *testing.T) {
		err := MD5("d41d8cd98f00b204e9800998ecf8427") // Only 31 characters
		assert.Error(t, err)
		assert.Equal(t, "invalid MD5 value: d41d8cd98f00b204e9800998ecf8427", err.Error())
	})

	t.Run("Invalid MD5 - non-hexadecimal characters", func(t *testing.T) {
		err := MD5("z41d8cd98f00b204e9800998ecf8427e") // Invalid hex character 'z'
		assert.Error(t, err)
		assert.Equal(t, "invalid MD5 value: z41d8cd98f00b204e9800998ecf8427e", err.Error())
	})

	t.Run("Invalid input type - integer", func(t *testing.T) {
		err := MD5(12345)
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got int", err.Error())
	})

	t.Run("Invalid input type - struct", func(t *testing.T) {
		err := MD5(struct{}{})
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got struct {}", err.Error())
	})
}
