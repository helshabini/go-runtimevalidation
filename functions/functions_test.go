package functions

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetLen(t *testing.T) {
	t.Run("StringLength", func(t *testing.T) {
		length, err := GetLen("hello")
		assert.NoError(t, err)
		assert.Equal(t, 5, length)
	})

	t.Run("SliceLength", func(t *testing.T) {
		length, err := GetLen([]int{1, 2, 3})
		assert.NoError(t, err)
		assert.Equal(t, 3, length)
	})

	t.Run("ArrayLength", func(t *testing.T) {
		length, err := GetLen([3]int{1, 2, 3})
		assert.NoError(t, err)
		assert.Equal(t, 3, length)
	})

	t.Run("MapLength", func(t *testing.T) {
		length, err := GetLen(map[string]int{"a": 1, "b": 2})
		assert.NoError(t, err)
		assert.Equal(t, 2, length)
	})

	t.Run("EmptyString", func(t *testing.T) {
		length, err := GetLen("")
		assert.NoError(t, err)
		assert.Equal(t, 0, length)
	})

	t.Run("EmptySlice", func(t *testing.T) {
		length, err := GetLen([]int{})
		assert.NoError(t, err)
		assert.Equal(t, 0, length)
	})

	t.Run("UnsupportedType", func(t *testing.T) {
		length, err := GetLen(123)
		assert.Error(t, err)
		assert.Equal(t, 0, length)
	})
}

func TestGetIntAny(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
		value, err := GetInt(42)
		assert.NoError(t, err)
		assert.Equal(t, int64(42), value)
	})

	t.Run("Int64", func(t *testing.T) {
		value, err := GetInt(int64(12345))
		assert.NoError(t, err)
		assert.Equal(t, int64(12345), value)
	})

	t.Run("Uint", func(t *testing.T) {
		value, err := GetInt(uint(123))
		assert.NoError(t, err)
		assert.Equal(t, int64(123), value)
	})

	t.Run("StringInt", func(t *testing.T) {
		value, err := GetInt("12345")
		assert.NoError(t, err)
		assert.Equal(t, int64(12345), value)
	})

	t.Run("InvalidString", func(t *testing.T) {
		value, err := GetInt("invalid")
		assert.Error(t, err)
		assert.Equal(t, int64(0), value)
	})

	t.Run("TimeDuration", func(t *testing.T) {
		value, err := GetInt(time.Duration(3 * time.Second))
		assert.NoError(t, err)
		assert.Equal(t, int64(3_000_000_000), value)
	})

	t.Run("TimeObject", func(t *testing.T) {
		value, err := GetInt(time.Date(2024, time.October, 5, 15, 4, 5, 0, time.UTC))
		assert.NoError(t, err)
		assert.Equal(t, int64(1728140645), value)
	})

	t.Run("RFC3339String", func(t *testing.T) {
		value, err := GetInt("2024-10-05T15:04:05Z")
		assert.NoError(t, err)
		assert.Equal(t, int64(1728140645), value)
	})

	t.Run("UnsupportedType", func(t *testing.T) {
		value, err := GetInt(3.14)
		assert.Error(t, err)
		assert.Equal(t, int64(0), value)
	})
}
