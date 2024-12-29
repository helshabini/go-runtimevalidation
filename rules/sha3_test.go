package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA3(t *testing.T) {
	t.Run("Valid SHA3 Hash (56 chars)", func(t *testing.T) {
		err := SHA3("a3dcb4d229de6fde0db5686dee47145d5b914a565a9c196d3bf56f8a")
		assert.NoError(t, err)
	})

	t.Run("Valid SHA3 Hash (64 chars)", func(t *testing.T) {
		err := SHA3("a3dcb4d229de6fde0db5686dee47145d5b914a565a9c196d3bf56f8a88583e80")
		assert.NoError(t, err)
	})

	t.Run("Valid SHA3 Hash (96 chars)", func(t *testing.T) {
		err := SHA3("b5a219edd7823edf81f18b32b296b17810005bf782cf0d70658d423dcc01edb431ee91e6764109ff9cccd0fe688e9bc1")
		assert.NoError(t, err)
	})

	t.Run("Valid SHA3 Hash (128 chars)", func(t *testing.T) {
		err := SHA3("fdd2a83bc4ddba0f6ce67ddf9b3fdbc14f9f6a53faefa5b486b4c7eabf7b8dd2fbe75d01e25fa0c76c8658d56ed7b2079c393efdb04e62394936572e9c0b76b2")
		assert.NoError(t, err)
	})

	t.Run("Invalid SHA3 Hash (too short)", func(t *testing.T) {
		err := SHA3("123")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA3 value: 123", err.Error())
	})

	t.Run("Invalid SHA3 Hash (non-hex characters)", func(t *testing.T) {
		err := SHA3("g3dcb4d229de6fde0db5686dee47145d5b914a565a9c196d3bf56f8a88583e80")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA3 value: g3dcb4d229de6fde0db5686dee47145d5b914a565a9c196d3bf56f8a88583e80", err.Error())
	})

	t.Run("Invalid SHA3 Hash (input is not a string)", func(t *testing.T) {
		err := SHA3(12345)
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got int", err.Error())
	})
}
