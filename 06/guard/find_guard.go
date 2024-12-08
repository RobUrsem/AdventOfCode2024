package guard

func FindGuard(labMap LabMap) (int, int) {
	for r, line := range labMap {
		for c, char := range line {
			if char == GUARD_DOWN || char == GUARD_LEFT || char == GUARD_RIGHT || char == GUARD_UP {
				return r, c
			}
		}
	}

	return -1, -1
}
