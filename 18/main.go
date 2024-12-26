package main

import (
	"18/puzzle"
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

	mem := puzzle.MakeMemory(71, 71)
	mem.LoadBytes(lines)

	start := time.Now()
	mem.Simulate(1024)
	_, cost := mem.SolveMaze()
	elapsed := time.Since(start)
	mem.Print()
	fmt.Printf("Cost to the maze: %v\n", cost)
	fmt.Printf("Elapsed: %v\n", elapsed)
}
