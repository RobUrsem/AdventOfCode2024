package main

import (
	"16/puzzle"
	"advent/shared"
	"fmt"
	"log"
	"path/filepath"

	"github.com/nsf/termbox-go"
)

func main() {
	filePath := filepath.Join("data", "input.txt")

	lines, err := shared.ReadInput(filePath)
	if err != nil {
		log.Fatalf("Error reading [%v]: %v", filePath, err)
	}

	err = termbox.Init()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer termbox.Close()

	maze := puzzle.MakeMaze(lines)
	cost := maze.FindShortestPathLength()
	fmt.Println(maze)
	fmt.Printf("Cost: %v\n", cost)
}
