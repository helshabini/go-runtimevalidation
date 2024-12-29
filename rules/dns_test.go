package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDNS(t *testing.T) {
	t.Run("Valid DNS label", func(t *testing.T) {
		err := DNS("example")
		assert.NoError(t, err)
	})

	t.Run("Valid DNS with hyphens", func(t *testing.T) {
		err := DNS("my-domain")
		assert.NoError(t, err)
	})

	t.Run("Valid DNS with numbers", func(t *testing.T) {
		err := DNS("123domain")
		assert.NoError(t, err)
	})

	t.Run("Invalid DNS - starts with hyphen", func(t *testing.T) {
		err := DNS("-example")
		assert.Error(t, err)
		assert.Equal(t, "invalid dns: -example", err.Error())
	})

	t.Run("Invalid DNS - ends with hyphen", func(t *testing.T) {
		err := DNS("example-")
		assert.Error(t, err)
		assert.Equal(t, "invalid dns: example-", err.Error())
	})

	t.Run("Invalid DNS - too long", func(t *testing.T) {
		err := DNS("thisisaverylongdnsnamethatexceedstheallowablelimitof63characters")
		assert.Error(t, err)
		assert.Equal(t, "invalid dns: thisisaverylongdnsnamethatexceedstheallowablelimitof63characters", err.Error())
	})

	t.Run("Invalid DNS - contains spaces", func(t *testing.T) {
		err := DNS("example domain")
		assert.Error(t, err)
		assert.Equal(t, "invalid dns: example domain", err.Error())
	})

	t.Run("Invalid DNS - only hyphens", func(t *testing.T) {
		err := DNS("---")
		assert.Error(t, err)
		assert.Equal(t, "invalid dns: ---", err.Error())
	})
}
