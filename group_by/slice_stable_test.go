package group_by

import "testing"

func TestSliceStable(t *testing.T) {
	input := []string{"Quang", "Pho", "Tuan", "Quan", "Tuan", "Quan", "Pho", "Tuan"}
	output := SliceStable(input, func(index int) interface{} { return input[index] })
	t.Log(output)

	if length := len(output); length != 4 {
		t.Errorf("Length of map is 4 instead of %v", length)
	}
}
