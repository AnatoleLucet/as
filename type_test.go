package as

import (
	"testing"
)

func TestType(t *testing.T) {
	t.Run("convert string to int", func(t *testing.T) {
		result, err := Type[int]("123")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != 123 {
			t.Errorf("Expected 123, got %v", result)
		}
	})

	t.Run("convert with custom type", func(t *testing.T) {
		type MyInt int
		result, err := Type[MyInt]("42")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != MyInt(42) {
			t.Errorf("Expected MyInt(42), got %v", result)
		}
	})

	t.Run("convert incompatible types returns error", func(t *testing.T) {
		result, err := Type[int]("not a number")

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != 0 {
			t.Errorf("Expected 0 result, got %v", result)
		}
	})
}
