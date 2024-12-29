package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEndsNotWith(t *testing.T) {
	t.Run("Valid string input not ending with suffix", func(t *testing.T) {
		input := "HelloWorld"
		args := map[string]args.Arg{
			"suffix": {Value: "Planet"},
		}
		err := EndsNotWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Valid int input not ending with suffix", func(t *testing.T) {
		input := 12345
		args := map[string]args.Arg{
			"suffix": {Value: "45"},
		}
		err := EndsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "endsnotwith validation failed: 12345 ends with 45")
	})

	t.Run("Valid float input not ending with suffix", func(t *testing.T) {
		input := 123.456
		args := map[string]args.Arg{
			"suffix": {Value: "56"},
		}
		err := EndsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "endsnotwith validation failed: 123.456 ends with 56")
	})

	t.Run("Valid bool input not ending with suffix", func(t *testing.T) {
		input := true
		args := map[string]args.Arg{
			"suffix": {Value: "false"},
		}
		err := EndsNotWith(input, nil, args)
		assert.NoError(t, err)
	})

	t.Run("Invalid string input ending with suffix", func(t *testing.T) {
		input := "HelloWorld"
		args := map[string]args.Arg{
			"suffix": {Value: "World"},
		}
		err := EndsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "endsnotwith validation failed: HelloWorld ends with World")
	})

	t.Run("Invalid int input ending with suffix", func(t *testing.T) {
		input := 12345
		args := map[string]args.Arg{
			"suffix": {Value: "5"},
		}
		err := EndsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "endsnotwith validation failed: 12345 ends with 5")
	})

	t.Run("Invalid float input ending with suffix", func(t *testing.T) {
		input := 123.456
		args := map[string]args.Arg{
			"suffix": {Value: "456"},
		}
		err := EndsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.EqualError(t, err, "endsnotwith validation failed: 123.456 ends with 456")
	})

	t.Run("Unsupported input type", func(t *testing.T) {
		input := struct{}{}
		args := map[string]args.Arg{
			"suffix": {Value: "anything"},
		}
		err := EndsNotWith(input, nil, args)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unsupported type for input field")
	})
}
