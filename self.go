package as

func Self[T any](v T) (T, error) {
	return v, nil
}
