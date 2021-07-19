package godash

import "reflect"

func convertSlice(in interface{}) (out []interface{}, ok bool) {
	value, ok := validateType(in, reflect.Slice)
	if !ok {
		return
	}

	length := value.Len()
	out = make([]interface{}, length)
	for i := 0; i < length; i++ {
		out[i] = value.Index(i).Interface()
	}

	return
}

func validateType(value interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(value)
	if val.Kind() == kind {
		ok = true
		return
	}

	ok = false

	return
}