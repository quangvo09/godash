package group_by

func SliceString(slice []string) map[string][]string {
	output := make(map[string][]string)

	for _, v := range slice {
		if value, ok := output[v]; ok {
			output[v] = append(value, v)
		} else {
			output[v] = []string{v}
		}
	}

	return output
}
