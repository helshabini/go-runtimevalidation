package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestISSN(t *testing.T) {
	t.Run("Valid ISSN", func(t *testing.T) {
		err := ISSN("1234-5679")
		assert.NoError(t, err)
	})

	t.Run("Valid ISSN with 'X' checksum", func(t *testing.T) {
		err := ISSN("1234-567X")
		assert.NoError(t, err)
	})

	t.Run("Invalid ISSN - incorrect format", func(t *testing.T) {
		err := ISSN("1234-567")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid issn value: 1234-567")
	})

	t.Run("Invalid ISSN - non-numeric characters", func(t *testing.T) {
		err := ISSN("12A4-567B")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid issn value: 12A4-567B")
	})

	t.Run("Invalid ISSN - empty string", func(t *testing.T) {
		err := ISSN("")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid issn value: ")
	})

	t.Run("Invalid ISSN - wrong type (integer)", func(t *testing.T) {
		err := ISSN(12345679)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})

	t.Run("Invalid ISSN - wrong type (float)", func(t *testing.T) {
		err := ISSN(1234.5679)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got float64")
	})
}
