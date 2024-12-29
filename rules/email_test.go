package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test cases for the Email function
func TestEmail(t *testing.T) {
	t.Run("Valid email", func(t *testing.T) {
		err := Email("test@example.com")
		assert.NoError(t, err)
	})

	t.Run("Valid email with '+' and domain suffix", func(t *testing.T) {
		err := Email("user+label@domain.co.uk")
		assert.NoError(t, err)
	})

	t.Run("Valid email with subdomain", func(t *testing.T) {
		err := Email("user.name@sub.domain.com")
		assert.NoError(t, err)
	})

	t.Run("Valid email with numeric local part", func(t *testing.T) {
		err := Email("12345@example.io")
		assert.NoError(t, err)
	})

	t.Run("Missing '@' and domain", func(t *testing.T) {
		err := Email("invalid-email")
		assert.EqualError(t, err, "invalid email: invalid-email")
	})

	t.Run("Missing local part", func(t *testing.T) {
		err := Email("@missinglocal.com")
		assert.EqualError(t, err, "invalid email: @missinglocal.com")
	})

	t.Run("Invalid domain part", func(t *testing.T) {
		err := Email("missingdomain@.com")
		assert.EqualError(t, err, "invalid email: missingdomain@.com")
	})

	t.Run("Non-string input", func(t *testing.T) {
		err := Email(12345)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("Invalid multiple '@'", func(t *testing.T) {
		err := Email("user@domain@extra.com")
		assert.EqualError(t, err, "invalid email: user@domain@extra.com")
	})

	t.Run("Invalid domain with consecutive dots", func(t *testing.T) {
		err := Email("user@domain..com")
		assert.EqualError(t, err, "invalid email: user@domain..com")
	})
}
