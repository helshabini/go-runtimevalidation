package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseValue(t *testing.T) {
	t.Run("should return a number", func(t *testing.T) {
		value, err := parseValue("1")
		assert.NoError(t, err)
		assert.Equal(t, 1, value)
	})

	t.Run("should return a float", func(t *testing.T) {
		value, err := parseValue("1.1")
		assert.NoError(t, err)
		assert.Equal(t, 1.1, value)
	})

	t.Run("should return a boolean", func(t *testing.T) {
		value, err := parseValue("true")
		assert.NoError(t, err)
		assert.Equal(t, true, value)
	})

	t.Run("should return an array", func(t *testing.T) {
		value, err := parseValue("[1,2,3]")
		assert.NoError(t, err)
		assert.Equal(t, []interface{}{1, 2, 3}, value)
	})

	t.Run("should return an array (spaced)", func(t *testing.T) {
		value, err := parseValue("[1 , 2 , 3]")
		assert.NoError(t, err)
		assert.Equal(t, []interface{}{1, 2, 3}, value)
	})

	t.Run("should return a map", func(t *testing.T) {
		value, err := parseValue(`{"key1":"value1","key2":2}`)
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"key1": "value1", "key2": 2}, value)
	})

	t.Run("should return a map (spaced)", func(t *testing.T) {
		value, err := parseValue("{\"key1\" : \"value1\" ,\"key2\" : 2 }")
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"key1": "value1", "key2": 2}, value)
	})

	t.Run("should return a string", func(t *testing.T) {
		value, err := parseValue("hello")
		assert.NoError(t, err)
		assert.Equal(t, "hello", value)
	})

	t.Run("should return a complex nested array", func(t *testing.T) {
		value, err := parseValue("[[1,2],[3,4]]")
		assert.NoError(t, err)
		assert.Equal(t, []interface{}{[]interface{}{1, 2}, []interface{}{3, 4}}, value)
	})

	t.Run("should return a complex nested map", func(t *testing.T) {
		value, err := parseValue(`{"outerKey":{"innerKey": "innerValue"}}`)
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{"outerKey": map[string]interface{}{"innerKey": "innerValue"}}, value)
	})

	t.Run("should return a nil value", func(t *testing.T) {
		value, err := parseValue("")
		assert.Error(t, err)
		assert.Nil(t, value)
	})

	t.Run("should return a boolean false", func(t *testing.T) {
		value, err := parseValue("false")
		assert.NoError(t, err)
		assert.Equal(t, false, value)
	})

	t.Run("should return an empty array", func(t *testing.T) {
		value, err := parseValue("[]")
		assert.NoError(t, err)
		assert.Equal(t, []interface{}{}, value)
	})

	t.Run("should return an empty map", func(t *testing.T) {
		value, err := parseValue("{}")
		assert.NoError(t, err)
		assert.Equal(t, map[string]interface{}{}, value)
	})
}

func TestParseFunc(t *testing.T) {
	t.Run("should return a function (single arg)", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func(a)")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 1)
		assert.Equal(t, "a", args[0].Value)
	})

	t.Run("should return a function (double args)", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func(a, b)")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 2)
		assert.Equal(t, "a", args[0].Value)
		assert.Equal(t, "b", args[1].Value)
	})

	t.Run("should return a function (trible args)", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func(a, b, c)")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 3)
		assert.Equal(t, "a", args[0].Value)
		assert.Equal(t, "b", args[1].Value)
		assert.Equal(t, "c", args[2].Value)
	})

	t.Run("should return a function (different arg basic types)", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func(a, 1, true)")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 3)
		assert.Equal(t, "a", args[0].Value)
		assert.Equal(t, 1, args[1].Value)
		assert.Equal(t, true, args[2].Value)
	})

	t.Run("should return a function (single field arg)", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func($Name)")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 1)
		assert.Equal(t, FieldArg, args[0].Type)
		assert.Equal(t, "Name", args[0].Field)
	})

	t.Run("should return a function (nested func arg)", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func($len($Name))")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 1)
		assert.Equal(t, FunctionArg, args[0].Type)
		assert.Equal(t, "len", args[0].Function.Name)
		assert.Equal(t, FieldArg, args[0].Function.Args[0].Type)
		assert.Equal(t, "Name", args[0].Function.Args[0].Field)
	})

	t.Run("should return a function (array arg)", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func([1,2,3])")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Equal(t, []interface{}{1, 2, 3}, args[0].Value)
	})

	t.Run("should return a function (map arg)", func(t *testing.T) {
		fn, args, err := parseFunctionCall(`$func({"key1":"value1","key2":2})`)
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Equal(t, map[string]interface{}{"key1": "value1", "key2": 2}, args[0].Value)
	})

	t.Run("should return a function with a condition arg", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func($Name == \"John\")")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 1)
		assert.Equal(t, ConditionArg, args[0].Type)
		assert.Equal(t, "Name", args[0].Condition.Lhs.Field)
		assert.Equal(t, "John", args[0].Condition.Rhs.Value)
		assert.Equal(t, "==", args[0].Condition.Operator)
	})

	t.Run("should handle multiple condition args", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func($Age >= 18, $Country == \"USA\")")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 2)

		assert.Equal(t, ConditionArg, args[0].Type)
		assert.Equal(t, "Age", args[0].Condition.Lhs.Field)
		assert.Equal(t, 18, args[0].Condition.Rhs.Value)
		assert.Equal(t, ">=", args[0].Condition.Operator)

		assert.Equal(t, ConditionArg, args[1].Type)
		assert.Equal(t, "Country", args[1].Condition.Lhs.Field)
		assert.Equal(t, "USA", args[1].Condition.Rhs.Value)
		assert.Equal(t, "==", args[1].Condition.Operator)
	})

	t.Run("should return a function with mixed args", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func($Name, $len($Items), 5)")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 3)

		assert.Equal(t, FieldArg, args[0].Type)
		assert.Equal(t, "Name", args[0].Field)

		assert.Equal(t, FunctionArg, args[1].Type)
		assert.Equal(t, "len", args[1].Function.Name)
		assert.Equal(t, FieldArg, args[1].Function.Args[0].Type)
		assert.Equal(t, "Items", args[1].Function.Args[0].Field)

		assert.Equal(t, 5, args[2].Value)
	})
}

func TestParseCondition(t *testing.T) {
	t.Run("should parse simple condition", func(t *testing.T) {
		condition, err := parseCondition("$len($Name)>3")
		assert.NoError(t, err)
		assert.Equal(t, Condition{
			Lhs:      &Arg{Type: FunctionArg, Function: Function{Name: "len", Args: []Arg{{Type: FieldArg, Field: "Name"}}}},
			Rhs:      &Arg{Type: ValueArg, Value: 3},
			Operator: ">",
		}, condition)
	})

	t.Run("should parse simple condition (spaced)", func(t *testing.T) {
		condition, err := parseCondition("$Age >= 18")
		assert.NoError(t, err)
		assert.Equal(t, Condition{
			Lhs:      &Arg{Type: FieldArg, Field: "Age"},
			Rhs:      &Arg{Type: ValueArg, Value: 18},
			Operator: ">=",
		}, condition)
	})

	t.Run("should parse condition with function on left side", func(t *testing.T) {
		condition, err := parseCondition("$len($Name) == 5")
		assert.NoError(t, err)
		assert.Equal(t, Condition{
			Lhs: &Arg{
				Type: FunctionArg,
				Function: Function{
					Name: "len",
					Args: []Arg{{Type: FieldArg, Field: "Name"}},
				},
			},
			Rhs:      &Arg{Type: ValueArg, Value: 5},
			Operator: "==",
		}, condition)
	})

	t.Run("should parse condition with function on right side", func(t *testing.T) {
		condition, err := parseCondition("$Age == $len($Name)")
		assert.NoError(t, err)
		assert.Equal(t, Condition{
			Lhs: &Arg{Type: FieldArg, Field: "Age"},
			Rhs: &Arg{
				Type: FunctionArg,
				Function: Function{
					Name: "len",
					Args: []Arg{{Type: FieldArg, Field: "Name"}},
				},
			},
			Operator: "==",
		}, condition)
	})
}

func TestParseArgWithCondition(t *testing.T) {
	t.Run("should handle function with condition argument", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func($Name == \"John\")")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 1)
		assert.Equal(t, ConditionArg, args[0].Type)
		// You can add more checks to validate the condition structure
	})

	t.Run("should handle multiple condition args", func(t *testing.T) {
		fn, args, err := parseFunctionCall("$func($Age >= 18, $City == \"London\")")
		assert.NoError(t, err)
		assert.Equal(t, "func", fn)
		assert.Len(t, args, 2)
		assert.Equal(t, ConditionArg, args[0].Type)
		assert.Equal(t, ConditionArg, args[1].Type)
		// You can add more checks to validate the condition structures
	})
}
func TestParseArgs(t *testing.T) {
	t.Run("should parse single value argument", func(t *testing.T) {
		args, err := ParseArgs("1")
		assert.NoError(t, err)
		assert.Len(t, args, 1)
		assert.Equal(t, 1, args["1"].Value)
	})

	t.Run("should parse multiple value arguments", func(t *testing.T) {
		args, err := ParseArgs("1, 2, 3")
		assert.NoError(t, err)
		assert.Len(t, args, 3)
		assert.Equal(t, 1, args["1"].Value)
		assert.Equal(t, 2, args["2"].Value)
		assert.Equal(t, 3, args["3"].Value)
	})

	t.Run("should parse field argument", func(t *testing.T) {
		args, err := ParseArgs("$Age")
		assert.NoError(t, err)
		assert.Len(t, args, 1)
		assert.Equal(t, FieldArg, args["$Age"].Type)
		assert.Equal(t, "Age", args["$Age"].Field)
	})

	t.Run("should parse function call argument", func(t *testing.T) {
		args, err := ParseArgs("$len($Name)")
		assert.NoError(t, err)
		assert.Len(t, args, 1)
		assert.Equal(t, FunctionArg, args["$len($Name)"].Type)
		assert.Equal(t, "len", args["$len($Name)"].Function.Name)
		assert.Len(t, args["$len($Name)"].Function.Args, 1)
		assert.Equal(t, FieldArg, args["$len($Name)"].Function.Args[0].Type)
		assert.Equal(t, "Name", args["$len($Name)"].Function.Args[0].Field)
	})

	t.Run("should parse condition argument", func(t *testing.T) {
		args, err := ParseArgs("$Age >= 18")
		assert.NoError(t, err)
		assert.Len(t, args, 1)
		assert.Equal(t, ConditionArg, args["$Age >= 18"].Type)
		assert.Equal(t, "Age", args["$Age >= 18"].Condition.Lhs.Field)
		assert.Equal(t, 18, args["$Age >= 18"].Condition.Rhs.Value)
		assert.Equal(t, ">=", args["$Age >= 18"].Condition.Operator)
	})

	t.Run("should parse mixed arguments", func(t *testing.T) {
		args, err := ParseArgs("$Name, 1, $len($Items), $Age >= 18")
		assert.NoError(t, err)
		assert.Len(t, args, 4)

		assert.Equal(t, FieldArg, args["$Name"].Type)
		assert.Equal(t, "Name", args["$Name"].Field)

		assert.Equal(t, ValueArg, args["1"].Type)
		assert.Equal(t, 1, args["1"].Value)

		assert.Equal(t, FunctionArg, args["$len($Items)"].Type)
		assert.Equal(t, "len", args["$len($Items)"].Function.Name)
		assert.Len(t, args["$len($Items)"].Function.Args, 1)
		assert.Equal(t, FieldArg, args["$len($Items)"].Function.Args[0].Type)
		assert.Equal(t, "Items", args["$len($Items)"].Function.Args[0].Field)

		assert.Equal(t, ConditionArg, args["$Age >= 18"].Type)
		assert.Equal(t, "Age", args["$Age >= 18"].Condition.Lhs.Field)
		assert.Equal(t, 18, args["$Age >= 18"].Condition.Rhs.Value)
		assert.Equal(t, ">=", args["$Age >= 18"].Condition.Operator)
	})

	t.Run("should parse escaped value argument", func(t *testing.T) {
		args, err := ParseArgs("\"escaped\"")
		assert.NoError(t, err)
		assert.Len(t, args, 1)
		assert.Equal(t, ValueArg, args["\"escaped\""].Type)
		assert.Equal(t, "escaped", args["\"escaped\""].Value)
	})
}
