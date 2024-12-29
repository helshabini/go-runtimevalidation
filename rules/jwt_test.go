package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	t.Run("valid JWT", func(t *testing.T) {
		err := JWT("valid.jwt.token")
		assert.NoError(t, err)
	})

	t.Run("empty JWT", func(t *testing.T) {
		err := JWT("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid jwt")
	})

	t.Run("unsupported type", func(t *testing.T) {
		err := JWT(12345)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid jwt")
	})

	t.Run("invalid JWT format", func(t *testing.T) {
		err := JWT("invalid.jwt")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid jwt")
	})
}
