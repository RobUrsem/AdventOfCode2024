package main

import (
	"09/puzzle"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	filePath := filepath.Join("data", "input.txt")

	lines, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	if len(lines) > 1 {
		fmt.Printf("Expected single line input")
		return
	}

	output := puzzle.Expand(lines[0])
	compressed := puzzle.Compress(output)
	checksum := puzzle.Checksum(compressed)
	fmt.Printf("Checksum: %v\n", checksum)
}
