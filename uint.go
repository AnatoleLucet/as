package as

import (
	"encoding/binary"
	"errors"
	"math"
	"strconv"
	"unicode/utf8"
)

func Uint64(v any) (uint64, error) {
	if v == nil {
		return 0, nil
	}

	switch val := v.(type) {
	case uint:
		return uint64(val), nil
	case uint8:
		return uint64(val), nil
	case uint16:
		return uint64(val), nil
	case uint32:
		return uint64(val), nil
	case uint64:
		return val, nil
	case string:
		return strconv.ParseUint(val, 10, 64)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		if utf8.Valid(val) {
			return strconv.ParseUint(string(val), 10, 64)
		}
		if len(val) >= 8 {
			return binary.BigEndian.Uint64(val), nil
		}
		return 0, errors.New("byte slice too short")
	}

	// fallback to Int64 conversion for signed types
	i64, err := Int64(v)
	if err != nil {
		return 0, err
	}

	if i64 < 0 {
		return 0, errors.New("cannot convert negative value to uint64")
	}

	return uint64(i64), nil
}

func Uint32(v any) (uint32, error) {
	u64, err := Uint64(v)
	if err != nil {
		return 0, err
	}

	if u64 > math.MaxUint32 {
		return 0, errors.New("value out of range for uint32")
	}

	return uint32(u64), nil
}

func Uint16(v any) (uint16, error) {
	u64, err := Uint64(v)
	if err != nil {
		return 0, err
	}

	if u64 > math.MaxUint16 {
		return 0, errors.New("value out of range for uint16")
	}

	return uint16(u64), nil
}

func Uint8(v any) (uint8, error) {
	u64, err := Uint64(v)
	if err != nil {
		return 0, err
	}

	if u64 > math.MaxUint8 {
		return 0, errors.New("value out of range for uint8")
	}

	return uint8(u64), nil
}

func Uint(v any) (uint, error) {
	u64, err := Uint64(v)
	if err != nil {
		return 0, err
	}

	if u64 > math.MaxUint {
		return 0, errors.New("value out of range for uint")
	}

	return uint(u64), nil
}
