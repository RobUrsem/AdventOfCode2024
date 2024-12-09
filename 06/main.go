package main

import (
	"06/guard"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	filePath := filepath.Join("data", "test_input.txt")

	lines, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	labMap, err := guard.ConstructMap(lines)
	if err != nil {
		fmt.Println("Could not construct the map")
		return
	}

	guardLeaves := false
	lastMoveWasTurn := false
	moveCounter := 0
	for {
		// fmt.Printf("\nMove: %v\n", moveCounter)
		// guard.PrintMap(labMap)

		r, c := guard.FindGuard(labMap)
		if r == -1 || c == -1 {
			fmt.Printf("Could not find a guard")
			return
		}

		guardLeaves, lastMoveWasTurn = guard.MoveGuard(labMap, r, c, lastMoveWasTurn)
		if guardLeaves {
			break
		}
		moveCounter++

		if moveCounter > 25000 {
			break
		}
	}

	fmt.Println("\nFinal location")
	guard.PrintMap(labMap)
	fmt.Printf("The guard made %v moves\n", moveCounter)
	visited := guard.CountVisited(labMap)
	fmt.Printf("The guard visited %v positions before leaving", visited+1)
}
