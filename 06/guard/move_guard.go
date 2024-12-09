package guard

func leaveMap(labMap LabMap, r, c int, guardPos int) bool {
	switch guardPos {
	case GUARD_DOWN:
		numRows := len(labMap)
		return r == numRows-1
	case GUARD_UP:
		return r == 0
	case GUARD_LEFT:
		return c == 0
	case GUARD_RIGHT:
		numCols := len(labMap[0])
		return c == numCols-1
	default:
		return false
	}
}

type MoveInfo struct {
	inFront int
	r, c    int
}

func getNextMove(guard int, r, c int, labMap LabMap) MoveInfo {
	switch guard {
	case GUARD_DOWN:
		return MoveInfo{
			r:       r + 1,
			c:       c,
			inFront: labMap[r+1][c],
		}
	case GUARD_LEFT:
		return MoveInfo{
			r:       r,
			c:       c - 1,
			inFront: labMap[r][c-1],
		}
	case GUARD_RIGHT:
		return MoveInfo{
			r:       r,
			c:       c + 1,
			inFront: labMap[r][c+1],
		}
	case GUARD_UP:
		return MoveInfo{
			r:       r - 1,
			c:       c,
			inFront: labMap[r-1][c],
		}
	}

	return MoveInfo{}
}

func markPreviousPos(labMap LabMap, r, c int, guardDir int, lastMoveWasTurn bool) {
	if lastMoveWasTurn {
		labMap[r][c] = VISITED_BOTH
	} else {
		switch guardDir {
		case GUARD_DOWN:
			labMap[r][c] = VISITED_VERTICAL
		case GUARD_UP:
			labMap[r][c] = VISITED_VERTICAL
		case GUARD_LEFT:
			labMap[r][c] = VISITED_HORIZONTAL
		case GUARD_RIGHT:
			labMap[r][c] = VISITED_HORIZONTAL
		}
	}
}

func MoveGuard(labMap LabMap, r, c int, lastMoveNeedsCross bool) (bool, bool, bool) {
	guard := labMap[r][c]
	moveInfo := getNextMove(guard, r, c, labMap)

	nextMoveNeedsCross := false
	switch moveInfo.inFront {
	case EMPTY:
		markPreviousPos(labMap, r, c, guard, lastMoveNeedsCross)
		labMap[moveInfo.r][moveInfo.c] = guard
		nextMoveNeedsCross = false
	case OBSTACLE:
		//--- turn right
		switch guard {
		case GUARD_DOWN:
			labMap[r][c] = GUARD_LEFT
		case GUARD_LEFT:
			labMap[r][c] = GUARD_UP
		case GUARD_UP:
			labMap[r][c] = GUARD_RIGHT
		case GUARD_RIGHT:
			labMap[r][c] = GUARD_DOWN
		}
		return leaveMap(labMap, r, c, labMap[r][c]), true, true
	case VISITED: //--- General "visited" items don't capture direction
		markPreviousPos(labMap, r, c, guard, lastMoveNeedsCross)
		labMap[moveInfo.r][moveInfo.c] = guard
		nextMoveNeedsCross = true
	case VISITED_HORIZONTAL: //--- We're crossing a horizontal line
		markPreviousPos(labMap, r, c, guard, lastMoveNeedsCross)
		labMap[moveInfo.r][moveInfo.c] = guard
		nextMoveNeedsCross = true
	case VISITED_VERTICAL: //--- We're crossing a vertical line
		markPreviousPos(labMap, r, c, guard, lastMoveNeedsCross)
		labMap[moveInfo.r][moveInfo.c] = guard
		nextMoveNeedsCross = true
	case VISITED_BOTH: //--- This already is an intersection or a turn
		markPreviousPos(labMap, r, c, guard, lastMoveNeedsCross)
		labMap[moveInfo.r][moveInfo.c] = guard
		nextMoveNeedsCross = false
	}

	return leaveMap(labMap, moveInfo.r, moveInfo.c, guard), nextMoveNeedsCross, false
}
