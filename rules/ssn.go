package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// SSN validates whether the input is a valid U.S. Social Security Number (SSN).
// It first checks if the input is a string, and then verifies if it matches
// the format of a valid SSN using a regular expression.
//
// An SSN follows the format "XXX-XX-XXXX", where X is a digit.
//
// Parameters:
// - input: The value being validated, expected to be a string.
//
// Returns nil if the input is a valid SSN, or an error if:
// - The input is not a string
// - The input does not match the SSN format
//
// Example:
//
//	err := SSN("123-45-6789")  // err will be nil
//	err := SSN("invalid-ssn")  // err will not be nil
func SSN(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid SSN value
	if !regex.SSNRegex().MatchString(value) {
		return fmt.Errorf("invalid SSN value: %s", value)
	}

	return nil
}
