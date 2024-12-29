package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// SHA256 validates whether the input is a valid SHA256 hash string.
//
// It checks if the input is a string, and if the string matches the SHA256 hash format.
// A valid SHA256 hash is 64 hexadecimal characters long.
//
// Parameters:
//
//	input: the value to be validated (expected to be a string).
//
// Returns:
//
//	error: nil if the input is a valid SHA256 hash, otherwise an error indicating invalid input.
func SHA256(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid SHA256 value
	if !regex.Sha256Regex().MatchString(value) {
		return fmt.Errorf("invalid SHA256 value: %s", value)
	}

	return nil
}
