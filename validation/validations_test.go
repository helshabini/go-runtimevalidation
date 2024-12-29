package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("Empty Rule", func(t *testing.T) {
		ruletext := ""
		rules, err := Parse(ruletext)

		assert.Error(t, err, "Expected error when parsing an empty rule")
		assert.NotNil(t, rules, "Expected non-nil rules for empty input")
		assert.Len(t, rules, 1, "Expected 1 group for empty input")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for empty input")
		assert.NotNil(t, rules[0][0].Error, "Expected parsing error for empty input, got none")
	})

	t.Run("Malformed Rule (Misspelled)", func(t *testing.T) {
		ruletext := "requirred"
		rules, err := Parse(ruletext)

		assert.Error(t, err, "Expected error when parsing malformed rule (misspelled)")
		assert.NotNil(t, rules, "Expected non-nil rules for misspelled input")
		assert.Len(t, rules, 1, "Expected 1 group for misspelled input")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for misspelled input")
		assert.NotNil(t, rules[0][0].Error, "Expected parsing error for misspelled input, got none")
	})

	t.Run("Rule with Leading Spaces", func(t *testing.T) {
		ruletext := "   required"
		rules, err := Parse(ruletext)

		assert.NoError(t, err, "Expected no error when parsing rule with leading spaces")
		assert.NotNil(t, rules, "Expected non-nil rules for input with leading spaces")
		assert.Len(t, rules, 1, "Expected 1 group for input with leading spaces")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for input with leading spaces")
		assert.Nil(t, rules[0][0].Error, "Expected no parsing error for input with leading spaces")
	})

	t.Run("Rule with Trailing Spaces", func(t *testing.T) {
		ruletext := "required   "
		rules, err := Parse(ruletext)

		assert.NoError(t, err, "Expected no error when parsing rule with trailing spaces")
		assert.NotNil(t, rules, "Expected non-nil rules for input with trailing spaces")
		assert.Len(t, rules, 1, "Expected 1 group for input with trailing spaces")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for input with trailing spaces")
		assert.Nil(t, rules[0][0].Error, "Expected no parsing error for input with trailing spaces")
	})

	t.Run("Rule with Extra Spaces", func(t *testing.T) {
		ruletext := "required &&   "
		rules, err := Parse(ruletext)

		assert.Error(t, err, "Expected error when parsing rule with extra spaces")
		assert.NotNil(t, rules, "Expected non-nil rules for input with extra spaces")
		assert.Len(t, rules, 2, "Expected 2 groups for input with extra spaces")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for input with extra spaces")
		assert.Len(t, rules[1], 1, "Expected 1 rule in group 1 for input with extra spaces")
		assert.Nil(t, rules[0][0].Error, "Expected no parsing error for group 0")
		assert.NotNil(t, rules[1][0].Error, "Expected parsing error for group 1")
	})

	t.Run("Multiple Rules with One Malformed", func(t *testing.T) {
		ruletext := "required && invalid_rule"
		rules, err := Parse(ruletext)

		assert.Error(t, err, "Expected error when parsing multiple rules with one malformed")
		assert.NotNil(t, rules, "Expected non-nil rules for input with one malformed rule")
		assert.Len(t, rules, 2, "Expected 2 groups for input with one malformed rule")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for input with one malformed rule")
		assert.Len(t, rules[1], 1, "Expected 1 rule in group 1 for input with one malformed rule")
		assert.Nil(t, rules[0][0].Error, "Expected no parsing error for group 0")
		assert.NotNil(t, rules[1][0].Error, "Expected parsing error for group 1")
	})

	t.Run("Multiple Rules with Two Malformed", func(t *testing.T) {
		ruletext := "invalid_rule && required && invalid_rule"
		rules, err := Parse(ruletext)

		assert.Error(t, err, "Expected error when parsing multiple rules with two malformed")
		assert.NotNil(t, rules, "Expected non-nil rules for input with two malformed rules")
		assert.Len(t, rules, 3, "Expected 3 groups for input with two malformed rules")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for input with two malformed rules")
		assert.Len(t, rules[1], 1, "Expected 1 rule in group 1 for input with two malformed rules")
		assert.Len(t, rules[2], 1, "Expected 1 rule in group 2 for input with two malformed rules")
		assert.NotNil(t, rules[0][0].Error, "Expected parsing error for group 0")
		assert.Nil(t, rules[1][0].Error, "Expected no parsing error for group 1")
		assert.NotNil(t, rules[2][0].Error, "Expected parsing error for group 2")
	})

	t.Run("Valid Rule with Malformed OR Input", func(t *testing.T) {
		ruletext := "required && ||"
		rules, err := Parse(ruletext)

		assert.Error(t, err, "Expected error when parsing rule with malformed OR input")
		assert.NotNil(t, rules, "Expected non-nil rules for input with malformed OR")
		assert.Len(t, rules, 2, "Expected 2 groups for input with malformed OR")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for input with malformed OR")
		assert.Len(t, rules[1], 2, "Expected 2 rules in group 1 for input with malformed OR")
		assert.Nil(t, rules[0][0].Error, "Expected no parsing error for group 0")
		assert.NotNil(t, rules[1][0].Error, "Expected parsing error for first rule in group 1")
		assert.NotNil(t, rules[1][1].Error, "Expected parsing error for second rule in group 1")
	})

	t.Run("Valid Required Rule", func(t *testing.T) {
		ruletext := "required"
		rules, err := Parse(ruletext)

		assert.NoError(t, err, "Expected no error parsing valid rule")
		assert.NotNil(t, rules, "Expected non-nil rules for valid input")
		assert.Len(t, rules, 1, "Expected 1 group for valid input")
		assert.Len(t, rules[0], 1, "Expected 1 rule in group 0 for valid input")
		assert.Nil(t, rules[0][0].Error, "Expected no parsing error for valid input")
	})

	t.Run("Complex Rule with Multiple Groups", func(t *testing.T) {
		// Test input string containing validation rules
		rulesString := "required||requiredif:$Name=John,$len($Name)>3&&between:18,$len($MaxSize)||min:18||max:30||oneof:usa,egypt,canada||length:10||endswith:.com||email"

		// Call Parse function
		groups, err := Parse(rulesString)

		// Ensure no error is returned
		assert.NoError(t, err)

		// Ensure correct number of rule groups are parsed
		actualGroups := len(groups)
		assert.Equal(t, 2, actualGroups, "Unexpected number of groups parsed")

		// Validate the number of rules in each group
		actualNumRules1 := len(groups[0])
		actualNumRules2 := len(groups[1])
		assert.Equal(t, 2, actualNumRules1, "Unexpected number of rules parsed in group 1")
		assert.Equal(t, 7, actualNumRules2, "Unexpected number of rules parsed in group 2")

		// Group 1 rules validation
		assert.Equal(t, "required", groups[0][0].Tag)
		assert.NotNil(t, groups[0][0].Validate)

		// Validate "requiredif" rule and its condition
		conditionRule := groups[0][1]
		assert.Equal(t, "requiredif", conditionRule.Tag)
		assert.Equal(t, "requiredif:$Name=John,$len($Name)>3", conditionRule.Text)
		assert.NotNil(t, conditionRule.Validate)

		// Group 2 rules validation

		// Validate "between" rule with its arguments
		betweenRule := groups[1][0]
		assert.Equal(t, "between", betweenRule.Tag)
		assert.Equal(t, "between:18,$len($MaxSize)", betweenRule.Text)
		assert.NotNil(t, betweenRule.Validate)

		// Validate "min" rule with its argument
		minRule := groups[1][1]
		assert.Equal(t, "min", minRule.Tag)
		assert.Equal(t, "min:18", minRule.Text)
		assert.NotNil(t, minRule.Validate)

		// Validate "max" rule with its argument
		maxRule := groups[1][2]
		assert.Equal(t, "max", maxRule.Tag)
		assert.Equal(t, "max:30", maxRule.Text)
		assert.NotNil(t, maxRule.Validate)

		// Validate "oneof" rule with its arguments
		oneofRule := groups[1][3]
		assert.Equal(t, "oneof", oneofRule.Tag)
		assert.Equal(t, "oneof:usa,egypt,canada", oneofRule.Text)
		assert.NotNil(t, oneofRule.Validate)

		// Validate "length" rule with its argument
		lengthRule := groups[1][4]
		assert.Equal(t, "length", lengthRule.Tag)
		assert.Equal(t, "length:10", lengthRule.Text)
		assert.NotNil(t, lengthRule.Validate)

		// Validate "endswith" rule with its argument
		endswithRule := groups[1][5]
		assert.Equal(t, "endswith", endswithRule.Tag)
		assert.Equal(t, "endswith:.com", endswithRule.Text)
		assert.NotNil(t, endswithRule.Validate)
	})

}
