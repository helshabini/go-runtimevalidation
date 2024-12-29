package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndsWith(t *testing.T) {
	t.Run("Valid string input and suffix", func(t *testing.T) {
		input := "HelloWorld"
		args := map[string]args.Arg{
			"suffix": {Value: "World"},
		}
		err := EndsWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid int input and string suffix", func(t *testing.T) {
		input := 12345
		args := map[string]args.Arg{
			"suffix": {Value: "45"},
		}
		err := EndsWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid float input and string suffix", func(t *testing.T) {
		input := 123.456
		args := map[string]args.Arg{
			"suffix": {Value: "456"},
		}
		err := EndsWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid bool input and string suffix", func(t *testing.T) {
		input := true
		args := map[string]args.Arg{
			"suffix": {Value: "true"},
		}
		err := EndsWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Invalid string input and suffix", func(t *testing.T) {
		input := "HelloWorld"
		args := map[string]args.Arg{
			"suffix": {Value: "Planet"},
		}
		err := EndsWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "endswith validation failed: HelloWorld does not end with Planet")
	})

	t.Run("Invalid int input and suffix", func(t *testing.T) {
		input := 12345
		args := map[string]args.Arg{
			"suffix": {Value: "46"},
		}
		err := EndsWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "endswith validation failed: 12345 does not end with 46")
	})

	t.Run("Unsupported input type", func(t *testing.T) {
		input := struct{}{}
		args := map[string]args.Arg{
			"suffix": {Value: "anything"},
		}
		err := EndsWith(input, nil, args)
		assert.Error(t, err)
	})
}
