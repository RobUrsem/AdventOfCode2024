package main

import (
	"14/puzzle"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"
)

/*
------------------------------------------------------------
We're looking for the image of a christmas tree:

	1111111111111111111111111111111
	1                             1
	1                             1
	1                             1
	1                             1
	1              1              1
	1             111             1
	1            11111            1
	1           1111111           1
	1          111111111          1
	1            11111            1
	1           1111111           1
	1          111111111          1
	1         11111111111         1
	1        1111111111111        1
	1          111111111          1
	1         11111111111         1
	1        1111111111111        1
	1       111111111111111       1
	1      11111111111111111      1
	1        1111111111111        1
	1       111111111111111       1
	1      11111111111111111      1
	1     1111111111111111111     1
	1    111111111111111111111    1
	1             111             1
	1             111             1
	1             111             1
	1                             1
	1                             1
	1                             1
	1                             1
	1111111111111111111111111111111
	Test if we can find a consecutive string of 20 occupied tiles

-------------------------------------------------------------
*/
func FindTree(bathroom puzzle.Bathroom) bool {
	tiles := bathroom.GetTiles()
	for _, line := range tiles {
		total := 0
		for _, v := range line {
			if v > 0 {
				total++
			} else {
				total = 0
			}
			if total > 20 {
				return true
			}
		}
	}
	return false
}

func main() {
	filePath := filepath.Join("data", "input.txt")

	lines, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	// Part A
	robots := puzzle.GetInput(lines)
	bathRoom := puzzle.NewBathroom(103, 101)
	bathRoom.AddRobots(robots)

	for tick := 0; tick < 100; tick++ {
		bathRoom.Tick()
	}

	safety := bathRoom.SafetyScore()
	fmt.Printf("Safety score: %v\n", safety)

	// Part B
	robots = puzzle.GetInput(lines)
	bathRoom = puzzle.NewBathroom(103, 101)
	bathRoom.AddRobots(robots)

	tick := 0
	for {
		bathRoom.Tick()

		if FindTree(bathRoom) {
			fmt.Printf("Found Christmas tree at tick %d\n", tick+1)
			fmt.Print(bathRoom)
			break
		}

		tick++
	}
}
