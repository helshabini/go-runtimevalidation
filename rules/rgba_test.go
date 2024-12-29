package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRGBA(t *testing.T) {
	t.Run("valid black color with 0 alpha", func(t *testing.T) {
		err := RGBA("rgba(0, 0, 0, 0)")
		assert.NoError(t, err)
	})

	t.Run("valid white color with full alpha", func(t *testing.T) {
		err := RGBA("rgba(255, 255, 255, 1)")
		assert.NoError(t, err)
	})

	t.Run("valid white color with 10% alpha", func(t *testing.T) {
		err := RGBA("rgba(255, 255, 255, 0.1)")
		assert.NoError(t, err)
	})

	t.Run("valid white color with 50% alpha", func(t *testing.T) {
		err := RGBA("rgba(255, 255, 255, 0.5)")
		assert.NoError(t, err)
	})

	t.Run("valid white color with 100% alpha", func(t *testing.T) {
		err := RGBA("rgba(255, 255, 255, 1.0)")
		assert.NoError(t, err)
	})

	t.Run("invalid alpha value 1.1", func(t *testing.T) {
		err := RGBA("rgba(255, 255, 255, 1.1)")
		assert.EqualError(t, err, "invalid rgba value: rgba(255, 255, 255, 1.1)")
	})

	t.Run("invalid alpha value 1.01", func(t *testing.T) {
		err := RGBA("rgba(255, 255, 255, 1.01)")
		assert.EqualError(t, err, "invalid rgba value: rgba(255, 255, 255, 1.01)")
	})

	t.Run("valid white color with 0 alpha", func(t *testing.T) {
		err := RGBA("rgba(255, 255, 255, 0.0)")
		assert.NoError(t, err)
	})

	t.Run("valid white color with 0 alpha (integer)", func(t *testing.T) {
		err := RGBA("rgba(255, 255, 255, 0)")
		assert.NoError(t, err)
	})

	t.Run("invalid RGB value 256", func(t *testing.T) {
		err := RGBA("rgba(256, 0, 0, 0)")
		assert.EqualError(t, err, "invalid rgba value: rgba(256, 0, 0, 0)")
	})

	t.Run("invalid alpha value 300", func(t *testing.T) {
		err := RGBA("rgba(0, 0, 0, 300)")
		assert.EqualError(t, err, "invalid rgba value: rgba(0, 0, 0, 300)")
	})

	t.Run("invalid alpha value -10", func(t *testing.T) {
		err := RGBA("rgba(0, 0, 0, -10)")
		assert.EqualError(t, err, "invalid rgba value: rgba(0, 0, 0, -10)")
	})

	t.Run("invalid format", func(t *testing.T) {
		err := RGBA("invalid")
		assert.EqualError(t, err, "invalid rgba value: invalid")
	})

	t.Run("empty string", func(t *testing.T) {
		err := RGBA("")
		assert.EqualError(t, err, "invalid rgba value: ")
	})

	t.Run("nil input", func(t *testing.T) {
		err := RGBA(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("non-string type", func(t *testing.T) {
		err := RGBA(123)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
