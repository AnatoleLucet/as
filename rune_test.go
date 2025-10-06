package as

import (
	"math"
	"testing"
)

func TestRune(t *testing.T) {
	tests := []struct {
		input    any
		expected rune
		hasError bool
	}{
		{input: 'A', expected: 'A', hasError: false},
		{input: rune('Z'), expected: 'Z', hasError: false},
		{input: int32('B'), expected: 'B', hasError: false},
		{input: int64('C'), expected: 'C', hasError: false},
		{input: uint32('D'), expected: 'D', hasError: false},
		{input: uint64('E'), expected: 'E', hasError: false},
		{input: 70, expected: 'F', hasError: false},
		{input: 71.0, expected: 'G', hasError: false},
		{input: "H", expected: 'H', hasError: false},
		{input: "long", expected: 'l', hasError: false},
		{input: []byte("I"), expected: 'I', hasError: false},
		// error cases
		{input: "", expected: 0, hasError: true},
		{input: math.MaxInt32 + 1, expected: 0, hasError: true},
		{input: struct{}{}, expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Rune(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Rune(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Rune(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
