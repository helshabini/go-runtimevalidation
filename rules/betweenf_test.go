package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBetweenF(t *testing.T) {
	// Test valid inputs within bounds
	t.Run("Valid input within bounds", func(t *testing.T) {
		input := 15.5
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 15.1},
			"upperBound": {Value: 16},
		}
		err := BetweenF(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test valid input equal to lower bound
	t.Run("Valid input equal to lower bound", func(t *testing.T) {
		input := 10.0
		obj := "nil"
		args := map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		}
		err := BetweenF(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test valid input equal to upper bound
	t.Run("Valid input equal to upper bound", func(t *testing.T) {
		input := 20.0
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		}
		err := BetweenF(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test valid input exactly between bounds
	t.Run("Valid input exactly between bounds", func(t *testing.T) {
		input := 15.0
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		}
		err := BetweenF(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input below the lower bound
	t.Run("Input below lower bound", func(t *testing.T) {
		input := 9.5
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		}
		err := BetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test input above the upper bound
	t.Run("Input above upper bound", func(t *testing.T) {
		input := 20.5
		obj := ""
		args := map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		}
		err := BetweenF(input, obj, args)
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
		err := BetweenF(input, obj, args)
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
		err := BetweenF(input, obj, args)
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
		err := BetweenF(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test field reference resolves to integers
	t.Run("Field reference resolves to floats", func(t *testing.T) {
		obj := struct {
			Lower float64
			Upper float64
		}{Lower: 10.2, Upper: 16}

		err := BetweenF(15, obj, map[string]args.Arg{
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

		err := BetweenF(15, obj, map[string]args.Arg{
			"lowerBound": {Type: args.FieldArg, Field: "StrField"},
			"upperBound": {Value: 20},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to parse \"test\" of type string as float64")
	})
}
