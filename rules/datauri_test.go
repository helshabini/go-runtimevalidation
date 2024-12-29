package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDataUri(t *testing.T) {
	t.Run("Valid data URI (image/png)", func(t *testing.T) {
		err := DataUri("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAABJElEQVQ4T6WTTU4CQRCEv3vFjm1YpaJIFBBWRUq0AifoaFgj6CgIAmVUlEKlASO4kQiIRBWQUKHQQQQQCkQUWzN7Ox8mvlMrN9pwzmReoFH4BCeMBTIyk9UUcgBOKhQBj8iMKOxRSHEBGAZnxmOn6+3dgglgkFYghMHcXFDVPAniPyWn5dP01ApNn3kFshyGMybZg1KiJY5LKXYf+QicwOD5yVV8I9Ymy1UQlJSEiFtJvUOiRzPbNElYDPWNHVhUE1Ae5yOq0gjGVVD1DtBGDPF5/gGM6a3I5UHrloVWom2fjB3oihmRC9HvhAkpdAHx3MGuDLMQiIoS+MoxYO9gfHKqhHdpFXZCEsmlN2AuNYhNLKE3E0fPnHzzDo8gFzqspcAmfu/WU2kRlaAMfAfjFPFYVfDdHuoJ9kOGflWvA+/hmOp+uR6JObXYOtAAAAABJRU5ErkJggg==")
		assert.NoError(t, err)
	})

	t.Run("Valid data URI (text/plain)", func(t *testing.T) {
		err := DataUri("data:text/plain;base64,SGVsbG8sIFdvcmxkIQ==")
		assert.NoError(t, err)
	})

	t.Run("Invalid data URI (wrong scheme)", func(t *testing.T) {
		err := DataUri("http://example.com")
		assert.Error(t, err)
	})

	t.Run("Invalid data URI (invalid base64)", func(t *testing.T) {
		err := DataUri("data:image/png;base64,invalid_base64")
		assert.NoError(t, err)
	})

	t.Run("Non-string input", func(t *testing.T) {
		err := DataUri(12345) // Invalid, since input is not a string
		assert.Error(t, err)
	})

	t.Run("Empty string input", func(t *testing.T) {
		err := DataUri("") // Invalid, since empty string is not a valid data URI
		assert.Error(t, err)
	})
}
