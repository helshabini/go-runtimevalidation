package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Email validates whether the input is a valid email address.
// The function first checks if the input is of type string, and then verifies
// whether it matches a regular expression designed to follow the standard email format.
//
// Parameters:
//   - input: The value to be checked. It should be a string representing an email address.
//
// Returns:
//   - An error if the input is not a string or if it does not match the email format.
//
// The regular expression used for validation allows emails containing letters, digits,
// certain special characters (like '.', '_', '%', '+', and '-'), and ensures there is a valid domain part.
//
// Example:
//
//	err := Email("test@example.com") // Returns: nil (valid email)
//	err := Email("invalid-email")    // Returns: error ("invalid email: invalid-email")
//	err := Email(12345)              // Returns: error ("expected a string, got int")
func Email(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the email is valid
	if !regex.EmailRegex().MatchString(value) {
		return fmt.Errorf("invalid email: %s", value)
	}

	return nil
}
