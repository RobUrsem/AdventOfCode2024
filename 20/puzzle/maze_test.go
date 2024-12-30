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
func TestLongCheat(t *testing.T) {
	testCases := []struct {
		name         string
		input        []string
		expectedCost int
		cheats       map[int]int
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
			cheats: map[int]int{
				50: 32,
				52: 31,
				54: 29,
				56: 39,
				58: 25,
				60: 23,
				62: 20,
				64: 19,
				66: 12,
				68: 14,
				70: 12,
				72: 22,
				74: 4,
				76: 3,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			maze := MakeMaze(tc.input)
			numCheats := maze.Part2(50)

			expected := 0
			for _, v := range tc.cheats {
				expected += v
			}

			if numCheats != expected {
				t.Errorf("Expected %v cheats but got %v", expected, numCheats)
			}
		})
	}
}
