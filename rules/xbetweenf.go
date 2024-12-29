package rules

import (
	"fmt"

	"go-runtimevalidation/args"
	"go-runtimevalidation/functions"
)

// XBetweenF validates that the input is between two specified bounds.
// Unlike BetweenF, XBetweenF checks for exclusive bounds.
// The input must be a number or a type that can be converted into a float.
// It checks against exactly two arguments provided in the args map.
//
// Parameters:
// - input: The value being validated, expected to be convertible to a float.
// - obj: The object containing additional data (can be used for Field references within the args).
// - args: A map of arguments where exactly two entries are expected to specify the bounds.
//
// Returns nil if the input is valid, or an error if:
// - The number of arguments is not equal to 2
// - The input cannot be converted to a float
// - Any of the arguments cannot be converted to a float
// - The input is not exclusively between the two specified bounds.
//
// Example:
//
//	input := 15.5
//	obj := nil
//	args := map[string]Arg{
//	    "lowerBound": Arg{Value: 10},
//	    "upperBound": Arg{Value: 15.6},
//	}
//	err := XBetweenF(input, obj, args)  // err will be nil
func XBetweenF(input any, obj any, arguments map[string]args.Arg) error {
	if len(arguments) != 2 {
		return fmt.Errorf("xbetweenf expects exactly 2 arguments, got %d", len(arguments))
	}

	// Extract the two arguments
	var lhsArg, rhsArg args.Arg
	i := 0
	for _, v := range arguments {
		if i == 0 {
			lhsArg = v
		} else {
			rhsArg = v
		}
		i++
	}

	// Evaluate the arguments
	lhsEval, err := lhsArg.Evaluate(obj)
	if err != nil {
		return err
	}

	rhsEval, err := rhsArg.Evaluate(obj)
	if err != nil {
		return err
	}

	// Get the value of the input
	inputVal, err := functions.GetFloat(input)
	if err != nil {
		return fmt.Errorf("unsupported type for input field: %w", err)
	}

	// Get the values to compare against
	lhs, err := functions.GetFloat(lhsEval)
	if err != nil {
		return fmt.Errorf("unsupported type for lower bound argument: %w", err)
	}

	rhs, err := functions.GetFloat(rhsEval)
	if err != nil {
		return fmt.Errorf("unsupported type for upper bound argument: %w", err)
	}

	// Compare values to determine if input is exclusively between bounds
	if (inputVal <= lhs && inputVal <= rhs) || (inputVal >= lhs && inputVal >= rhs) {
		return fmt.Errorf("xbetweenf validation failed: %f is not exclusively between %f and %f", inputVal, lhs, rhs)
	}

	// Validation passed
	return nil
}
