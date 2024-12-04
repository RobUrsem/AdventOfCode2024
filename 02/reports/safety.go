package reports

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getDistance(a int, b int) int {
	return abs(a - b)
}

func getDirection(a int, b int) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func hasDistanceViolation(a int, b int) bool {
	dist := getDistance(a, b)
	return dist < 1 || dist > 3
}

func hasDirectionViolation(a int, b int, direction *int) bool {
	if *direction == 0 {
		*direction = getDirection(a, b)
	} else if *direction != getDirection(a, b) {
		return true
	}
	return false
}

func removeElement(arr []int, elm int) ([]int, error) {
	if elm < 0 || elm >= len(arr) {
		return nil, fmt.Errorf("illegal index %v", elm)
	}

	copySlice := make([]int, len(arr))
	copy(copySlice, arr)

	if elm == len(copySlice)-1 {
		return copySlice[:elm], nil
	}

	return append(copySlice[:elm], copySlice[elm+1:]...), nil
}

func reportIsSafe(report []int) (bool, int) {
	direction := 0

	for i := 0; i < len(report)-1; i++ {
		if hasDistanceViolation(report[i], report[i+1]) ||
			hasDirectionViolation(report[i], report[i+1], &direction) {
			return false, i
		}
	}

	return true, -1
}

func FindSafeReports(reports [][]int) (int, error) {
	numSafeReports := 0

	for _, report := range reports {
		safe, idx := reportIsSafe(report)
		if safe {
			// fmt.Printf("%v : Safe without removing\n", report)
			numSafeReports++
		} else {
			isSafe := false
			for offset := -1; offset <= 1; offset++ {
				retry, err := removeElement(report, idx+offset)
				if safe, _ := reportIsSafe(retry); err == nil && safe {
					// fmt.Printf("%v : Safe by removing elm %v\n", report, idx+offset+1)
					isSafe = true
					break
				}
			}

			if isSafe {
				numSafeReports++
				// } else {
				// 	fmt.Printf("%v : Unsafe\n", report)
			}
		}
	}

	return numSafeReports, nil
}
