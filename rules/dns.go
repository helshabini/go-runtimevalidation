package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// DNS validates if the input is a valid DNS label based on RFC 1035.
//
// DNS labels are separated by dots (.) and consist of alphanumeric characters
// and hyphens (-), but labels cannot start or end with a hyphen. Each label can
// be up to 63 characters long.
//
// Parameters:
// - input (any): The value to be validated. It should be convertible to a string.
//
// Returns:
// - error: If the DNS label is invalid, it returns an error with details.
func DNS(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid dns
	if !regex.DnsRegex().MatchString(value) {
		return fmt.Errorf("invalid dns: %s", value)
	}

	return nil
}
