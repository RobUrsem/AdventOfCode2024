package puzzle

import "testing"

func TestMaze(t *testing.T) {
	testCases := []struct {
		name         string
		input        []string
		expectedCost int
	}{
		{
			name: "1",
			input: []string{
				"###############",
				"#.......#....E#",
				"#.#.###.#.###.#",
				"#.....#.#...#.#",
				"#.###.#####.#.#",
				"#.#.#.......#.#",
				"#.#.#####.###.#",
				"#...........#.#",
				"###.#.#####.#.#",
				"#...#.....#.#.#",
				"#.#.#.###.#.#.#",
				"#.....#...#.#.#",
				"#.###.#.#.#.#.#",
				"#S..#.....#...#",
				"###############",
			},
			expectedCost: 7036,
		},
		{
			name: "2",
			input: []string{
				"#################",
				"#...#...#...#..E#",
				"#.#.#.#.#.#.#.#.#",
				"#.#.#.#...#...#.#",
				"#.#.#.#.###.#.#.#",
				"#...#.#.#.....#.#",
				"#.#.#.#.#.#####.#",
				"#.#...#.#.#.....#",
				"#.#.#####.#.###.#",
				"#.#.#.......#...#",
				"#.#.###.#####.###",
				"#.#.#...#.....#.#",
				"#.#.#.#####.###.#",
				"#.#.#.........#.#",
				"#.#.#.#########.#",
				"#S#.............#",
				"#################",
			},
			expectedCost: 11048,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			maze := MakeMaze(tc.input)
			cost := maze.FindShortestPathLength()
			if cost != tc.expectedCost {
				t.Errorf("Expected total to be %v, but got %v", tc.expectedCost, cost)
			}
		})
	}
}
