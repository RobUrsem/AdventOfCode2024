package puzzle

import (
	"container/heap"
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

func (m Maze) SolveMaze() ([]Point, int) {
	start := Point{m.start[0], m.start[1]}
	goal := Point{m.end[0], m.end[1]}

	numRows, numCols := len(m.grid), len(m.grid[0])
	costs := make([][]int, numRows)
	for i := range costs {
		costs[i] = make([]int, numCols)
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
			return m.reconstructPath(costs), curCost
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

	return nil, -1 // No path found
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
