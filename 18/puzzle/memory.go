package puzzle

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	SPACE = '.'
	WALL  = '#'
	PATH  = 'O'
)

type Location struct {
	R, C int
}

func (l *Location) Add(other Location) {
	l.R += other.R
	l.C += other.C
}

func Combine(a, b Location) Location {
	return Location{a.R + b.R, a.C + b.C}
}

func MakeLocation(line string) Location {
	var location Location
	parts := strings.Split(line, ",")
	if len(parts) == 2 {
		c, _ := strconv.Atoi(parts[0])
		r, _ := strconv.Atoi(parts[1])
		location.C = c
		location.R = r
	}
	return location
}

type Memory struct {
	grid       [][]rune
	bytes      []Location
	start, end Location
	fallen     int
}

func MakeMemory(width, height int) Memory {
	m := Memory{}
	m.grid = make([][]rune, height)
	for r := 0; r < height; r++ {
		m.grid[r] = make([]rune, width)
	}
	m.start = Location{0, 0}
	m.end = Location{height - 1, width - 1}
	return m
}

func (m *Memory) Reset() {
	for r := range m.grid {
		for c := range m.grid[r] {
			m.grid[r][c] = SPACE
		}
	}
}

func (m *Memory) LoadBytes(lines []string) {
	for _, line := range lines {
		m.bytes = append(m.bytes, MakeLocation(line))
	}
}

func (m *Memory) Simulate(numSteps int) {
	m.Reset()
	steps := min(numSteps, len(m.bytes))
	for step := 0; step < steps; step++ {
		m.NextByte()
	}
}

func (m *Memory) NextByte() {
	if m.fallen < len(m.bytes) {
		loc := m.bytes[m.fallen]
		m.grid[loc.R][loc.C] = WALL
		m.fallen++
	}
}

func (m Memory) LastFallenByte() Location {
	return m.bytes[m.fallen-1]
}

func (m *Memory) Print() {
	for _, line := range m.grid {
		for _, char := range line {
			fmt.Printf("%v", string(char))
		}
		fmt.Println()
	}
}
