package main

import (
	"19/puzzle"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"
	"time"
)

func main() {
	filePath := filepath.Join("data", "input.txt")

	lines, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	start := time.Now()
	onsen := puzzle.ReadPatterns(lines)
	numPossible := onsen.FindPossible()
	elapsed := time.Since(start)

	fmt.Printf("Found %v possible patterns\n", numPossible)
	fmt.Printf("Elapsed: %v\n", elapsed)
}
