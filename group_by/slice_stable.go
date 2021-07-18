package group_by

import "reflect"

func SliceStable(in interface{}, groupKey func(index int) interface{}) map[interface{}][]interface{} {
	slice, ok := convertSlice(in)
	if !ok {
		panic("The input should be a slice")
	}

	output := make(map[interface{}][]interface{})

	for i, v := range slice {
		key := groupKey(i)
		if value, ok := output[key]; ok {
			output[key] = append(value, v)
		} else {
			output[key] = []interface{}{v}
		}
	}

	return output
}

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
