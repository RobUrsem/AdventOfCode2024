package main

import (
	"05/ordering"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"
)

func getMiddle(slice []int) (int, bool) {
	length := len(slice)

	if length == 0 {
		return 0, false // Return false if the slice is empty
	}

	mid := length / 2

	if length%2 == 0 {
		// Even-length slice: Return the first of the two middle numbers
		return slice[mid-1], true
	} else {
		// Odd-length slice: Return the middle number
		return slice[mid], true
	}
}

func main() {
	filePath := filepath.Join("data", "input.txt")

	lines, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	rulebook, err := ordering.ConstructRulebook(lines)
	if err != nil {
		log.Fatalf("Error constructing rulebook: %v", err)
	}

	updates, err := ordering.GetUpdates(lines)
	if err != nil {
		log.Fatalf("Error getting updates: %v", err)
	}

	validUpdates, invalidUpdates := ordering.FilterUpdates(updates, rulebook)

	total := 0
	for _, update := range validUpdates {
		middle, success := getMiddle(update)
		if success {
			total += middle
		}
	}

	fmt.Printf("Sum of middle numbers for correct orders: %v\n", total)

	correctedUpdates := ordering.FixUpdates(invalidUpdates, rulebook)

	total = 0
	for _, update := range correctedUpdates {
		middle, success := getMiddle(update)
		if success {
			total += middle
		}
	}
	fmt.Printf("Sum of middle numbers for corrected orders: %v\n", total)
}
