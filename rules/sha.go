package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// SHA validates whether the input is a valid SHA hash string.
//
// The function expects the input to be a string and checks it against
// a regex pattern for valid SHA hash formats (40 to 128 hexadecimal characters).
// If the input is not a string, or if it doesn't match the SHA format,
// an error is returned.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
//   - error: an error if the input is not a valid SHA hash or if it is not a string.
//     Returns nil if the input is valid.
func SHA(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid SHA value
	if !regex.ShaRegex().MatchString(value) {
		return fmt.Errorf("invalid SHA value: %s", value)
	}

	return nil
}
