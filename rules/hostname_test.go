package rules

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHostname(t *testing.T) {
	t.Run("Valid hostname with letters", func(t *testing.T) {
		err := Hostname("example.com")
		assert.NoError(t, err)
	})

	t.Run("Valid hostname with digits", func(t *testing.T) {
		err := Hostname("123example.com")
		assert.NoError(t, err)
	})

	t.Run("Valid hostname with hyphens", func(t *testing.T) {
		err := Hostname("test-example.com")
		assert.NoError(t, err)
	})

	t.Run("Invalid hostname - starts with hyphen", func(t *testing.T) {
		err := Hostname("-example.com")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid hostname: -example.com")
	})

	t.Run("Invalid hostname - label ends with hyphen", func(t *testing.T) {
		err := Hostname("example-.com")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid hostname: example-.com")
	})

	t.Run("Invalid hostname - ends with hyphen", func(t *testing.T) {
		err := Hostname("example.com-")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid hostname: example.com-")
	})

	t.Run("Invalid hostname - contains invalid characters", func(t *testing.T) {
		err := Hostname("ex@mple.com")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid hostname: ex@mple.com")
	})

	t.Run("Invalid hostname - too long", func(t *testing.T) {
		longHostname := strings.Repeat("a", 64) + ".com"
		err := Hostname(longHostname)
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid hostname: "+longHostname)
	})

	t.Run("Invalid hostname - empty string", func(t *testing.T) {
		err := Hostname("")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid hostname: ")
	})

	t.Run("Invalid hostname - wrong type (integer)", func(t *testing.T) {
		err := Hostname(12345)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
