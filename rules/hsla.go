package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// HSLA validates whether the input is a valid HSLA color value.
// It accepts HSLA values in the format hsla(hue, saturation, lightness, alpha).
//
// Parameters:
//   - input: The value to be validated. It must be a string.
//
// Returns:
//   - An error if the input is not a string or if it does not match the HSLA color value format.
func HSLA(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid HSLA color value
	if !regex.HslaRegex().MatchString(value) {
		return fmt.Errorf("invalid hsla value: %s", value)
	}

	return nil
}
