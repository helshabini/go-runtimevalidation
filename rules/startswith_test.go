package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartsWith(t *testing.T) {
	t.Run("Valid string input and prefix", func(t *testing.T) {
		input := "HelloWorld"
		args := map[string]args.Arg{
			"prefix": {Value: "Hello"},
		}
		err := StartsWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid int input and string prefix", func(t *testing.T) {
		input := 12345
		args := map[string]args.Arg{
			"prefix": {Value: "12"},
		}
		err := StartsWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid float input and string prefix", func(t *testing.T) {
		input := 123.456
		args := map[string]args.Arg{
			"prefix": {Value: "123"},
		}
		err := StartsWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid bool input and string prefix", func(t *testing.T) {
		input := true
		args := map[string]args.Arg{
			"prefix": {Value: "tru"},
		}
		err := StartsWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Invalid string input and prefix", func(t *testing.T) {
		input := "HelloWorld"
		args := map[string]args.Arg{
			"prefix": {Value: "Planet"},
		}
		err := StartsWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "startswith validation failed: HelloWorld does not start with Planet")
	})

	t.Run("Invalid int input and prefix", func(t *testing.T) {
		input := 12345
		args := map[string]args.Arg{
			"prefix": {Value: "23"},
		}
		err := StartsWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "startswith validation failed: 12345 does not start with 23")
	})

	t.Run("Unsupported input type", func(t *testing.T) {
		input := struct{}{}
		args := map[string]args.Arg{
			"prefix": {Value: "anything"},
		}
		err := StartsWith(input, nil, args)
		assert.Error(t, err)
	})
}
