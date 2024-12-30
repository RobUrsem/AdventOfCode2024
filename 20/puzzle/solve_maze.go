package puzzle

import (
	"container/heap"
	"fmt"
	"math"
	"sort"
)

const (
	UP int = iota
	RIGHT
	DOWN
	LEFT
)

type Node struct {
	Location  Location
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

func (m Maze) isSafe(p Location) bool {
	return p.R >= 0 && p.R < len(m.grid) &&
		p.C >= 0 && p.C < len(m.grid[0]) &&
		m.grid[p.R][p.C] != WALL &&
		m.visited[p.R][p.C] == NOPE
}

func (m Maze) GetCheatLocations() []Location {
	locations := []Location{}
	for r, line := range m.grid {
		for c, char := range line {
			if r > 0 && r < len(m.grid) &&
				c > 0 && c < len(m.grid[0]) &&
				char == WALL {
				locations = append(locations, Location{r, c})
			}
		}
	}
	return locations
}

func (m Maze) GetLongCheatLocations() []Location {
	locations := []Location{}
	for r, line := range m.grid {
		for c, char := range line {
			if r > 0 && r < len(m.grid) &&
				c > 0 && c < len(m.grid[0]) &&
				char != WALL {
				locations = append(locations, Location{r, c})
			}
		}
	}
	return locations
}

func (m *Maze) Part1(minSaving int) int {
	numCheats := 0

	m.SolveMaze()
	cheats := m.GetCheatLocations()
	for _, cheat := range cheats {
		saving := m.AddCheat(cheat)
		if saving >= minSaving {
			numCheats++
		}
	}
	return numCheats
}

func sortedKeys(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	// Sort the keys
	sort.Ints(keys)
	return keys
}

func (m *Maze) Part2(minSaving int) int {
	m.SolveMaze()
	cheats := m.GetLongCheatLocations()
	results := map[string]int{}
	for _, cheat := range cheats {
		m.AddLongerCheat(cheat, results)
	}

	cheatsByLength := map[int]int{}
	for _, v := range results {
		cheatsByLength[v]++
	}

	numCheats := 0
	for _, k := range sortedKeys(cheatsByLength) {
		if k >= minSaving {
			// fmt.Printf("%6d cheats that save %4d steps\n", cheatsByLength[k], k)
			numCheats += cheatsByLength[k]
		}
	}

	return numCheats
}

// Directions: up, right, down, left
var directions = []Location{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func (m *Maze) SolveMaze() int {
	start := m.start
	goal := m.end

	numRows, numCols := len(m.grid), len(m.grid[0])
	m.visited = make([][]rune, numRows)
	m.costs = make([][]int, numRows)
	for i := range m.costs {
		m.costs[i] = make([]int, numCols)
		m.visited[i] = make([]rune, numCols)
		for j := range m.costs[i] {
			m.costs[i][j] = math.MaxInt
		}
	}

	pq := &PriorityQueue{}
	heap.Init(pq)

	// Initialize start Location
	for dir := 0; dir < 4; dir++ {
		startNode := &Node{Location: start, cost: 0, direction: dir}
		heap.Push(pq, startNode)
		m.costs[start.R][start.C] = 0
	}

	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Node)
		curLocation, curCost := current.Location, current.cost

		if curLocation == goal {
			return curCost
		}

		// Explore neighbors
		for newDir, move := range directions {
			newLocation := Combine(curLocation, move)
			if !m.isSafe(newLocation) {
				continue
			}

			newCost := curCost + 1
			if newCost < m.costs[newLocation.R][newLocation.C] {
				m.costs[newLocation.R][newLocation.C] = newCost
				heap.Push(pq, &Node{Location: newLocation, cost: newCost, direction: newDir})
			}
		}
	}

	return -1 // No path found
}

func (m *Maze) ReconstructPath(costs [][]int) []Location {
	end := m.end
	path := []Location{}

	start := m.start
	for current := end; current != start; {
		path = append(path, current)
		currentCost := costs[current.R][current.C]
		nextDir := 0
		for newDir, move := range directions {
			newLocation := Combine(current, move)
			newCost := costs[newLocation.R][newLocation.C]
			if newCost < currentCost {
				currentCost = newCost
				nextDir = newDir
			}
		}
		current = Combine(current, directions[nextDir])
		m.visited[current.R][current.C] = '+'
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

func PrintCosts(costs [][]int) {
	for _, line := range costs {
		for _, cost := range line {
			if cost == math.MaxInt {
				fmt.Print("#####,")
			} else if cost < 0 {
				fmt.Print("  O  ,")
			} else {
				fmt.Printf("%5d,", cost)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}
