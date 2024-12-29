package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// AsciiPrint validates whether the input contains only printable ASCII characters.
//
// The function expects the input to be a string and checks it against a regex
// pattern for valid printable ASCII characters. If the input is not a string,
// or if it contains non-printable characters, an error is returned.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
//   - error: an error if the input is not valid printable ASCII or not a string.
//     Returns nil if the input is valid.
func AsciiPrint(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only Printable Ascii characters
	if !regex.PrintableASCIIRegex().MatchString(value) {
		return fmt.Errorf("invalid Printable ASCII: %s", value)
	}

	return nil
}
