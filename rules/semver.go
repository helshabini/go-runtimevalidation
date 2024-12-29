package rules

import (
	"fmt"
	"go-runtimevalidation/functions"
	"go-runtimevalidation/regex"
)

// SemVer validates if the input is a valid Semantic Version (SemVer).
//
// A valid SemVer string has the following format:
// - MAJOR.MINOR.PATCH (e.g., 1.0.0)
// - Optionally followed by pre-release labels (e.g., -alpha.1)
// - Optionally followed by build metadata (e.g., +build.001)
//
// Parameters:
// - input (any): The value to be validated. It should be convertible to a string.
//
// Returns:
// - error: If the SemVer is invalid, it returns an error with details.
func SemVer(input any) error {
	// Check if the input is a string
	value, err := functions.GetString(input)
	if err != nil {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid semver
	if !regex.SemverRegex().MatchString(value) {
		return fmt.Errorf("invalid semver: %s", value)
	}

	return nil
}
