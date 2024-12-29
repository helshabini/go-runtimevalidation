package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Longitude checks if the input string represents a valid longitude value.
//
// The function expects the input to be a string representing a valid longitude.
// Longitude values should be between -180 and 180 degrees, and can be provided
// in integer or decimal form.
//
// Note: The regex used only checks for the format and range of the value,
// it doesn't validate any additional context like precision or geographical significance.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
// - error: an error if validation fails or if the input type is incorrect. Returns nil if validation passes.
func Longitude(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid longitude value
	if !regex.LongitudeRegex().MatchString(value) {
		return fmt.Errorf("invalid longitude value: %s", value)
	}

	return nil
}
