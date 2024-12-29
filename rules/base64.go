package rules

import (
	"encoding/base64"
	"fmt"
)

// Base64 checks whether the input string is a valid Base64 encoded string.
// If the input is not a string, or the Base64 string is invalid, it returns an error.
func Base64(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Validate the Base64 string by actually decoding it
	_, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return fmt.Errorf("invalid Base64 string: %s", value)
	}

	return nil
}
