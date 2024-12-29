package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// CVE validates if the input is a valid CVE identifier.
//
// CVE identifiers follow the format: CVE-YYYY-NNNN, where:
// - "CVE" is the fixed prefix.
// - "YYYY" is a 4-digit year.
// - "NNNN" is a number that can have 4 or more digits.
//
// Parameters:
// - input (any): The value to be validated. It should be convertible to a string.
//
// Returns:
// - error: If the CVE identifier is invalid, it returns an error with details.
func CVE(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid cve
	if !regex.CveRegex().MatchString(value) {
		return fmt.Errorf("invalid cve: %s", value)
	}

	return nil
}
