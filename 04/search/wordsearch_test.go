package search

import "testing"

func TestSearchHorizontal(t *testing.T) {
	needle := "XMAS"
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{"0 times", []string{"MMMSMMAMXX"}, 0},
		{"1 time", []string{"XMASMMAMXX"}, 1},
		{"2 times", []string{"MXMASXMASX"}, 2},
		{"forward and backward", []string{"MXMASXSAMX"}, 2},
		{"overlap", []string{"XMASAMX"}, 2},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			num := SearchForWord(needle, testCase.input)
			if num != testCase.expected {
				t.Errorf("%v: expected %v but got %v", testCase.name, testCase.expected, num)
			}
		})
	}
}
