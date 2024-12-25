package main

import (
	"17/puzzle"
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

	computer := puzzle.NewComputer(lines)
	computer.Run()

	output := computer.Output()
	fmt.Printf("Output: [%v]\n", output)
}
