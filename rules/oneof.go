package rules

import (
	"fmt"
	"go-runtimevalidation/args"
	"reflect"
)

// OneOf validates that the input is equal to one of the provided arguments.
// The input can be of any type, and the function will check if it matches
// any of the evaluated arguments in the args map.
//
// Parameters:
// - input: The value being validated. Can be of any type.
// - obj: The object containing additional data, which can be used for field references or function calls in the args.
// - args: A map of arguments where each entry will be evaluated, and the input will be compared against each.
//
// Returns nil if the input matches one of the provided arguments, or an error if:
// - No arguments are provided.
// - Any argument evaluation results in an error.
// - The input does not match any of the provided arguments.
//
// Example:
//
//	input := "apple"
//	obj := nil
//	args := map[string]Arg{
//	    "option1": {Value: "apple"},
//	    "option2": {Value: "banana"},
//	    "option3": {Value: "cherry"},
//	}
//	err := OneOf(input, obj, args)  // err will be nil, since input matches "apple"
func OneOf(input any, obj any, arguments map[string]args.Arg) error {
	if len(arguments) == 0 {
		return fmt.Errorf("oneof expects atleast 1 argument, got %d", len(arguments))
	}

	for _, v := range arguments {
		arg, err := v.Evaluate(obj)
		if err != nil {
			return err
		}

		if reflect.DeepEqual(input, arg) {
			return nil
		}
	}

	return fmt.Errorf("oneof validation failed: %v is not one of %v", input, arguments)
}
