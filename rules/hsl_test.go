package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHSL(t *testing.T) {
	t.Run("valid black color", func(t *testing.T) {
		err := HSL("hsl(0, 0%, 0%)")
		assert.NoError(t, err)
	})

	t.Run("valid white color", func(t *testing.T) {
		err := HSL("hsl(360, 100%, 100%)")
		assert.NoError(t, err)
	})

	t.Run("valid green color", func(t *testing.T) {
		err := HSL("hsl(120, 50%, 50%)")
		assert.NoError(t, err)
	})

	t.Run("valid blue color", func(t *testing.T) {
		err := HSL("hsl(240, 100%, 50%)")
		assert.NoError(t, err)
	})

	t.Run("valid teal color", func(t *testing.T) {
		err := HSL("hsl(180, 50%, 25%)")
		assert.NoError(t, err)
	})

	t.Run("invalid hue 361", func(t *testing.T) {
		err := HSL("hsl(361, 100%, 100%)")
		assert.EqualError(t, err, "invalid hsl value: hsl(361, 100%, 100%)")
	})

	t.Run("invalid saturation 101%", func(t *testing.T) {
		err := HSL("hsl(360, 101%, 100%)")
		assert.EqualError(t, err, "invalid hsl value: hsl(360, 101%, 100%)")
	})

	t.Run("invalid lightness 101%", func(t *testing.T) {
		err := HSL("hsl(360, 100%, 101%)")
		assert.EqualError(t, err, "invalid hsl value: hsl(360, 100%, 101%)")
	})

	t.Run("invalid hue -1", func(t *testing.T) {
		err := HSL("hsl(-1, 100%, 50%)")
		assert.EqualError(t, err, "invalid hsl value: hsl(-1, 100%, 50%)")
	})

	t.Run("invalid format", func(t *testing.T) {
		err := HSL("invalid")
		assert.EqualError(t, err, "invalid hsl value: invalid")
	})

	t.Run("empty string", func(t *testing.T) {
		err := HSL("")
		assert.EqualError(t, err, "invalid hsl value: ")
	})

	t.Run("nil input", func(t *testing.T) {
		err := HSL(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("non-string type", func(t *testing.T) {
		err := HSL(123)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
