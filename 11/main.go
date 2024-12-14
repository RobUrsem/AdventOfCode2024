package main

import (
	"11/puzzle"
	"advent/shared"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

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

	numBlinks, err := strconv.Atoi(os.Args[1])
	if err != nil {
		numBlinks = 25
	}

	start := time.Now()
	counter := puzzle.MakeStoneCounter(input)
	for i := 0; i < numBlinks; i++ {
		counter.Blink()
	}
	elapsed := time.Since(start)

	fmt.Printf("Got %v stones\n", counter.Total())
	fmt.Printf("Elapsed time: %v\n", elapsed)
}