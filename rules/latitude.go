package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Latitude checks if the provided input is a valid latitude value.
//
// This function expects the input to be a string representing a latitude.
// It uses a regular expression to validate that the string follows the expected
// latitude format, i.e., it must be a numeric value between -90 and 90 (inclusive),
// with optional decimal places.
//
// If the input is not a string or does not match the expected latitude format,
// the function returns an error.
//
// Parameters:
// - input: the value to be validated (expected to be a string representing a latitude).
//
// Returns:
// - error: nil if the input is a valid latitude, otherwise an error.
func Latitude(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid latitude value
	if !regex.LatitudeRegex().MatchString(value) {
		return fmt.Errorf("invalid latitude value: %s", value)
	}

	return nil
}
