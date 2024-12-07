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
		{"1 time down with false start", []string{
			".X.",
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
		{"1 time up with false start", []string{
			"..S",
			"..A",
			"..M",
			"..X",
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

func TestSearchDiagonal(t *testing.T) {
	needle := "XMAS"
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{"0 times", []string{
			"MXMah",
			"XAXbg",
			"SXMcf",
			"MSXde",
		}, 0},
		{"1 time down NW-SE", []string{
			"X...",
			".M..",
			"..A.",
			"...S",
		}, 1},
		{"1 time down NE-SW", []string{
			"...X",
			"..M.",
			".A..",
			"S...",
		}, 1},
		{"1 time up SW-NE", []string{
			"...S",
			"..A.",
			".M..",
			"X...",
		}, 1},
		{"1 time up SE-NW", []string{
			"S...",
			".A..",
			"..M.",
			"...X",
		}, 1},
		{"2 times down", []string{
			"X....",
			".M...",
			"X.A..",
			".M.S.",
			"..A..",
			"...S.",
		}, 2},
		{"down and up", []string{
			"X..S",
			".MA.",
			".MA.",
			"X..S",
		}, 2},
		{"overlap", []string{
			"X......",
			".M.....",
			"..A....",
			"...S...",
			"....A..",
			".....M.",
			"......X",
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

func TestFullGrid(t *testing.T) {
	needle := "XMAS"
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{"full grid", []string{
			"MMMSXXMASM",
			"MSAMXMSMSA",
			"AMXSXMAAMM",
			"MSAMASMSMX",
			"XMASAMXAMM",
			"XXAMMXXAMA",
			"SMSMSASXSS",
			"SAXAMASAAA",
			"MAMMMXMMMM",
			"MXMXAXMASX",
		}, 18},
		{"full grid with dots", []string{
			//0123456789
			"....XXMAS.", // 0
			".SAMXMS...", // 1
			"...S..A...", // 2
			"..A.A.MS.X", // 3
			"XMASAMX.MM", // 4
			"X.....XA.A", // 5
			"S.S.S.S.SS", // 6
			".A.A.A.A.A", // 7
			"..M.M.M.MM", // 8
			".X.X.XMASX", // 9
		}, 18},
	}

	// Horizontal : 5
	// Vertical   : 3 [99-69U, 39-69D, 46-16U]
	// Diagonal  NWSE : 1
	// Diagonal  NESW : 1
	// Diagonal  SWNE : 4
	// Diagonal  SENW : 4
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			num := SearchForWord(needle, testCase.input)
			if num != testCase.expected {
				t.Errorf("%v: expected %v but got %v", testCase.name, testCase.expected, num)
			}
		})
	}
}
