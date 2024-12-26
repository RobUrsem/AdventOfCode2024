package puzzle

import (
	"container/heap"
	"fmt"
	"math"
)

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

type Node struct {
	point Location
	cost  int
	index int // index in the priority queue
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

func (m Memory) isSafe(loc Location) bool {
	return loc.R >= 0 && loc.R < len(m.grid) &&
		loc.C >= 0 && loc.C < len(m.grid[0]) &&
		m.grid[loc.R][loc.C] != WALL
}

// Directions: up, right, down, left
var directions = []Location{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (m *Memory) SolveMaze() ([]Location, int) {
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
		startNode := &Node{point: m.start, cost: 0}
		heap.Push(pq, startNode)
		costs[m.start.R][m.start.C] = 0
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		// fmt.Printf("Current: %v -> \n", current)
		curPoint, curCost := current.point, current.cost

		if curPoint == m.end {
			// PrintCosts(costs)
			// m.Print()
			return m.reconstructPath(costs), curCost
		}

		// Explore neighbors
		for _, move := range directions {
			newPoint := Combine(curPoint, move)
			if !m.isSafe(newPoint) {
				// fmt.Printf("\t%v not safe\n", newPoint)
				continue
			}

			newCost := curCost + 1
			if newCost < costs[newPoint.R][newPoint.C] {
				costs[newPoint.R][newPoint.C] = newCost
				// fmt.Printf("\t%v %v\n", newPoint, newCost)
				heap.Push(pq, &Node{
					point: newPoint,
					cost:  newCost,
				})
				// } else {
				// 	fmt.Printf("\t%v visited\n", newPoint)
			}
		}
	}

	return nil, -1 // No path found
}

func (m *Memory) reconstructPath(costs [][]int) []Location {
	path := []Location{}

	for current := m.end; current != m.start; {
		path = append(path, current)
		currentCost := costs[current.R][current.C]
		nextDir := 0
		for newDir, move := range directions {
			newPoint := Combine(current, move)
			if !m.isSafe(newPoint) {
				continue
			}

			newCost := costs[newPoint.R][newPoint.C]
			if newCost < currentCost {
				currentCost = newCost
				nextDir = newDir
			}
		}
		current = Combine(current, directions[nextDir])
	}

	for _, p := range path {
		m.grid[p.R][p.C] = PATH
	}

	return path
}

func PrintCosts(costs [][]int) {
	for _, line := range costs {
		for _, cost := range line {
			if cost == math.MaxInt {
				fmt.Print("  #,")
			} else if cost < 0 {
				fmt.Print("  O,")
			} else {
				fmt.Printf("%3d,", cost)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
