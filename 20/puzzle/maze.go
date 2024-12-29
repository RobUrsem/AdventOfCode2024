package puzzle

import (
	"fmt"
)

const (
	SPACE = '.'
	WALL  = '#'
	START = 'S'
	END   = 'E'
	PATH  = 'O'
)

const (
	NOPE = iota
)

type Location struct {
	R, C int
}

func Combine(a, b Location) Location {
	return Location{a.R + b.R, a.C + b.C}
}

type Maze struct {
	grid       []string
	visited    [][]rune
	start, end Location
}

func MakeMaze(lines []string) Maze {
	maze := Maze{
		grid:    lines,
		visited: make([][]rune, len(lines)),
	}

	for r, line := range lines {
		maze.visited[r] = make([]rune, len(line))
		for c, char := range line {
			if char == START {
				maze.start = Location{r, c}
				maze.visited[r][c] = PATH
			} else if char == END {
				maze.end = Location{r, c}
			}
		}
	}

	return maze
}

func (m Maze) Print() {
	for r, line := range m.grid {
		for c, char := range line {
			switch char {
			case WALL, START, END:
				fmt.Print(string(char))
			case SPACE:
				if m.visited[r][c] != NOPE {
					fmt.Print(string(m.visited[r][c]))
				} else {
					fmt.Print(" ")
				}
			}
			// fmt.Print(",")
		}
		fmt.Println()
	}
}
