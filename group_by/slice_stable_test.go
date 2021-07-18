package group_by

import "testing"

func TestSliceStable(t *testing.T) {
	input := []string{"Quang", "Pho", "Tuan", "Quan", "Tuan", "Quan", "Pho", "Tuan"}
	output := SliceStable(input, func(obj T) T { return obj })
	t.Log(output)

	if length := len(output); length != 4 {
		t.Errorf("Length of map is 4 instead of %v", length)
	}
}
