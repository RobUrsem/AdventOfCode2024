package puzzle

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	ROBOT = '@'
	WALL  = '#'
	SPACE = '.'
	BOX   = 'O'
)

type Grid struct {
	grid  [][]rune
	robot []int
}

func (g Grid) String() string {
	var builder strings.Builder
	for _, row := range g.grid {
		for _, cell := range row {
			builder.WriteString(fmt.Sprintf("%v", string(cell)))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (g Grid) updateRobotPos(r, c int) {
	g.grid[g.robot[0]][g.robot[1]] = SPACE
	g.robot[0] = r
	g.robot[1] = c
	g.grid[r][c] = ROBOT
}

func (g Grid) move(dr, dc int) {
	r := g.robot[0] + dr
	c := g.robot[1] + dc

	if g.grid[r][c] == WALL {
		//--- Can't move a wall, return
		return
	}

	if g.grid[r][c] == SPACE {
		g.updateRobotPos(r, c)
		return
	}

	if g.grid[r][c] == BOX {
		//--- See if we can push the box(es)
		//--- find max num boxes to push
		foundSpace := false
		foundWall := false
		moveUntil := 0
		for ; !foundSpace && !foundWall; moveUntil++ {
			if g.grid[r+dr*moveUntil][c+dc*moveUntil] == SPACE {
				foundSpace = true
			}
			if g.grid[r+dr*moveUntil][c+dc*moveUntil] == WALL {
				foundWall = true
			}
		}

		if foundSpace {
			for j := moveUntil - 1; j >= 0; j-- {
				g.grid[r+dr*j][c+dc*j] = BOX
			}
			g.updateRobotPos(r, c)
		}
	}
}

func (g Grid) MoveRobot(moves string) {
	for _, move := range moves {
		switch move {
		case '^':
			g.move(-1, 0)
		case '>':
			g.move(0, 1)
		case 'v':
			g.move(1, 0)
		case '<':
			g.move(0, -1)
		}
		// fmt.Printf("Move: %v \n", string(move))
		// fmt.Println(g)
	}
}

func (g Grid) CalculateGPS() int {
	total := 0
	for r, line := range g.grid {
		for c, char := range line {
			if char == BOX {
				total += 100*r + c
			}
		}
	}
	return total
}

func ReadGrid(lines []string) Grid {
	grid := Grid{}
	for r, line := range lines {
		if len(line) > 0 && line[0] == WALL {
			gridRow := make([]rune, len(line))
			for c, char := range line {
				gridRow[c] = char
				if char == ROBOT {
					grid.robot = []int{r, c}
				}
			}
			grid.grid = append(grid.grid, gridRow)
		}
	}
	return grid
}

func ReadMoves(lines []string) string {
	rx := regexp.MustCompile(`[v>^<]`)
	moves := ""
	for _, line := range lines {
		if rx.MatchString(line) {
			moves += line
		}
	}
	return moves
}
