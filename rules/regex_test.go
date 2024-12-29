package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegex(t *testing.T) {
	t.Run("Valid input matches regex pattern", func(t *testing.T) {
		err := Regex("hello123", nil, map[string]args.Arg{"pattern": {Value: "^hello[0-9]+$"}})
		assert.NoError(t, err)
	})

	t.Run("Input does not match regex pattern", func(t *testing.T) {
		err := Regex("helloWorld", nil, map[string]args.Arg{"pattern": {Value: "^hello[0-9]+$"}})
		assert.Error(t, err)
	})

	t.Run("Valid input matches regex with special characters", func(t *testing.T) {
		err := Regex("abc@domain.com", nil, map[string]args.Arg{"pattern": {Value: `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`}})
		assert.NoError(t, err)
	})

	t.Run("Invalid regex pattern", func(t *testing.T) {
		err := Regex("hello123", nil, map[string]args.Arg{"pattern": {Value: "("}})
		assert.Error(t, err)
	})

	t.Run("Invalid input type", func(t *testing.T) {
		err := Regex(12345, nil, map[string]args.Arg{"pattern": {Value: "^\\d+$"}})
		assert.Error(t, err)
	})

	t.Run("Non-string pattern (converted to string)", func(t *testing.T) {
		err := Regex("hello123", nil, map[string]args.Arg{"pattern": {Value: 123}})
		assert.NoError(t, err) // Since "hello123" contains "123", this should pass.
	})

	t.Run("Empty input string", func(t *testing.T) {
		err := Regex("", nil, map[string]args.Arg{"pattern": {Value: "^hello[0-9]+$"}})
		assert.Error(t, err)
	})

	t.Run("No matching pattern provided", func(t *testing.T) {
		err := Regex("hello123", nil, map[string]args.Arg{})
		assert.Error(t, err)
	})
}
