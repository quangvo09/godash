package group_by

import (
	"reflect"
	"testing"
)

func TestGroupSliceInt(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 4, 3, 2, 1}
	output := SliceInt(input)
	t.Log(output)

	if length := len(output); length != 5 {
		t.Errorf("Length of map is 5 instead of %v", length)
	}

	group1 := []int{1, 1}
	res1 := reflect.DeepEqual(output[1], group1)
	if !res1 {
		t.Errorf("Group key `1` shouuld be %v", group1)
	}

	group2 := []int{2, 2}
	res2 := reflect.DeepEqual(output[2], group2)
	if !res2 {
		t.Errorf("Group key `1` shouuld be %v", group2)
	}
}
