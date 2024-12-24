package puzzle

import (
	"container/heap"
	"fmt"
	"math"
)

type Point struct {
	r, c int
}

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

type Node struct {
	point     Point
	cost      int
	direction int // 0=up, 1=right, 2=down, 3=left
	index     int // index in the priority queue
}

type PriorityQueue []*Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := x.(*Node)
	n.index = len(*pq)
	*pq = append(*pq, n)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := old[len(old)-1]
	n.index = -1 // for safety
	*pq = old[:len(old)-1]
	return n
}

func (m Maze) isSafe(x, y int) bool {
	return x >= 0 && x < len(m.grid) &&
		y >= 0 && y < len(m.grid[0]) &&
		m.grid[x][y] != WALL && m.visited[x][y] == NOPE
}

// Directions: up, right, down, left
var directions = []Point{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (m *Maze) SolveMaze() ([]Point, int, int) {
	start := Point{m.start[0], m.start[1]}
	goal := Point{m.end[0], m.end[1]}

	numRows, numCols := len(m.grid), len(m.grid[0])
	m.visited = make([][]rune, numRows)
	costs := make([][]int, numRows)
	for i := range costs {
		costs[i] = make([]int, numCols)
		m.visited[i] = make([]rune, numCols)
		for j := range costs[i] {
			costs[i][j] = math.MaxInt
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	// Initialize start point
	for dir := 0; dir < 4; dir++ {
		// As per directions we start moving to the east
		startNode := &Node{point: start, cost: 0, direction: RIGHT}
		heap.Push(pq, startNode)
		costs[start.r][start.c] = 0
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		curPoint, curCost, curDir := current.point, current.cost, current.direction

		if curPoint == goal {
			BlockNonOptimalPaths(costs, curCost)
			// PrintCosts(costs)
			numSeats := m.FindBestSeats(costs, Point{m.end[0], m.end[1]})
			// m.Print()
			// fmt.Printf("Num best seats: %v\n", numSeats)
			return m.reconstructPath(costs), curCost, numSeats
		}

		// Explore neighbors
		for newDir, move := range directions {
			newPoint := Point{curPoint.r + move.r, curPoint.c + move.c}
			if !m.isSafe(newPoint.r, newPoint.c) {
				continue
			}

			turnCost := 0
			if curDir != newDir {
				turnCost = 1000
			}

			newCost := curCost + 1 + turnCost
			if newCost < costs[newPoint.r][newPoint.c] {
				costs[newPoint.r][newPoint.c] = newCost
				heap.Push(pq, &Node{point: newPoint, cost: newCost, direction: newDir})
			}
		}
	}

	return nil, -1, 0 // No path found
}

func (m *Maze) reconstructPath(costs [][]int) []Point {
	end := Point{m.end[0], m.end[1]}
	path := []Point{}

	start := Point{m.start[0], m.start[1]}
	for current := end; current != start; {
		path = append(path, current)
		currentCost := costs[current.r][current.c]
		nextDir := 0
		for newDir, move := range directions {
			newPoint := Point{current.r + move.r, current.c + move.c}
			newCost := costs[newPoint.r][newPoint.c]
			if newCost < currentCost {
				currentCost = newCost
				nextDir = newDir
			}
		}
		current = Point{current.r + directions[nextDir].r, current.c + directions[nextDir].c}
		m.visited[current.r][current.c] = '+'
	}

	return path
}

func BlockNonOptimalPaths(costs [][]int, maxCost int) {
	for r, line := range costs {
		for c, cost := range line {
			if cost > maxCost && cost != math.MaxInt {
				// fmt.Printf("Blocking %v,%v\n", r, c)
				costs[r][c] = math.MaxInt
			}
		}
	}
}

func DecodeCost(cost int) (int, int) {
	turns := cost / 1000
	steps := cost - turns*1000
	return turns, steps
}

func (m Maze) FindBestSeats(costs [][]int, start Point) int {
	// fmt.Printf("FindBestSeats(%v)\n", start)
	end := Point{m.start[0], m.start[1]}

	totalSeats := 0
	for current := start; current != end; {
		paths := []Point{}
		if m.visited[current.r][current.c] != BEST {
			currentCost := costs[current.r][current.c]
			// fmt.Printf("Best seat: %v\n", current)
			m.visited[current.r][current.c] = BEST
			totalSeats++
			_, currentSteps := DecodeCost(currentCost)

			for _, move := range directions {
				newPoint := Point{current.r + move.r, current.c + move.c}
				newCost := costs[newPoint.r][newPoint.c]

				if newCost != math.MaxInt &&
					m.visited[newPoint.r][newPoint.c] != BEST {
					_, steps := DecodeCost(newCost)
					if steps < currentSteps {
						paths = append(paths, newPoint)
					}
				}
			}
		}

		if len(paths) > 1 {
			for _, point := range paths {
				totalSeats += m.FindBestSeats(costs, point)
			}
		} else if len(paths) > 0 {
			current = paths[0]
			if current == end {
				totalSeats++
			}
		} else {
			break
		}
	}

	return totalSeats
}

func PrintCosts(costs [][]int) {
	for _, line := range costs {
		for _, cost := range line {
			if cost == math.MaxInt {
				fmt.Print("#,")
			} else if cost < 0 {
				fmt.Print("O,")
			} else {
				fmt.Printf("%5d,", cost)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
