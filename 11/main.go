package main

import (
	"11/puzzle"
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

	if len(lines) != 1 {
		log.Fatalf("Expected single line input")
	}

	input, err := shared.TextToIntArray(lines[0])
	if err != nil {
		log.Fatalf("Error converting [%v]: %v", lines[0], err)
	}

	outcome := input
	for i := 0; i < 25; i++ {
		outcome = puzzle.Blink(outcome)
	}
	fmt.Printf("Got %v stones\n", len(outcome))
}
