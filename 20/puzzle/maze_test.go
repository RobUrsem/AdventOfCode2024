package puzzle

import (
	"fmt"
	"sort"
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
				"#...#...#.....#",
				"#.#.#.#.#.###.#",
				"#S#...#.#.#...#",
				"#######.#.#.###",
				"#######.#.#...#",
				"#######.#.###.#",
				"###..E#...#...#",
				"###.#######.###",
				"#...###...#...#",
				"#.#####.#.###.#",
				"#.#...#.#.#...#",
				"#.#.#.#.#.#.###",
				"#...#...#...###",
				"###############",
			},
			expectedCost: 84,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			maze := MakeMaze(tc.input)
			basecost := maze.SolveMaze()

			if basecost != tc.expectedCost {
				t.Errorf("Expected total to be %v, but got %v", tc.expectedCost, basecost)
			}

			counts := map[int][]Location{}
			cheats := maze.GetCheatLocations()
			for _, cheat := range cheats {
				saving := maze.AddCheat(cheat)
				counts[saving] = append(counts[saving], cheat)
			}

			keys := make([]int, 0, len(counts))
			for k := range counts {
				keys = append(keys, k)
			}

			// Sort the keys
			sort.Ints(keys)
			for _, k := range keys {
				if k == 0 {
					continue
				}
				fmt.Printf("%2d steps: %2d cheats\n", k, len(counts[k]))
			}

		})
	}
}
