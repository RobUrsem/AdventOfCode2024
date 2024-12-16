package main

import (
	"13/puzzle"
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

	games := puzzle.GetInput(lines)

	totalTokens := 0
	for _, game := range games {
		cost := game.Cost()
		if cost > 0 {
			totalTokens += cost
		}
	}

	fmt.Printf("Total cost for winning games: %v tokens\n", totalTokens)
}
