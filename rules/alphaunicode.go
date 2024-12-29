package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// AlphaUnicode validates whether the input string contains only alphabetic characters,
// including Unicode letters from different languages (e.g., accented characters, Cyrillic letters).
// It ensures that the string is composed solely of alphabetic Unicode characters.
//
// This function is particularly useful for validating names or fields that may include international characters.
//
// Parameters:
//   - input: The value to be validated, expected to be of type `string`.
//
// Returns:
//   - `nil` if the input is a valid Unicode alphabetic string.
//   - An error if the input is not a string or contains non-alphabetic characters.
//
// The regular expression used for validation allows all Unicode letters: `^[\\p{L}]+$`.
//
// Example:
//
//	AlphaUnicode("JohnDoe")         // Returns: nil
//	AlphaUnicode("José")            // Returns: nil
//	AlphaUnicode("Сергей")          // Returns: nil (Cyrillic alphabet)
//	AlphaUnicode("123ABC")          // Returns: error ("invalid alpha unicode: 123ABC")
//	AlphaUnicode("John_Doe")        // Returns: error ("invalid alpha unicode: John_Doe")
//	AlphaUnicode(12345)             // Returns: error ("expected a string, got int")
func AlphaUnicode(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only alphabetic characters
	if !regex.AlphaUnicodeRegex().MatchString(value) {
		return fmt.Errorf("invalid alpha unicode: %s", value)
	}

	return nil
}
