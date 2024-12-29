package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongitude(t *testing.T) {
	t.Run("Valid longitude - integer value", func(t *testing.T) {
		err := Longitude("45")
		assert.NoError(t, err)
	})

	t.Run("Valid longitude - positive decimal value", func(t *testing.T) {
		err := Longitude("45.123")
		assert.NoError(t, err)
	})

	t.Run("Valid longitude - negative decimal value", func(t *testing.T) {
		err := Longitude("-45.123")
		assert.NoError(t, err)
	})

	t.Run("Valid longitude - boundary positive", func(t *testing.T) {
		err := Longitude("180")
		assert.NoError(t, err)
	})

	t.Run("Valid longitude - boundary negative", func(t *testing.T) {
		err := Longitude("-180")
		assert.NoError(t, err)
	})

	t.Run("Invalid longitude - out of range positive", func(t *testing.T) {
		err := Longitude("181")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid longitude value: 181")
	})

	t.Run("Invalid longitude - out of range negative", func(t *testing.T) {
		err := Longitude("-181")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid longitude value: -181")
	})

	t.Run("Invalid longitude - non-numeric string", func(t *testing.T) {
		err := Longitude("abc")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid longitude value: abc")
	})

	t.Run("Invalid longitude - empty string", func(t *testing.T) {
		err := Longitude("")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid longitude value: ")
	})

	t.Run("Invalid longitude - wrong type (integer)", func(t *testing.T) {
		err := Longitude(45)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("Invalid longitude - wrong type (float)", func(t *testing.T) {
		err := Longitude(45.123)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got float64")
	})
}
