package as

func toSlice[T, U any](converter func(v any) (T, error), v []U) ([]T, error) {
	result := make([]T, len(v))

	for i, item := range v {
		converted, err := converter(item)
		if err != nil {
			return nil, err
		}

		result[i] = converted
	}

	return result, nil
}

// Slice attempts to convert the input to a slice of type T.
func Slice[T any](converter func(v any) (T, error)) func(v any) ([]T, error) {
	// FIXME: idealy this should be a generic function but Go does not support that yet
	return func(v any) ([]T, error) {
		switch val := any(v).(type) {
		case nil:
			return nil, nil
		case []T:
			return val, nil
		case []any:
			return toSlice(converter, val)
		case []string:
			return toSlice(converter, val)
		case []int:
			return toSlice(converter, val)
		case []int8:
			return toSlice(converter, val)
		case []int16:
			return toSlice(converter, val)
		case []int64:
			return toSlice(converter, val)
		case []uint:
			return toSlice(converter, val)
		case []uint8:
			return toSlice(converter, val)
		case []uint16:
			return toSlice(converter, val)
		case []uint32:
			return toSlice(converter, val)
		case []uint64:
			return toSlice(converter, val)
		case []float32:
			return toSlice(converter, val)
		case []float64:
			return toSlice(converter, val)
		case []bool:
			return toSlice(converter, val)
		case []rune:
			return toSlice(converter, val)
		case [][]byte:
			return toSlice(converter, val)
		default:
			converted, err := converter(val)
			if err != nil {
				return nil, err
			}

			return []T{converted}, nil
		}
	}
}
