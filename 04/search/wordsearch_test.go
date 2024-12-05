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

func TestSearchVertical(t *testing.T) {
	needle := "XMAS"
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{"0 times", []string{
			"MXM",
			"XAX",
			"SXM",
			"MSX",
		}, 0},
		{"1 time down", []string{
			".X.",
			".M.",
			".A.",
			".S.",
		}, 1},
		{"1 time up", []string{
			"..S",
			"..A",
			"..M",
			"..X",
		}, 1},
		{"2 times down", []string{
			".X.",
			".M.",
			"XA.",
			"MS.",
			"A..",
			"S..",
		}, 2},
		{"down and up", []string{
			".XS",
			".MA",
			".AM",
			".SX",
		}, 2},
		{"overlap", []string{
			".X.",
			".M.",
			".A.",
			".S.",
			".A.",
			".M.",
			".X.",
		}, 2},
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
