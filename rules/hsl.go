package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// HSL validates whether the input is a valid HSL color value.
// It accepts HSL values in the format hsl(hue, saturation, lightness).
//
// Parameters:
//   - input: The value to be validated. It must be a string.
//
// Returns:
//   - An error if the input is not a string or if it does not match the HSL color value format.
func HSL(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid HSL color value
	if !regex.HslRegex().MatchString(value) {
		return fmt.Errorf("invalid hsl value: %s", value)
	}

	return nil
}
