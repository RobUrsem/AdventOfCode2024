package main

import (
	"16/puzzle"
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

	maze := puzzle.MakeMaze(lines)

	start := time.Now()
	_, cost, numSeats := maze.SolveMaze()
	elapsed := time.Since(start)

	fmt.Printf("Cost: %v\n", cost)
	fmt.Printf("Best seats: %v\n", numSeats)
	fmt.Printf("Elapsed: %v\n", elapsed)
}
