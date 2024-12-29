package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// FQDN checks if the input string is a fully qualified domain name (FQDN).
//
// This function validates the input based on the same rules as RFC 1123 for hostnames,
// but additionally requires a non-numeric top-level domain (TLD). The FQDN may end with a period ('.').
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
// - error: an error if validation fails or if the input type is incorrect. Returns nil if validation passes.
func FQDN(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid fqdn code
	if !regex.FqdnRegex().MatchString(value) {
		return fmt.Errorf("invalid fqdn: %s", value)
	}

	return nil
}
