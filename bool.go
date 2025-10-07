package as

import (
	"fmt"
	"strconv"
)

func stringToBool(s string) (bool, error) {
	if s == "true" || s == "yes" || s == "on" || s == "1" {
		return true, nil
	}
	if s == "false" || s == "no" || s == "off" || s == "0" {
		return false, nil
	}

	return false, fmt.Errorf("cannot convert string %q to bool", s)
}

func Bool[T any](v T) (bool, error) {
	switch val := any(v).(type) {
	case nil:
		return false, nil
	case bool:
		return val, nil
	case rune:
		return stringToBool(strconv.Itoa(int(val)))
	case int, int8, int16, int64, uint, uint8, uint16, uint32, uint64, float32, float64, string, []byte:
		s, err := String(val)
		if err != nil {
			return false, err
		}
		return stringToBool(s)
	default:
		return false, fmt.Errorf("cannot convert %T to bool", v)
	}
}
