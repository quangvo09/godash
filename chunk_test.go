package godash

import (
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	output := Chunk(input, 2)
	t.Log(output)

	if len(output) != 3 {
		t.Errorf("The length of output must be 3")
	}

	chunks := [][]interface{}{
		[]interface{}{1, 2},
		[]interface{}{3, 4},
		[]interface{}{5, 6},
	}
	res := reflect.DeepEqual(output, chunks)
	if !res {
		t.Errorf("Chunks should be %v", chunks)
	}
}