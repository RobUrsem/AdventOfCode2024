package main

import (
	"14/puzzle"
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

	robots := puzzle.GetInput(lines)

	bathRoom := puzzle.NewBathroom(103, 101)
	bathRoom.AddRobots(robots)

	for tick := 0; tick < 100; tick++ {
		bathRoom.Tick()
	}

	fmt.Println(bathRoom)
	safety := bathRoom.SafetyScore()
	fmt.Printf("Safety score: %v\n", safety)
}
