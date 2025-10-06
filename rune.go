package as

import (
	"errors"
	"unicode/utf8"
)

func Rune(v any) (rune, error) {
	if v == nil {
		return 0, nil
	}

	switch val := v.(type) {
	case rune:
		return val, nil
	case string:
		if val == "" {
			return 0, errors.New("cannot convert empty string to rune")
		}
		r, size := utf8.DecodeRuneInString(val)
		if r == utf8.RuneError && size == 1 {
			return 0, errors.New("invalid UTF-8 encoding")
		}
		return r, nil
	case []byte:
		if len(val) == 0 {
			return 0, errors.New("cannot convert empty byte slice to rune")
		}
		r, size := utf8.DecodeRune(val)
		if r == utf8.RuneError && size == 1 {
			return 0, errors.New("invalid UTF-8 encoding")
		}
		return r, nil
	case bool:
		if val {
			return '1', nil
		}
		return '0', nil
	}

	// fallback to Int64 conversion for numeric types
	i64, err := Int64(v)
	if err != nil {
		return 0, err
	}

	if i64 < 0 || i64 > utf8.MaxRune {
		return 0, errors.New("integer value out of rune range")
	}

	return rune(i64), nil
}
