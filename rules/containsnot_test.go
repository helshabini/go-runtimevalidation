package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNot(t *testing.T) {
	t.Run("Valid string does not contain substring", func(t *testing.T) {
		err := ContainsNot("hello world", nil, map[string]args.Arg{"arg": {Value: "planet"}})
		assert.NoError(t, err)
	})

	t.Run("Valid string does not contain single character", func(t *testing.T) {
		err := ContainsNot("hello world", nil, map[string]args.Arg{"arg": {Value: "x"}})
		assert.NoError(t, err)
	})

	t.Run("Valid string with number does not contain substring", func(t *testing.T) {
		err := ContainsNot("12345", nil, map[string]args.Arg{"arg": {Value: "678"}})
		assert.NoError(t, err)
	})

	t.Run("Invalid string contains substring", func(t *testing.T) {
		err := ContainsNot("hello world", nil, map[string]args.Arg{"arg": {Value: "world"}})
		assert.Error(t, err)
	})

	t.Run("Valid string with symbols does not contain substring", func(t *testing.T) {
		err := ContainsNot("hello@world.com", nil, map[string]args.Arg{"arg": {Value: "#"}})
		assert.NoError(t, err)
	})

	t.Run("Invalid argument type", func(t *testing.T) {
		err := ContainsNot("hello world", nil, map[string]args.Arg{"arg": {Value: 123}})
		assert.NoError(t, err)
	})

	t.Run("Invalid input type", func(t *testing.T) {
		err := ContainsNot(12345, nil, map[string]args.Arg{"arg": {Value: "435"}})
		assert.NoError(t, err)
	})

	t.Run("ContainsNot fails with empty string", func(t *testing.T) {
		err := ContainsNot("hello world", nil, map[string]args.Arg{"arg": {Value: ""}})
		assert.Error(t, err)
	})

	t.Run("Empty input and empty substring", func(t *testing.T) {
		err := ContainsNot("", nil, map[string]args.Arg{"arg": {Value: ""}})
		assert.Error(t, err)
	})
}
