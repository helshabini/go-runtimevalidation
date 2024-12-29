package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneOf(t *testing.T) {
	// Test valid input matching one of the provided arguments
	t.Run("Valid input matching argument", func(t *testing.T) {
		input := "apple"
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: "apple"},
			"option2": {Value: "banana"},
			"option3": {Value: "cherry"},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input not matching any argument
	t.Run("Input not matching any argument", func(t *testing.T) {
		input := "grape"
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: "apple"},
			"option2": {Value: "banana"},
			"option3": {Value: "cherry"},
		}
		err := OneOf(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		} else {
			assert.Contains(t, err.Error(), "oneof validation failed: grape is not one of")
		}
	})

	// Test with no arguments
	t.Run("No arguments provided", func(t *testing.T) {
		input := "apple"
		obj := ""
		args := map[string]args.Arg{}
		err := OneOf(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		} else {
			assert.EqualError(t, err, "oneof expects atleast 1 argument, got 0")
		}
	})

	// Test input matching an integer argument
	t.Run("Input matching integer argument", func(t *testing.T) {
		input := 42
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: 10},
			"option2": {Value: 20},
			"option3": {Value: 42},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input matching a boolean argument
	t.Run("Input matching boolean argument", func(t *testing.T) {
		input := true
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: false},
			"option2": {Value: true},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input not matching a boolean argument
	t.Run("Input not matching boolean argument", func(t *testing.T) {
		input := false
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: true},
		}
		err := OneOf(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test input matching a field reference
	t.Run("Input matching a field reference", func(t *testing.T) {
		obj := struct {
			Fruit string
		}{Fruit: "banana"}

		err := OneOf("banana", obj, map[string]args.Arg{
			"option1": {Type: args.FieldArg, Field: "Fruit"},
			"option2": {Value: "apple"},
		})
		assert.NoError(t, err)
	})

	// Test input not matching any field reference
	t.Run("Input not matching any field reference", func(t *testing.T) {
		obj := struct {
			Fruit string
		}{Fruit: "banana"}

		err := OneOf("grape", obj, map[string]args.Arg{
			"option1": {Type: args.FieldArg, Field: "Fruit"},
			"option2": {Value: "apple"},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "oneof validation failed: grape is not one of")
	})

	// Test with non-convertible types
	t.Run("Non-convertible argument types", func(t *testing.T) {
		input := "42"
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: 42}, // Integer vs string
		}
		err := OneOf(input, obj, args)
		if err == nil {
			t.Errorf("expected an error but got nil")
		}
	})

	// Test input matching a string argument
	t.Run("Input matching string argument", func(t *testing.T) {
		input := "banana"
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: "apple"},
			"option2": {Value: "banana"},
			"option3": {Value: "cherry"},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input matching an integer argument
	t.Run("Input matching integer argument", func(t *testing.T) {
		input := 42
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: 10},
			"option2": {Value: 20},
			"option3": {Value: 42},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input matching a float argument
	t.Run("Input matching float argument", func(t *testing.T) {
		input := 15.75
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: 10.5},
			"option2": {Value: 15.75},
			"option3": {Value: 20.2},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input matching a boolean argument
	t.Run("Input matching boolean argument", func(t *testing.T) {
		input := false
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: true},
			"option2": {Value: false},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input matching an array argument
	t.Run("Input matching array argument", func(t *testing.T) {
		input := []int{1, 2, 3}
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: []int{4, 5, 6}},
			"option2": {Value: []int{1, 2, 3}},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input matching a map argument
	t.Run("Input matching map argument", func(t *testing.T) {
		input := map[string]int{"a": 1, "b": 2}
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: map[string]int{"x": 10, "y": 20}},
			"option2": {Value: map[string]int{"a": 1, "b": 2}},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

	// Test input matching a key (e.g., matching a struct field)
	t.Run("Input matching struct field", func(t *testing.T) {
		obj := struct {
			Option1 string
			Option2 int
		}{Option1: "keyMatch", Option2: 10}

		err := OneOf("keyMatch", obj, map[string]args.Arg{
			"option1": {Type: args.FieldArg, Field: "Option1"},
			"option2": {Type: args.FieldArg, Field: "Option2"},
		})
		assert.NoError(t, err)
	})

	// Test input matching a slice argument
	t.Run("Input matching slice argument", func(t *testing.T) {
		input := []string{"apple", "banana"}
		obj := ""
		args := map[string]args.Arg{
			"option1": {Value: []string{"apple", "banana"}},
			"option2": {Value: []string{"cherry", "date"}},
		}
		err := OneOf(input, obj, args)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
