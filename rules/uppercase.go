package rules

import (
	"fmt"
	"unicode"
)

// Uppercase validates whether the input string consists entirely of uppercase letters.
//
// The function checks each character in the input. If a character is a letter and it is not
// uppercase, it returns an error. Non-letter characters are allowed and do not affect validation.
//
// Parameters:
// - input: the value to be validated (expected to be a string).
//
// Returns:
// - error: nil if the input is entirely uppercase letters or non-letter characters, otherwise an error indicating the failure.
func Uppercase(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	for _, r := range value {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return fmt.Errorf("invalid uppercase: %s", value)
		}
	}

	return nil
}
