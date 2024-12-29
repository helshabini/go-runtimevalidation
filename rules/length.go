package rules

import (
	"fmt"
	"go-runtimevalidation/args"
	"go-runtimevalidation/functions"
)

// Length validates that the length of the input matches the length specified in the argument.
// It is expected to be used in validation rules such as `length:10` or `length:$len($Password)`.
//
// Parameters:
// - input: The value whose length will be validated. This can be a string, array, map, slice, or any type that supports length.
// - obj: The object containing the data for field or function evaluations. This is required when using dynamic evaluations like `$len($Password)`.
// - args: A map containing a single `Arg` that specifies the expected length. This can either be a constant or a dynamic value evaluated from obj.
//
// Returns:
// - An error if the length of the input does not match the expected length or if any other error occurs during evaluation.
// - `nil` if the length of the input matches the expected length.
func Length(input any, obj any, arguments map[string]args.Arg) error {
	// Ensure only one argument is passed
	if len(arguments) != 1 {
		return fmt.Errorf("length expects exactly 1 argument, got %d", len(arguments))
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

	// Get the length of the input
	lhs, err := functions.GetLen(input)
	if err != nil {
		return err
	}

	// Get the length to compare against
	rhs, err := functions.GetInt(eval)
	if err != nil {
		return fmt.Errorf("unsupported type for length argument: %w", err)
	}

	// Compare lengths
	if int64(lhs) != rhs {
		return fmt.Errorf("length mismatch: %d != %d", lhs, rhs)
	}

	return nil
}
