package guard

import (
	"fmt"
)

func filterLocations(base Locations, toRemove Locations) Locations {
	// Create a map to track items to be removed
	removeMap := make(map[string]struct{})
	for _, item := range toRemove {
		key := fmt.Sprintf("%v", item) // Use a string representation of the slice as a key
		removeMap[key] = struct{}{}
	}

	// Filter the base slice
	result := Locations{}
	for _, item := range base {
		key := fmt.Sprintf("%v", item)
		if _, exists := removeMap[key]; !exists {
			result = append(result, item) // Only add if not in removeMap
		}
	}

	return result
}

func BruteForceObstructions(labMap LabMap) Locations {
	copyMap := CopyMap(labMap)

	turns, _ := DoWalk(labMap)
	visited := FindVisited(labMap)

	filterLocations(visited, turns)

	var obstructions Locations
	for _, loc := range visited {
		theMap := CopyMap(copyMap)
		theMap[loc[0]][loc[1]] = OBSTRUCTION
		_, err := DoWalk(theMap)
		if err == ErrInfiniteLoop {
			obstructions = append(obstructions, loc)
		}
	}

	return obstructions
}
