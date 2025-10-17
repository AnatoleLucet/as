package as

import "testing"

func TestMap(t *testing.T) {
	t.Run("convert map[string]string to map[string]int", func(t *testing.T) {
		input := map[string]string{"one": "1", "two": "2", "three": "3"}
		expected := map[string]int{"one": 1, "two": 2, "three": 3}
		result, err := Map(Self, Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("At key %s: expected %d, got %d", k, v, result[k])
			}
		}
	})

	t.Run("convert map[string]string to map[int]int", func(t *testing.T) {
		input := map[string]string{"1": "10", "2": "20", "3": "30"}
		expected := map[int]int{1: 10, 2: 20, 3: 30}
		result, err := Map(Int, Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("At key %d: expected %d, got %d", k, v, result[k])
			}
		}
	})

	t.Run("convert map[string]int to itself using Self", func(t *testing.T) {
		input := map[string]int{"one": 1, "two": 2, "three": 3}
		expected := map[string]int{"one": 1, "two": 2, "three": 3}
		result, err := Map(Self, Self, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(result) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(result))
		}

		for k, v := range expected {
			if result[k] != v {
				t.Errorf("At key %s: expected %d, got %d", k, v, result[k])
			}
		}
	})

	t.Run("convert map with incompatible key types", func(t *testing.T) {
		input := map[string]string{"one": "1", "two": "2"}
		expected := map[int]int{}
		result, err := Map(Int, Int, input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if len(result) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(result))
		}
	})

	t.Run("convert map with incompatible value types", func(t *testing.T) {
		input := map[string]string{"1": "one", "2": "two"}
		expected := map[int]int{}
		result, err := Map(Int, Int, input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if len(result) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(result))
		}
	})

	t.Run("input is nil", func(t *testing.T) {
		var input map[string]string = nil
		var expected map[string]int = nil
		result, err := Map(Self, Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if len(result) != len(expected) {
			t.Errorf("Expected length %d, got %d", len(expected), len(result))
		}
	})
}
