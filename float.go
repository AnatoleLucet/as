package as

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func Float64[T any](v T) (float64, error) {
	switch val := any(v).(type) {
	case nil:
		return 0, nil
	case int:
		return float64(val), nil
	case int8:
		return float64(val), nil
	case int16:
		return float64(val), nil
	case int32:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case uint:
		return float64(val), nil
	case uint8:
		return float64(val), nil
	case uint16:
		return float64(val), nil
	case uint32:
		return float64(val), nil
	case uint64:
		return float64(val), nil
	case float32:
		return float64(val), nil
	case float64:
		return val, nil
	case string:
		return strconv.ParseFloat(val, 64)
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case []byte:
		if utf8.Valid(val) {
			return strconv.ParseFloat(string(val), 64)
		}
		if len(val) >= 8 {
			bits := binary.BigEndian.Uint64(val)
			return float64(bits), nil
		}
		return 0, errors.New("byte slice too short")
	default:
		return 0, fmt.Errorf("cannot convert %T to float64", v)
	}
}

func Float32[T any](v T) (float32, error) {
	f64, err := Float64(v)
	if err != nil {
		return 0, err
	}

	if f64 > math.MaxFloat32 || f64 < -math.MaxFloat32 {
		return 0, errors.New("value out of range for float32")
	}

	return float32(f64), nil
}

func Float(v any) (float64, error) {
	return Float64(v)
}
