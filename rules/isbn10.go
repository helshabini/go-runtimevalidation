package rules

import (
	"fmt"
	"go-runtimevalidation/regex"
)

// Isbn10 validates whether the input is a valid ISBN-10 string.
// The input must be a string that matches the ISBN-10 format (10 digits, optionally with dashes).
//
// Returns nil if the input is a valid ISBN-10 string, or an error if:
// - The input is not a string
// - The input does not match the ISBN-10 format
//
// Parameters:
// - input: The value to be validated, expected to be a string.
//
// Example:
//
//	err := Isbn10("0-306-40615-2")  // err will be nil if the string is a valid ISBN-10
//	err := Isbn10(12345)            // err will not be nil, as the input is not a string
//	err := Isbn10("invalid123")     // err will not be nil, as the string is not a valid ISBN-10
func Isbn10(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Remove any dashes or spaces to count the digits
	digitOnly := regex.ExtractDigitsRegex().ReplaceAllString(value, "")
	if len(digitOnly) != 10 {
		return fmt.Errorf("invalid ISBN10: must contain exactly 10 digits")
	}

	// Check if the string is a valid ISBN10
	var sum int
	for i, r := range digitOnly {
		digit := int(r - '0')
		sum += digit * (10 - i)
	}

	if sum%11 != 0 {
		return fmt.Errorf("invalid ISBN10: checksum failed")
	}

	return nil
}
