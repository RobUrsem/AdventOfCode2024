package numbers

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CalcDistances(a []int, b []int) (int, error) {
	distance := 0

	if len(a) != len(b) {
		return 0, fmt.Errorf("lists must be of equal length")
	}

	//--- Sort both lists
	sort.Ints(a)
	sort.Ints(b)

	for i := range a {
		dist := abs(a[i] - b[i])
		distance = distance + dist
	}

	return distance, nil
}
