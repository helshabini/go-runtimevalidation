package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRGB(t *testing.T) {
	t.Run("valid black color", func(t *testing.T) {
		err := RGB("rgb(0, 0, 0)")
		assert.NoError(t, err)
	})

	t.Run("valid white color", func(t *testing.T) {
		err := RGB("rgb(255, 255, 255)")
		assert.NoError(t, err)
	})

	t.Run("valid gray color", func(t *testing.T) {
		err := RGB("rgb(128, 128, 128)")
		assert.NoError(t, err)
	})

	t.Run("valid red color", func(t *testing.T) {
		err := RGB("rgb(255, 0, 0)")
		assert.NoError(t, err)
	})

	t.Run("valid green color", func(t *testing.T) {
		err := RGB("rgb(0, 255, 0)")
		assert.NoError(t, err)
	})

	t.Run("valid blue color", func(t *testing.T) {
		err := RGB("rgb(0, 0, 255)")
		assert.NoError(t, err)
	})

	t.Run("invalid RGB value 256", func(t *testing.T) {
		err := RGB("rgb(256, 0, 0)")
		assert.EqualError(t, err, "invalid rgb value: rgb(256, 0, 0)")
	})

	t.Run("invalid RGB value -1", func(t *testing.T) {
		err := RGB("rgb(0, 0, -1)")
		assert.EqualError(t, err, "invalid rgb value: rgb(0, 0, -1)")
	})

	t.Run("invalid format with four values", func(t *testing.T) {
		err := RGB("rgb(0, 0, 0, 0)")
		assert.EqualError(t, err, "invalid rgb value: rgb(0, 0, 0, 0)")
	})

	t.Run("invalid format with four values no spaces", func(t *testing.T) {
		err := RGB("rgb(0,0,0,0)")
		assert.EqualError(t, err, "invalid rgb value: rgb(0,0,0,0)")
	})

	t.Run("invalid format string", func(t *testing.T) {
		err := RGB("invalid")
		assert.EqualError(t, err, "invalid rgb value: invalid")
	})

	t.Run("empty string", func(t *testing.T) {
		err := RGB("")
		assert.EqualError(t, err, "invalid rgb value: ")
	})

	t.Run("nil input", func(t *testing.T) {
		err := RGB(nil)
		assert.EqualError(t, err, "expected a string, got <nil>")
	})

	t.Run("non-string type", func(t *testing.T) {
		err := RGB(123)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
