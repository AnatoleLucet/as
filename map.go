package as

func Map[OldK comparable, NewK comparable, OldV any, NewV any](
	keyConverter func(OldK) (NewK, error),
	valConverter func(OldV) (NewV, error),
	v map[OldK]OldV,
) (map[NewK]NewV, error) {
	result := make(map[NewK]NewV, len(v))

	for k, v := range v {
		newK, err := keyConverter(k)
		if err != nil {
			return nil, err
		}

		newV, err := valConverter(v)
		if err != nil {
			return nil, err
		}

		result[newK] = newV
	}

	return result, nil
}
