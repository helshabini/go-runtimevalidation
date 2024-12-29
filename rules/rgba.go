package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// RGBA validates whether the input is a valid RGBA color value.
// It accepts RGB values in integer (0-255) format,
// and alpha values can be in decimal (0 to 1) format.
//
// Parameters:
//   - input: The value to be validated. It must be a string.
//
// Returns:
//   - An error if the input is not a string or if it does not match the RGBA color value format.
func RGBA(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid RGBA color value
	if !regex.RgbaRegex().MatchString(value) {
		return fmt.Errorf("invalid rgba value: %s", value)
	}

	return nil
}
