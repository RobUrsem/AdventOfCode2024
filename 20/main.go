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
	steps := maze.SolveMaze()
	elapsed := time.Since(start)

	fmt.Printf("Steps to solve the maze: %v\n", steps)
	fmt.Printf("Elapsed: %v\n", elapsed)

	start = time.Now()
	numCheats := maze.Part1(100)
	elapsed = time.Since(start)

	fmt.Printf("Num cheats with >= 100 saving: %v\n", numCheats)
	fmt.Printf("Elapsed: %v\n", elapsed)

	start = time.Now()
	numCheats = maze.Part2(100)
	elapsed = time.Since(start)

	fmt.Printf("Num long cheats with >= 100 saving: %v\n", numCheats)
	fmt.Printf("Elapsed: %v\n", elapsed)
}
