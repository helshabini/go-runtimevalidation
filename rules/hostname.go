package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Hostname checks if the input string is a valid hostname.
//
// This function validates the input based on the hostname rules defined in RFC 1123,
// allowing hostnames to start with a digit. The function checks the input using a regex
// for format validity, but does not check whether the hostname actually resolves or is in use.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
// - error: an error if validation fails or if the input type is incorrect. Returns nil if validation passes.
func Hostname(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid hostname code
	if !regex.HostnameRegex().MatchString(value) {
		return fmt.Errorf("invalid hostname: %s", value)
	}

	return nil
}
