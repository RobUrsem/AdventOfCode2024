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
	towels := puzzle.MakeTowels(lines)
	numPossible := towels.Part1()
	numWays := towels.Part2()
	elapsed := time.Since(start)
	fmt.Printf("Elapsed: %v\n", elapsed)

	fmt.Printf("Found %v possible patterns\n", numPossible)
	fmt.Printf("Found %v total ways\n", numWays)
	fmt.Printf("Elapsed: %v\n", elapsed)
}
