package rules

import (
	"encoding/base32"
	"fmt"
)

// Base32Hex checks whether the input string is a valid Base32Hex encoded string.
// If the input is not a string, or the Base32Hex string is invalid, it returns an error.
func Base32Hex(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Validate the Base32Hex string by actually decoding it
	_, err := base32.HexEncoding.DecodeString(value)
	if err != nil {
		return fmt.Errorf("invalid Base32Hex string: %s", value)
	}

	return nil
}
