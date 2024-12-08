package main

import (
	"advent/04/search"
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

	needle := "XMAS"
	numTimes := search.SearchForWord(needle, lines)
	fmt.Printf("Found %v a total of %v times\n", needle, numTimes)

	needle = "MAS"
	numTimes, err = search.SearchForCross(needle, lines)
	if err != nil {
		fmt.Printf("Oops, encountered an error: %v", err)
	}
	fmt.Printf("Found cross a total of %v times\n", numTimes)
}
