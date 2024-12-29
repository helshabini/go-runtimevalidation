package rules

import (
	"fmt"
	"unicode"
)

// Lowercase checks if all the letters in the input string are lowercase.
//
// It validates the input to ensure that every letter is a lowercase character. Non-letter
// characters are ignored during validation. If any uppercase letters are found, the function
// returns an error indicating that the input string is invalid.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
// - error: nil if all letters are lowercase, otherwise an error indicating invalid input.
func Lowercase(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	for _, r := range value {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return fmt.Errorf("invalid lowercase: %s", value)
		}
	}

	return nil
}
