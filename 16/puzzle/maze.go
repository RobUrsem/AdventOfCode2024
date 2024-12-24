package puzzle

import (
	"fmt"
)

const (
	SPACE = '.'
	WALL  = '#'
	START = 'S'
	END   = 'E'
	NORTH = '^'
	EAST  = '>'
	SOUTH = 'v'
	WEST  = '<'
	BEST  = 'O'
)

const (
	NOPE = iota
	STEP
	TURN
)

type Maze struct {
	grid       []string
	visited    [][]rune
	start, end []int
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
				maze.start = []int{r, c}
				maze.visited[r][c] = EAST
			} else if char == END {
				maze.end = []int{r, c}
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
