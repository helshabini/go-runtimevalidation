package rules

import (
	"fmt"
	"go-runtimevalidation/args"
	"go-runtimevalidation/functions"
	"regexp"
)

// Regex checks if the input string matches a given regular expression pattern.
//
// The function expects exactly one argument (the regex pattern) provided in the `args` map.
// It evaluates the argument, converts it into a valid regex pattern, and checks whether the input string matches it.
// If the input is not a valid string, or if the regex pattern is invalid, or if the input does not match the pattern,
// an error is returned.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
// - obj : The struct object whose fields will be compared against the values in the args map.
// - args: a map of arguments (expects one argument, the regex pattern).
//
// Returns:
// - error: an error if validation fails or if the argument count/type is incorrect, or if the regex pattern is invalid.
// Returns nil if validation passes.
func Regex(input any, obj any, arguments map[string]args.Arg) error {
	// Check if the args map contains exactly one argument
	if len(arguments) != 1 {
		return fmt.Errorf("regex expects exactly 1 argument, got %d", len(arguments))
	}

	// Get the argument value
	arg, err := arguments["pattern"].Evaluate(obj)
	if err != nil {
		return err
	}

	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	expString, err := functions.GetString(arg)
	if err != nil {
		return err
	}

	exp, err := regexp.Compile(expString)
	if err != nil {
		return fmt.Errorf("invalid regex: %s", value)
	}

	// Check if the string contains only numeric characters
	if !exp.MatchString(value) {
		return fmt.Errorf("value %s does not match regex %s", value, expString)
	}

	return nil
}
