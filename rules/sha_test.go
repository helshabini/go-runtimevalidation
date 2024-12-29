package rules

import (
	"fmt"
	"testing"
)

func TestSHA(t *testing.T) {
	t.Run("valid SHA string (SHA-1)", func(t *testing.T) {
		input := "da39a3ee5e6b4b0d3255bfef95601890afd80709" // Valid SHA-1
		err := SHA(input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("valid SHA string (SHA-256)", func(t *testing.T) {
		input := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" // Valid SHA-256
		err := SHA(input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("valid SHA string (SHA-512)", func(t *testing.T) {
		input := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d70eaf42a945ff77b2b5b635b72bfe9d2a22bcbaac7f8965c5bbac2e0b7a" // Valid SHA-512
		err := SHA(input)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
	})

	t.Run("invalid SHA string", func(t *testing.T) {
		input := "invalid-sha-hash"
		err := SHA(input)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("non-string input", func(t *testing.T) {
		input := 12345
		err := SHA(input)
		if err == nil {
			t.Error("expected error, got nil")
		}
		expected := fmt.Sprintf("expected a string, got %T", input)
		if err.Error() != expected {
			t.Errorf("expected %s, got %v", expected, err)
		}
	})
}
