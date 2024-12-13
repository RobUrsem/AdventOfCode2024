package main

import (
	"11/puzzle"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"
	"time"
)

func doBlink(input []int, ch chan int) {
	outcome := input
	for i := 0; i < 75; i++ {
		outcome = puzzle.Blink(outcome)
	}
	ch <- len(outcome)
}

func main() {
	filePath := filepath.Join("data", "input.txt")

	lines, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	if len(lines) != 1 {
		log.Fatalf("Expected single line input")
	}

	input, err := shared.TextToIntArray(lines[0])
	if err != nil {
		log.Fatalf("Error converting [%v]: %v", lines[0], err)
	}

	start := time.Now()
	ch := make(chan int, len(input))
	for _, v := range input {
		go doBlink([]int{v}, ch)
	}

	totalStones := 0
	for range input {
		totalStones += <-ch
	}
	fmt.Printf("Got %v stones\n", totalStones)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %v\n", elapsed)

	// outcome := input
	// for i := 0; i < 75; i++ {
	// 	outcome = puzzle.Blink(outcome)
	// }
	// fmt.Printf("Got %v stones\n", len(outcome))
}
