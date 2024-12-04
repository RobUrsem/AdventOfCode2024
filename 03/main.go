package main

import (
	"03/operations"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	filePath := filepath.Join("data", "input.txt")

	memory, err := operations.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	total := 0
	for _, dump := range memory {
		ops, err := operations.FindOperations(dump)

		if err != nil {
			fmt.Printf("Fatal error finding operations: %v\n", err)
			break
		}
		for _, operation := range ops {
			total += operations.Execute(operation)
		}
	}

	fmt.Printf("Total from the operations: %v\n", total)
}
