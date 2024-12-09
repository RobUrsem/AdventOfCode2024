package guard

func CountVisited(labMap LabMap) int {
	count := 0

	for _, row := range labMap {
		for _, char := range row {
			if char == VISITED || char == VISITED_BOTH || char == VISITED_HORIZONTAL || char == VISITED_VERTICAL {
				count++
			}
		}
	}

	return count
}
