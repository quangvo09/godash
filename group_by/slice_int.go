package group_by

func SliceInt(slice []int) map[int][]int {
	output := make(map[int][]int)

	for _, v := range slice {
		if value, ok := output[v]; ok {
			output[v] = append(value, v)
		} else {
			output[v] = []int{v}
		}
	}

	return output
}
