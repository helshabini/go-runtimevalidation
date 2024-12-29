package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// UUID validates that the input is a valid UUID value.
// The input must be a string matching the RFC 4122 standard for UUIDs.
//
// Parameters:
// - input: The value being validated, expected to be a string.
//
// Returns nil if the input is a valid UUID, or an error if:
// - The input is not a string
// - The input does not match the UUID RFC 4122 format.
//
// Example:
//
//	err := UUID("123e4567-e89b-12d3-a456-426614174000") // err will be nil for a valid UUID
func UUID(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid UUID value
	if !regex.UUIDRFC4122Regex().MatchString(value) {
		return fmt.Errorf("invalid UUID value: %s", value)
	}

	return nil
}
