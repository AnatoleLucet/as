package as

import (
	"reflect"
	"testing"
)

func TestT(t *testing.T) {
	t.Run("convert string to int", func(t *testing.T) {
		result, err := T[int]("123")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != 123 {
			t.Errorf("Expected 123, got %v", result)
		}
	})

	t.Run("convert with custom type", func(t *testing.T) {
		type MyInt int
		result, err := T[MyInt]("42")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != MyInt(42) {
			t.Errorf("Expected MyInt(42), got %v", result)
		}
	})

	t.Run("convert incompatible types returns error", func(t *testing.T) {
		result, err := T[int]("not a number")

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != 0 {
			t.Errorf("Expected 0 result, got %v", result)
		}
	})
}

func TestType(t *testing.T) {
	t.Run("convert string to int type", func(t *testing.T) {
		result, err := Type(reflect.TypeOf(int(0)), "123")
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

	t.Run("convert int to string type", func(t *testing.T) {
		result, err := Type(reflect.TypeOf(""), 456)
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

	t.Run("convert to same type returns unchanged", func(t *testing.T) {
		result, err := Type(reflect.TypeOf(int(0)), 789)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if result != 789 {
			t.Errorf("Expected 789, got %v", result)
		}
	})

	t.Run("convert incompatible types returns error", func(t *testing.T) {
		result, err := Type(reflect.TypeOf(int(0)), "not a number")
		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != 0 {
			t.Errorf("Expected 0 result, got %v", result)
		}
	})

	t.Run("convert slice delegates to TypeSlice", func(t *testing.T) {
		input := []string{"1", "2", "3"}
		result, err := Type(reflect.TypeOf([]int{0}), input)
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

	t.Run("convert map delegates to TypeMap", func(t *testing.T) {
		input := map[string]string{"one": "1", "two": "2"}
		result, err := Type(reflect.TypeOf(map[string]int{}), input)
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

func TestTypeSlice(t *testing.T) {
	t.Run("convert []string to []int", func(t *testing.T) {
		input := []string{"1", "2", "3"}
		result, err := TypeSlice(reflect.TypeOf(int(0)), input)
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

	t.Run("convert slice to same type returns unchanged", func(t *testing.T) {
		input := []int{1, 2, 3}
		result, err := TypeSlice(reflect.TypeOf(int(0)), input)
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

	t.Run("wrap single value into slice", func(t *testing.T) {
		result, err := TypeSlice(reflect.TypeOf(""), "hello")
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
		result, err := TypeSlice(reflect.TypeOf(int(0)), "123")
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

	t.Run("convert slice with incompatible types returns error", func(t *testing.T) {
		input := []string{"one", "two", "three"}
		result, err := TypeSlice(reflect.TypeOf(int(0)), input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}

func TestTypeMap(t *testing.T) {
	t.Run("convert map[string]string to map[string]int", func(t *testing.T) {
		input := map[string]string{"one": "1", "two": "2", "three": "3"}
		result, err := TypeMap(reflect.TypeOf(""), reflect.TypeOf(int(0)), input)
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

	t.Run("convert map[string]string to map[int]int", func(t *testing.T) {
		input := map[string]string{"1": "10", "2": "20", "3": "30"}
		result, err := TypeMap(reflect.TypeOf(int(0)), reflect.TypeOf(int(0)), input)
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

	t.Run("convert map to same type returns unchanged", func(t *testing.T) {
		input := map[string]int{"one": 1, "two": 2}
		result, err := TypeMap(reflect.TypeOf(""), reflect.TypeOf(int(0)), input)
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
		result, err := TypeMap(reflect.TypeOf(""), reflect.TypeOf(int(0)), input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})

	t.Run("convert map with incompatible value types returns error", func(t *testing.T) {
		input := map[string]string{"1": "one", "2": "two"}
		result, err := TypeMap(reflect.TypeOf(int(0)), reflect.TypeOf(int(0)), input)

		if err == nil {
			t.Fatalf("Expected error but got none")
		}

		if result != nil {
			t.Errorf("Expected nil result, got %v", result)
		}
	})
}
