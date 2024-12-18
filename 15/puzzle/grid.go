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
	BOX_L = '['
	BOX_R = ']'
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

func (g Grid) moveBox(r, c, dr, dc int) bool {
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
			g.grid[r+dr*j][c+dc*j] = g.grid[r+dr*(j-1)][c+dc*(j-1)]
			g.grid[r+dr*j][c+dc*j] = g.grid[r+dr*(j-1)][c+dc*(j-1)]
		}
		return true
	}
	return false
}

func (g Grid) move(r, c, dr, dc int) bool {
	r = r + dr
	c = c + dc

	if g.grid[r][c] == WALL {
		//--- Can't move a wall, return
		return false
	}

	if g.grid[r][c] == SPACE {
		return true
	}

	if g.grid[r][c] == BOX {
		// narrow boxes can be moved in both directions
		return g.moveBox(r, c, dr, dc)
	}

	return false
}

func (g Grid) MoveRobot(moves string) {
	for _, move := range moves {
		r := g.robot[0]
		c := g.robot[1]
		dr := 0
		dc := 0

		switch move {
		case '^':
			dr = -1
			dc = 0
		case '>':
			dr = 0
			dc = 1
		case 'v':
			dr = 1
			dc = 0
		case '<':
			dr = 0
			dc = -1
		}

		if g.move(r, c, dr, dc) {
			g.updateRobotPos(r+dr, c+dc)
		}

		// fmt.Printf("Move: %v \n", string(move))
		// fmt.Println(g)
	}
}

func (g Grid) CalculateGPS() int {
	total := 0
	for r, line := range g.grid {
		for c, char := range line {
			if char == BOX || char == BOX_L {
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
