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

type Cheat struct {
	a     Location
	costs [][]int
}

func Combine(a, b Location) Location {
	return Location{a.R + b.R, a.C + b.C}
}

type Maze struct {
	grid       []string
	visited    [][]rune
	costs      [][]int
	start, end Location
	cheat      Cheat
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

func (m Maze) withinWalls(loc Location) bool {
	return loc.R > 0 && loc.R < len(m.grid) &&
		loc.C > 0 && loc.C < len(m.grid[0])
}

func copy(a [][]int) [][]int {
	b := make([][]int, len(a))
	for i, line := range a {
		b[i] = make([]int, len(line))
		for j, v := range line {
			b[i][j] = v
		}
	}
	return b
}

func (m *Maze) AddCheat(loc Location) bool {
	if !m.withinWalls(loc) {
		return false
	}
	m.cheat.a = loc
	m.cheat.costs = copy(m.costs)

	return true
}

func (m *Maze) RemoveCheat() {
	m.costs = copy(m.cheat.costs)
	m.cheat = Cheat{}
}

func (m *Maze) isCheat(l Location) bool {
	if l.R == m.cheat.a.R && l.C == m.cheat.a.C {
		return true
	}
	return false
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
