package rules

import (
	"go-runtimevalidation/args"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the RequiredIf function
func TestRequiredIf(t *testing.T) {
	obj := struct {
		Name    string
		Age     int
		MaxSize int
	}{
		Name:    "John",
		Age:     25,
		MaxSize: 100,
	}

	t.Run("RequiredIf All Conditions True", func(t *testing.T) {
		args := map[string]args.Arg{
			"NameCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Name"}, Rhs: &args.Arg{Value: "John"}, Operator: "=="},
			},
			"AgeCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Age"}, Rhs: &args.Arg{Value: 25}, Operator: "=="},
			},
		}

		err := RequiredIf("some input", obj, args)
		assert.NoError(t, err)
	})

	t.Run("RequiredIf Some Condition False", func(t *testing.T) {
		args := map[string]args.Arg{
			"NameCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Name"}, Rhs: &args.Arg{Value: "John"}, Operator: "=="},
			},
			"AgeCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Age"}, Rhs: &args.Arg{Value: 30}, Operator: "=="},
			},
		}

		err := RequiredIf("some input", obj, args)
		assert.NoError(t, err) // Required check is skipped since AgeCheck is false
	})

	t.Run("RequiredIf Condition False and Input Required", func(t *testing.T) {
		args := map[string]args.Arg{
			"NameCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Name"}, Rhs: &args.Arg{Value: "John"}, Operator: "=="},
			},
			"AgeCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Age"}, Rhs: &args.Arg{Value: 30}, Operator: "=="},
			},
		}

		err := RequiredIf("", obj, args)
		assert.NoError(t, err) // Required check skipped since a condition is false
	})

	t.Run("RequiredIf All Conditions True and Input Empty", func(t *testing.T) {
		args := map[string]args.Arg{
			"NameCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Name"}, Rhs: &args.Arg{Value: "John"}, Operator: "=="},
			},
			"AgeCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Age"}, Rhs: &args.Arg{Value: 25}, Operator: "=="},
			},
		}

		err := RequiredIf("", obj, args)
		assert.Error(t, err)
		assert.Equal(t, "value is required", err.Error())
	})

	t.Run("RequiredIf Object Nil", func(t *testing.T) {
		args := map[string]args.Arg{
			"NameCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Name"}, Rhs: &args.Arg{Value: "John"}, Operator: "=="},
			},
		}

		err := RequiredIf("some input", nil, args)
		assert.Error(t, err)
		assert.Equal(t, "object is nil", err.Error())
	})

	t.Run("RequiredIf Empty Args", func(t *testing.T) {
		err := RequiredIf("some input", obj, map[string]args.Arg{})
		assert.NoError(t, err)
	})

	t.Run("RequiredIf Invalid Condition Operator", func(t *testing.T) {
		args := map[string]args.Arg{
			"AgeCheck": {
				Type:      args.ConditionArg,
				Condition: args.Condition{Lhs: &args.Arg{Type: args.FieldArg, Field: "Age"}, Rhs: &args.Arg{Value: 25}, Operator: "invalid"},
			},
		}

		err := RequiredIf("some input", obj, args)
		assert.Error(t, err)
		assert.Equal(t, "unknown operator: invalid", err.Error())
	})
}
