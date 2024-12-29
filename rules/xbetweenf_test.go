package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXBetweenF(t *testing.T) {
	// Test valid inputs exclusively within bounds
	t.Run("Valid input exclusively within bounds", func(t *testing.T) {
		input := 15.5
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 15.0},
			"upperBound": {Value: 16.0},
		}
		err := XBetweenF(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input equal to lower bound
	t.Run("Input equal to lower bound", func(t *testing.T) {
		input := 15.0
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 15.0},
			"upperBound": {Value: 16.0},
		}
		err := XBetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test input equal to upper bound
	t.Run("Input equal to upper bound", func(t *testing.T) {
		input := 16.0
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 15.0},
			"upperBound": {Value: 16.0},
		}
		err := XBetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test input exactly between bounds
	t.Run("Valid input exactly between bounds", func(t *testing.T) {
		input := 15.5
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 15.0},
			"upperBound": {Value: 16.0},
		}
		err := XBetweenF(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input below the lower bound
	t.Run("Input below lower bound", func(t *testing.T) {
		input := 14.9
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 15.0},
			"upperBound": {Value: 16.0},
		}
		err := XBetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test input above the upper bound
	t.Run("Input above upper bound", func(t *testing.T) {
		input := 16.1
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 15.0},
			"upperBound": {Value: 16.0},
		}
		err := XBetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test with invalid number of arguments
	t.Run("Invalid number of arguments", func(t *testing.T) {
		input := 15.5
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 10},
		}
		err := XBetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test with non-convertible input type
	t.Run("Non-convertible input type", func(t *testing.T) {
		input := "invalid"
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		}
		err := XBetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test with non-convertible argument type
	t.Run("Non-convertible argument type", func(t *testing.T) {
		input := 15.0
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: "invalid"},
			"upperBound": {Value: 20},
		}
		err := XBetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test field reference resolves to floats
	t.Run("Field reference resolves to floats", func(t *testing.T) {
		obj := struct {
			Lower float64
			Upper float64
		}{Lower: 10.2, Upper: 16.0}

		err := XBetweenF(15.5, obj, map[string]args.Arg{
			"lowerBound": {Type: args.FieldArg, Field: "Lower"},
			"upperBound": {Type: args.FieldArg, Field: "Upper"},
		})
		assert.NoError(t, err)
	})

	// Test one argument resolves to a string
	t.Run("Field reference resolves to string", func(t *testing.T) {
		obj := struct {
			StrField string
		}{StrField: "test"}

		err := XBetweenF(15.0, obj, map[string]args.Arg{
			"lowerBound": {Type: args.FieldArg, Field: "StrField"},
			"upperBound": {Value: 20},
		})
		assert.Error(t, err)
		assert.EqualError(t, err, "unsupported type for lower bound argument: failed to parse \"test\" of type string as float64")
	})
}
