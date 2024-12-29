package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSemVer(t *testing.T) {
	t.Run("Valid SemVer (major.minor.patch)", func(t *testing.T) {
		err := SemVer("1.0.0")
		assert.NoError(t, err)
	})

	t.Run("Valid SemVer with pre-release", func(t *testing.T) {
		err := SemVer("1.0.0-alpha.1")
		assert.NoError(t, err)
	})

	t.Run("Valid SemVer with build metadata", func(t *testing.T) {
		err := SemVer("1.0.0+build.001")
		assert.NoError(t, err)
	})

	t.Run("Valid SemVer with pre-release and build metadata", func(t *testing.T) {
		err := SemVer("1.0.0-alpha.1+build.001")
		assert.NoError(t, err)
	})

	t.Run("Invalid SemVer (missing patch version)", func(t *testing.T) {
		err := SemVer("1.0")
		assert.Error(t, err)
		assert.Equal(t, "invalid semver: 1.0", err.Error())
	})

	t.Run("Invalid SemVer (letters in patch)", func(t *testing.T) {
		err := SemVer("1.0.a")
		assert.Error(t, err)
		assert.Equal(t, "invalid semver: 1.0.a", err.Error())
	})

	t.Run("Invalid SemVer (negative version numbers)", func(t *testing.T) {
		err := SemVer("-1.0.0")
		assert.Error(t, err)
		assert.Equal(t, "invalid semver: -1.0.0", err.Error())
	})
}
