package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHSLA(t *testing.T) {
	t.Run("valid black color with 0 alpha", func(t *testing.T) {
		err := HSLA("hsla(0, 0%, 0%, 0)")
		assert.NoError(t, err)
	})

	t.Run("valid white color with 1 alpha", func(t *testing.T) {
		err := HSLA("hsla(360, 100%, 100%, 1)")
		assert.NoError(t, err)
	})

	t.Run("valid green color with 0.5 alpha", func(t *testing.T) {
		err := HSLA("hsla(120, 50%, 50%, 0.5)")
		assert.NoError(t, err)
	})

	t.Run("valid blue color with 1.0 alpha", func(t *testing.T) {
		err := HSLA("hsla(240, 100%, 50%, 1.0)")
		assert.NoError(t, err)
	})

	t.Run("valid teal color with 0.1 alpha", func(t *testing.T) {
		err := HSLA("hsla(180, 50%, 25%, 0.1)")
		assert.NoError(t, err)
	})

	t.Run("invalid hue", func(t *testing.T) {
		err := HSLA("hsla(361, 100%, 100%, 1)")
		assert.EqualError(t, err, "invalid hsla value: hsla(361, 100%, 100%, 1)")
	})

	t.Run("invalid saturation", func(t *testing.T) {
		err := HSLA("hsla(360, 101%, 100%, 1)")
		assert.EqualError(t, err, "invalid hsla value: hsla(360, 101%, 100%, 1)")
	})

	t.Run("invalid lightness", func(t *testing.T) {
		err := HSLA("hsla(360, 100%, 101%, 1)")
		assert.EqualError(t, err, "invalid hsla value: hsla(360, 100%, 101%, 1)")
	})

	t.Run("invalid alpha 1.1", func(t *testing.T) {
		err := HSLA("hsla(360, 100%, 50%, 1.1)")
		assert.EqualError(t, err, "invalid hsla value: hsla(360, 100%, 50%, 1.1)")
	})

	t.Run("invalid alpha 1.01", func(t *testing.T) {
		err := HSLA("hsla(360, 100%, 50%, 1.01)")
		assert.EqualError(t, err, "invalid hsla value: hsla(360, 100%, 50%, 1.01)")
	})

	t.Run("invalid hue -1", func(t *testing.T) {
		err := HSLA("hsla(-1, 100%, 50%, 0)")
		assert.EqualError(t, err, "invalid hsla value: hsla(-1, 100%, 50%, 0)")
	})

	t.Run("invalid format", func(t *testing.T) {
		err := HSLA("invalid")
		assert.EqualError(t, err, "invalid hsla value: invalid")
	})

	t.Run("empty string", func(t *testing.T) {
		err := HSLA("")
		assert.EqualError(t, err, "invalid hsla value: ")
	})

	t.Run("nil input", func(t *testing.T) {
		err := HSLA(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("non-string type", func(t *testing.T) {
		err := HSLA(123)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
