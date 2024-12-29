package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// E164 validates if the given input is a valid E.164 phone number.
// E.164 phone numbers are formatted with a leading '+' followed by 7 to 15 digits.
// The phone number should not have any spaces, dashes, or special characters.
// The input should be a string, and it is required to start with a '+' symbol.
//
// Examples of valid E.164 numbers:
// +1234567890123
// +19876543210
//
// If the input is valid, the function returns nil. Otherwise, it returns an error.
func E164(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid E.164 formatted phone number
	if !regex.E164Regex().MatchString(value) {
		return fmt.Errorf("invalid E.164 phone number: %s", value)
	}

	return nil
}
