package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatitude(t *testing.T) {
	t.Run("Valid latitude - integer value", func(t *testing.T) {
		err := Latitude("45")
		assert.NoError(t, err)
	})

	t.Run("Valid latitude - positive decimal value", func(t *testing.T) {
		err := Latitude("45.123")
		assert.NoError(t, err)
	})

	t.Run("Valid latitude - negative decimal value", func(t *testing.T) {
		err := Latitude("-45.123")
		assert.NoError(t, err)
	})

	t.Run("Valid latitude - boundary positive", func(t *testing.T) {
		err := Latitude("90")
		assert.NoError(t, err)
	})

	t.Run("Valid latitude - boundary negative", func(t *testing.T) {
		err := Latitude("-90")
		assert.NoError(t, err)
	})

	t.Run("Invalid latitude - out of range positive", func(t *testing.T) {
		err := Latitude("91")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid latitude value: 91")
	})

	t.Run("Invalid latitude - out of range negative", func(t *testing.T) {
		err := Latitude("-91")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid latitude value: -91")
	})

	t.Run("Invalid latitude - non-numeric string", func(t *testing.T) {
		err := Latitude("abc")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid latitude value: abc")
	})

	t.Run("Invalid latitude - empty string", func(t *testing.T) {
		err := Latitude("")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid latitude value: ")
	})

	t.Run("Invalid latitude - wrong type (integer)", func(t *testing.T) {
		err := Latitude(45)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("Invalid latitude - wrong type (float)", func(t *testing.T) {
		err := Latitude(45.123)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got float64")
	})
}
