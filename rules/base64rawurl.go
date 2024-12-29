package rules

import (
	"encoding/base64"
	"fmt"
)

// Base64RawUrl checks whether the input string is a valid Base64 URL encoded string without padding.
// If the input is not a string, or the Base64 URL string is invalid, it returns an error.
func Base64RawUrl(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Validate the Base64 string by actually decoding it
	_, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil {
		return fmt.Errorf("invalid Base64RawUrl string: %s", value)
	}

	return nil
}
