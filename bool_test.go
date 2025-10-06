package as

import "testing"

func TestBool(t *testing.T) {
	tests := []struct {
		input    any
		expected bool
		hasError bool
	}{
		{input: true, expected: true, hasError: false},
		{input: false, expected: false, hasError: false},
		{input: "true", expected: true, hasError: false},
		{input: "false", expected: false, hasError: false},
		{input: "yes", expected: true, hasError: false},
		{input: "no", expected: false, hasError: false},
		{input: "on", expected: true, hasError: false},
		{input: "off", expected: false, hasError: false},
		{input: "1", expected: true, hasError: false},
		{input: "0", expected: false, hasError: false},
		{input: 1, expected: true, hasError: false},
		{input: 0, expected: false, hasError: false},
		{input: int8(1), expected: true, hasError: false},
		{input: int8(0), expected: false, hasError: false},
		{input: int16(1), expected: true, hasError: false},
		{input: int16(0), expected: false, hasError: false},
		{input: int32(1), expected: true, hasError: false},
		{input: int32(0), expected: false, hasError: false},
		{input: int64(1), expected: true, hasError: false},
		{input: int64(0), expected: false, hasError: false},
		{input: uint(1), expected: true, hasError: false},
		{input: uint(0), expected: false, hasError: false},
		{input: uint8(1), expected: true, hasError: false},
		{input: uint8(0), expected: false, hasError: false},
		{input: uint16(1), expected: true, hasError: false},
		{input: uint16(0), expected: false, hasError: false},
		{input: uint32(1), expected: true, hasError: false},
		{input: uint32(0), expected: false, hasError: false},
		{input: uint64(1), expected: true, hasError: false},
		{input: uint64(0), expected: false, hasError: false},
		{input: float32(1.0), expected: true, hasError: false},
		{input: float32(0.0), expected: false, hasError: false},
		{input: float64(1.0), expected: true, hasError: false},
		{input: float64(0.0), expected: false, hasError: false},
		{input: []byte("true"), expected: true, hasError: false},
		{input: []byte("false"), expected: false, hasError: false},
		{input: []byte("yes"), expected: true, hasError: false},
		{input: []byte("no"), expected: false, hasError: false},
		{input: []byte("on"), expected: true, hasError: false},
		{input: []byte("off"), expected: false, hasError: false},
		{input: []byte("1"), expected: true, hasError: false},
		{input: []byte("0"), expected: false, hasError: false},
		{input: nil, expected: false, hasError: false},
		// error cases
		{input: "invalid", expected: false, hasError: true},
		{input: 2, expected: false, hasError: true},
		{input: struct{}{}, expected: false, hasError: true},
	}

	for _, test := range tests {
		result, err := Bool(test.input)
		if (err != nil) != test.hasError {
			t.Errorf("Bool(%v) unexpected error state: got error '%v'", test.input, err)
			continue
		}
		if err == nil && result != test.expected {
			t.Errorf("Bool(%v) = %v; want %v", test.input, result, test.expected)
		}
	}
}
