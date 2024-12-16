package puzzle

import (
	"fmt"
	"testing"
)

func TestRobot(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		numticks int
		expected [][]int
		safety   int
	}{
		{
			name: "single",
			input: []string{
				"p=2,4 v=2,-3",
			},
			numticks: 4,
			expected: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			},
			safety: 0,
		},
		{
			name: "multiple",
			input: []string{
				"p=0,4 v=3,-3",
				"p=6,3 v=-1,-3",
				"p=10,3 v=-1,2",
				"p=2,0 v=2,-1",
				"p=0,0 v=1,3",
				"p=3,0 v=-2,-2",
				"p=7,6 v=-1,-3",
				"p=3,0 v=-1,-2",
				"p=9,3 v=2,3",
				"p=7,3 v=-1,2",
				"p=2,4 v=2,-3",
				"p=9,5 v=-3,-3",
			},
			numticks: 100,
			expected: [][]int{
				{0, 0, 0, 0, 0, 0, 2, 0, 0, 1, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 2, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0},
			},
			safety: 12,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			robots := GetInput(tc.input)

			bathRoom := NewBathroom(7, 11)
			bathRoom.AddRobots(robots)
			// fmt.Println(bathRoom)

			for tick := 0; tick < tc.numticks; tick++ {
				bathRoom.Tick()
				// fmt.Println(bathRoom)
			}

			fmt.Println(bathRoom)
			safety := bathRoom.SafetyScore()
			if safety != tc.safety {
				t.Errorf("Expected safety score of %v but got %v", tc.safety, safety)
			}
		})
	}
}
