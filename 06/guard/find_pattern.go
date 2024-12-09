package guard

import "sort"

//--- The pattern is a series of 3 obstacles
//--- for which there is a square of paths
// .#...
// .+-+#
// .|.|.
// O--+.
// ...#.
// Adding the 'O' at 3,0 makes the pattern
// an infinite loop
type Obstruction [2]int // r,c for obstruction

func isCorner(labMap LabMap, r, c int) bool {
	numRows := len(labMap)
	numCols := len(labMap[0])

	if r > 0 && labMap[r-1][c] == OBSTACLE {
		return true
	}
	if c > 0 && labMap[r][c-1] == OBSTACLE {
		return true
	}
	if r < numRows && labMap[r+1][c] == OBSTACLE {
		return true
	}
	if c < numCols && labMap[r][c+1] == OBSTACLE {
		return true
	}
	return false
}

func findCorner(corners []Obstruction, r, c int) bool {
	for _, corner := range corners {
		if corner[0] == r && corner[1] == c {
			return true
		}
	}

	return false
}

func createObstruction(labMap LabMap, r, c int) Obstruction {

	/*---
	After 3 turns, is there a location that would turn back to the first turn?

	circle 1
	turns: [(1,1), (1,3), (3,3)] => next turn should be (3,1)
	say (3,1) is missing turn
	We came from (3,3) so travel is horizontal to the left
	Obstacle should be placed at (3,0)

	circle 2a
	turns: [(1,3), (3,3), (3,1)] => next turn should be (1,1)
	We came from (3,1) so travel is vertical to the top
	Obstacle should be placed at (0,1)

	---*/
	return Obstruction{3, 0}
}

func turnsFormBox(a, b, c Turn) (bool, Turn) {
	/*
			     1         2           3         4
			----------+----------+---------------------
		      a    b  |       a  |  c       |  b    c
		              |          |          |
		           c  |  c    b  |  b    a  |  a
	*/

	//--- Case 1 & 3: D = (c0, a1)
	if a[0] == b[0] && b[1] == c[1] {
		return true, Turn{c[0], a[1]}
	}

	//--- Case 2 & 4: D = (a0, c1)
	if a[1] == b[1] && b[0] == c[0] {
		return true, Turn{a[0], c[1]}
	}

	return false, nil
}

func getDirection(a, b Turn) []int {
	dr := 0
	if b[0] < a[0] {
		dr = -1
	} else if b[0] > a[0] {
		dr = 1
	}

	dc := 0
	if b[1] < a[1] {
		dc = -1
	} else if b[1] > a[1] {
		dc = 1
	}

	return []int{dr, dc}
}

func isVisited(labMap LabMap, pos []int) bool {
	switch labMap[pos[0]][pos[1]] {
	case VISITED:
		return true
	case VISITED_HORIZONTAL:
		return true
	case VISITED_VERTICAL:
		return true
	case VISITED_BOTH:
		return true
	}
	return false
}

func obstaclesInTheWay(labmap LabMap, a, b []int) bool {
	direction := getDirection(a, b)

	if direction[0] != 0 {
		for r := a[0]; r != b[0]; r += direction[0] {
			if labmap[r][a[1]] == OBSTACLE {
				return true
			}
		}
		return false
	} else {
		for c := a[1]; c != b[1]; c += direction[1] {
			if labmap[a[0]][c] == OBSTACLE {
				return true
			}
		}
		return false
	}
}

func sortObstructions(obs []Obstruction) {
	sort.Slice(obs, func(i, j int) bool {
		if obs[i][0] == obs[j][0] {
			return obs[i][1] < obs[j][1]
		}
		return obs[i][0] < obs[j][0]
	})
}

func FindPattern(labMap LabMap, turns TurnLocations) []Obstruction {
	var obstructions []Obstruction

	for i := 0; i < len(turns)-2; i++ {
		if ok, missing := turnsFormBox(turns[i], turns[i+1], turns[i+2]); ok {
			if isVisited(labMap, missing) && !obstaclesInTheWay(labMap, turns[i+2], missing) {
				direction := getDirection(turns[i+2], missing)
				obstructions = append(obstructions,
					Obstruction{missing[0] + direction[0], missing[1] + direction[1]})
			}
		}
	}

	sortObstructions(obstructions)
	return obstructions
}
