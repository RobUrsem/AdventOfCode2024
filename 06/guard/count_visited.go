package guard

func FindVisited(labMap LabMap) Locations {
	var visited Locations

	for r, row := range labMap {
		for c, char := range row {
			if char == VISITED ||
				char == VISITED_BOTH ||
				char == VISITED_HORIZONTAL ||
				char == VISITED_VERTICAL ||
				char == GUARD_DOWN ||
				char == GUARD_LEFT ||
				char == GUARD_RIGHT ||
				char == GUARD_UP {
				visited = append(visited, Location{r, c})
			}
		}
	}

	return visited
}

func CountVisited(labMap LabMap) int {
	visited := FindVisited(labMap)
	return len(visited)
}
