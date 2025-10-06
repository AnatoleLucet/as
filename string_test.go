package as

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []struct {
		input    any
		expected string
		hasError bool
	}{
		{input: "test", expected: "test", hasError: false},
		{input: []byte("test"), expected: "test", hasError: false},
		{input: rune('A'), expected: "A", hasError: false},
		{input: 123, expected: "123", hasError: false},
		{input: int8(-12), expected: "-12", hasError: false},
		{input: int16(1234), expected: "1234", hasError: false},
		{input: int64(-1234567890), expected: "-1234567890", hasError: false},
		{input: uint(456), expected: "456", hasError: false},
		{input: uint8(78), expected: "78", hasError: false},
		{input: uint16(910), expected: "910", hasError: false},
		{input: uint32(1122), expected: "1122", hasError: false},
		{input: uint64(3344556677), expected: "3344556677", hasError: false},
		{input: 45.67, expected: "45.67", hasError: false},
		{input: float32(-0.123), expected: "-0.123", hasError: false},
		{input: float64(3.14159), expected: "3.14159", hasError: false},
		{input: true, expected: "true", hasError: false},
		{input: false, expected: "false", hasError: false},
		{input: nil, expected: "", hasError: false},
		{input: struct{}{}, expected: "{}", hasError: false},
		{input: []int{1, 2, 3}, expected: "[1 2 3]", hasError: false},
	}

	for _, test := range tests {
		result, err := String(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("String(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("String(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
