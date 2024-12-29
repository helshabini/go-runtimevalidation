package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequired(t *testing.T) {
	t.Run("Valid non-zero integer", func(t *testing.T) {
		err := Required(42)
		assert.NoError(t, err)
	})

	t.Run("Zero integer", func(t *testing.T) {
		err := Required(0)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Valid non-empty string", func(t *testing.T) {
		err := Required("hello")
		assert.NoError(t, err)
	})

	t.Run("Empty string", func(t *testing.T) {
		err := Required("")
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Valid non-nil slice", func(t *testing.T) {
		err := Required([]int{1, 2, 3})
		assert.NoError(t, err)
	})

	t.Run("Nil slice", func(t *testing.T) {
		var s []int
		err := Required(s)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Valid non-nil map", func(t *testing.T) {
		err := Required(map[string]int{"key": 1})
		assert.NoError(t, err)
	})

	t.Run("Nil map", func(t *testing.T) {
		var m map[string]int
		err := Required(m)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Valid non-nil pointer", func(t *testing.T) {
		val := 10
		err := Required(&val)
		assert.NoError(t, err)
	})

	t.Run("Nil pointer", func(t *testing.T) {
		var ptr *int
		err := Required(ptr)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Valid non-nil interface", func(t *testing.T) {
		var iface any = "non-empty"
		err := Required(iface)
		assert.NoError(t, err)
	})

	t.Run("Nil interface", func(t *testing.T) {
		var iface any
		err := Required(iface)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Non-zero struct", func(t *testing.T) {
		type Example struct {
			Field int
		}
		err := Required(Example{Field: 1})
		assert.NoError(t, err)
	})

	t.Run("Zero struct", func(t *testing.T) {
		type Example struct {
			Field int
		}
		err := Required(Example{})
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Valid non-nil channel", func(t *testing.T) {
		ch := make(chan int)
		err := Required(ch)
		assert.NoError(t, err)
	})

	t.Run("Nil channel", func(t *testing.T) {
		var ch chan int
		err := Required(ch)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Non-zero float", func(t *testing.T) {
		err := Required(3.14)
		assert.NoError(t, err)
	})

	t.Run("Zero float", func(t *testing.T) {
		err := Required(0.0)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("Non-nil function", func(t *testing.T) {
		fn := func() {}
		err := Required(fn)
		assert.NoError(t, err)
	})

	t.Run("Nil function", func(t *testing.T) {
		var fn func()
		err := Required(fn)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})
}
