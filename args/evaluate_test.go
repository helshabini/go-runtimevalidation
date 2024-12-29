package args

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name    string
	Age     int
	MaxSize int
}

// Test the Evaluate function for different Arg types
func TestEvaluate(t *testing.T) {
	obj := TestStruct{Name: "John", Age: 25, MaxSize: 100}

	t.Run("Evaluate FieldArg", func(t *testing.T) {
		arg := Arg{
			Type:  FieldArg,
			Field: "Name",
		}

		result, err := arg.Evaluate(obj)
		assert.NoError(t, err)
		assert.Equal(t, "John", result)
	})

	t.Run("Evaluate ConditionArg Equal", func(t *testing.T) {
		condition := Condition{
			Lhs:      &Arg{Type: FieldArg, Field: "Age"},
			Rhs:      &Arg{Value: 25},
			Operator: "==",
		}
		arg := Arg{
			Type:      ConditionArg,
			Condition: condition,
		}

		result, err := arg.Evaluate(obj)
		assert.NoError(t, err)
		assert.Equal(t, true, result)
	})

	t.Run("Evaluate ConditionArg Greater Than", func(t *testing.T) {
		condition := Condition{
			Lhs:      &Arg{Type: FieldArg, Field: "Age"},
			Rhs:      &Arg{Value: 20},
			Operator: ">",
		}
		arg := Arg{
			Type:      ConditionArg,
			Condition: condition,
		}

		result, err := arg.Evaluate(obj)
		assert.NoError(t, err)
		assert.Equal(t, true, result)
	})

	t.Run("Evaluate FunctionArg Len", func(t *testing.T) {
		function := Function{
			Name: "len",
			Args: []Arg{{Type: FieldArg, Field: "Name"}},
		}
		arg := Arg{
			Type:     FunctionArg,
			Function: function,
		}

		result, err := arg.Evaluate(obj)
		assert.NoError(t, err)
		assert.Equal(t, 4, result)
	})

	t.Run("Evaluate Invalid FieldArg", func(t *testing.T) {
		arg := Arg{
			Type:  FieldArg,
			Field: "InvalidField",
		}

		result, err := arg.Evaluate(obj)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

	t.Run("Evaluate Null Object for FieldArg", func(t *testing.T) {
		arg := Arg{
			Type:  FieldArg,
			Field: "Name",
		}

		result, err := arg.Evaluate(nil)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "object is nil", err.Error())
	})
}

// Test the EvaluateFunctionCall function
func TestEvaluateFunctionCall(t *testing.T) {
	obj := TestStruct{Name: "John", Age: 25}

	t.Run("Evaluate Len Function", func(t *testing.T) {
		function := Function{
			Name: "len",
			Args: []Arg{{Type: FieldArg, Field: "Name"}},
		}

		result, err := EvaluateFunctionCall(function, obj)
		assert.NoError(t, err)
		assert.Equal(t, 4, result)
	})

	t.Run("Evaluate Len Function with Invalid Arg", func(t *testing.T) {
		function := Function{
			Name: "len",
			Args: []Arg{{Type: FieldArg, Field: "InvalidField"}},
		}

		result, err := EvaluateFunctionCall(function, obj)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

// Test the compare function for various operators
func TestCompare(t *testing.T) {
	t.Run("Compare Equal Integers", func(t *testing.T) {
		result, err := compare(5, 5, "==")
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("Compare Not Equal Integers", func(t *testing.T) {
		result, err := compare(5, 10, "!=")
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("Compare Greater Than", func(t *testing.T) {
		result, err := compare(10, 5, ">")
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("Compare Less Than", func(t *testing.T) {
		result, err := compare(5, 10, "<")
		assert.NoError(t, err)
		assert.True(t, result)
	})

	t.Run("Compare Invalid Operator", func(t *testing.T) {
		result, err := compare(5, 5, "invalid")
		assert.Error(t, err)
		assert.False(t, result)
	})

	t.Run("Compare Type Mismatch", func(t *testing.T) {
		result, err := compare(5, "5", "==")
		assert.Error(t, err)
		assert.False(t, result)
	})
}
