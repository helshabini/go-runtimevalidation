package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// UUID3 validates if the input is a valid UUID version 3 string.
// It checks whether the input is a string and matches the UUID3 regex pattern.
// A UUID3 follows the pattern of the RFC 4122 specification.
//
// Parameters:
// - input: The value being validated, expected to be a string.
//
// Returns nil if the input is valid, or an error if:
// - The input is not a string.
// - The input does not match the UUID3 regex pattern.
//
// Example:
//
//	err := UUID3("f47ac10b-58cc-3bf1-8a9a-1234567890ab")  // err will be nil
func UUID3(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid UUID3 value
	if !regex.UUID3RFC4122Regex().MatchString(value) {
		return fmt.Errorf("invalid UUID3 value: %s", value)
	}

	return nil
}
