package rules

import (
	"fmt"
	"go-runtimevalidation/functions"
	"go-runtimevalidation/regex"
)

// BIC validates if the input is a valid Bank Identifier Code (BIC).
//
// A valid BIC has 8 or 11 characters, where:
// - The first 4 characters are alphabetic (Bank Code).
// - The next 2 characters are alphabetic (Country Code).
// - The next 2 characters are alphanumeric (Location Code).
// - If present, the last 3 characters are alphanumeric (Branch Code).
//
// Parameters:
// - input (any): The value to be validated. It should be convertible to a string.
//
// Returns:
// - error: If the BIC is invalid, it returns an error with details.
func BIC(input any) error {
	// Check if the input is a string
	value, err := functions.GetString(input)
	if err != nil {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid bic
	if !regex.BicRegex().MatchString(value) {
		return fmt.Errorf("invalid bic: %s", value)
	}

	return nil
}
