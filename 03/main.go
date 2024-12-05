package main

import (
	"advent/03/operations"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	filePath := filepath.Join("data", "input.txt")

	memory, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	var ops []operations.Operation
	for _, dump := range memory {
		ops = append(ops, operations.FindOperations(dump)...)
	}
	total := operations.ExecuteOperations(ops)

	fmt.Printf("Total from the operations: %v\n", total)
}
