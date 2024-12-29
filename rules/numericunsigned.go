package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// NumericUnsigned validates whether the input is a valid unsigned numeric string.
// The string must consist of digits, optionally followed by a decimal point and more digits.
// This function is useful for validation rules where non-negative numbers are required,
// such as in financial applications.
//
// Parameters:
//   - input: The value to be validated. It must be a string.
//
// Returns:
//   - An error if the input is not a string or if it does not match the unsigned numeric format.
//
// Example:
//
//	NumericUnsigned("123")     // Returns: nil
//	NumericUnsigned("123.45")  // Returns: nil
//	NumericUnsigned("-123")     // Returns: error (invalid)
//	NumericUnsigned("abc")      // Returns: error (invalid)
//	NumericUnsigned("12.34.56") // Returns: error (invalid)
func NumericUnsigned(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only unsigned numeric characters
	if !regex.NumericUnsignedRegex().MatchString(value) {
		return fmt.Errorf("invalid unsigned numeric: %s", value)
	}

	return nil
}
