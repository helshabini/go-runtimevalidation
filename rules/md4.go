package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// MD4 checks if the input is a valid MD4 hash.
// The input should be a string and must match the regex for an MD4 hash (32 hexadecimal characters).
//
// Parameters:
//   - input: the value to validate, expected to be a string.
//
// Returns:
//   - error: nil if the input is a valid MD4 hash, or an error if the input is invalid.
//
// Errors:
//   - If the input is not a string, an error will be returned indicating the expected and received types.
//   - If the input does not match the MD4 regex (i.e., it is not a valid MD4), an error will be returned.
func MD4(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid MD4 value
	if !regex.Md4Regex().MatchString(value) {
		return fmt.Errorf("invalid MD4 value: %s", value)
	}

	return nil
}
