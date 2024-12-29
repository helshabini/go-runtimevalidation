package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase32(t *testing.T) {
	t.Run("valid Base32 string", func(t *testing.T) {
		err := Base32("JBSWY3DPEHPK3PXP") // valid Base32 encoded string
		assert.NoError(t, err)
	})

	t.Run("invalid Base32 string", func(t *testing.T) {
		err := Base32("InvalidBase32String!")
		assert.EqualError(t, err, "invalid Base32 string: InvalidBase32String!")
	})

	t.Run("non-string input", func(t *testing.T) {
		err := Base32(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
