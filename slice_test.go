package as

import "testing"

func TestSlice(t *testing.T) {
	t.Run("convert []string to []int", func(t *testing.T) {
		input := []string{"1", "2", "3", "4", "5"}
		expected := []int{1, 2, 3, 4, 5}
		result, err := Slice(Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for i := range expected {
			if result[i] != expected[i] {
				t.Errorf("At index %d: expected %d, got %d", i, expected[i], result[i])
			}
		}
	})

	t.Run("convert []any with incompatible types", func(t *testing.T) {
		input := []any{"one", "two", "three"}
		expected := []int{}
		result, err := Slice(Int, input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if len(result) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(result))
		}
	})

	t.Run("input is nil", func(t *testing.T) {
		var input []any = nil
		var expected []int = nil
		result, err := Slice(Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(result) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(result))
		}
	})
}
