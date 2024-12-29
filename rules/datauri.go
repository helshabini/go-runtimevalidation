package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// DataUri checks if the provided input is a valid Data URI.
//
// This function expects the input to be a string. It validates the format
// of the Data URI using a regular expression, which checks for the general structure of
// the Data URI (e.g., "data:[<mediatype>][;base64],<data>").
//
// Note: This function does not verify the validity or integrity of the content within
// the Data URI, only the structural format.
//
// If the input is not a string or does not match the expected Data URI pattern,
// the function returns an error.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
// - error: nil if the input is a valid Data URI, otherwise an error.
func DataUri(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid data URI
	if !regex.DataURIRegex().MatchString(value) {
		return fmt.Errorf("invalid data uri: %s", value)
	}

	return nil
}
