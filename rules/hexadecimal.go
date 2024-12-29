package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Hexadecimal validates whether the input is a valid hexadecimal string.
// The string may optionally start with "0x" or "0X" followed by one or more hexadecimal digits.
// This function is useful for validation rules where hexadecimal values are required,
// such as in color codes or memory addresses.
//
// Parameters:
//   - input: The value to be validated. It must be a string.
//
// Returns:
//   - An error if the input is not a string or if it does not match the hexadecimal format.
//
// Example:
//
//	Hexadecimal("0x1A3F")    // Returns: nil
//	Hexadecimal("1A3F")       // Returns: nil
//	Hexadecimal("0x")         // Returns: error (invalid)
//	Hexadecimal("GHIJK")      // Returns: error (invalid)
//	Hexadecimal("123")        // Returns: nil
func Hexadecimal(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid hexadecimal
	if !regex.HexadecimalRegex().MatchString(value) {
		return fmt.Errorf("invalid hexadecimal: %s", value)
	}

	return nil
}
