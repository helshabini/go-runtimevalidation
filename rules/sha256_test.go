package rules

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSHA256(t *testing.T) {
	t.Run("ValidSHA256Hash", func(t *testing.T) {
		// Valid SHA256 hash string (64 characters long)
		input := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
		err := SHA256(input)
		assert.NoError(t, err)
	})

	t.Run("InvalidSHA256Hash_ShortLength", func(t *testing.T) {
		// SHA256 hash too short (less than 64 characters)
		input := "e3b0c44298fc1c149afbf4c8996fb924"
		err := SHA256(input)
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA256 value: e3b0c44298fc1c149afbf4c8996fb924", err.Error())
	})

	t.Run("InvalidSHA256Hash_TooLong", func(t *testing.T) {
		// SHA256 hash too long (more than 64 characters)
		input := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855e3b0"
		err := SHA256(input)
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA256 value: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855e3b0", err.Error())
	})

	t.Run("InvalidSHA256Hash_NonHexCharacters", func(t *testing.T) {
		// SHA256 hash with non-hexadecimal characters
		input := "zzzzc44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
		err := SHA256(input)
		assert.Error(t, err)
		assert.Equal(t, "invalid SHA256 value: zzzzc44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", err.Error())
	})

	t.Run("InvalidInput_NotAString", func(t *testing.T) {
		// Input is not a string
		input := 123456
		err := SHA256(input)
		assert.Error(t, err)
		assert.Equal(t, "expected a string, got int", err.Error())
	})
}
