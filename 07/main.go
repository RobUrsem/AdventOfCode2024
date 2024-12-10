package main

import (
	"07/puzzle"
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

	equations := puzzle.GetEquations(lines)
	fmt.Printf("Got %v equations\n", len(equations))

	total := int64(0)
	for _, equation := range equations {
		solved := puzzle.SolveEquation(equation)

		if solved.Valid {
			total += solved.Answer
		}
	}

	fmt.Printf("Total of the valid equations: %v\n", total)
}
