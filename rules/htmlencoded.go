package rules

import (
	"fmt"
	"go-runtimevalidation/functions"
	"go-runtimevalidation/regex"
)

// HTMLEncoded validates whether the given input is a properly HTML-encoded string.
// The function checks if the input string contains valid HTML entities such as
// `&amp;`, `&lt;`, etc. It does not check the validity of the content itself but
// ensures that the string follows the proper format for HTML encoding.
//
// Example of valid HTML-encoded strings: "&amp;", "&lt;", "&#39;", etc.
//
// Parameters:
// - input: The value to validate (expected to be convertible to string).
//
// Returns:
//   - error: If the input is not valid or cannot be converted to a string,
//     or if it is not a valid HTML-encoded string.
//
// Note: This function assumes `hTMLEncodedRegex` is defined for checking the
//
//	HTML encoding format.
func HTMLEncoded(input any) error {
	// Check if the input is a string
	value, err := functions.GetString(input)
	if err != nil {
		return fmt.Errorf("unsupported type for input field: %w", err)
	}

	if len(value) == 0 {
		return fmt.Errorf("invalid html encoded: %s", value)
	}

	// Check if the string is a valid html encoded
	if !regex.HTMLEncodedRegex().MatchString(value) {
		return fmt.Errorf("invalid html encoded: %s", value)
	}

	return nil
}
