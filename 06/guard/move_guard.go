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

func MoveGuard(labMap LabMap, r, c int) bool {

	var moveInfo MoveInfo
	guard := labMap[r][c]
	switch guard {
	case GUARD_DOWN:
		moveInfo = MoveInfo{
			r:       r + 1,
			c:       c,
			inFront: labMap[r+1][c],
		}
	case GUARD_LEFT:
		moveInfo = MoveInfo{
			r:       r,
			c:       c - 1,
			inFront: labMap[r][c-1],
		}
	case GUARD_RIGHT:
		moveInfo = MoveInfo{
			r:       r,
			c:       c + 1,
			inFront: labMap[r][c+1],
		}
	case GUARD_UP:
		moveInfo = MoveInfo{
			r:       r - 1,
			c:       c,
			inFront: labMap[r-1][c],
		}
	}

	switch moveInfo.inFront {
	case EMPTY:
		if labMap[r][c] == EMPTY {
			labMap[r][c] = VISITED
		}
		labMap[moveInfo.r][moveInfo.c] = guard
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
	case VISITED:
		labMap[moveInfo.r][moveInfo.c] = guard
	}

	return leaveMap(labMap, moveInfo.r, moveInfo.c, guard)
}
