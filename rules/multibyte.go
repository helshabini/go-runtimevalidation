package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// MultiByte validates whether the input contains only MultiByte characters.
//
// The function expects the input to be a string and checks it against a regex
// pattern for valid MultiByte characters (e.g., non-ASCII characters such as
// Chinese, Japanese, Korean, or other multi-byte characters). If the input is
// not a string, or if it contains characters outside the multi-byte range (except white-space),
// an error is returned.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
//   - error: an error if the input is not valid MultiByte or not a string.
//     Returns nil if the input is valid.
func MultiByte(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only MultiByte characters
	if !regex.MultibyteRegex().MatchString(value) {
		return fmt.Errorf("invalid MultiByte: %s", value)
	}

	return nil
}
