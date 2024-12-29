package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// SHA224 validates whether the input is a valid SHA224 hash string.
//
// It checks if the input is a string, and if the string matches the SHA224 hash format.
// A valid SHA224 hash is 56 hexadecimal characters long.
//
// Parameters:
//
//	input: the value to be validated (expected to be a string).
//
// Returns:
//
//	error: nil if the input is a valid SHA224 hash, otherwise an error indicating invalid input.
func SHA224(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid SHA224 value
	if !regex.Sha224Regex().MatchString(value) {
		return fmt.Errorf("invalid SHA224 value: %s", value)
	}

	return nil
}
