package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// AlphaNumericUnicode validates whether the input string contains only alphabetic Unicode characters and numbers.
// It allows for Unicode letters from different languages as well as numeric digits (0-9).
//
// This function is useful for validating fields like usernames, product codes, or other identifiers that may contain
// international alphabetic characters and numbers.
//
// Parameters:
//   - input: The value to be validated, expected to be of type `string`.
//
// Returns:
//   - `nil` if the input is a valid Unicode alphabetic-numeric string.
//   - An error if the input is not a string or contains characters other than letters or numbers.
//
// The regular expression used for validation allows Unicode letters and numbers: `^[\\p{L}\\p{N}]+$`.
//
// Example:
//
//	AlphaNumericUnicode("John123")         // Returns: nil
//	AlphaNumericUnicode("Producto2023")    // Returns: nil
//	AlphaNumericUnicode("User@Name")       // Returns: error ("invalid alpha unicode numeric: User@Name")
//	AlphaNumericUnicode(12345)             // Returns: error ("expected a string, got int")
func AlphaNumericUnicode(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only alphabetic Unicode characters and numbers
	if !regex.AlphaNumericUnicodeRegex().MatchString(value) {
		return fmt.Errorf("invalid alpha unicode numeric: %s", value)
	}

	return nil
}
