package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumericUnsigned(t *testing.T) {
	t.Run("valid unsigned integer", func(t *testing.T) {
		err := NumericUnsigned("123")
		assert.NoError(t, err)
	})

	t.Run("valid zero", func(t *testing.T) {
		err := NumericUnsigned("0")
		assert.NoError(t, err)
	})

	t.Run("valid unsigned float", func(t *testing.T) {
		err := NumericUnsigned("123.45")
		assert.NoError(t, err)
	})

	t.Run("valid unsigned float with two decimals", func(t *testing.T) {
		err := NumericUnsigned("12.34")
		assert.NoError(t, err)
	})

	t.Run("valid unsigned float with zero decimals", func(t *testing.T) {
		err := NumericUnsigned("0.00")
		assert.NoError(t, err)
	})

	t.Run("valid large unsigned integer", func(t *testing.T) {
		err := NumericUnsigned("123456789")
		assert.NoError(t, err)
	})

	t.Run("valid unsigned float with many decimals", func(t *testing.T) {
		err := NumericUnsigned("123.456789")
		assert.NoError(t, err)
	})

	t.Run("invalid empty string", func(t *testing.T) {
		err := NumericUnsigned("")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid unsigned numeric: ")
	})

	t.Run("invalid contains letters", func(t *testing.T) {
		err := NumericUnsigned("abc")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid unsigned numeric: abc")
	})

	t.Run("invalid negative number", func(t *testing.T) {
		err := NumericUnsigned("-123")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid unsigned numeric: -123")
	})

	t.Run("invalid multiple decimal points", func(t *testing.T) {
		err := NumericUnsigned("12.34.56")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid unsigned numeric: 12.34.56")
	})

	t.Run("invalid nil input", func(t *testing.T) {
		err := NumericUnsigned(nil)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("invalid non-string type", func(t *testing.T) {
		err := NumericUnsigned(123)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
