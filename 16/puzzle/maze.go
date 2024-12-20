package puzzle

import (
	"advent/shared"
	"fmt"
	"math"

	"github.com/nsf/termbox-go"
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
)

const (
	NOPE = iota
	STEP
	TURN
)

type Maze struct {
	grid       []string
	visited    [][]rune
	moves      shared.Stack
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

func (m Maze) isSafe(x, y int) bool {
	return x >= 0 && x < len(m.grid) &&
		y >= 0 && y < len(m.grid[0]) &&
		m.grid[x][y] != WALL && m.visited[x][y] == NOPE
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
		}
		fmt.Println()
	}
}

func (m Maze) Draw() {
	err := termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	if err != nil {
		fmt.Println("Error clearing termbox:", err)
		return
	}

	for y, row := range m.grid {
		for x, cell := range row {
			visited := m.visited[y][x]
			color := termbox.ColorWhite
			if cell == '#' {
				color = termbox.ColorRed
			} else if visited == '<' || visited == '>' || visited == '^' || visited == 'v' {
				cell = visited
				color = termbox.ColorGreen
			} else if cell == '.' {
				color = termbox.ColorBlack
			}
			termbox.SetCell(x, y, cell, color, termbox.ColorDefault)
		}
	}

	termbox.Flush()
}

func (m *Maze) updateMoves(dir1, dir2 rune) {
	if dir1 == dir2 {
		m.moves.Push(STEP)
	} else {
		m.moves.Push(TURN)
	}
}

func (m *Maze) findShortestPath(i, j int, dir rune, min_cost *int, dist int) {
	if i == m.end[0] && j == m.end[1] {
		cost := m.CalcCosts()
		if cost < *min_cost {
			*min_cost = cost
			fmt.Println()
			fmt.Printf("Distance: %v\n", dist)
			fmt.Printf("Cost: %v\n", cost)
			// m.Print()
			fmt.Println()
		} else {
			fmt.Printf("More expensive: %v %v\n", dist, cost)
		}
		return
	}

	m.visited[i][j] = dir

	// We start to the east as per description
	if m.isSafe(i, j+1) {
		m.updateMoves(dir, EAST)
		m.findShortestPath(i, j+1, EAST, min_cost, dist+1)
		m.moves.Pop()
	}

	if m.isSafe(i+1, j) {
		m.updateMoves(dir, SOUTH)
		m.findShortestPath(i+1, j, SOUTH, min_cost, dist+1)
		m.moves.Pop()
	}

	if m.isSafe(i-1, j) {
		m.updateMoves(dir, NORTH)
		m.findShortestPath(i-1, j, NORTH, min_cost, dist+1)
		m.moves.Pop()
	}

	if m.isSafe(i, j-1) {
		m.updateMoves(dir, WEST)
		m.findShortestPath(i, j-1, WEST, min_cost, dist+1)
		m.moves.Pop()
	}

	//backtrack
	m.visited[i][j] = NOPE
	m.Draw()
}

func (m Maze) FindShortestPathLength() int {
	cost := math.MaxInt32
	m.findShortestPath(m.start[0], m.start[1], EAST, &cost, 0)

	if cost != math.MaxInt32 {
		return cost
	}
	return -1
}

func (m Maze) CalcCosts() int {
	cost := 0

	for _, move := range m.moves.Items {
		switch move {
		case STEP:
			cost++
		case TURN: // a turn includes a turn
			cost += 1001
		}
	}
	return cost
}
