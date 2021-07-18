package group_by

import "reflect"

func SliceStable(in T, groupKey func(obj T) T) map[T][]T {
	slice, ok := convertSlice(in)
	if !ok {
		panic("The input should be a slice")
	}

	output := make(map[T][]T)

	for _, v := range slice {
		key := groupKey(v)
		if value, ok := output[key]; ok {
			output[key] = append(value, v)
		} else {
			output[key] = []T{v}
		}
	}

	return output
}

func convertSlice(in T) (out []T, ok bool) {
	value, ok := validateType(in, reflect.Slice)
	if !ok {
		return
	}

	length := value.Len()
	out = make([]T, length)
	for i := 0; i < length; i++ {
		out[i] = value.Index(i).Interface()
	}

	return
}

func validateType(value T, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(value)
	if val.Kind() == kind {
		ok = true
		return
	}

	ok = false

	return
}
