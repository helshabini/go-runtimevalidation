package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA224(t *testing.T) {
	t.Run("Valid SHA224 Hash", func(t *testing.T) {
		err := SHA224("d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42f")
		assert.NoError(t, err)
	})

	t.Run("Invalid SHA224 Hash (too short)", func(t *testing.T) {
		err := SHA224("d14a028c2a3a2bc9476102bb")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA224 value: d14a028c2a3a2bc9476102bb", err.Error())
	})

	t.Run("Invalid SHA224 Hash (too long)", func(t *testing.T) {
		err := SHA224("d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42fffff")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA224 value: d14a028c2a3a2bc9476102bb288234c415a2b01f828ea62ac5b3e42fffff", err.Error())
	})

	t.Run("Invalid SHA224 Hash (non-hex characters)", func(t *testing.T) {
		err := SHA224("d14a028c2a3a2bc9476102bb288234c415a2b01g828ea62ac5b3e42f")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA224 value: d14a028c2a3a2bc9476102bb288234c415a2b01g828ea62ac5b3e42f", err.Error())
	})

	t.Run("Invalid SHA224 Hash (input is not a string)", func(t *testing.T) {
		err := SHA224(12345)
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got int", err.Error())
	})
}
