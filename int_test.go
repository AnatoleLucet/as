package as

import (
	"math"
	"strconv"
	"testing"
)

func TestInt(t *testing.T) {
	tests := []struct {
		input    any
		expected int
		hasError bool
	}{
		{input: 0, expected: 0, hasError: false},
		{input: -0, expected: -0, hasError: false},
		{input: int(123), expected: 123, hasError: false},
		{input: -int(123), expected: -123, hasError: false},
		{input: int8(45), expected: 45, hasError: false},
		{input: int8(-45), expected: -45, hasError: false},
		{input: int16(6789), expected: 6789, hasError: false},
		{input: int16(-6789), expected: -6789, hasError: false},
		{input: int32(123456), expected: 123456, hasError: false},
		{input: int32(-123456), expected: -123456, hasError: false},
		{input: int64(9876543210), expected: 9876543210, hasError: false},
		{input: int64(-9876543210), expected: -9876543210, hasError: false},
		{input: uint(123), expected: 123, hasError: false},
		{input: uint8(45), expected: 45, hasError: false},
		{input: uint16(6789), expected: 6789, hasError: false},
		{input: uint32(123456), expected: 123456, hasError: false},
		{input: uint64(9876543210), expected: 9876543210, hasError: false},
		{input: float32(123.0), expected: 123, hasError: false},
		{input: float32(-123.0), expected: -123, hasError: false},
		{input: float64(456.0), expected: 456, hasError: false},
		{input: float64(-456.0), expected: -456, hasError: false},
		{input: "789", expected: 789, hasError: false},
		{input: []byte("101112"), expected: 101112, hasError: false},
		{input: rune('A'), expected: 65, hasError: false},
		{input: true, expected: 1, hasError: false},
		{input: false, expected: 0, hasError: false},
		{input: nil, expected: 0, hasError: false},
		// error cases
		{input: "invalid", expected: 0, hasError: true},
		{input: struct{}{}, expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Int(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Int(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Int(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestInt8(t *testing.T) {
	overflow := int64(math.MaxInt8) + 1
	underflow := int64(math.MinInt8) - 1
	tests := []struct {
		input    any
		expected int8
		hasError bool
	}{
		{input: int(overflow), expected: 0, hasError: true},
		{input: int16(overflow), expected: 0, hasError: true},
		{input: int32(overflow), expected: 0, hasError: true},
		{input: int64(overflow), expected: 0, hasError: true},
		{input: uint(overflow), expected: 0, hasError: true},
		{input: uint8(overflow), expected: 0, hasError: true},
		{input: uint16(overflow), expected: 0, hasError: true},
		{input: uint32(overflow), expected: 0, hasError: true},
		{input: uint64(overflow), expected: 0, hasError: true},
		{input: int(underflow), expected: 0, hasError: true},
		{input: int16(underflow), expected: 0, hasError: true},
		{input: int32(underflow), expected: 0, hasError: true},
		{input: int64(underflow), expected: 0, hasError: true},
		{input: float32(overflow), expected: 0, hasError: true},
		{input: float64(overflow), expected: 0, hasError: true},
		{input: float32(underflow), expected: 0, hasError: true},
		{input: float64(underflow), expected: 0, hasError: true},
		{input: strconv.Itoa(int(overflow)), expected: 0, hasError: true},
		{input: strconv.Itoa(int(underflow)), expected: 0, hasError: true},
		{input: []byte(strconv.Itoa(int(overflow))), expected: 0, hasError: true},
		{input: []byte(strconv.Itoa(int(underflow))), expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Int8(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Int8(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Int8(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestInt16(t *testing.T) {
	overflow := int64(math.MaxInt16) + 1
	underflow := int64(math.MinInt16) - 1
	tests := []struct {
		input    any
		expected int16
		hasError bool
	}{
		{input: int(overflow), expected: 0, hasError: true},
		{input: int32(overflow), expected: 0, hasError: true},
		{input: int64(overflow), expected: 0, hasError: true},
		{input: uint(overflow), expected: 0, hasError: true},
		{input: uint16(overflow), expected: 0, hasError: true},
		{input: uint32(overflow), expected: 0, hasError: true},
		{input: uint64(overflow), expected: 0, hasError: true},
		{input: int(underflow), expected: 0, hasError: true},
		{input: int32(underflow), expected: 0, hasError: true},
		{input: int64(underflow), expected: 0, hasError: true},
		{input: float32(overflow), expected: 0, hasError: true},
		{input: float64(overflow), expected: 0, hasError: true},
		{input: float32(underflow), expected: 0, hasError: true},
		{input: float64(underflow), expected: 0, hasError: true},
		{input: strconv.Itoa(int(overflow)), expected: 0, hasError: true},
		{input: strconv.Itoa(int(underflow)), expected: 0, hasError: true},
		{input: []byte(strconv.Itoa(int(overflow))), expected: 0, hasError: true},
		{input: []byte(strconv.Itoa(int(underflow))), expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Int16(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Int16(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Int16(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestInt32(t *testing.T) {
	overflow := int64(math.MaxInt32) + 1
	underflow := int64(math.MinInt32) - 1
	tests := []struct {
		input    any
		expected int32
		hasError bool
	}{
		{input: int(overflow), expected: 0, hasError: true},
		{input: int64(overflow), expected: 0, hasError: true},
		{input: uint(overflow), expected: 0, hasError: true},
		{input: uint32(overflow), expected: 0, hasError: true},
		{input: uint64(overflow), expected: 0, hasError: true},
		{input: int(underflow), expected: 0, hasError: true},
		{input: int64(underflow), expected: 0, hasError: true},
		{input: float32(overflow), expected: 0, hasError: true},
		{input: float64(overflow), expected: 0, hasError: true},
		{input: float64(underflow), expected: 0, hasError: true},
		{input: strconv.FormatInt(overflow, 10), expected: 0, hasError: true},
		{input: strconv.FormatInt(underflow, 10), expected: 0, hasError: true},
		{input: []byte(strconv.FormatInt(overflow, 10)), expected: 0, hasError: true},
		{input: []byte(strconv.FormatInt(underflow, 10)), expected: 0, hasError: true},
	}

	for _, test := range tests {
		result, err := Int32(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Int32(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Int32(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
