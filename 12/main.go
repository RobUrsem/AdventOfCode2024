package main

import (
	"12/puzzle"
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
	regions := puzzle.Segment(lines)
	cost := puzzle.CalcCost(regions)
	elapsed := time.Since(start)

	fmt.Printf("Total cost with perimeter: %v\n", cost)
	fmt.Printf("Elapsed time: %v\n", elapsed)

	start = time.Now()
	cost = puzzle.CalcSides(regions)
	elapsed = time.Since(start)

	fmt.Printf("Total cost with sides: %v\n", cost)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
