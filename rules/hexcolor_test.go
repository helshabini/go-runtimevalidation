package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexColor(t *testing.T) {
	t.Run("valid 3-digit hex color", func(t *testing.T) {
		err := HexColor("#FFF")
		assert.NoError(t, err)
	})

	t.Run("valid 4-digit hex color", func(t *testing.T) {
		err := HexColor("#FFFF")
		assert.NoError(t, err)
	})

	t.Run("valid 6-digit hex color", func(t *testing.T) {
		err := HexColor("#FFFFFF")
		assert.NoError(t, err)
	})

	t.Run("valid 8-digit hex color", func(t *testing.T) {
		err := HexColor("#FFFFFFFF")
		assert.NoError(t, err)
	})

	t.Run("valid 3-digit hex color", func(t *testing.T) {
		err := HexColor("#123")
		assert.NoError(t, err)
	})

	t.Run("valid 4-digit hex color", func(t *testing.T) {
		err := HexColor("#1234")
		assert.NoError(t, err)
	})

	t.Run("valid 6-digit hex color", func(t *testing.T) {
		err := HexColor("#123456")
		assert.NoError(t, err)
	})

	t.Run("valid 8-digit hex color", func(t *testing.T) {
		err := HexColor("#12345678")
		assert.NoError(t, err)
	})

	t.Run("invalid: contains invalid characters", func(t *testing.T) {
		err := HexColor("#XYZ")
		assert.EqualError(t, err, "invalid hex color: #XYZ")
	})

	t.Run("invalid: contains invalid characters", func(t *testing.T) {
		err := HexColor("#1234567G")
		assert.EqualError(t, err, "invalid hex color: #1234567G")
	})

	t.Run("invalid: too many digits", func(t *testing.T) {
		err := HexColor("#123456789")
		assert.EqualError(t, err, "invalid hex color: #123456789")
	})

	t.Run("invalid: empty string", func(t *testing.T) {
		err := HexColor("")
		assert.EqualError(t, err, "invalid hex color: ")
	})

	t.Run("invalid: nil input", func(t *testing.T) {
		err := HexColor(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("invalid: non-string type", func(t *testing.T) {
		err := HexColor(123)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
