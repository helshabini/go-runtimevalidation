package rules

import (
	"fmt"
	"go-runtimevalidation/functions"
	"go-runtimevalidation/regex"
	"strings"

	"golang.org/x/net/html"
)

// HTML checks if the input is a valid HTML string.
// This function is intended to validate if the given input contains valid HTML content, such as tags or entities.
// It does not fully verify the correctness of the HTML structure (e.g., whether tags are correctly nested),
// but ensures that the input at least resembles basic HTML content.
//
// The input can be either a string or any type that can be converted to a string. If the input is not
// of a convertible type, an error is returned. If the input is an empty string, it will also return an error.
//
// The validation uses a regular expression that checks for the presence of typical HTML elements,
// like opening and closing tags (e.g., <p></p>), self-closing tags (e.g., <img />), or special HTML entities
// (e.g., &lt;, &amp;). Additionally, it check html integrity by parsing the input string. The parsing function is forgiving,
// so it will not return an error for every possible invalid HTML structure.
//
// Parameters:
// - input (any): The value to be validated. It should be convertible to a string.
//
// Returns:
// - error: If the input is not a valid HTML string or if it cannot be converted to a string, an error is returned.
func HTML(input any) error {
	// Check if the input is a string
	value, err := functions.GetString(input)
	if err != nil {
		return fmt.Errorf("unsupported type for input field: %w", err)
	}

	if len(value) == 0 {
		return fmt.Errorf("invalid html: %s", value)
	}

	// Check if the string is a valid html
	if !regex.HTMLRegex().MatchString(value) {
		return fmt.Errorf("invalid html: %s", value)
	}

	// Parse the HTML string and check for errors
	doc := strings.NewReader(value)
	_, err = html.Parse(doc)
	if err != nil {
		return fmt.Errorf("invalid html: %s", value)
	}

	return nil
}
