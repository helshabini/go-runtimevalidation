package rules

import (
	"fmt"
	"go-runtimevalidation/args"
	"go-runtimevalidation/functions"
)

// Min validates that the input is greater than or equal to a minimum value.
// The input must be an integer or a type that can be converted into an integer.
// It checks against exactly one argument provided in the args map.
//
// Parameters:
// - input: The value being validated, expected to be convertible to an integer.
// - obj: The object containing additional data (can be used for Field references within the args).
// - args: A map of arguments where exactly one entry is expected to specify the minimum value.
//
// Returns nil if the input is valid, or an error if:
// - The number of arguments is not equal to 1
// - The input cannot be converted to an integer
// - The argument cannot be converted to an integer
// - The input is less than the specified minimum value.
//
// Example:
//
//	input := 5
//	obj := nil
//	args := map[string]Arg{
//	    "minValue": Arg{Value: 10},
//	}
//	err := Min(input, obj, args)  // err will be: "min validation failed: 5 < 10"
func Min(input any, obj any, arguments map[string]args.Arg) error {
	if len(arguments) != 1 {
		return fmt.Errorf("min expects exactly 1 argument, got %d", len(arguments))
	}

	// Get the first (and only) argument
	var arg args.Arg
	for _, v := range arguments {
		arg = v
		break
	}

	// Evaluate the argument
	eval, err := arg.Evaluate(obj)
	if err != nil {
		return err
	}

	// Get the value of the input
	lhs, err := functions.GetInt(input)
	if err != nil {
		return fmt.Errorf("unsupported type for input field: %w", err)
	}

	// Get the value to compare against
	rhs, err := functions.GetInt(eval)
	if err != nil {
		return fmt.Errorf("unsupported type for min argument: %w", err)
	}

	// Compare values
	if lhs < rhs {
		return fmt.Errorf("min validation failed: %d < %d", lhs, rhs)
	}

	// Validation passed
	return nil
}
