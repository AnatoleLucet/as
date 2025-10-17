package as

import (
	"fmt"
	"reflect"
)

// T converts the given value to the specified type T.
func T[T any](v any) (T, error) {
	typ := findReflectType(new(T))
	if reflect.TypeOf(v) == typ {
		return v.(T), nil
	}

	result, err := Type(typ, v)
	if err != nil {
		var zero T
		return zero, err
	}

	return reflect.ValueOf(result).Convert(typ).Interface().(T), nil
}

// Type converts the given value to the specified reflect.Type.
func Type[T any](target reflect.Type, v T) (any, error) {
	typ := findReflectType(v)

	if isSameReflectType(typ, target) {
		return v, nil
	}

	switch target.Kind() {
	case reflect.Bool:
		return Bool(v)
	case reflect.Int:
		return Int(v)
	case reflect.Int8:
		return Int8(v)
	case reflect.Int16:
		return Int16(v)
	case reflect.Int32:
		return Int32(v)
	case reflect.Int64:
		return Int64(v)
	case reflect.Uint:
		return Uint(v)
	case reflect.Uint8:
		return Uint8(v)
	case reflect.Uint16:
		return Uint16(v)
	case reflect.Uint32:
		return Uint32(v)
	case reflect.Uint64:
		return Uint64(v)
	case reflect.Float32:
		return Float32(v)
	case reflect.Float64:
		return Float64(v)
	case reflect.String:
		return String(v)
	case reflect.Interface:
		return Any(v), nil
	case reflect.Slice:
		return TypeSlice(target.Elem(), v)
	case reflect.Map:
		if typ.Kind() == reflect.Map {
			return TypeMap(target.Key(), target.Elem(), v)
		}
	}

	return nil, fmt.Errorf("cannot convert type %s to kind %s", typ.Kind(), target.String())
}

// TypeSlice converts elements of a slice to the specified reflect.Type.
func TypeSlice[T any](target reflect.Type, v T) (any, error) {
	typ := findReflectType(v)
	val := findReflectValue(v)

	if typ.Kind() == reflect.Slice && isSameReflectType(typ.Elem(), target) {
		return v, nil
	}

	// wrap non-slice value into a slice
	if typ.Kind() != reflect.Slice {
		val = reflect.ValueOf([]T{v})
		typ = val.Type()
	}

	slice := reflect.MakeSlice(reflect.SliceOf(target), 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		elem, err := Type(target, val.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		slice = reflect.Append(slice, reflect.ValueOf(elem))
	}

	return slice.Interface(), nil
}

// TypeMap converts keys and values of a map to the specified reflect.Types.
func TypeMap[T any](keyTarget, elemTarget reflect.Type, v T) (any, error) {
	typ := findReflectType(v)
	val := findReflectValue(v)

	if typ.Kind() == reflect.Map && isSameReflectType(typ.Key(), keyTarget) && isSameReflectType(typ.Elem(), elemTarget) {
		return v, nil
	}

	if typ.Kind() != reflect.Map {
		return nil, fmt.Errorf("cannot convert non-map type %s to map", typ.Kind())
	}

	result := reflect.MakeMap(reflect.MapOf(keyTarget, elemTarget))
	for _, key := range val.MapKeys() {
		newKey, err := Type(keyTarget, key.Interface())
		if err != nil {
			return nil, err
		}

		newElem, err := Type(elemTarget, val.MapIndex(key).Interface())
		if err != nil {
			return nil, err
		}

		result.SetMapIndex(reflect.ValueOf(newKey), reflect.ValueOf(newElem))
	}

	return result.Interface(), nil
}

func findReflectType(v any) reflect.Type {
	typ := reflect.TypeOf(v)
	if typ.Kind() == reflect.Pointer {
		typ = typ.Elem()
	}

	return typ
}

func findReflectValue(v any) reflect.Value {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}

func isSameReflectType(a, b reflect.Type) bool {
	if a.Kind() != b.Kind() {
		return false
	}

	if a.Kind() == reflect.Map {
		return isSameReflectType(a.Key(), b.Key()) && isSameReflectType(a.Elem(), b.Elem())
	}

	if a.Kind() == reflect.Slice {
		return isSameReflectType(a.Elem(), b.Elem())
	}

	return true
}
