package puzzle

import (
	"testing"
)

func AreEqual(a, b string) bool {
	return a != b
}

func TestGrid(t *testing.T) {
	testCases := []struct {
		name      string
		input     []string
		grid      Grid
		moves     string
		expected  string
		gpscoords int
	}{
		{
			name: "0",
			input: []string{
				"########",
				"#..O.O.#",
				"##@.O..#",
				"#...O..#",
				"#.#.O..#",
				"#...O..#",
				"#......#",
				"########",
				"",
				"<^^>>>vv<v>>v<<",
			},
			grid: Grid{
				grid: [][]rune{
					{'#', '#', '#', '#', '#', '#', '#', '#'},
					{'#', '.', '.', 'O', '.', 'O', '.', '#'},
					{'#', '#', '@', '.', 'O', '.', '.', '#'},
					{'#', '.', '.', '.', 'O', '.', '.', '#'},
					{'#', '.', '#', '.', 'O', '.', '.', '#'},
					{'#', '.', '.', '.', 'O', '.', '.', '#'},
					{'#', '.', '.', '.', '.', '.', '.', '#'},
					{'#', '#', '#', '#', '#', '#', '#', '#'},
				},
			},
			moves:     "<^^>>>vv<v>>v<<",
			expected:  "########\n#....OO#\n##.....#\n#.....O#\n#.#O@..#\n#...O..#\n#...O..#\n########",
			gpscoords: 2028,
		},
		{
			name: "1",
			input: []string{
				"##########",
				"#..O..O.O#",
				"#......O.#",
				"#.OO..O.O#",
				"#..O@..O.#",
				"#O#..O...#",
				"#O..O..O.#",
				"#.OO.O.OO#",
				"#....O...#",
				"##########",
				"",
				"<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^",
				"vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v",
				"><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<",
				"<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^",
				"^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><",
				"^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^",
				">^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^",
				"<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>",
				"^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>",
				"v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
			},
			grid: Grid{
				grid: [][]rune{
					{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
					{'#', '.', '.', 'O', '.', '.', 'O', '.', 'O', '#'},
					{'#', '.', '.', '.', '.', '.', '.', 'O', '.', '#'},
					{'#', '.', 'O', 'O', '.', '.', 'O', '.', 'O', '#'},
					{'#', '.', '.', 'O', '@', '.', '.', 'O', '.', '#'},
					{'#', 'O', '#', '.', '.', 'O', '.', '.', '.', '#'},
					{'#', 'O', '.', '.', 'O', '.', '.', 'O', '.', '#'},
					{'#', '.', 'O', 'O', '.', 'O', '.', 'O', 'O', '#'},
					{'#', '.', '.', '.', '.', 'O', '.', '.', '.', '#'},
					{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#'},
				},
			},
			moves:     "<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^",
			expected:  "##########\n#.O.O.OOO#\n#........#\n#OO......#\n#OO@.....#\n#O#.....O#\n#O.....OO#\n#O.....OO#\n#OO....OO#\n##########",
			gpscoords: 10092,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			grid := ReadGrid(tc.input)
			if len(grid.grid) != len(tc.grid.grid) {
				t.Errorf("Error reading grid")
			}

			moves := ReadMoves(tc.input)
			if len(moves) != len(tc.moves) {
				t.Errorf("Error reading moves")
			}

			grid.MoveRobot(tc.moves)

			if !AreEqual(grid.String(), tc.expected) {
				t.Errorf("Wrong grid at the end")
			}

			gps := grid.CalculateGPS()
			if gps != tc.gpscoords {
				t.Errorf("Expected GPS coordinates of %v but got %v", tc.gpscoords, gps)
			}
		})
	}
}
