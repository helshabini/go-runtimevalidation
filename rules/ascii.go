package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Ascii validates whether the input is a valid ASCII string.
//
// It checks if the input is a string and contains only ASCII characters
// (characters in the range of 0 to 127). If the input is not a string
// or if it contains non-ASCII characters, an error is returned.
//
// Parameters:
//
//	input: the value to be validated (expected to be a string).
//
// Returns:
//
//	error: nil if the input contains only ASCII characters, otherwise
//	an error indicating invalid input.
func Ascii(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only Ascii characters
	if !regex.ASCIIRegex().MatchString(value) {
		return fmt.Errorf("invalid ASCII: %s", value)
	}

	return nil
}
