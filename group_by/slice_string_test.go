package group_by

import (
	"reflect"
	"testing"
)

func TestGroupSliceString(t *testing.T) {
	input := []string{"Quang", "Pho", "Tuan", "Quan", "Tuan", "Quan", "Pho", "Tuan"}
	output := SliceString(input)
	t.Log(output)

	if length := len(output); length != 4 {
		t.Errorf("Length of map is 4 instead of %v", length)
	}

	group := map[string][]string{
		"Tuan":  {"Tuan", "Tuan", "Tuan"},
		"Pho":   {"Pho", "Pho"},
		"Quan":  {"Quan", "Quan"},
		"Quang": {"Quang"},
	}
	res5 := reflect.DeepEqual(output, group)
	if !res5 {
		t.Errorf("Group should be %v", group)
	}
}
