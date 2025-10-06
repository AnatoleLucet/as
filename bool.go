package as

import (
	"errors"
	"strconv"
)

func stringToBool(s string) (bool, error) {
	if s == "true" || s == "yes" || s == "on" || s == "1" {
		return true, nil
	}
	if s == "false" || s == "no" || s == "off" || s == "0" {
		return false, nil
	}
	return false, errors.New("cannot convert string to bool")
}

func Bool(v any) (bool, error) {
	if v == nil {
		return false, nil
	}

	switch val := v.(type) {
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
		return false, errors.New("cannot convert to bool")
	}
}
