package main

import (
	"17/puzzle"
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

	computer := puzzle.NewComputer(lines)

	// Part A
	begin := time.Now()
	computer.Run()
	elapsed := time.Since(begin)

	output := computer.Print()
	fmt.Printf("Output: [%v]\n", output)
	fmt.Printf("Elapsed: %v\n", elapsed)

	fmt.Println()

	// Part B
	begin = time.Now()
	computer.RunReverse()
	elapsed = time.Since(begin)

	fmt.Printf("Corrected value of register A: %v\n", computer.CorrectedA)
	fmt.Printf("Elapsed: %v\n", elapsed)
}
