package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// MD5 checks if the input is a valid MD5 hash.
// The input should be a string and must match the regex for an MD5 hash (32 hexadecimal characters).
//
// Parameters:
//   - input: the value to validate, expected to be a string.
//
// Returns:
//   - error: nil if the input is a valid MD5 hash, or an error if the input is invalid.
//
// Errors:
//   - If the input is not a string, an error will be returned indicating the expected and received types.
//   - If the input does not match the MD5 regex (i.e., it is not a valid MD5), an error will be returned.
func MD5(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid MD5 value
	if !regex.Md5Regex().MatchString(value) {
		return fmt.Errorf("invalid MD5 value: %s", value)
	}

	return nil
}
