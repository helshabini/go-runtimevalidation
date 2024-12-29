package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase32Hex(t *testing.T) {
	t.Run("valid Base32Hex string", func(t *testing.T) {
		err := Base32Hex("91IMOR3F41BMUSJCCG======") // valid Base32Hex encoded string
		assert.NoError(t, err)
	})

	t.Run("invalid Base32Hex string", func(t *testing.T) {
		err := Base32Hex("InvalidBase32HexString!")
		assert.EqualError(t, err, "invalid Base32Hex string: InvalidBase32HexString!")
	})

	t.Run("non-string input", func(t *testing.T) {
		err := Base32Hex(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
