package rules

import (
	"fmt"
	"go-runtimevalidation/args"
	"go-runtimevalidation/functions"
	"strings"
)

// Contains checks if the input string contains the given substring.
//
// The function expects exactly one argument (substring) provided in the `args` map.
// It evaluates the argument and checks whether the input string contains the evaluated substring.
// If the input or substring is not a valid string, or if the input string does not contain the substring,
// an error is returned.
//
// Parameters:
// - input: the value to be validated (expected to be convertible to a string).
// - obj: the object used for argument evaluation.
// - args: a map of arguments (expects one argument, the substring).
//
// Returns:
// - error: an error if validation fails or if the argument count/type is incorrect. Returns nil if validation passes.
func Contains(input any, obj any, arguments map[string]args.Arg) error {
	if len(arguments) != 1 {
		return fmt.Errorf("contains expects 1 argument, got %d", len(arguments))
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
	value, err := functions.GetString(input)
	if err != nil {
		return fmt.Errorf("unsupported type for input field: %w", err)
	}

	// Get the value to compare against
	contained, err := functions.GetString(eval)
	if err != nil {
		return fmt.Errorf("unsupported type for contains argument: %w", err)
	}

	if len(value) == 0 || len(contained) == 0 {
		return fmt.Errorf("contains validation failed: empty string")
	}

	// Compare values
	if !strings.Contains(value, contained) {
		return fmt.Errorf("contains validation failed: %s does not contain %s", value, contained)
	}

	return nil
}
