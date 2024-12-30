package puzzle

import (
	"fmt"
	"math"
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
	costs      [][]int
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

func (m Maze) withinWalls(loc Location) bool {
	return loc.R > 0 && loc.R < len(m.grid) &&
		loc.C > 0 && loc.C < len(m.grid[0])
}

func (m *Maze) AddCheat(loc Location) int {
	if !m.withinWalls(loc) {
		return 0
	}

	minCost := math.MaxInt
	maxCost := 0
	for dir := 0; dir < 4; dir++ {
		l := Combine(loc, directions[dir])
		if m.withinWalls(l) {
			cost := m.costs[l.R][l.C]
			if cost != math.MaxInt {
				minCost = min(minCost, cost)
				maxCost = max(maxCost, cost)
			}
		}
	}

	const COST_OF_CHEAT = 2
	gain := maxCost - minCost - COST_OF_CHEAT
	if gain < 0 || gain == math.MaxInt {
		return 0
	}

	return gain
}

func locKey(a, b Location) string {
	// if a.R < b.R && a.C < b.C {
	// 	return fmt.Sprintf("%v->%v", a, b)
	// }
	return fmt.Sprintf("%v->%v", a, b)
}
func Distance(a, b Location) int {
	dr := a.R - b.R
	dc := a.C - b.C
	if dr < 0 {
		dr = -dr
	}
	if dc < 0 {
		dc = -dc
	}
	return dr + dc
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (m Maze) AddLongerCheat(loc Location, result map[string]int) {
	if !m.withinWalls(loc) {
		return
	}

	// Get cost before this cheat starts
	startLoc := loc
	startCost := m.costs[loc.R][loc.C]

	maxDistance := 20
	for dr := -maxDistance; dr <= maxDistance; dr++ {
		max_c := maxDistance - Abs(dr)
		for dc := -max_c; dc <= max_c; dc++ {
			l := Combine(startLoc, Location{dr, dc})
			if m.withinWalls(l) {
				cost := m.costs[l.R][l.C]
				if cost != math.MaxInt && cost > startCost {
					costOfCheat := Distance(startLoc, l)
					gain := cost - startCost - costOfCheat
					if gain > 0 {
						result[locKey(startLoc, l)] = gain
					}
				}
			}
		}
	}
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
