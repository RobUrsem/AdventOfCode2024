package main

import (
	"15/puzzle"
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

	grid := puzzle.ReadGrid(lines)
	moves := puzzle.ReadMoves(lines)
	grid.MoveRobot(moves)
	gps := grid.CalculateGPS()

	fmt.Printf("Final grid:\n")
	fmt.Println(grid)
	fmt.Printf("GPS Coordinates sum: %v\n", gps)
}
