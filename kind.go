package as

import (
	"fmt"
	"reflect"
)

// Kind converts the given value to the specified reflect.Kind.
//
// For scalar values, converts directly to the target kind (e.g., "123" → int(123)).
// For slices, converts all elements to the target kind (e.g., []string{"1"} → []int{1}).
// For maps, converts all values to the target kind, preserving key types.
// If the target is reflect.Slice, it wraps the value in a slice (e.g., 42 → []int{42}).
func Kind[T any](target reflect.Kind, v T) (any, error) {
	typ := findReflectType(v)

	if typ.Kind() == target {
		return v, nil
	}

	if typ.Kind() == reflect.Slice {
		return KindSlice(target, v)
	}
	if typ.Kind() == reflect.Map {
		return KindMap(typ.Key().Kind(), target, v)
	}

	switch target {
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
		slice := reflect.MakeSlice(reflect.SliceOf(typ), 0, 1)
		slice = reflect.Append(slice, findReflectValue(v))
		return slice.Interface(), nil
	}

	return nil, fmt.Errorf("cannot convert type %s to kind %s", typ.Kind(), target.String())
}

// KindSlice converts elements of a slice to the specified reflect.Kind.
//
// If v is a slice, converts all elements to the target kind.
// If v is not a slice, wraps it into a single-element slice first, then converts.
//
// Examples:
//   - KindSlice(reflect.Int, []string{"1", "2"}) → []int{1, 2}
//   - KindSlice(reflect.String, "hello") → []string{"hello"}
//   - KindSlice(reflect.Int, "123") → []int{123}
func KindSlice[T any](target reflect.Kind, v T) (any, error) {
	val := findReflectValue(v)
	typ := findReflectType(v)

	if typ.Kind() == reflect.Slice && typ.Elem().Kind() == target {
		return v, nil
	}

	// wrap non-slice into slice
	if typ.Kind() != reflect.Slice {
		val = reflect.ValueOf([]T{v})
		typ = reflect.TypeOf([]T{v})
	}

	result := reflect.MakeSlice(reflect.SliceOf(toReflectType(target)), 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		elem, err := Kind(target, val.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		result = reflect.Append(result, reflect.ValueOf(elem))
	}

	return result.Interface(), nil
}

// KindMap converts keys and values of a map to the specified reflect.Kinds.
//
// Returns error if v is not a map, or if any key/value conversion fails.
//
// Example:
//   - KindMap(reflect.Int, reflect.Int, map[string]string{"1": "10"}) → map[int]int{1: 10}
func KindMap[T any](keyTarget reflect.Kind, valTarget reflect.Kind, v T) (any, error) {
	typ := findReflectType(v)
	if typ.Kind() != reflect.Map {
		return nil, fmt.Errorf("cannot convert a non-map type %s to a map", typ.Kind())
	}

	if typ.Key().Kind() == keyTarget && typ.Elem().Kind() == valTarget {
		return v, nil
	}

	result := reflect.MakeMap(reflect.MapOf(toReflectType(keyTarget), toReflectType(valTarget)))

	iter := findReflectValue(v).MapRange()
	for iter.Next() {
		key, err := Kind(keyTarget, iter.Key().Interface())
		if err != nil {
			return nil, err
		}

		val, err := Kind(valTarget, iter.Value().Interface())
		if err != nil {
			return nil, err
		}

		result.SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(val))
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

func toReflectType(k reflect.Kind) reflect.Type {
	switch k {
	case reflect.Bool:
		return reflect.TypeOf(false)
	case reflect.Int:
		return reflect.TypeOf(int(0))
	case reflect.Int8:
		return reflect.TypeOf(int8(0))
	case reflect.Int16:
		return reflect.TypeOf(int16(0))
	case reflect.Int32:
		return reflect.TypeOf(int32(0))
	case reflect.Int64:
		return reflect.TypeOf(int64(0))
	case reflect.Uint:
		return reflect.TypeOf(uint(0))
	case reflect.Uint8:
		return reflect.TypeOf(uint8(0))
	case reflect.Uint16:
		return reflect.TypeOf(uint16(0))
	case reflect.Uint32:
		return reflect.TypeOf(uint32(0))
	case reflect.Uint64:
		return reflect.TypeOf(uint64(0))
	case reflect.Float32:
		return reflect.TypeOf(float32(0))
	case reflect.Float64:
		return reflect.TypeOf(float64(0))
	case reflect.String:
		return reflect.TypeOf("")
	default:
		return nil
	}
}
