package rules

import (
	"fmt"
	"go-runtimevalidation/functions"
	"go-runtimevalidation/regex"
)

// JWT validates whether the given input is a properly formatted JWT token.
// The function checks if the input string contains a valid JWT token. This is a regex-based structural validation,
// and it does not verify the correctness of the token itself.
//
// Example
//   - "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
//   - "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ"
//
// Parameters:
//   - input (any): The value to validate (expected to be convertible to string).
//
// Returns:
//   - error: If the input is not valid or cannot be converted to a string,
//     or if it is not a valid JWT token.
func JWT(input any) error {
	// Check if the input is a string
	value, err := functions.GetString(input)
	if err != nil {
		return fmt.Errorf("unsupported type for input field: %w", err)
	}

	if len(value) == 0 {
		return fmt.Errorf("invalid jwt: %s", value)
	}

	// Check if the string is a valid html
	if !regex.JWTRegex().MatchString(value) {
		return fmt.Errorf("invalid jwt: %s", value)
	}

	return nil
}
