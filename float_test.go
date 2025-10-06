package as

import (
	"math"
	"strconv"
	"testing"
)

func TestFloat(t *testing.T) {
	tests := []struct {
		input    any
		expected float64
		hasError bool
	}{
		{input: 0.0, expected: 0.0, hasError: false},
		{input: -1.0, expected: -1.0, hasError: false},
		{input: 45.67, expected: 45.67, hasError: false},
		{input: float64(3.14159), expected: 3.14159, hasError: false},
		{input: "123.456", expected: 123.456, hasError: false},
		{input: "-78.9", expected: -78.9, hasError: false},
		{input: "0.001", expected: 0.001, hasError: false},
		{input: 'A', expected: 65.0, hasError: false},
		{input: 123, expected: 123.0, hasError: false},
		{input: int8(-12), expected: -12.0, hasError: false},
		{input: int16(1234), expected: 1234.0, hasError: false},
		{input: int64(-1234567890), expected: -1234567890.0, hasError: false},
		{input: uint(456), expected: 456.0, hasError: false},
		{input: uint8(78), expected: 78.0, hasError: false},
		{input: uint16(910), expected: 910.0, hasError: false},
		{input: uint32(1122), expected: 1122.0, hasError: false},
		{input: uint64(3344556677), expected: 3344556677.0, hasError: false},
		{input: []byte("3.14"), expected: 3.14, hasError: false},
		{input: []byte("-0.001"), expected: -0.001, hasError: false},
		{input: true, expected: 1.0, hasError: false},
		{input: false, expected: 0.0, hasError: false},
		{input: nil, expected: 0.0, hasError: false},
		// error cases
		{input: "invalid", expected: 0, hasError: true},
		{input: struct{}{}, expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Float(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Float(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Float(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestFloat32(t *testing.T) {
	overflow := math.MaxFloat32 * 2
	tests := []struct {
		input    any
		expected float32
		hasError bool
	}{
		{input: overflow, expected: 0, hasError: true},
		{input: strconv.FormatFloat(overflow, 'f', -1, 64), expected: 0, hasError: true},
		{input: []byte(strconv.FormatFloat(overflow, 'f', -1, 64)), expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Float32(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Float32(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Float32(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
