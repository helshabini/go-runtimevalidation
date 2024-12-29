package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUrlEncoded(t *testing.T) {
	t.Run("Valid URL-encoded string", func(t *testing.T) {
		err := UrlEncoded("hello%20world")
		assert.NoError(t, err)
	})

	t.Run("Valid URL-encoded string with special characters", func(t *testing.T) {
		err := UrlEncoded("foo%3Dbar%26baz")
		assert.NoError(t, err)
	})

	t.Run("Valid URL-encoded with numbers and letters", func(t *testing.T) {
		err := UrlEncoded("a%2Fb%3F123")
		assert.NoError(t, err)
	})

	t.Run("Invalid URL-encoded string - incomplete encoding", func(t *testing.T) {
		err := UrlEncoded("hello%2")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid url encoded value: hello%2")
	})

	t.Run("Invalid URL-encoded string - non-encoded special characters", func(t *testing.T) {
		err := UrlEncoded("hello world")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid url encoded value: hello world")
	})

	t.Run("Invalid URL-encoded string - empty string", func(t *testing.T) {
		err := UrlEncoded("")
		assert.Error(t, err)
		assert.EqualError(t, err, "invalid url encoded value: ")
	})

	t.Run("All numbers", func(t *testing.T) {
		err := UrlEncoded(12345)
		assert.NoError(t, err)
	})
}
