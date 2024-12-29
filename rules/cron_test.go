package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCRON(t *testing.T) {
	t.Run("Valid CRON expression - 5 fields", func(t *testing.T) {
		err := Cron("0 12 * * MON-FRI")
		assert.NoError(t, err)
	})

	t.Run("Valid CRON expression - 6 fields with year", func(t *testing.T) {
		err := Cron("0 12 * * MON-FRI 2024")
		assert.NoError(t, err)
	})

	t.Run("Invalid CRON - missing a field", func(t *testing.T) {
		err := Cron("0 12 * *")
		assert.Error(t, err)
		assert.Equal(t, "invalid cron: 0 12 * *", err.Error())
	})

	t.Run("Invalid CRON - invalid characters", func(t *testing.T) {
		err := Cron("0 12 * * M0N")
		assert.Error(t, err)
		assert.Equal(t, "invalid cron: 0 12 * * M0N", err.Error())
	})

	t.Run("Valid CRON expression with step values", func(t *testing.T) {
		err := Cron("*/15 0 1,15 * 1-5")
		assert.NoError(t, err)
	})

	t.Run("Invalid CRON - too many fields", func(t *testing.T) {
		err := Cron("*/15 0 1,15 * 1-5 2024 extra")
		assert.Error(t, err)
		assert.Equal(t, "invalid cron: */15 0 1,15 * 1-5 2024 extra", err.Error())
	})
}
