package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMax(t *testing.T) {
	// Test case 1: Integer input with a constant argument
	t.Run("Integer input with constant argument", func(t *testing.T) {
		err := Max(5, nil, map[string]args.Arg{
			"arg1": {Value: 10},
		})
		assert.NoError(t, err)
	})

	// Test case 2: Integer input greater than the argument
	t.Run("Integer input greater than constant argument", func(t *testing.T) {
		err := Max(15, nil, map[string]args.Arg{
			"arg1": {Value: 10},
		})
		assert.Error(t, err)
		assert.EqualError(t, err, "max validation failed: 15 > 10")
	})

	// Test case 3: Input is not an integer
	t.Run("Non-integer input", func(t *testing.T) {
		err := Max("text", nil, map[string]args.Arg{
			"arg1": {Value: 10},
		})
		assert.Error(t, err)
		assert.EqualError(t, err, "unsupported type for input field: failed to parse \"text\" of type string as int64")
	})

	// Test case 4: Argument is not an integer
	t.Run("Non-integer argument", func(t *testing.T) {
		err := Max(10, nil, map[string]args.Arg{
			"arg1": {Value: "text"},
		})
		assert.Error(t, err)
		assert.EqualError(t, err, "unsupported type for max argument: failed to parse \"text\" of type string as int64")
	})

	// Test case 5: Field reference resolves to an integer
	t.Run("Field reference resolves to integer", func(t *testing.T) {
		obj := struct {
			Number int
		}{Number: 10}

		err := Max(5, obj, map[string]args.Arg{
			"arg1": {Type: args.FieldArg, Field: "Number"},
		})
		assert.NoError(t, err)
	})

	// Test case 6: Field reference resolves to a string
	t.Run("Field reference resolves to string", func(t *testing.T) {
		obj := struct {
			StrField string
		}{StrField: "test"}

		err := Max(20, obj, map[string]args.Arg{
			"arg1": {Type: args.FieldArg, Field: "StrField"},
		})
		assert.Error(t, err)
		assert.EqualError(t, err, "unsupported type for max argument: failed to parse \"test\" of type string as int64")
	})
}
