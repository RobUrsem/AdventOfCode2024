package puzzle

import (
	"fmt"
	"strings"
)

type WideGrid struct {
	grid  [][]rune
	robot []int
}

func (g WideGrid) String() string {
	var builder strings.Builder
	for _, row := range g.grid {
		for _, cell := range row {
			builder.WriteString(fmt.Sprintf("%v", string(cell)))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (g WideGrid) updateRobotPos(r, c int) {
	g.grid[g.robot[0]][g.robot[1]] = SPACE
	g.robot[0] = r
	g.robot[1] = c
	g.grid[r][c] = ROBOT
}

func (g WideGrid) moveBoxHorizontally(r, c, dc int) bool {
	//--- See if we can push the box(es)
	//--- find max num boxes to push
	foundSpace := false
	foundWall := false
	moveUntil := 0
	for ; !foundSpace && !foundWall; moveUntil++ {
		if g.grid[r][c+dc*moveUntil] == SPACE {
			foundSpace = true
		}
		if g.grid[r][c+dc*moveUntil] == WALL {
			foundWall = true
		}
	}

	if foundSpace {
		for j := moveUntil - 1; j >= 0; j-- {
			g.grid[r][c+dc*j] = g.grid[r][c+dc*(j-1)]
		}
		return true
	}
	return false
}

func (g WideGrid) canBeMovedInto(r, c, dr int) bool {
	// fmt.Printf("canBeMovedInto(%v,%v,%v) => ", r, c, dr)
	cell := g.grid[r][c]
	switch cell {
	case SPACE:
		// fmt.Println("YES")
		return true
	case WALL:
		// fmt.Println("NO")
		return false
	case BOX_L:
		return g.canBeMovedInto(r+dr, c, dr) && g.canBeMovedInto(r+dr, c+1, dr)
	case BOX_R:
		return g.canBeMovedInto(r+dr, c-1, dr) && g.canBeMovedInto(r+dr, c, dr)
	}
	fmt.Println("HUH?")
	return false
}

func (g WideGrid) canBoxMoveVertically(r, c, dr int) bool {
	// r,c should be BOX_L
	moveable := [2]bool{
		g.canBeMovedInto(r+dr, c, dr),
		g.canBeMovedInto(r+dr, c+1, dr),
	}

	canMove := moveable[0] && moveable[1]
	// fmt.Printf("canBoxMoveVertically(%v,%v,%v) => %v\n", r, c, dr, canMove)
	return canMove
}
func (g WideGrid) moveBoxVertically(r, c, dr int) bool {
	for offset := 0; offset < 2; offset++ {
		dest := g.grid[r+dr][c+offset]
		switch dest {
		case BOX_L:
			g.moveBoxVertically(r+dr, c+offset, dr)
		case BOX_R:
			g.moveBoxVertically(r+dr, c-1+offset, dr)
		}
	}

	g.grid[r+dr][c] = g.grid[r][c]
	g.grid[r+dr][c+1] = g.grid[r][c+1]

	g.grid[r][c] = SPACE
	g.grid[r][c+1] = SPACE

	return true
}

func (g WideGrid) move(r, c, dr, dc int) bool {
	cell := g.grid[r+dr][c+dc]
	if cell == WALL {
		return false
	}
	if cell == SPACE {
		return true
	}

	if dr == 0 { // horizontal move
		return g.moveBoxHorizontally(r, c+dc, dc)
	}

	offset := 0
	if cell == BOX_R {
		offset = -1
	}

	if g.canBoxMoveVertically(r+dr, c+offset, dr) {
		return g.moveBoxVertically(r+dr, c+offset, dr)
	}

	return false
}

func (g WideGrid) MoveRobot(moves string) {
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

func (g WideGrid) CalculateGPS() int {
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

func MakeWideGrid(lines []string) WideGrid {
	grid := WideGrid{}
	for r, line := range lines {
		if len(line) > 0 && line[0] == WALL {
			gridRow := make([]rune, len(line)*2)
			for c, char := range line {
				switch char {
				case WALL:
					gridRow[2*c] = WALL
					gridRow[2*c+1] = WALL
				case SPACE:
					gridRow[2*c] = SPACE
					gridRow[2*c+1] = SPACE
				case BOX:
					gridRow[2*c] = BOX_L
					gridRow[2*c+1] = BOX_R
				case ROBOT:
					gridRow[2*c] = ROBOT
					gridRow[2*c+1] = SPACE
				}
				if char == ROBOT {
					grid.robot = []int{r, 2 * c}
				}
			}
			grid.grid = append(grid.grid, gridRow)
		}
	}
	return grid
}
