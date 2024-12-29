package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE164(t *testing.T) {
	t.Run("Valid E.164 phone number", func(t *testing.T) {
		err := E164("+1234567890123")
		assert.Nil(t, err, "Expected no error for valid E.164 phone number")
	})

	t.Run("Valid E.164 phone number with 7 digits", func(t *testing.T) {
		err := E164("+1234567")
		assert.Nil(t, err, "Expected no error for valid E.164 phone number with 7 digits")
	})

	t.Run("Valid E.164 phone number with 14 digits", func(t *testing.T) {
		err := E164("+12345678901234")
		assert.Nil(t, err, "Expected no error for valid E.164 phone number with 14 digits")
	})

	t.Run("Invalid E.164 phone number without leading +", func(t *testing.T) {
		err := E164("1234567890123")
		assert.NotNil(t, err, "Expected error for missing '+' in E.164 phone number")
		assert.EqualError(t, err, "invalid E.164 phone number: 1234567890123")
	})

	t.Run("Invalid E.164 phone number with special characters", func(t *testing.T) {
		err := E164("+1234-5678")
		assert.NotNil(t, err, "Expected error for E.164 phone number with special characters")
		assert.EqualError(t, err, "invalid E.164 phone number: +1234-5678")
	})

	t.Run("Invalid E.164 phone number with too many digits", func(t *testing.T) {
		err := E164("+123456789012345")
		assert.Nil(t, err, "Expected no error for valid E.164 phone number with 15 digits")
	})

	t.Run("Invalid E.164 phone number with too few digits", func(t *testing.T) {
		err := E164("+123456")
		assert.NotNil(t, err, "Expected error for E.164 phone number with too few digits")
		assert.EqualError(t, err, "invalid E.164 phone number: +123456")
	})
}
