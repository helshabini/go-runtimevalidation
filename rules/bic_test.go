package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBIC(t *testing.T) {
	t.Run("Valid BIC (8 characters)", func(t *testing.T) {
		err := BIC("DEUTDEFF")
		assert.NoError(t, err)
	})

	t.Run("Valid BIC (11 characters)", func(t *testing.T) {
		err := BIC("DEUTDEFF500")
		assert.NoError(t, err)
	})

	t.Run("Invalid BIC (too short)", func(t *testing.T) {
		err := BIC("DEUTDE")
		assert.Error(t, err)
		assert.Equal(t, "invalid bic: DEUTDE", err.Error())
	})

	t.Run("Invalid BIC (invalid characters)", func(t *testing.T) {
		err := BIC("DEUT@EFF")
		assert.Error(t, err)
		assert.Equal(t, "invalid bic: DEUT@EFF", err.Error())
	})

	t.Run("Invalid BIC (too long)", func(t *testing.T) {
		err := BIC("DEUTDEFF500000")
		assert.Error(t, err)
		assert.Equal(t, "invalid bic: DEUTDEFF500000", err.Error())
	})
}
