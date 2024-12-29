package rules

import (
	"testing"

	"go-runtimevalidation/args"

	"github.com/stretchr/testify/assert"
)

func TestXBetween(t *testing.T) {
	// Test case 1: Input is exclusively between two constant arguments
	t.Run("Input exclusively between constant arguments", func(t *testing.T) {
		err := XBetween(15, nil, map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		})
		assert.NoError(t, err)
	})

	// Test case 2: Input is equal to the lower bound
	t.Run("Input equals lower bound", func(t *testing.T) {
		err := XBetween(10, nil, map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "exclusive between validation failed: 10 is not exclusively between")
	})

	// Test case 3: Input is equal to the upper bound
	t.Run("Input equals upper bound", func(t *testing.T) {
		err := XBetween(20, nil, map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "exclusive between validation failed: 20 is not exclusively between")
	})

	// Test case 4: Input is less than the lower bound
	t.Run("Input less than lower bound", func(t *testing.T) {
		err := XBetween(5, nil, map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "exclusive between validation failed: 5 is not exclusively between")
	})

	// Test case 5: Input is greater than the upper bound
	t.Run("Input greater than upper bound", func(t *testing.T) {
		err := XBetween(25, nil, map[string]args.Arg{
			"lowerBound": {Value: 10},
			"upperBound": {Value: 20},
		})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "exclusive between validation failed: 25 is not exclusively between")
	})

	// Test case 6: Field reference resolves to integers
	t.Run("Field reference resolves to integers", func(t *testing.T) {
		obj := struct {
			Lower int
			Upper int
		}{Lower: 10, Upper: 20}

		err := XBetween(15, obj, map[string]args.Arg{
			"lowerBound": {Type: args.FieldArg, Field: "Lower"},
			"upperBound": {Type: args.FieldArg, Field: "Upper"},
		})
		assert.NoError(t, err)
	})

	// Test case 7: One argument resolves to a string
	t.Run("Field reference resolves to string", func(t *testing.T) {
		obj := struct {
			StrField string
		}{StrField: "test"}

		err := XBetween(15, obj, map[string]args.Arg{
			"lowerBound": {Type: args.FieldArg, Field: "StrField"},
			"upperBound": {Value: 20},
		})
		assert.Error(t, err)
		assert.ErrorContains(t, err, "unsupported type for")
	})
}
