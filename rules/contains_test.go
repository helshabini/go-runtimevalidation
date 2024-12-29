package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	t.Run("Valid string contains substring", func(t *testing.T) {
		err := Contains("hello world", nil, map[string]args.Arg{"arg": {Value: "world"}})
		assert.NoError(t, err)
	})

	t.Run("Valid string contains single character", func(t *testing.T) {
		err := Contains("hello world", nil, map[string]args.Arg{"arg": {Value: "o"}})
		assert.NoError(t, err)
	})

	t.Run("Valid string with number contains substring", func(t *testing.T) {
		err := Contains("12345", nil, map[string]args.Arg{"arg": {Value: "345"}})
		assert.NoError(t, err)
	})

	t.Run("Invalid string does not contain substring", func(t *testing.T) {
		err := Contains("hello world", nil, map[string]args.Arg{"arg": {Value: "planet"}})
		assert.Error(t, err)
	})

	t.Run("Valid string with symbols contains substring", func(t *testing.T) {
		err := Contains("hello@world.com", nil, map[string]args.Arg{"arg": {Value: "@world"}})
		assert.NoError(t, err)
	})

	t.Run("Invalid argument type", func(t *testing.T) {
		err := Contains("hello world", nil, map[string]args.Arg{"arg": {Value: 123}})
		assert.Error(t, err)
	})

	t.Run("Invalid input type", func(t *testing.T) {
		err := Contains(12345, nil, map[string]args.Arg{"arg": {Value: "45"}})
		assert.NoError(t, err)
	})

	t.Run("Contains validation fails with empty string", func(t *testing.T) {
		err := Contains("hello world", nil, map[string]args.Arg{"arg": {Value: ""}})
		assert.Error(t, err)
	})

	t.Run("Empty input and empty substring", func(t *testing.T) {
		err := Contains("", nil, map[string]args.Arg{"arg": {Value: ""}})
		assert.Error(t, err)
	})
}
