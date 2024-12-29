package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartsNotWith(t *testing.T) {
	t.Run("Valid string input not starting with prefix", func(t *testing.T) {
		input := "HelloWorld"
		args := map[string]args.Arg{
			"prefix": {Value: "Planet"},
		}
		err := StartsNotWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid int input not starting with prefix", func(t *testing.T) {
		input := 12345
		args := map[string]args.Arg{
			"prefix": {Value: "23"},
		}
		err := StartsNotWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid float input not starting with prefix", func(t *testing.T) {
		input := 123.456
		args := map[string]args.Arg{
			"prefix": {Value: "12.3"},
		}
		err := StartsNotWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid bool input not starting with prefix", func(t *testing.T) {
		input := true
		args := map[string]args.Arg{
			"prefix": {Value: "false"},
		}
		err := StartsNotWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Invalid string input starting with prefix", func(t *testing.T) {
		input := "HelloWorld"
		args := map[string]args.Arg{
			"prefix": {Value: "Hello"},
		}
		err := StartsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "startsnotwith validation failed: HelloWorld starts with Hello")
	})

	t.Run("Invalid int input starting with prefix", func(t *testing.T) {
		input := 12345
		args := map[string]args.Arg{
			"prefix": {Value: "123"},
		}
		err := StartsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "startsnotwith validation failed: 12345 starts with 123")
	})

	t.Run("Invalid float input starting with prefix", func(t *testing.T) {
		input := 123.456
		args := map[string]args.Arg{
			"prefix": {Value: "123"},
		}
		err := StartsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "startsnotwith validation failed: 123.456 starts with 123")
	})

	t.Run("Unsupported input type", func(t *testing.T) {
		input := struct{}{}
		args := map[string]args.Arg{
			"prefix": {Value: "anything"},
		}
		err := StartsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type for input field")
	})
}
