package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Numeric validates that the input is a numeric value.
// It checks if the input is a string that can represent an integer or a floating-point number,
// including optional signs for positive or negative values.
//
// The function utilizes a regular expression to match the following patterns:
// - Optional leading '+' or '-' sign
// - One or more digits (0-9)
// - An optional decimal point followed by one or more digits
//
// Parameters:
//   - input: The value to be validated. It must be of type string.
//
// Returns:
//   - An error if the input is not a string or if it does not match the numeric pattern,
//     indicating the reason for the failure.
//
// Example usage:
//
//	err := Numeric("-123.45")  // Returns: nil (valid)
//	err := Numeric("42")        // Returns: nil (valid)
//	err := Numeric("abc")       // Returns: error (invalid)
//	err := Numeric("12.34.56")  // Returns: error (invalid)
//
// The regex pattern used for validation is:
// "^[-+]?[0-9]+(?:\\.[0-9]+)?$"
func Numeric(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string contains only numeric characters
	if !regex.NumericRegex().MatchString(value) {
		return fmt.Errorf("invalid numeric: %s", value)
	}

	return nil
}
