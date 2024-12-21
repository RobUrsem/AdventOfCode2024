package puzzle

import (
	"fmt"
	"testing"
)

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

	// err := termbox.Init()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer termbox.Close()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			maze := MakeMaze(tc.input)

			// path := maze.AStarSolve()
			// fmt.Printf("Path length: %v\n", len(path))
			// _, cost := maze.SolveMaze()
			path, cost := maze.SolveMaze()

			for i, p := range path {
				fmt.Printf("%3d: %v\n", i, p)
			}

			// cost := maze.FindShortestPathLength()
			if cost != tc.expectedCost {
				t.Errorf("Expected total to be %v, but got %v", tc.expectedCost, cost)
			}
		})
	}
}
