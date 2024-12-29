package rules

import (
	"fmt"
	"go-runtimevalidation/args"
	"go-runtimevalidation/functions"
	"strings"
)

// StartsNotWith checks whether the input doesn't start with the value of the argument provided.
//
// The function expects the input and the argument to be convertible to strings.
// It first evaluates the argument using the obj and converts both the input and
// the evaluated argument to strings using the getString function. Then it checks
// if the input string doesn't start with the evaluated argument string. If it does, it returns
// an error indicating the failure.
//
// Parameters:
// - input: the value to be validated (can be any type that can be converted to a string).
// - obj: the object context (used to evaluate dynamic arguments).
// - args: a map of argument names to argument values. This should contain exactly one argument.
//
// Returns:
// - error: nil if the validation passes, otherwise an error indicating the validation failure.
func StartsNotWith(input any, obj any, arguments map[string]args.Arg) error {
	if len(arguments) != 1 {
		return fmt.Errorf("startsnotwith expects 1 argument, got %d", len(arguments))
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
	prefix, err := functions.GetString(eval)
	if err != nil {
		return fmt.Errorf("unsupported type for startsnotwith argument: %w", err)
	}

	// Compare values
	if strings.HasPrefix(value, prefix) {
		return fmt.Errorf("startsnotwith validation failed: %s starts with %s", value, prefix)
	}

	return nil
}
