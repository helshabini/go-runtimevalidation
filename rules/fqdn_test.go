package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFQDN(t *testing.T) {
	t.Run("Valid FQDN with non-numeric TLD", func(t *testing.T) {
		err := FQDN("example.com")
		assert.NoError(t, err)
	})

	t.Run("Valid FQDN with dot at the end", func(t *testing.T) {
		err := FQDN("example.com.")
		assert.NoError(t, err)
	})

	t.Run("Valid FQDN with subdomain", func(t *testing.T) {
		err := FQDN("sub.example.com")
		assert.NoError(t, err)
	})

	t.Run("Invalid FQDN - numeric TLD", func(t *testing.T) {
		err := FQDN("example.123")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid fqdn: example.123")
	})

	t.Run("Invalid FQDN - starts with hyphen", func(t *testing.T) {
		err := FQDN("-example.com")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid fqdn: -example.com")
	})

	t.Run("Invalid FQDN - ends with hyphen", func(t *testing.T) {
		err := FQDN("example-.com")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid fqdn: example-.com")
	})

	t.Run("Invalid FQDN - contains invalid characters", func(t *testing.T) {
		err := FQDN("ex@mple.com")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid fqdn: ex@mple.com")
	})

	t.Run("Invalid FQDN - empty string", func(t *testing.T) {
		err := FQDN("")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid fqdn: ")
	})

	t.Run("Invalid FQDN - wrong type (integer)", func(t *testing.T) {
		err := FQDN(12345)
		assert.Error(t, err)
		assert.EqualError(t, err, "expected a string, got int")
	})
}
