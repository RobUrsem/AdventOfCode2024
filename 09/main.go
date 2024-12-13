package main

import (
	"09/puzzle"
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

	if len(lines) > 1 {
		fmt.Printf("Expected single line input")
		return
	}

	disk := puzzle.Analyze(lines[0])
	fast := puzzle.FastCompress(disk)
	checksum := puzzle.FastChecksum(fast)
	fmt.Printf("Checksum: %v\n", checksum)

	disk = puzzle.Analyze(lines[0])
	defrag := puzzle.Defrag(disk)
	checksum = puzzle.FastChecksum(defrag)
	fmt.Printf("Defrag checksum: %v\n", checksum)
}
