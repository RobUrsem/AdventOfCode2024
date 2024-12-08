package guard

import "testing"

func TestMoveGuard(t *testing.T) {
	testCases := []struct {
		name             string
		input            LabMap
		start_r, start_c int
		expected         LabMap
		left             bool
	}{
		{
			"Leave on up",
			LabMap{
				[]int{EMPTY},
				[]int{GUARD_UP},
			},
			1, 0,
			LabMap{
				[]int{GUARD_UP},
				[]int{VISITED},
			},
			true,
		},
		{
			"Leave on right",
			LabMap{
				[]int{GUARD_RIGHT, EMPTY},
			},
			0, 0,
			LabMap{
				[]int{VISITED, GUARD_RIGHT},
			},
			true,
		},
		{
			"Leave on down",
			LabMap{
				[]int{GUARD_DOWN},
				[]int{EMPTY},
			},
			0, 0,
			LabMap{
				[]int{VISITED},
				[]int{GUARD_DOWN},
			},
			true,
		},
		{
			"Leave on left",
			LabMap{
				[]int{EMPTY, GUARD_LEFT},
			},
			0, 1,
			LabMap{
				[]int{GUARD_LEFT, VISITED},
			},
			true,
		},
		{
			"Turn up to right and leave",
			LabMap{
				[]int{EMPTY, EMPTY},
				[]int{EMPTY, OBSTACLE},
				[]int{EMPTY, GUARD_UP},
			},
			2, 1,
			LabMap{
				[]int{EMPTY, EMPTY},
				[]int{EMPTY, OBSTACLE},
				[]int{EMPTY, GUARD_RIGHT},
			},
			true,
		},
		{
			"Turn up to right and not leave",
			LabMap{
				[]int{EMPTY, EMPTY},
				[]int{OBSTACLE, EMPTY},
				[]int{GUARD_UP, EMPTY},
			},
			2, 0,
			LabMap{
				[]int{EMPTY, EMPTY},
				[]int{OBSTACLE, EMPTY},
				[]int{GUARD_RIGHT, EMPTY},
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r, c := FindGuard(tc.input)
			if r != tc.start_r || c != tc.start_c {
				t.Errorf("Expected start pos to be (%v,%v) but got (%v,%v)", tc.start_r, tc.start_c, r, c)
			}

			leavesMap := MoveGuard(tc.input, r, c)
			if leavesMap != tc.left {
				t.Errorf("Expected leave map to be: %v but got %v", tc.left, leavesMap)
			}

			if !areEqual(tc.expected, tc.input) {
				t.Errorf("Move not correct")
			}
		})
	}
}
