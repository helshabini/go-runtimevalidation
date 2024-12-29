package rules

import (
	"encoding/base32"
	"fmt"
)

// Base32 checks whether the input string is a valid Base32 encoded string.
// If the input is not a string, or the Base32 string is invalid, it returns an error.
func Base32(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Validate the Base32 string by actually decoding it
	_, err := base32.StdEncoding.DecodeString(value)
	if err != nil {
		return fmt.Errorf("invalid Base32 string: %s", value)
	}

	return nil
}
