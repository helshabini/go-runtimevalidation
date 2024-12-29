package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// UUID5 validates whether the input is a valid UUID5 string.
// The input must be a string in the correct UUID5 format.
// A UUID5 follows the pattern of the RFC 4122 specification.
//
// Parameters:
// - input: The value being validated, expected to be a string.
//
// Returns an error if:
// - The input is not a string
// - The input string does not match the UUID5 pattern.
//
// Example:
//
//	input := "550e8400-e29b-41d4-a716-446655440000"
//	err := UUID5(input)  // err will be nil if the input is a valid UUID5
func UUID5(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid UUID5 value
	if !regex.UUID5RFC4122Regex().MatchString(value) {
		return fmt.Errorf("invalid UUID5 value: %s", value)
	}

	return nil
}
