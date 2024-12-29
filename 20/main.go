package main

import (
	"20/puzzle"
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
	maze := puzzle.MakeMaze(lines)
	_, steps := maze.SolveMaze()
	elapsed := time.Since(start)

	fmt.Printf("Steps to solve the maze: %v\n", steps)
	fmt.Printf("Elapsed: %v\n", elapsed)
}
