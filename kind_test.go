package as

import (
	"reflect"
	"testing"
)

func TestKind(t *testing.T) {
	t.Run("convert string to int kind", func(t *testing.T) {
		result, err := Kind(reflect.Int, "123")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != 123 {
			t.Errorf("Expected 123, got %v", result)
		}

		if reflect.TypeOf(result).Kind() != reflect.Int {
			t.Errorf("Expected kind Int, got %v", reflect.TypeOf(result).Kind())
		}
	})

	t.Run("convert int to string kind", func(t *testing.T) {
		result, err := Kind(reflect.String, 456)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != "456" {
			t.Errorf("Expected \"456\", got %v", result)
		}

		if reflect.TypeOf(result).Kind() != reflect.String {
			t.Errorf("Expected kind String, got %v", reflect.TypeOf(result).Kind())
		}
	})

	t.Run("convert string to bool kind", func(t *testing.T) {
		result, err := Kind(reflect.Bool, "true")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != true {
			t.Errorf("Expected true, got %v", result)
		}

		if reflect.TypeOf(result).Kind() != reflect.Bool {
			t.Errorf("Expected kind Bool, got %v", reflect.TypeOf(result).Kind())
		}
	})

	t.Run("convert to same kind returns unchanged", func(t *testing.T) {
		result, err := Kind(reflect.Int, 789)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != 789 {
			t.Errorf("Expected 789, got %v", result)
		}
	})

	t.Run("convert to interface kind", func(t *testing.T) {
		result, err := Kind(reflect.Interface, "hello")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != "hello" {
			t.Errorf("Expected \"hello\", got %v", result)
		}
	})

	t.Run("convert incompatible types returns error", func(t *testing.T) {
		result, err := Kind(reflect.Int, "not a number")
		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != 0 {
			t.Errorf("Expected 0 result, got %v", result)
		}
	})

	t.Run("wrap value into slice", func(t *testing.T) {
		result, err := Kind(reflect.Slice, 42)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultSlice, ok := result.([]int)
		if !ok {
			t.Fatalf("Expected []int, got %T", result)
		}

		if len(resultSlice) != 1 {
			t.Fatalf("Expected length 1, got %d", len(resultSlice))
		}

		if resultSlice[0] != 42 {
			t.Errorf("Expected [42], got %v", resultSlice)
		}
	})

	t.Run("slice to slice returns unchanged", func(t *testing.T) {
		input := []int{1, 2, 3}
		result, err := Kind(reflect.Slice, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultSlice, ok := result.([]int)
		if !ok {
			t.Fatalf("Expected []int, got %T", result)
		}

		if len(resultSlice) != len(input) {
			t.Fatalf("Expected length %d, got %d", len(input), len(resultSlice))
		}

		for i := range input {
			if resultSlice[i] != input[i] {
				t.Errorf("At index %d: expected %d, got %d", i, input[i], resultSlice[i])
			}
		}
	})

	t.Run("convert slice delegates to KindSlice", func(t *testing.T) {
		input := []string{"1", "2", "3"}
		result, err := Kind(reflect.Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultSlice, ok := result.([]int)
		if !ok {
			t.Fatalf("Expected []int, got %T", result)
		}

		expected := []int{1, 2, 3}
		if len(resultSlice) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(resultSlice))
		}

		for i := range expected {
			if resultSlice[i] != expected[i] {
				t.Errorf("At index %d: expected %d, got %d", i, expected[i], resultSlice[i])
			}
		}
	})

	t.Run("convert map delegates to KindMap", func(t *testing.T) {
		input := map[string]string{"one": "1", "two": "2"}
		result, err := Kind(reflect.Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultMap, ok := result.(map[string]int)
		if !ok {
			t.Fatalf("Expected map[string]int, got %T", result)
		}

		expected := map[string]int{"one": 1, "two": 2}
		if len(resultMap) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(resultMap))
		}

		for k, v := range expected {
			if resultMap[k] != v {
				t.Errorf("At key %s: expected %d, got %d", k, v, resultMap[k])
			}
		}
	})
}

func TestKindSlice(t *testing.T) {
	t.Run("convert []string to []int kind", func(t *testing.T) {
		input := []string{"1", "2", "3"}
		result, err := KindSlice(reflect.Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultSlice, ok := result.([]int)
		if !ok {
			t.Fatalf("Expected []int, got %T", result)
		}

		expected := []int{1, 2, 3}
		if len(resultSlice) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(resultSlice))
		}

		for i := range expected {
			if resultSlice[i] != expected[i] {
				t.Errorf("At index %d: expected %d, got %d", i, expected[i], resultSlice[i])
			}
		}
	})

	t.Run("convert []int to []string kind", func(t *testing.T) {
		input := []int{10, 20, 30}
		result, err := KindSlice(reflect.String, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultSlice, ok := result.([]string)
		if !ok {
			t.Fatalf("Expected []string, got %T", result)
		}

		expected := []string{"10", "20", "30"}
		if len(resultSlice) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(resultSlice))
		}

		for i := range expected {
			if resultSlice[i] != expected[i] {
				t.Errorf("At index %d: expected %s, got %s", i, expected[i], resultSlice[i])
			}
		}
	})

	t.Run("convert slice to same kind returns unchanged", func(t *testing.T) {
		input := []int{1, 2, 3}
		result, err := KindSlice(reflect.Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultSlice, ok := result.([]int)
		if !ok {
			t.Fatalf("Expected []int, got %T", result)
		}

		if len(resultSlice) != len(input) {
			t.Fatalf("Expected length %d, got %d", len(input), len(resultSlice))
		}

		for i := range input {
			if resultSlice[i] != input[i] {
				t.Errorf("At index %d: expected %d, got %d", i, input[i], resultSlice[i])
			}
		}
	})

	t.Run("convert slice with incompatible types returns error", func(t *testing.T) {
		input := []string{"one", "two", "three"}
		result, err := KindSlice(reflect.Int, input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})

	t.Run("wrap single value into slice", func(t *testing.T) {
		result, err := KindSlice(reflect.String, "hello")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultSlice, ok := result.([]string)
		if !ok {
			t.Fatalf("Expected []string, got %T", result)
		}

		if len(resultSlice) != 1 {
			t.Fatalf("Expected length 1, got %d", len(resultSlice))
		}

		if resultSlice[0] != "hello" {
			t.Errorf("Expected [\"hello\"], got %v", resultSlice)
		}
	})

	t.Run("wrap and convert single value", func(t *testing.T) {
		result, err := KindSlice(reflect.Int, "123")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultSlice, ok := result.([]int)
		if !ok {
			t.Fatalf("Expected []int, got %T", result)
		}

		if len(resultSlice) != 1 {
			t.Fatalf("Expected length 1, got %d", len(resultSlice))
		}

		if resultSlice[0] != 123 {
			t.Errorf("Expected [123], got %v", resultSlice)
		}
	})

	t.Run("wrap with incompatible conversion returns error", func(t *testing.T) {
		result, err := KindSlice(reflect.Int, "not a number")

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}

func TestKindMap(t *testing.T) {
	t.Run("convert map[string]string to map[string]int kind", func(t *testing.T) {
		input := map[string]string{"one": "1", "two": "2", "three": "3"}
		result, err := KindMap(reflect.String, reflect.Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultMap, ok := result.(map[string]int)
		if !ok {
			t.Fatalf("Expected map[string]int, got %T", result)
		}

		expected := map[string]int{"one": 1, "two": 2, "three": 3}
		if len(resultMap) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(resultMap))
		}

		for k, v := range expected {
			if resultMap[k] != v {
				t.Errorf("At key %s: expected %d, got %d", k, v, resultMap[k])
			}
		}
	})

	t.Run("convert map[string]string to map[int]int kind", func(t *testing.T) {
		input := map[string]string{"1": "10", "2": "20", "3": "30"}
		result, err := KindMap(reflect.Int, reflect.Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultMap, ok := result.(map[int]int)
		if !ok {
			t.Fatalf("Expected map[int]int, got %T", result)
		}

		expected := map[int]int{1: 10, 2: 20, 3: 30}
		if len(resultMap) != len(expected) {
			t.Fatalf("Expected length %d, got %d", len(expected), len(resultMap))
		}

		for k, v := range expected {
			if resultMap[k] != v {
				t.Errorf("At key %d: expected %d, got %d", k, v, resultMap[k])
			}
		}
	})

	t.Run("convert map to same kind returns unchanged", func(t *testing.T) {
		input := map[string]int{"one": 1, "two": 2}
		result, err := KindMap(reflect.String, reflect.Int, input)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		resultMap, ok := result.(map[string]int)
		if !ok {
			t.Fatalf("Expected map[string]int, got %T", result)
		}

		if len(resultMap) != len(input) {
			t.Fatalf("Expected length %d, got %d", len(input), len(resultMap))
		}

		for k, v := range input {
			if resultMap[k] != v {
				t.Errorf("At key %s: expected %d, got %d", k, v, resultMap[k])
			}
		}
	})

	t.Run("convert non-map type returns error", func(t *testing.T) {
		input := "not a map"
		result, err := KindMap(reflect.String, reflect.Int, input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})

	t.Run("convert map with incompatible key types returns error", func(t *testing.T) {
		input := map[string]string{"one": "1", "two": "2"}
		result, err := KindMap(reflect.Int, reflect.Int, input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})

	t.Run("convert map with incompatible value types returns error", func(t *testing.T) {
		input := map[string]string{"1": "one", "2": "two"}
		result, err := KindMap(reflect.Int, reflect.Int, input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}
