package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Alpha validates whether the input contains only alphabetic characters (A-Z, a-z).
// It checks that the input is a string and that all characters in the string are alphabetic.
//
// This function is useful in validation rules where fields are expected to contain only letters,
// such as a name or a code that cannot have numbers or special characters.
//
// Parameters:
//   - input: The value to be validated, expected to be of type `string`.
//
// Returns:
//   - `nil` if the input is a valid alphabetic string.
//   - An error if the input is not a string or contains non-alphabetic characters.
//
// Example:
//
//	Alpha("Hello")  // Returns: nil
//	Alpha("12345")  // Returns: error ("invalid alpha: 12345")
//	Alpha("Hello123")  // Returns: error ("invalid alpha: Hello123")
//	Alpha(12345)  // Returns: error ("expected a string, got int")
func Alpha(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only alphabetic characters
	if !regex.AlphaRegex().MatchString(value) {
		return fmt.Errorf("invalid alpha: %s", value)
	}

	return nil
}
