package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Isbn13 validates whether the input is a valid ISBN-13 string.
// It expects a string input and checks if it matches the ISBN-13 format using a regular expression.
//
// Parameters:
// - input: The value being validated, expected to be a string.
//
// Returns nil if the input is a valid ISBN-13, or an error if:
// - The input is not a string
// - The input does not match the ISBN-13 format
//
// Example:
//
//	err := Isbn13("9783161484100")  // err will be nil for a valid ISBN-13
//	err := Isbn13("invalid")        // err will be non-nil for an invalid ISBN-13
//	err := Isbn13(12345)            // err will be non-nil as the input is not a string
func Isbn13(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Remove any dashes or spaces to count the digits
	digitOnly := regex.ExtractDigitsRegex().ReplaceAllString(value, "")
	if len(digitOnly) != 13 {
		return fmt.Errorf("invalid ISBN13: must contain exactly 13 digits")
	}

	// Check if the string is a valid ISBN13
	var sum int
	for i, r := range digitOnly {
		digit := int(r - '0')
		if i%2 == 0 {
			sum += digit
		} else {
			sum += digit * 3
		}
	}

	if sum%10 != 0 {
		return fmt.Errorf("invalid ISBN13: checksum failed")
	}

	return nil
}
