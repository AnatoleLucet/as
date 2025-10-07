package as

// Slice attempts to convert the input to a slice of type T.
func Slice[T any](converter func(v any) (T, error)) func(v any) ([]T, error) {
	// FIXME: idealy this should be a generic function but Go does not support that yet
	return func(v any) ([]T, error) {
		switch val := any(v).(type) {
		case []T:
			return val, nil
		case []any:
			result := make([]T, len(val))

			for i, item := range val {
				converted, err := converter(item)
				if err != nil {
					return nil, err
				}

				result[i] = converted
			}

			return result, nil
		default:
			converted, err := converter(val)
			if err != nil {
				return nil, err
			}

			return []T{converted}, nil
		}
	}
}
