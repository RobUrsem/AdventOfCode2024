package puzzle

import (
	"fmt"
	"strings"
)

type Bathroom struct {
	size   []int
	robots Robots
}

func NewBathroom(numRows, numCols int) Bathroom {
	return Bathroom{
		size: []int{numRows, numCols},
	}
}

func (b Bathroom) wrap(r Robot, idx int) {
	if r.Pos[idx] < 0 {
		r.Pos[idx] += b.size[idx]
	}
	if r.Pos[idx] >= b.size[idx] {
		r.Pos[idx] -= b.size[idx]
	}
}

func (b *Bathroom) AddRobots(robots Robots) {
	b.robots = robots
}

func (b Bathroom) Tick() {
	for _, robot := range b.robots {
		robot.Move()
		b.wrap(robot, 0)
		b.wrap(robot, 1)
	}
}

func (b Bathroom) GetTiles() [][]int {
	tiles := make([][]int, b.size[0])
	for r := 0; r < b.size[0]; r++ {
		tiles[r] = make([]int, b.size[1])
	}

	for _, robot := range b.robots {
		tiles[robot.Pos[0]][robot.Pos[1]]++
	}

	return tiles
}

func (b Bathroom) String() string {
	tiles := b.GetTiles()
	var builder strings.Builder
	for _, row := range tiles {
		for _, cell := range row {
			if cell == 0 {
				builder.WriteString(" ")
			} else {
				builder.WriteString(fmt.Sprintf("%d", cell))
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (b Bathroom) SafetyScore() int {
	tiles := b.GetTiles()

	halfRow := b.size[0] / 2
	halfCol := b.size[1] / 2

	// Sum each quadrant
	quadrants := make([]int, 4)

	//Q1
	for r := 0; r < halfRow; r++ {
		for c := 0; c < halfCol; c++ {
			quadrants[0] += tiles[r][c]
		}
	}
	//Q2
	for r := 0; r < halfRow; r++ {
		for c := halfCol + 1; c < b.size[1]; c++ {
			quadrants[1] += tiles[r][c]
		}
	}
	//Q3
	for r := halfRow + 1; r < b.size[0]; r++ {
		for c := 0; c < halfCol; c++ {
			quadrants[2] += tiles[r][c]
		}
	}
	//Q4
	for r := halfRow + 1; r < b.size[0]; r++ {
		for c := halfCol + 1; c < b.size[1]; c++ {
			quadrants[3] += tiles[r][c]
		}
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}
