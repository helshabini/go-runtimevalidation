package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// AlphaNumeric validates whether the input contains only alphanumeric characters (A-Z, a-z, 0-9).
// It checks that the input is a string and that all characters in the string are alphanumeric,
// allowing only letters and numbers.
//
// This function is useful in validation rules where fields must consist of alphanumeric characters,
// such as usernames or product codes that cannot have special characters.
//
// Parameters:
//   - input: The value to be validated, expected to be of type `string`.
//
// Returns:
//   - `nil` if the input is a valid alphanumeric string.
//   - An error if the input is not a string or contains non-alphanumeric characters.
//
// Example:
//
//	AlphaNumeric("User123")  // Returns: nil
//	AlphaNumeric("12345")    // Returns: nil
//	AlphaNumeric("Hello!@#") // Returns: error ("invalid alphanumeric: Hello!@#")
//	AlphaNumeric(12345)      // Returns: error ("expected a string, got int")
func AlphaNumeric(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only alphanumeric characters
	if !regex.AlphaNumericRegex().MatchString(value) {
		return fmt.Errorf("invalid alphanumeric: %s", value)
	}

	return nil
}
