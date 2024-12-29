package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumeric(t *testing.T) {
	t.Run("valid negative float", func(t *testing.T) {
		err := Numeric("-123.45")
		assert.NoError(t, err)
	})

	t.Run("valid positive integer with sign", func(t *testing.T) {
		err := Numeric("+42")
		assert.NoError(t, err)
	})

	t.Run("valid integer", func(t *testing.T) {
		err := Numeric("0")
		assert.NoError(t, err)
	})

	t.Run("valid positive integer", func(t *testing.T) {
		err := Numeric("123")
		assert.NoError(t, err)
	})

	t.Run("valid negative float", func(t *testing.T) {
		err := Numeric("-0.001")
		assert.NoError(t, err)
	})

	t.Run("valid positive float", func(t *testing.T) {
		err := Numeric("3.14")
		assert.NoError(t, err)
	})

	t.Run("valid positive float", func(t *testing.T) {
		err := Numeric("123.456")
		assert.NoError(t, err)
	})

	t.Run("valid positive float with sign", func(t *testing.T) {
		err := Numeric("+1.0")
		assert.NoError(t, err)
	})

	t.Run("valid negative float with sign", func(t *testing.T) {
		err := Numeric("-1.0")
		assert.NoError(t, err)
	})

	t.Run("invalid: contains letters", func(t *testing.T) {
		err := Numeric("123.45abc")
		assert.EqualError(t, err, "invalid numeric: 123.45abc")
	})

	t.Run("invalid: starts with letters", func(t *testing.T) {
		err := Numeric("abc123")
		assert.EqualError(t, err, "invalid numeric: abc123")
	})

	t.Run("invalid negative float (missing leading zero)", func(t *testing.T) {
		err := Numeric("-.5")
		assert.EqualError(t, err, "invalid numeric: -.5")
	})

	t.Run("invalid: multiple decimal points", func(t *testing.T) {
		err := Numeric("12.34.56")
		assert.EqualError(t, err, "invalid numeric: 12.34.56")
	})

	t.Run("invalid: multiple signs", func(t *testing.T) {
		err := Numeric("++123")
		assert.EqualError(t, err, "invalid numeric: ++123")
	})

	t.Run("invalid: multiple signs", func(t *testing.T) {
		err := Numeric("--123")
		assert.EqualError(t, err, "invalid numeric: --123")
	})

	t.Run("invalid: letters mixed with numbers", func(t *testing.T) {
		err := Numeric("12a34")
		assert.EqualError(t, err, "invalid numeric: 12a34")
	})

	t.Run("invalid: empty string", func(t *testing.T) {
		err := Numeric("")
		assert.EqualError(t, err, "invalid numeric: ")
	})

	t.Run("invalid: nil input", func(t *testing.T) {
		err := Numeric(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("invalid: non-string type", func(t *testing.T) {
		err := Numeric(123)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
