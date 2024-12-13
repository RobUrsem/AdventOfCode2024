package shared

import (
	"fmt"
	"strings"
)

type Location struct {
	R, C int
}
type Locations []Location

func NewLocation(r, c int) Location {
	return Location{R: r, C: c}
}

func AreEqual(a, b Locations) bool {
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

func (locs Locations) Copy() Locations {
	copyLocs := make(Locations, len(locs))
	copy(copyLocs, locs)
	return copyLocs
}

func (locs Locations) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, loc := range locs {
		sb.WriteString(fmt.Sprintf("%v", loc))
		if i < len(locs)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func RemoveDuplicates(locations Locations) Locations {
	uniqueMap := make(map[string]struct{})
	var uniqueLocations Locations

	for _, loc := range locations {
		locKey := fmt.Sprint(loc)
		if _, exists := uniqueMap[locKey]; !exists {
			uniqueMap[locKey] = struct{}{}
			uniqueLocations = append(uniqueLocations, loc)
		}
	}

	// if len(locations) != len(uniqueLocations) {
	// 	fmt.Printf("Orig: %v, unique %v\n", len(locations), len(uniqueLocations))
	// }

	return uniqueLocations
}
