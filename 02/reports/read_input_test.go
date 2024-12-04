package reports

import (
	"path/filepath"
	"testing"
)

func TestReadInput(t *testing.T) {
	filePath := filepath.Join("..", "data", "test_input.txt")
	expected := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	records, err := ReadInput(filePath)
	if err != nil {
		t.Fatalf("Could not read file [%v]: %v", filePath, err)
	}

	if len(records) != len(expected) {
		t.Fatalf("Number of records is %d but expected %d", len(records), len(expected))
	}

	for i, record := range records {
		for j, num := range record {
			if num != expected[i][j] {
				t.Errorf("Expected %d at index %d,%d but got %d", expected[i][j], i, j, num)
			}
		}
	}
}
