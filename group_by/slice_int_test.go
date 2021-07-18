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
		t.Errorf("Group key `2` shouuld be %v", group2)
	}

	group3 := []int{3, 3}
	res3 := reflect.DeepEqual(output[3], group3)
	if !res3 {
		t.Errorf("Group key `3` shouuld be %v", group3)
	}
	group4 := []int{4, 4}
	res4 := reflect.DeepEqual(output[4], group4)
	if !res4 {
		t.Errorf("Group key `4` shouuld be %v", group4)
	}
	group5 := []int{5}
	res5 := reflect.DeepEqual(output[5], group5)
	if !res5 {
		t.Errorf("Group key `5` shouuld be %v", group5)
	}
}
