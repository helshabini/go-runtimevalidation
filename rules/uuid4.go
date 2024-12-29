package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// UUID4 validates whether the input is a valid UUID4 string.
// The input must be a string in the correct UUID4 format.
// A UUID4 follows the pattern of the RFC 4122 specification.
//
// Parameters:
// - input: The value being validated, expected to be a string.
//
// Returns an error if:
// - The input is not a string
// - The input string does not match the UUID4 pattern.
//
// Example:
//
//	input := "550e8400-e29b-41d4-a716-446655440000"
//	err := UUID4(input)  // err will be nil if the input is a valid UUID4
func UUID4(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid UUID4 value
	if !regex.UUID4RFC4122Regex().MatchString(value) {
		return fmt.Errorf("invalid UUID4 value: %s", value)
	}

	return nil
}
