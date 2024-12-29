package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCVE(t *testing.T) {
	t.Run("Valid CVE format", func(t *testing.T) {
		err := CVE("CVE-2021-34527")
		assert.NoError(t, err)
	})

	t.Run("Valid CVE with more than 4 digits in the ID", func(t *testing.T) {
		err := CVE("CVE-2017-123456")
		assert.NoError(t, err)
	})

	t.Run("Invalid CVE - missing prefix", func(t *testing.T) {
		err := CVE("2021-34527")
		assert.Error(t, err)
		assert.Equal(t, "invalid cve: 2021-34527", err.Error())
	})

	t.Run("Invalid CVE - invalid year", func(t *testing.T) {
		err := CVE("CVE-21-34527")
		assert.Error(t, err)
		assert.Equal(t, "invalid cve: CVE-21-34527", err.Error())
	})

	t.Run("Invalid CVE - insufficient digits in ID", func(t *testing.T) {
		err := CVE("CVE-2021-123")
		assert.Error(t, err)
		assert.Equal(t, "invalid cve: CVE-2021-123", err.Error())
	})

	t.Run("Invalid CVE - non-numeric characters in ID", func(t *testing.T) {
		err := CVE("CVE-2021-ABCD")
		assert.Error(t, err)
		assert.Equal(t, "invalid cve: CVE-2021-ABCD", err.Error())
	})

	t.Run("Invalid CVE - random string", func(t *testing.T) {
		err := CVE("invalid-string")
		assert.Error(t, err)
		assert.Equal(t, "invalid cve: invalid-string", err.Error())
	})
}
