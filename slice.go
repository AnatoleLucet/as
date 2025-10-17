package as

func convertSlice[T, U any](converter func(v any) (U, error), v []T) ([]U, error) {
	result := make([]U, len(v))

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
func Slice[T, U any](converter func(v any) (U, error), v T) ([]U, error) {
	// FIXME: idealy this should be a generic function but Go does not support that yet
	switch val := any(v).(type) {
	case nil:
		return nil, nil
	case []U:
		return val, nil
	case []any:
		return convertSlice(converter, val)
	case []string:
		return convertSlice(converter, val)
	case []int:
		return convertSlice(converter, val)
	case []int8:
		return convertSlice(converter, val)
	case []int16:
		return convertSlice(converter, val)
	case []int64:
		return convertSlice(converter, val)
	case []uint:
		return convertSlice(converter, val)
	case []uint8:
		return convertSlice(converter, val)
	case []uint16:
		return convertSlice(converter, val)
	case []uint32:
		return convertSlice(converter, val)
	case []uint64:
		return convertSlice(converter, val)
	case []float32:
		return convertSlice(converter, val)
	case []float64:
		return convertSlice(converter, val)
	case []bool:
		return convertSlice(converter, val)
	case []rune:
		return convertSlice(converter, val)
	case [][]byte:
		return convertSlice(converter, val)
	default:
		converted, err := converter(val)
		if err != nil {
			return nil, err
		}

		return []U{converted}, nil
	}
}
