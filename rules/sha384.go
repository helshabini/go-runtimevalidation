package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// SHA384 validates whether the input is a valid SHA384 hash string.
//
// It checks if the input is a string, and if the string matches the SHA384 hash format.
// A valid SHA384 hash is 96 hexadecimal characters long.
//
// Parameters:
//
//	input: the value to be validated (expected to be a string).
//
// Returns:
//
//	error: nil if the input is a valid SHA384 hash, otherwise an error indicating invalid input.
func SHA384(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid SHA384 value
	if !regex.Sha384Regex().MatchString(value) {
		return fmt.Errorf("invalid SHA384 value: %s", value)
	}

	return nil
}
