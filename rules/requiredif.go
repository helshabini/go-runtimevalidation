package rules

import "go-runtimevalidation/args"

// RequiredIf checks if the input is non-nil, non-zero, or non-empty
// if all specified fields in the comparison object (obj) match the values
// provided in the args map.
//
// Parameters:
// - input: The value to be checked if required. This can be any type.
// - obj: The struct object whose fields will be compared against the values in the args map.
// - args: A map of field names and their expected values. These are used to compare against the fields in obj.
//
// Returns:
// - nil if the input is valid or if the condition in the args map is not met.
// - An error from the Required function if the input fails the required check.
func RequiredIf(input any, obj any, args map[string]args.Arg) error {
	for _, arg := range args {
		result, err := arg.Evaluate(obj)
		if err != nil {
			return err
		}
		// If any condition is false, skip the Required check
		if result == false {
			return nil
		}
	}

	// If all conditions are true, run the Required check
	return Required(input)
}
