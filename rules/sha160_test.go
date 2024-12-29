package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA160(t *testing.T) {
	t.Run("Valid SHA160", func(t *testing.T) {
		err := SHA160("a3f390d88e4c41f2747bfa2f1b5f87db39bc0ce5")
		assert.NoError(t, err)
	})

	t.Run("Invalid SHA160 - wrong length", func(t *testing.T) {
		err := SHA160("a3f390d88e4c41f2747bfa2f1b5f87db39bc0")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA160 value: a3f390d88e4c41f2747bfa2f1b5f87db39bc0", err.Error())
	})

	t.Run("Invalid SHA160 - contains non-hex characters", func(t *testing.T) {
		err := SHA160("a3f390d88e4c41g2747bfa2f1b5f87db39bc0ce5") // 'g' is not a valid hex character
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA160 value: a3f390d88e4c41g2747bfa2f1b5f87db39bc0ce5", err.Error())
	})

	t.Run("Invalid SHA160 - not a string", func(t *testing.T) {
		err := SHA160(12345)
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got int", err.Error())
	})
}
