package args

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ArgType int

const (
	ValueArg ArgType = iota
	FieldArg
	ConditionArg
	FunctionArg
)

type Arg struct { // Represents a field, function call, or value
	Type      ArgType
	Field     string
	Value     any
	Function  Function
	Condition Condition
}

type Function struct {
	Name    string
	Args    []Arg
	Returns reflect.Type
}

type Condition struct { // Evaluates to true or false
	Lhs      *Arg
	Rhs      *Arg
	Operator string
}

func ParseArgs(text string) (map[string]Arg, error) {
	// Prepare to hold the parsed arguments
	argsMap := make(map[string]Arg)

	// Split the input text into individual components by comma
	parts := splitAndHandleEscapes(text, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)

		// Case 1: Handle Condition (e.g. $Age>18, $Name=="John")
		if isCondition(part) {
			condition, err := parseCondition(part)
			if err != nil {
				return nil, err
			}

			// Store the condition argument
			argsMap[part] = Arg{
				Type:      ConditionArg,
				Condition: condition,
			}

			// Case 2: Handle Function Call (e.g. $len($Name))
		} else if isFunctionCall(part) {
			funcName, funcArgs, err := parseFunctionCall(part)
			if err != nil {
				return nil, err
			}

			// Store the function argument
			argsMap[part] = Arg{
				Type: FunctionArg,
				Function: Function{
					Name: funcName,
					Args: funcArgs,
				},
			}

			// Case 3: Handle Field Reference (e.g. $Age)
		} else if isField(part) {
			// Store the field argument
			argsMap[part] = Arg{
				Type:  FieldArg,
				Field: strings.TrimPrefix(part, "$"),
			}

			// Case 4: Handle Escaped Value
		} else if strings.HasPrefix(part, `\`) {
			// Remove escape characters and treat as a value
			value := unescapeText(part)
			argsMap[part] = Arg{
				Type:  ValueArg,
				Value: value,
			}

			// Case 5: Handle Value (e.g. 1, "John")
		} else {
			value, err := parseValue(part)
			if err != nil {
				return nil, err
			}

			// Store the value argument
			argsMap[part] = Arg{
				Type:  ValueArg,
				Value: value,
			}
		}
	}

	return argsMap, nil
}

// Helper function to split arguments while handling nested structures like arrays or maps
func splitAndHandleEscapes(s string, sep string) []string {
	var parts []string
	var currentPart strings.Builder
	nestingLevel := 0
	inQuotes := false

	for i := 0; i < len(s); i++ {
		c := s[i]

		switch c {
		case '[', '{':
			nestingLevel++
		case ']', '}':
			nestingLevel--
		case '"':
			inQuotes = !inQuotes // Toggle inQuotes on encountering a quote
		case ',':
			// If we encounter a comma and we are not inside a nested structure and not in quotes, split
			if nestingLevel == 0 && !inQuotes {
				parts = append(parts, strings.TrimSpace(currentPart.String()))
				currentPart.Reset()
				continue
			}
		case ':':
			// If we encounter a colon and we are not inside a nested structure and not in quotes
			// Check if we are parsing a map entry
			if nestingLevel == 0 && !inQuotes {
				// If we are parsing a map entry, split the key-value pair if we encounter a colon
				if sep == ":" {
					parts = append(parts, strings.TrimSpace(currentPart.String()))
					currentPart.Reset()
					continue
				}
				// Otherwise,
				// Since this is a map, we might want to just continue accumulating the part
				// rather than splitting here.
				// The colon indicates a key-value separation within the current part.
				// We don't want to split yet; we'll handle it later.
				// We still add the colon to the current part because without it the key-value pair is malformed.
				currentPart.WriteByte(c)
				continue
			}
		}
		// Add the current character to the current part
		currentPart.WriteByte(c)
	}

	// Add the last part
	if currentPart.Len() > 0 {
		parts = append(parts, strings.TrimSpace(currentPart.String()))
	}

	return parts
}

// Helper function to unescape text (remove backslashes)
func unescapeText(text string) string {
	var result string
	var escaped bool

	for _, ch := range text {
		if escaped {
			result += string(ch)
			escaped = false
		} else if ch == '\\' {
			escaped = true
		} else {
			result += string(ch)
		}
	}

	return result
}

// Helper function to parse a function call
func parseFunctionCall(text string) (string, []Arg, error) {
	if !strings.HasPrefix(text, "$") || !strings.Contains(text, "(") {
		return "", nil, fmt.Errorf("invalid function call: %s", text)
	}

	openParen := strings.Index(text, "(")
	closeParen := strings.LastIndex(text, ")")

	if openParen == -1 || closeParen == -1 || closeParen < openParen {
		return "", nil, fmt.Errorf("invalid function format: %s", text)
	}

	funcName := strings.TrimPrefix(text[:openParen], "$")
	argsText := text[openParen+1 : closeParen]

	argParts := splitAndHandleEscapes(argsText, ",")
	args := make([]Arg, 0, len(argParts))

	for _, argText := range argParts {
		argText = strings.TrimSpace(argText)

		// Parse argument as either field, function, or value
		arg, err := parseArg(argText)
		if err != nil {
			return "", nil, err
		}
		args = append(args, arg)
	}

	return funcName, args, nil
}

// Helper function to parse an individual argument (field, function, or value)
func parseArg(text string) (Arg, error) {
	if isField(text) {
		// Handle field argument
		return Arg{
			Type:  FieldArg,
			Field: strings.TrimPrefix(text, "$"),
		}, nil
	} else if isFunctionCall(text) {
		// Handle function call argument
		funcName, funcArgs, err := parseFunctionCall(text)
		if err != nil {
			return Arg{}, err
		}

		return Arg{
			Type: FunctionArg,
			Function: Function{
				Name: funcName,
				Args: funcArgs,
			},
		}, nil
	} else if isCondition(text) {
		// Handle condition argument
		condition, err := parseCondition(text)
		if err != nil {
			return Arg{}, err
		}

		return Arg{
			Type:      ConditionArg,
			Condition: condition,
		}, nil
	} else if strings.HasPrefix(text, `\`) {
		// Handle escaped text as a value argument
		return Arg{
			Type:  ValueArg,
			Value: unescapeText(text),
		}, nil
	} else {
		// Handle value argument
		value, err := parseValue(text)
		if err != nil {
			return Arg{}, err
		}

		return Arg{
			Type:  ValueArg,
			Value: value,
		}, nil
	}
}

func parseCondition(text string) (Condition, error) {
	// Detect comparison operators and split accordingly
	operators := []string{"<=", ">=", "==", "!=", "<", ">"}
	for _, op := range operators {
		if strings.Contains(text, op) {
			parts := strings.SplitN(text, op, 2)
			if len(parts) != 2 {
				return Condition{}, fmt.Errorf("invalid condition: %s", text)
			}

			lhsText := strings.TrimSpace(parts[0])
			rhsText := strings.TrimSpace(parts[1])

			// Parse the left-hand side argument (could be field, function, or value)
			lhsArg, err := parseArg(lhsText)
			if err != nil {
				return Condition{}, err
			}

			// Parse the right-hand side argument (could be field, function, or value)
			rhsArg, err := parseArg(rhsText)
			if err != nil {
				return Condition{}, err
			}

			return Condition{
				Lhs:      &lhsArg,
				Rhs:      &rhsArg,
				Operator: op,
			}, nil
		}
	}

	return Condition{}, fmt.Errorf("invalid condition: %s", text)
}

// parseValue is the main function to parse a string into its corresponding type.
func parseValue(s string) (interface{}, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return nil, fmt.Errorf("empty value")
	}

	if strings.HasPrefix(s, "{") {
		return parseMapValue(s)
	}

	if strings.HasPrefix(s, "[") {
		return parseArrayValue(s)
	}

	// Check for quoted strings
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return strings.Trim(s, "\""), nil
	}

	return parseValueType(s), nil
}

// parseMapValue parses a string representation of a map into a map[string]interface{}.
func parseMapValue(s string) (map[string]interface{}, error) {
	// Try to parse the value as a map (e.g., '{"key1":"value1","key2":2}')
	if strings.HasPrefix(s, "{") && strings.HasSuffix(s, "}") {
		mapText := s[1 : len(s)-1]
		mapParts := splitAndHandleEscapes(mapText, ",")
		resultMap := make(map[string]any)

		for _, part := range mapParts {
			kvPair := splitAndHandleEscapes(part, ":")
			if len(kvPair) != 2 {
				return nil, fmt.Errorf("invalid map entry: %s", part)
			}

			key := strings.TrimSpace(kvPair[0])
			valueText := strings.TrimSpace(kvPair[1])

			// Parse the value for the key-value pair
			value, err := parseValue(valueText)
			if err != nil {
				return nil, err
			}

			// Remove quotes around the key and value if necessary
			key = strings.Trim(key, `"`)
			if strValue, ok := value.(string); ok {
				value = strings.Trim(strValue, `"`)
			}

			resultMap[key] = value
		}

		return resultMap, nil
	}

	return nil, fmt.Errorf("invalid map format: %s", s)
}

// parseArrayValue parses a string representation of an array into a []interface{}.
func parseArrayValue(s string) ([]interface{}, error) {
	// Try to parse the value as an array (e.g., "[1,2,3]")
	if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
		// Remove the square brackets
		arrayText := s[1 : len(s)-1]
		arrayParts := splitAndHandleEscapes(arrayText, ",")
		array := make([]any, 0, len(arrayParts))

		for _, part := range arrayParts {
			part = strings.TrimSpace(part)
			value, err := parseValue(part)
			if err != nil {
				return nil, err
			}
			array = append(array, value)
		}

		return array, nil
	}

	return nil, fmt.Errorf("invalid array format: %s", s)
}

// parseValueType determines the type of value and returns it.
func parseValueType(s string) interface{} {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return strings.Trim(s, "\"")
	}

	// Try to parse the value as an integer
	if intValue, err := strconv.Atoi(s); err == nil {
		return intValue
	}

	// Try to parse the value as a float
	if floatValue, err := strconv.ParseFloat(s, 64); err == nil {
		return floatValue
	}

	// Try to parse the value as a boolean
	if boolValue, err := strconv.ParseBool(s); err == nil {
		return boolValue
	}

	return s // Return as string if no other type matched
}

// Helper function to determine if the input is a field reference (starts with $)
func isField(s string) bool {
	return strings.HasPrefix(s, "$") && !strings.Contains(s, "(") && !strings.Contains(s, "{") && !strings.Contains(s, "[") && !isCondition(s)
}

// Helper function to determine if the input is a function call (starts with $ and contains parentheses)
func isFunctionCall(s string) bool {
	return strings.HasPrefix(s, "$") && strings.Contains(s, "(")
}

// Helper function to determine if the input is a condition (contains an operator)
func isCondition(s string) bool {
	operators := []string{"<=", ">=", "==", "!=", "<", ">"}
	for _, op := range operators {
		if strings.Contains(s, op) {
			return true
		}
	}
	return false
}
