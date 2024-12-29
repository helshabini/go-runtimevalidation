package rules

import (
	"fmt"
	"go-runtimevalidation/args"
	"go-runtimevalidation/functions"
)

// Max validates that the input is less than or equal to the specified maximum value.
// The function takes an input of any type, an object of any type for evaluation,
// and a map of arguments that must contain exactly one argument specifying the maximum value.
//
// Parameters:
// - input: The value being validated, expected to be convertible to an integer.
// - obj: The object containing additional data (can be used for Field references within the args).
// - args: A map of arguments where exactly one entry is expected to specify the maximum value.
//
// Returns nil if the input is valid, or an error if:
// - The number of arguments is not equal to 1
// - The input cannot be converted to an integer
// - The argument cannot be converted to an integer
// - The input is greater than the specified maximum value.
//
// Example:
//
//	input := 10
//	obj := nil
//	args := map[string]Arg{
//	    "minValue": Arg{Value: 5},
//	}
//	err := Max(input, obj, args)  // err will be: "max validation failed: 10 > 5"
func Max(input any, obj any, arguments map[string]args.Arg) error {
	if len(arguments) != 1 {
		return fmt.Errorf("max expects exactly 1 argument, got %d", len(arguments))
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
		return fmt.Errorf("unsupported type for max argument: %w", err)
	}

	// Compare values
	if rhs < lhs {
		return fmt.Errorf("max validation failed: %d > %d", lhs, rhs)
	}

	// Validation passed
	return nil
}
