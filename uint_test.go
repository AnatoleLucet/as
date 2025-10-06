package as

import (
	"math"
	"strconv"
	"testing"
)

func TestUint(t *testing.T) {
	tests := []struct {
		input    any
		expected uint
		hasError bool
	}{
		{input: 0, expected: 0, hasError: false},
		{input: -0, expected: 0, hasError: false},
		{input: uint(123), expected: 123, hasError: false},
		{input: uint8(45), expected: 45, hasError: false},
		{input: uint16(6789), expected: 6789, hasError: false},
		{input: uint32(123456), expected: 123456, hasError: false},
		{input: uint64(9876543210), expected: 9876543210, hasError: false},
		{input: int(123), expected: 123, hasError: false},
		{input: int8(45), expected: 45, hasError: false},
		{input: int16(6789), expected: 6789, hasError: false},
		{input: int32(123456), expected: 123456, hasError: false},
		{input: int64(9876543210), expected: 9876543210, hasError: false},
		{input: float32(123.0), expected: 123, hasError: false},
		{input: float64(456.0), expected: 456, hasError: false},
		{input: "789", expected: 789, hasError: false},
		{input: []byte("101112"), expected: 101112, hasError: false},
		{input: rune('A'), expected: 65, hasError: false},
		{input: true, expected: 1, hasError: false},
		{input: false, expected: 0, hasError: false},
		{input: nil, expected: 0, hasError: false},
		// error cases
		{input: -1, expected: 0, hasError: true},
		{input: int8(-45), expected: 0, hasError: true},
		{input: int16(-6789), expected: 0, hasError: true},
		{input: int32(-123456), expected: 0, hasError: true},
		{input: int64(-9876543210), expected: 0, hasError: true},
		{input: float32(-123.0), expected: 0, hasError: true},
		{input: float64(-456.0), expected: 0, hasError: true},
		{input: "invalid", expected: 0, hasError: true},
		{input: struct{}{}, expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Uint(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Uint(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Uint(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestUint8(t *testing.T) {
	overflow := math.MaxUint8 + 1
	tests := []struct {
		input    any
		expected uint8
		hasError bool
	}{
		{input: uint(overflow), expected: 0, hasError: true},
		{input: uint16(overflow), expected: 0, hasError: true},
		{input: uint32(overflow), expected: 0, hasError: true},
		{input: uint64(overflow), expected: 0, hasError: true},
		{input: int(overflow), expected: 0, hasError: true},
		{input: int16(overflow), expected: 0, hasError: true},
		{input: int32(overflow), expected: 0, hasError: true},
		{input: int64(overflow), expected: 0, hasError: true},
		{input: float32(overflow), expected: 0, hasError: true},
		{input: float64(overflow), expected: 0, hasError: true},
		{input: strconv.Itoa(overflow), expected: 0, hasError: true},
		{input: []byte(strconv.Itoa(overflow)), expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Uint8(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Uint8(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Uint8(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestUint16(t *testing.T) {
	overflow := math.MaxUint16 + 1
	tests := []struct {
		input    any
		expected uint16
		hasError bool
	}{
		{input: uint(overflow), expected: 0, hasError: true},
		{input: uint32(overflow), expected: 0, hasError: true},
		{input: uint64(overflow), expected: 0, hasError: true},
		{input: int(overflow), expected: 0, hasError: true},
		{input: int32(overflow), expected: 0, hasError: true},
		{input: int64(overflow), expected: 0, hasError: true},
		{input: float32(overflow), expected: 0, hasError: true},
		{input: float64(overflow), expected: 0, hasError: true},
		{input: strconv.Itoa(overflow), expected: 0, hasError: true},
		{input: []byte(strconv.Itoa(overflow)), expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Uint16(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Uint16(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Uint16(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestUint32(t *testing.T) {
	overflow := uint64(math.MaxUint32) + 1
	tests := []struct {
		input    any
		expected uint32
		hasError bool
	}{
		{input: uint(overflow), expected: 0, hasError: true},
		{input: uint64(overflow), expected: 0, hasError: true},
		{input: int(overflow), expected: 0, hasError: true},
		{input: int64(overflow), expected: 0, hasError: true},
		{input: float32(overflow), expected: 0, hasError: true},
		{input: float64(overflow), expected: 0, hasError: true},
		{input: strconv.FormatUint(overflow, 10), expected: 0, hasError: true},
		{input: []byte(strconv.FormatUint(overflow, 10)), expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Uint32(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Uint32(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Uint32(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
