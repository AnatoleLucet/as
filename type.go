package as

import "reflect"

// Type converts the given value to the specified type T.
func Type[T any](v any) (T, error) {
	typ := findReflectType(new(T))
	if reflect.TypeOf(v) == typ {
		return v.(T), nil
	}

	result, err := Kind(typ.Kind(), v)
	if err != nil {
		var zero T
		return zero, err
	}

	return reflect.ValueOf(result).Convert(typ).Interface().(T), nil
}
