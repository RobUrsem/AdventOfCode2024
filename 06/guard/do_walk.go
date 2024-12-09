package guard

import (
	"errors"
	"fmt"
)

var ErrOutOfMoves = errors.New("out of moves")
var ErrNoGuard = errors.New("no guard")

type Turn []int
type TurnLocations []Turn

func DoWalk(labMap LabMap) (TurnLocations, error) {
	guardLeaves := false
	lastMoveWasTurn := false

	var turns TurnLocations

	const MAX_MOVES = 25000
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
			turns = append(turns, Turn{r, c})
		}
		moveCounter++

		if moveCounter > MAX_MOVES {
			break
		}
	}

	if moveCounter > MAX_MOVES {
		return nil, ErrOutOfMoves
	}

	fmt.Printf("Turns\n")
	for _, turn := range turns {
		fmt.Printf("turn: (%v, %v)\n", turn[0], turn[1])
	}

	return turns, nil
}
