package guard

import (
	"errors"
)

var ErrOutOfMoves = errors.New("out of moves")
var ErrNoGuard = errors.New("no guard")
var ErrInfiniteLoop = errors.New("infinite loop")

type Location []int
type Locations []Location

func locationsEqual(a, b Location) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Function to check if a Location exists in Locations
func containsLocation(locations Locations, loc Location) bool {
	for _, existing := range locations {
		if locationsEqual(existing, loc) {
			return true
		}
	}
	return false
}

func DoWalk(labMap LabMap) (Locations, error) {
	guardLeaves := false
	lastMoveWasTurn := false

	var turns Locations

	numDuplicateTurns := 0
	maxMoves := len(labMap) * len(labMap[0])
	moveCounter := 0
	for {
		r, c := FindGuard(labMap)
		if r == -1 || c == -1 {
			return nil, ErrNoGuard
		}

		guardLeaves, _, lastMoveWasTurn = MoveGuard(labMap, r, c, lastMoveWasTurn)
		if guardLeaves {
			break
		}

		if lastMoveWasTurn {
			turn := Location{r, c}
			if containsLocation(turns, turn) {
				numDuplicateTurns++
				if numDuplicateTurns == 4 {
					return nil, ErrInfiniteLoop
				}
			}
			turns = append(turns, turn)
		}
		moveCounter++

		if moveCounter > maxMoves {
			break
		}
	}

	if moveCounter > maxMoves {
		return nil, ErrOutOfMoves
	}

	// fmt.Printf("Turns\n")
	// for _, turn := range turns {
	// 	fmt.Printf("turn: (%v, %v)\n", turn[0], turn[1])
	// }

	return turns, nil
}
