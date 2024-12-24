package puzzle

import (
	"testing"
)

func TestMaze(t *testing.T) {
	testCases := []struct {
		name             string
		input            []string
		expectedCost     int
		expectedNumSeats int
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
			expectedCost:     7036,
			expectedNumSeats: 45,
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
			expectedCost:     11048,
			expectedNumSeats: 64,
		},
	}

	// err := termbox.Init()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer termbox.Close()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			maze := MakeMaze(tc.input)
			_, cost, numSeats := maze.SolveMaze()

			if cost != tc.expectedCost {
				t.Errorf("Expected total to be %v, but got %v", tc.expectedCost, cost)
			}

			if numSeats != tc.expectedNumSeats {
				t.Errorf("Expected total best seats to be %v, but got %v", tc.expectedNumSeats, numSeats)
			}
		})
	}
}
