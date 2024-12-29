package rules

import (
	"fmt"

	"github.com/adhocore/gronx"
)

// CRON validates if the input is a valid CRON expression.
//
// A CRON expression typically consists of 5 or 6 space-separated fields, where each field can be:
// - Minutes (0 - 59)
// - Hours (0 - 23)
// - Day of month (1 - 31)
// - Month (1 - 12 or JAN - DEC)
// - Day of week (0 - 6 or SUN - SAT)
// Optionally, there can be a sixth field representing the year.
//
// Parameters:
// - input (any): The value to be validated. It should be convertible to a string.
//
// Returns:
// - error: If the CRON expression is invalid, it returns an error with details.
func Cron(input any) error {
	// Check if the input is a string
	value, ok := input.(string)
	if !ok {
		return fmt.Errorf("expected a string, got %T", input)
	}

	// Check if the string is a valid cron
	gron := gronx.New()
	if !gron.IsValid(value) {
		return fmt.Errorf("invalid cron: %s", value)
	}

	return nil
}
