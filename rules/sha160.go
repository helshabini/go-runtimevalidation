package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// SHA160 validates whether the input is a valid SHA-0 or SHA-1 hash string.
//
// The function expects the input to be a string and checks it against
// a regex pattern for valid SHA-0 or SHA-1 hash formats (40 hexadecimal characters).
// If the input is not a string, or if it doesn't match the SHA format,
// an error is returned.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
//   - error: an error if the input is not a valid SHA-0 or SHA-1 hash or if it is not a string.
//     Returns nil if the input is valid.
func SHA160(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid SHA160 value
	if !regex.Sha160Regex().MatchString(value) {
		return fmt.Errorf("invalid SHA160 value: %s", value)
	}

	return nil
}
