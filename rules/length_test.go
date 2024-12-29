package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLength(t *testing.T) {
	// Test case: static length check with matching length
	t.Run("Static length match", func(t *testing.T) {
		input := "example" // length 7
		args := map[string]args.Arg{
			"length": {Type: args.ValueArg, Value: 7},
		}
		err := Length(input, nil, args)
		assert.NoError(t, err, "Expected no error for matching length")
	})

	// Test case: static length check with non-matching length
	t.Run("Static length mismatch", func(t *testing.T) {
		input := "example" // length 7
		args := map[string]args.Arg{
			"length": {Type: args.ValueArg, Value: 10},
		}
		err := Length(input, nil, args)
		assert.Error(t, err, "Expected an error for mismatched length")
		assert.Equal(t, "length mismatch: 7 != 10", err.Error())
	})

	// Test case: dynamic length check with matching length using function
	t.Run("Dynamic length match using function", func(t *testing.T) {
		input := "secretPassword" // length 14
		args := map[string]args.Arg{
			"length": {Type: args.FunctionArg, Function: args.Function{
				Name: "len",
				Args: []args.Arg{{Type: args.FieldArg, Field: "Password"}},
			}},
		}
		obj := struct {
			Password string
		}{
			Password: "secretPassword", // matching input
		}
		err := Length(input, obj, args)
		assert.NoError(t, err, "Expected no error for matching length with dynamic function")
	})

	// Test case: dynamic length check with non-matching length using function
	t.Run("Dynamic length mismatch using function", func(t *testing.T) {
		input := "short" // length 5
		args := map[string]args.Arg{
			"length": {Type: args.FunctionArg, Function: args.Function{
				Name: "len",
				Args: []args.Arg{{Type: args.FieldArg, Field: "Password"}},
			}},
		}
		obj := struct {
			Password string
		}{
			Password: "longPassword", // length 12
		}
		err := Length(input, obj, args)
		assert.Error(t, err, "Expected an error for mismatched length with dynamic function")
		assert.Equal(t, "length mismatch: 5 != 12", err.Error())
	})

	// Test case: length check with unsupported input type
	t.Run("Unsupported input type", func(t *testing.T) {
		input := 12345 // not a string, slice, or array
		args := map[string]args.Arg{
			"length": {Type: args.ValueArg, Value: 5},
		}
		err := Length(input, nil, args)
		assert.Error(t, err, "Expected an error for unsupported input type")
		assert.Contains(t, err.Error(), "unsupported type for len")
	})

	// Test case: argument type mismatch (non-integer length)
	t.Run("Argument type mismatch", func(t *testing.T) {
		input := "example" // length 7
		args := map[string]args.Arg{
			"length": {Type: args.ValueArg, Value: "ten"}, // non-integer
		}
		err := Length(input, nil, args)
		assert.Error(t, err, "Expected an error for non-integer length argument")
		assert.Contains(t, err.Error(), "unsupported type for length argument")
	})

	// Test case: empty input and length 0
	t.Run("Empty input length 0", func(t *testing.T) {
		input := "" // length 0
		args := map[string]args.Arg{
			"length": {Type: args.ValueArg, Value: 0},
		}
		err := Length(input, nil, args)
		assert.NoError(t, err, "Expected no error for empty input with length 0")
	})
}
