package rules

import (
	"fmt"
	"testing"
)

func TestSHA512(t *testing.T) {
	t.Run("valid SHA512 string", func(t *testing.T) {
		input := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d70eaf42a945ff77b2b5b635b72bfe9d2a22bcbaac7f8965c5bbac2e0b7a"
		err := SHA512(input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("invalid SHA512 string", func(t *testing.T) {
		input := "invalid-sha512"
		err := SHA512(input)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("non-string input", func(t *testing.T) {
		input := 12345
		err := SHA512(input)
		if err == nil {
			t.Error("expected error, got nil")
		}
		expected := fmt.Sprintf("expected a string, got %T", input)
		if err.Error() != expected {
			t.Errorf("expected %s, got %v", expected, err)
		}
	})
}
