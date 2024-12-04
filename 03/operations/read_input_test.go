package operations

import (
	"path/filepath"
	"testing"
)

func TestReadInput(t *testing.T) {
	filePath := filepath.Join("..", "data", "test_input.txt")
	input, err := ReadInput(filePath)
	if err != nil {
		t.Fatalf("Error reading test file: %v", err)
	}

	expected := 2
	if len(input) != expected {
		t.Fatalf("Expected %v lines but got %v", expected, len(input))
	}
}
