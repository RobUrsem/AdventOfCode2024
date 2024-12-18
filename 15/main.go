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

	moves := puzzle.ReadMoves(lines)

	//--- Part A
	grid := puzzle.ReadGrid(lines)
	grid.MoveRobot(moves)
	gps := grid.CalculateGPS()

	fmt.Printf("Final grid:\n")
	fmt.Println(grid)
	fmt.Printf("GPS Coordinates sum: %v\n", gps)

	//--- Part B
	widegrid := puzzle.MakeWideGrid(lines)
	widegrid.MoveRobot(moves)
	gps = widegrid.CalculateGPS()

	fmt.Printf("Final grid:\n")
	fmt.Println(widegrid)
	fmt.Printf("GPS Coordinates sum: %v\n", gps)
}
