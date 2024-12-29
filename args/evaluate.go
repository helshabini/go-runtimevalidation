package args

import (
	"fmt"
	"reflect"

	"go-runtimevalidation/functions"
)

// Evaluate function that traverses and evaluates based on the type of Arg
func (a Arg) Evaluate(obj any) (any, error) {
	if a.Type == FieldArg {
		if obj == nil {
			return nil, fmt.Errorf("object is nil")
		}
		fieldValue := reflect.ValueOf(obj).FieldByName(a.Field)
		if !fieldValue.IsValid() {
			return nil, fmt.Errorf("field %s not found in input", a.Field)
		}
		return fieldValue.Interface(), nil
	} else if a.Type == ConditionArg {
		// Evaluate condition
		lhsVal, err := a.Condition.Lhs.Evaluate(obj)
		if err != nil {
			return nil, err
		}
		rhsVal, err := a.Condition.Rhs.Evaluate(obj)
		if err != nil {
			return nil, err
		}
		return compare(lhsVal, rhsVal, a.Condition.Operator)
	} else if a.Type == FunctionArg {
		// Delegate function evaluation to EvaluateFunctionCall
		return EvaluateFunctionCall(a.Function, obj)
	}
	return a.Value, nil
}

// Separate function for evaluating function calls
func EvaluateFunctionCall(function Function, obj any) (any, error) {
	switch function.Name {
	case "len":
		if len(function.Args) != 1 {
			return nil, fmt.Errorf("len expects 1 argument")
		}
		argValue, err := function.Args[0].Evaluate(obj)
		if err != nil {
			return nil, err
		}
		return functions.GetLen(argValue)
	case "int":
		if len(function.Args) != 1 {
			return nil, fmt.Errorf("int expects 1 argument")
		}
		argValue, err := function.Args[0].Evaluate(obj)
		if err != nil {
			return nil, err
		}
		return functions.GetInt(argValue)
	case "float":
		if len(function.Args) != 1 {
			return nil, fmt.Errorf("float expects 1 argument")
		}
		argValue, err := function.Args[0].Evaluate(obj)
		if err != nil {
			return nil, err
		}
		return functions.GetFloat(argValue)
	// Add other function calls like "min", "max", etc.
	default:
		return nil, fmt.Errorf("unknown function: %s", function.Name)
	}
}

func compare(lhs, rhs interface{}, operator string) (bool, error) {
	if reflect.TypeOf(lhs) != reflect.TypeOf(rhs) {
		return false, fmt.Errorf("type mismatch: lhs is %T, rhs is %T", lhs, rhs)
	}
	switch operator {
	case "==":
		return reflect.DeepEqual(lhs, rhs), nil
	case "!=":
		return !reflect.DeepEqual(lhs, rhs), nil
	case ">":
		return compareNumbers(lhs, rhs, operator)
	case "<":
		return compareNumbers(lhs, rhs, operator)
	case ">=":
		return compareNumbers(lhs, rhs, operator)
	case "<=":
		return compareNumbers(lhs, rhs, operator)
	default:
		return false, fmt.Errorf("unknown operator: %s", operator)
	}
}

// Helper function to compare numbers
func compareNumbers(lhs, rhs interface{}, operator string) (bool, error) {
	switch val := any(lhs).(type) {
	case int, int8, int16, int32, int64:
		lhsInt := reflect.ValueOf(lhs).Int()
		rhsInt := reflect.ValueOf(rhs).Int()
		switch operator {
		case ">":
			return lhsInt > rhsInt, nil
		case "<":
			return lhsInt < rhsInt, nil
		case ">=":
			return lhsInt >= rhsInt, nil
		case "<=":
			return lhsInt <= rhsInt, nil
		}
	case uint, uint8, uint16, uint32, uint64:
		lhsUint := reflect.ValueOf(lhs).Uint()
		rhsUint := reflect.ValueOf(rhs).Uint()
		switch operator {
		case ">":
			return lhsUint > rhsUint, nil
		case "<":
			return lhsUint < rhsUint, nil
		case ">=":
			return lhsUint >= rhsUint, nil
		case "<=":
			return lhsUint <= rhsUint, nil
		}
	case float32, float64:
		lhsFloat := reflect.ValueOf(lhs).Float()
		rhsFloat := reflect.ValueOf(rhs).Float()
		switch operator {
		case ">":
			return lhsFloat > rhsFloat, nil
		case "<":
			return lhsFloat < rhsFloat, nil
		case ">=":
			return lhsFloat >= rhsFloat, nil
		case "<=":
			return lhsFloat <= rhsFloat, nil
		}
	default:
		return false, fmt.Errorf("unsupported type for comparison: %T", val)
	}

	return false, fmt.Errorf("unknown operator: %s", operator)
}
