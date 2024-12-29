package rules

import (
	"fmt"
	"go-runtimevalidation/functions"
	"go-runtimevalidation/regex"
)

// UrlEncoded checks if the input string is a valid URL-encoded value.
//
// The function validates that the input is a string and that it matches the expected URL-encoded format
// using a regular expression. It does not verify the content or correctness of the decoded value.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
// - error: an error if validation fails or if the input type is incorrect. Returns nil if validation passes.
func UrlEncoded(input any) error {
	// Check if the input is a string
	value, err := functions.GetString(input)
	if err != nil {
		return fmt.Errorf("unsupported type for input field: %w", err)
	}

	if len(value) == 0 {
		return fmt.Errorf("invalid url encoded value: %s", value)
	}

	// Check if the string is a valid url encoded value
	if !regex.URLEncodedRegex().MatchString(value) {
		return fmt.Errorf("invalid url encoded value: %s", value)
	}

	return nil
}
