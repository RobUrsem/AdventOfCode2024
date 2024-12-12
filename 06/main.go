package main

import (
	"06/guard"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"
	"time"
)

func main() {
	filePath := filepath.Join("data", "input.txt")

	lines, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	labMap, err := guard.ConstructMap(lines)
	if err != nil {
		fmt.Println("Could not construct the map")
		return
	}

	mapCopy := guard.CopyMap(labMap)

	_, err = guard.DoWalk(labMap)
	if err != nil {
		fmt.Printf("Error finding solution: %v", err)
	} else {
		fmt.Println("\nFinal location")
		guard.PrintMap(labMap)
		visited := guard.CountVisited(labMap)
		fmt.Printf("The guard visited %v positions before leaving\n", visited)
	}

	//--- Part 2
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	obstructions := guard.BruteForceObstructions(mapCopy)
	fmt.Printf("There are %v obstructions we can choose\n", len(obstructions))

	t = time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05"))
}
