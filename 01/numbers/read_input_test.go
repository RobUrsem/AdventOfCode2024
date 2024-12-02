package numbers

import (
	"path/filepath"
	"testing"
)

func TestReadInput(t *testing.T) {
	filePath := filepath.Join("..", "data", "test_input.txt")
	expected_left := []int{3, 4, 2, 1, 3, 3}
	expected_right := []int{4, 3, 5, 3, 9, 3}

	left, right, err := ReadInput(filePath)
	if err != nil {
		t.Fatalf("Could not read file [%v]: %v", filePath, err)
	}

	checkArray(t, "left", left, expected_left)
	checkArray(t, "right", right, expected_right)
}

func checkArray(t *testing.T, name string, actual []int, expected []int) {
	if len(actual) != len(expected) {
		t.Fatalf("Length of %v is %d but expected %d", name, len(actual), len(expected))
	}

	for i, num := range actual {
		if num != expected[i] {
			t.Errorf("Expected %d at index %d but got %d", expected[i], i, num)
		}
	}
}
