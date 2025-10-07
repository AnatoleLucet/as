package as

import (
	"encoding/binary"
	"errors"
	"math"
	"strconv"
	"unicode/utf8"
)

func Int64[T any](v T) (int64, error) {
	switch val := any(v).(type) {
	case nil:
		return 0, nil
	case int:
		return int64(val), nil
	case int8:
		return int64(val), nil
	case int16:
		return int64(val), nil
	case int32:
		return int64(val), nil
	case int64:
		return val, nil
	case uint:
		return int64(val), nil
	case uint8:
		return int64(val), nil
	case uint16:
		return int64(val), nil
	case uint32:
		return int64(val), nil
	case uint64:
		return int64(val), nil
	case float32:
		return int64(val), nil
	case float64:
		return int64(val), nil
	case string:
		return strconv.ParseInt(val, 10, 64)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		if utf8.Valid(val) {
			return strconv.ParseInt(string(val), 10, 64)
		}
		if len(val) >= 8 {
			return int64(binary.BigEndian.Uint64(val)), nil
		}
		return 0, errors.New("byte slice too short")
	default:
		return 0, errors.New("cannot convert to int64")
	}
}

func Int32[T any](v T) (int32, error) {
	i64, err := Int64(v)
	if err != nil {
		return 0, err
	}

	if i64 > math.MaxInt32 || i64 < math.MinInt32 {
		return 0, errors.New("value out of range for int32")
	}

	return int32(i64), nil
}

func Int16[T any](v T) (int16, error) {
	i64, err := Int64(v)
	if err != nil {
		return 0, err
	}

	if i64 > math.MaxInt16 || i64 < math.MinInt16 {
		return 0, errors.New("value out of range for int16")
	}

	return int16(i64), nil
}

func Int8[T any](v T) (int8, error) {
	i64, err := Int64(v)
	if err != nil {
		return 0, err
	}

	if i64 > math.MaxInt8 || i64 < math.MinInt8 {
		return 0, errors.New("value out of range for int8")
	}

	return int8(i64), nil
}

func Int[T any](v T) (int, error) {
	i64, err := Int64(v)
	if err != nil {
		return 0, err
	}

	return int(i64), nil
}
