package rules

import (
	"fmt"
	"go-runtimevalidation/args"
	"go-runtimevalidation/functions"
	"strings"
)

// EndsNotWith checks if the input string doesn't end with the given suffix.
//
// This function expects exactly one argument (suffix) provided in the `args` map.
// It evaluates the argument and checks whether the input string doesn't end with the evaluated suffix.
// If the input or suffix is not a valid string, or if the input string ends with the suffix,
// an error is returned.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
// - obj: the object used for argument evaluation.
// - args: a map of arguments (expects one argument, the suffix).
//
// Returns:
// - error: an error if validation fails or if the argument count/type is incorrect. Returns nil if validation passes.
func EndsNotWith(input any, obj any, arguments map[string]args.Arg) error {
	if len(arguments) != 1 {
		return fmt.Errorf("endsnotwith expects 1 argument, got %d", len(arguments))
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
	suffix, err := functions.GetString(eval)
	if err != nil {
		return fmt.Errorf("unsupported type for endsnotwith argument: %w", err)
	}

	// Compare values
	if strings.HasSuffix(value, suffix) {
		return fmt.Errorf("endsnotwith validation failed: %s ends with %s", value, suffix)
	}

	return nil
}
