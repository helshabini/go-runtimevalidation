package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// ULID validates if the input is a valid ULID (Universally Unique Lexicographically Sortable Identifier).
// A ULID is a 26-character string consisting of upper-case letters and numbers, following a specific format.
//
// Parameters:
// - input: The value to be validated, expected to be a string.
//
// Returns:
// - nil if the input is a valid ULID.
// - an error if the input is not a string or if it doesn't match the ULID format.
//
// Example:
//
//	input := "01ARZ3NDEKTSV4RRFFQ69G5FAV"
//	err := ULID(input)  // err will be nil
//
//	input = "invalid-ulid"
//	err = ULID(input)  // err will not be nil
func ULID(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid ULID value
	if !regex.ULIDRegex().MatchString(value) {
		return fmt.Errorf("invalid ULID value: %s", value)
	}

	return nil
}
