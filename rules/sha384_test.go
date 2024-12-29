package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA384(t *testing.T) {
	t.Run("valid SHA384 hash", func(t *testing.T) {
		err := SHA384("7d5b45282a57f6384a4883a48f0e9a26d9c9f52e458c892a403a0d5e43811a06131d77213c806b21d9dd4fda74f3ac2f")
		assert.NoError(t, err)
	})

	t.Run("invalid SHA384 hash with less than 96 characters", func(t *testing.T) {
		err := SHA384("0d1c3a4714dd9bfaea3b2f3c8edc65a3b2f2a29b564789")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA384 value: 0d1c3a4714dd9bfaea3b2f3c8edc65a3b2f2a29b564789", err.Error())
	})

	t.Run("invalid SHA384 hash with more than 96 characters", func(t *testing.T) {
		err := SHA384("7d5b45282a57f6384a4883a48f0e9a26d9c9f52e458c892a403a0d5e43811a06131d77213c806b21d9dd4fda74f3ac2f0d1c3a4714dd9bfaea3b2f3c8edc65a3b2f2a29b564789c08971a36f4b88cdcbbadd4d6dd4")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA384 value: 7d5b45282a57f6384a4883a48f0e9a26d9c9f52e458c892a403a0d5e43811a06131d77213c806b21d9dd4fda74f3ac2f0d1c3a4714dd9bfaea3b2f3c8edc65a3b2f2a29b564789c08971a36f4b88cdcbbadd4d6dd4", err.Error())
	})

	t.Run("non-hexadecimal characters in hash", func(t *testing.T) {
		err := SHA384("xyzb45282a57f6384a4883a48f0e9a26d9c9f52e458c892a403a0d5e43811a06131d77213c806b21d9dd4fda74f3ac2f")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA384 value: xyzb45282a57f6384a4883a48f0e9a26d9c9f52e458c892a403a0d5e43811a06131d77213c806b21d9dd4fda74f3ac2f", err.Error())
	})

	t.Run("input not a string", func(t *testing.T) {
		err := SHA384(12345)
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got int", err.Error())
	})

	t.Run("empty string", func(t *testing.T) {
		err := SHA384("")
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA384 value: ", err.Error())
	})
}
