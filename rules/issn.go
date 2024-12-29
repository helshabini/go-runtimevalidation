package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// ISSN checks if the input string is a valid ISSN (International Standard Serial Number).
//
// The function expects the input to be a string representing a valid ISSN, formatted
// as 'XXXX-XXXX', where each 'X' is a digit, and the last digit may be a number or 'X'.
// Note: This function only checks for format validity based on regex and does not validate
// whether the ISSN is officially assigned or correct in context.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
// - error: an error if validation fails or if the input type is incorrect. Returns nil if validation passes.
func ISSN(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid issn value
	if !regex.ISSNRegex().MatchString(value) {
		return fmt.Errorf("invalid issn value: %s", value)
	}

	return nil
}
