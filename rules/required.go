package rules

import (
	"fmt"
	"reflect"
)

// Required checks if the input value is non-nil, non-zero, or non-empty.
//
// Parameters:
// - input: The value to be validated. This can be of any type.
//
// Behavior:
//   - For reference types (e.g., slices, maps, pointers, interfaces, channels, functions),
//     it returns an error if the value is nil.
//   - For other types (e.g., integers, floats, structs, strings), it checks if the value is zero (e.g., 0, "", nil).
//   - If the value is a pointer to a non-nil value, the check will succeed.
//   - If the input is valid (i.e., non-zero or non-nil), it returns nil. Otherwise, it returns an error indicating that
//     the value is required.
//
// Returns:
// - An error if the input is nil, zero, or empty.
// - nil if the input passes the required check.
func Required(input any) error {
	value := reflect.ValueOf(input)
	switch value.Kind() {
	case reflect.Slice, reflect.Map, reflect.Ptr, reflect.Interface, reflect.Chan, reflect.Func:
		if value.IsNil() {
			return fmt.Errorf("value is required")
		}
		return nil
	default:
		if value.Kind() == reflect.Ptr && value.Interface() != nil {
			return nil
		} else if value.IsValid() && !value.IsZero() {
			return nil
		} else {
			return fmt.Errorf("value is required")
		}
	}
}
