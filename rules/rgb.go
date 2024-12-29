package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// RGB validates whether the input is a valid RGB color value.
// It accepts RGB values in integer (0-255) format.
//
// Parameters:
//   - input: The value to be validated. It must be a string.
//
// Returns:
//   - An error if the input is not a string or if it does not match the RGB color value format.
func RGB(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid RGB color value
	if !regex.RgbRegex().MatchString(value) {
		return fmt.Errorf("invalid rgb value: %s", value)
	}

	return nil
}
