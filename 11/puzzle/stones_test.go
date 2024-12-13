package puzzle

import (
	"advent/shared"
	"fmt"
	"testing"
)

func TestStones(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "rule 1",
			input: []int{0},
			expected: [][]int{
				{1},
			},
		},
		{
			name:  "rule 2",
			input: []int{12},
			expected: [][]int{
				{1, 2},
			},
		},
		{
			name:  "rule 2 - leading zeros",
			input: []int{1000},
			expected: [][]int{
				{10, 0},
			},
		},
		{
			name:  "1",
			input: []int{125, 17},
			expected: [][]int{
				{253000, 1, 7},
				{253, 0, 2024, 14168},
				{512072, 1, 20, 24, 28676032},
				{512, 72, 2024, 2, 0, 2, 4, 2867, 6032},
				{1036288, 7, 2, 20, 24, 4048, 1, 4048, 8096, 28, 67, 60, 32},
				{2097446912, 14168, 4048, 2, 0, 2, 4, 40, 48, 2024, 40, 48, 80, 96, 2, 8, 6, 7, 6, 0, 3, 2},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			outcome := tc.input
			for i, blink := range tc.expected {
				outcome = Blink(outcome)
				fmt.Printf("Got %v stones\n", len(outcome))
				if !shared.AreEqualInts(outcome, blink) {
					t.Errorf("Expected %v for blink %v but got %v", blink, i, outcome)
				}
			}
		})
	}
}
