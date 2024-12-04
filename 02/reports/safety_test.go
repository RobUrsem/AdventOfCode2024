package reports

import "testing"

func TestFindSafeReports(t *testing.T) {
	input := [][]int{
		{7, 6, 4, 2, 1},          //- Safe without removing any elm
		{9, 7, 6, 2, 1},          //- Unsafe regardless of what elm is removed
		{10, 13, 10, 7, 4, 3, 1}, //- Safe by removing the 1st elm, 1.
		{1, 3, 2, 4, 5},          //- Safe by removing the 2nd elm, 3.
		{8, 6, 4, 4, 1},          //- Safe by removing the 3rd elm, 4.
		{7, 6, 4, 7, 1},          //- Safe by removing the 4th elm, 7
		{7, 6, 4, 2, 8},          //- Safe by removing the 5th elm, 8
	}

	numSafeReports, err := FindSafeReports(input)
	if err != nil {
		t.Fatalf("Error running safety report: %v", err)
	}

	numExpected := 6
	if numSafeReports != numExpected {
		t.Fatalf("Expected %v safe reports, but got %v", numExpected, numSafeReports)
	}
}

func slicesEqual(a, b []int) bool {
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

func TestRemoveElement(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		remove   int
		expected []int
	}{
		{"negative", []int{1, 3, 2, 4, 5}, -1, nil},
		{"1", []int{1, 3, 2, 4, 5}, 0, []int{3, 2, 4, 5}},
		{"2", []int{1, 3, 2, 4, 5}, 1, []int{1, 2, 4, 5}},
		{"3", []int{1, 3, 2, 4, 5}, 2, []int{1, 3, 4, 5}},
		{"4", []int{1, 3, 2, 4, 5}, 3, []int{1, 3, 2, 5}},
		{"5", []int{1, 3, 2, 4, 5}, 4, []int{1, 3, 2, 4}},
		{"outside", []int{1, 3, 2, 4, 5}, 5, nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			output, _ := removeElement(tc.input, tc.remove)

			if !slicesEqual(output, tc.expected) {
				t.Errorf("Case %v failed", tc.name)
			}
		})
	}
}
