package main

import (
	"10/puzzle"
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

	trailHeads := puzzle.FindTrailHeads(lines)

	totalRating := 0
	totalScore := 0
	for _, th := range trailHeads {
		summitLocations := shared.Locations{}
		totalRating += puzzle.Venture(lines, shared.Locations{th}, &summitLocations)
		totalScore += len(summitLocations)
	}

	fmt.Printf("Total score: %v\n", totalScore)
	fmt.Printf("Total rating: %v\n", totalRating)
}
