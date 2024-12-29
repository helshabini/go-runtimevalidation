package functions

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// GetLen evaluates the length of a value and returns it as an integer.
// It supports types that have a defined length, such as strings, slices, arrays, and maps.
//
// This function is useful for validation rules where the length of a field needs to be checked,
// such as `length:$len($Password)`. It can also be used internally to assess the length of a value.
//
// Parameters:
//   - value: The input value for which the length is to be evaluated. It can be of type string, slice, array, or map.
//
// Returns:
//   - An integer representing the length of the input value.
//   - An error if the type of the value does not support a length property (e.g., int, float).
//
// Example:
//
//	GetLen("hello")  // Returns: 5, nil
//	GetLen([]int{1, 2, 3})  // Returns: 3, nil
//	GetLen(123)  // Returns: 0, error
func GetLen(value any) (int, error) {
	val := reflect.ValueOf(value)
	switch val.Kind() {
	case reflect.String, reflect.Slice, reflect.Array, reflect.Map:
		return val.Len(), nil
	default:
		return 0, fmt.Errorf("unsupported type for len: %s", val.Kind())
	}
}

// GetIntAny converts an input of various types into an int64 value.
// It supports conversion from multiple data types such as int, uint, and string,
// making it flexible for use in validation rules like `min:$int($DateField)`.
//
// It can also parse a time.Duration or a string that represents an integer.
//
// Parameters:
//   - input: The value to be converted. It can be of type int, uint, time.Duration, or string.
//
// Returns:
//   - The converted int64 value.
//   - An error if the input type is unsupported or the conversion fails.
//
// Supported types:
//   - int, int8, int16, int32, int64
//   - uint, uint8, uint16, uint32, uint64
//   - time.Duration (parsed as int64 based on the duration)
//   - string (parsed using getInt function)
//
// Example:
//
//	GetIntAny(42)  // Returns: 42, nil
//	GetIntAny("12345")  // Returns: 12345, nil
//	GetIntAny("invalid")  // Returns: 0, error
func GetInt(input any) (int64, error) {
	value := reflect.ValueOf(input)
	switch v := input.(type) {
	case int, int32, int8, int16, int64:
		return value.Int(), nil
	case uint, uint8, uint16, uint32, uint64:
		return int64(value.Uint()), nil
	case time.Duration:
		// Return the duration in nanoseconds
		return int64(value.Int()), nil
	case time.Time:
		// Return the Unix timestamp (in seconds) for time.Time
		return v.Unix(), nil
	case string:
		// Try to parse the string as an int64 first
		if i, err := strconv.ParseInt(value.String(), 0, 64); err == nil {
			return i, nil
		}
		// Try to parse the string as a time
		if t, err := time.Parse(time.RFC3339, value.String()); err == nil {
			return t.Unix(), nil
		}
		return 0, fmt.Errorf("failed to parse %q of type %T as int64", value, v)
	default:
		return 0, fmt.Errorf("failed to parse %q of type %T as int64", value, v)
	}
}

// GetFloat converts an input of various types into a float64 value.
// It supports conversion from multiple data types such as int, uint, and string,
// making it flexible for use in validation rules requiring floating-point values.
//
// Parameters:
//   - input: The value to be converted. It can be of type int, uint, float, or string.
//
// Returns:
//   - The converted float64 value.
//   - An error if the input type is unsupported or the conversion fails.
//
// Supported types:
//   - int, int8, int16, int32, int64
//   - uint, uint8, uint16, uint32, uint64
//   - float32, float64
//   - string (parsed using strconv.ParseFloat)
//
// Example:
//
//	GetFloat(42)  // Returns: 42.0, nil
//	GetFloat("123.45")  // Returns: 123.45, nil
//	GetFloat("invalid")  // Returns: 0, error
func GetFloat(input any) (float64, error) {
	value := reflect.ValueOf(input)
	switch v := input.(type) {
	case int, int32, int8, int16, int64:
		return float64(value.Int()), nil
	case uint, uint8, uint16, uint32, uint64:
		return float64(value.Uint()), nil
	case float32:
		return float64(value.Float()), nil
	case float64:
		return value.Float(), nil
	case string:
		// Try to parse the string as a float64
		if f, err := strconv.ParseFloat(value.String(), 64); err == nil {
			return f, nil
		}
		return 0, fmt.Errorf("failed to parse %q of type %T as float64", value, v)
	default:
		return 0, fmt.Errorf("failed to parse %q of type %T as float64", value, v)
	}
}

// GetString converts an input of various types into a string value.
// It supports conversion from multiple data types such as int, uint, and float,
// making it flexible for use in validation rules requiring string values.
//
// Parameters:
//   - input: The value to be converted. It can be of type int, uint, float, bool, or string.
//
// Returns:
//   - The converted string value.
//   - An error if the input type is unsupported or the conversion fails.
//
// Supported types:
//   - int, int8, int16, int32, int64
//   - uint, uint8, uint16, uint32, uint64
//   - float32, float64
//   - bool
//   - string
//
// Example:
//
//	GetString(42)  // Returns: "42", nil
//	GetString(true)  // Returns: "true", nil
//	GetString("hello")  // Returns: "hello", nil
func GetString(input any) (string, error) {
	value := reflect.ValueOf(input)
	switch v := input.(type) {
	case string:
		return value.String(), nil
	case int, int32, int8, int16, int64:
		return strconv.FormatInt(value.Int(), 10), nil
	case uint, uint8, uint16, uint32, uint64:
		return strconv.FormatUint(value.Uint(), 10), nil
	case float32:
		return strconv.FormatFloat(value.Float(), 'f', -1, 32), nil
	case float64:
		return strconv.FormatFloat(value.Float(), 'f', -1, 64), nil
	case bool:
		return strconv.FormatBool(value.Bool()), nil
	default:
		return "", fmt.Errorf("failed to parse %q of type %T as string", value, v)
	}
}
