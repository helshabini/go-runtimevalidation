package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// HexColor validates whether the input is a valid hexadecimal color code.
// A valid color code can be in the following formats:
// - 3 hex digits (e.g., #FFF)
// - 4 hex digits (e.g., #FFFF)
// - 6 hex digits (e.g., #FFFFFF)
// - 8 hex digits (e.g., #FFFFFFFF)
//
// This function is useful for validation rules where color codes are required,
// such as in web design or graphics applications.
//
// Parameters:
//   - input: The value to be validated. It must be a string.
//
// Returns:
//   - An error if the input is not a string or if it does not match the hexadecimal color code format.
//
// Example:
//
//	HexColor("#FFF")          // Returns: nil
//	HexColor("#FFFF")         // Returns: nil
//	HexColor("#FFFFFF")       // Returns: nil
//	HexColor("#FFFFFFFF")     // Returns: nil
//	HexColor("#XYZ")          // Returns: error (invalid)
//	HexColor("#12345678")     // Returns: error (invalid)
//	HexColor("#123")          // Returns: nil
func HexColor(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid hexadecimal color code
	if !regex.HexColorRegex().MatchString(value) {
		return fmt.Errorf("invalid hex color: %s", value)
	}

	return nil
}
