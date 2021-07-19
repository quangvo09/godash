package godash

func Chunk(input interface{}, number int) [][]interface{} {
	slice, ok := convertSlice(input)
	if !ok {
		panic("The input must be a slice")
	}

	length := len(slice)
	output := make([][]interface{}, 0)
	i := 0
	j := number
	for ; j < length; {
		output = append(output, slice[i:j])
		i = j
		j = j + number
	}

	if j != length - 1 {
		output = append(output, slice[i:])
	}

	return output
}
