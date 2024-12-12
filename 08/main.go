package main

import (
	"08/puzzle"
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

	theMap := puzzle.CreateMap(lines)

	//--- Set to true for part B
	useHarmonics := true
	theMap.Filter(useHarmonics)

	fmt.Println(theMap)
	fmt.Printf("Total antinodes: %v\n", theMap.CountAntiNodes())
}
