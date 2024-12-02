package main

import (
	"01/numbers"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	filePath := filepath.Join("data", "input.txt")

	left, right, err := numbers.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	distance, err := numbers.CalcDistances(left, right)
	if err != nil {
		log.Fatalf("Error calculating distance: %v", err)
	}

	fmt.Printf("Distance for the given input  : %v\n", distance)

	similarity := numbers.CalcSimilarity(left, right)
	fmt.Printf("Similarity for the given input: %v\n", similarity)
}
