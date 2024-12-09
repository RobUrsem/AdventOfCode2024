package guard

import "testing"

func TestWalkSample(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected []Obstruction
	}{
		{
			name: "circle 1",
			input: []string{
				//01234
				".#...", // 0
				"....#", // 1
				".^...", // 2
				".....", // 3
				"...#.", // 4
			},
			expected: []Obstruction{
				{3, 0},
			},
		},
		{
			name: "circle 2a",
			input: []string{
				".....",
				"..>.#",
				".....",
				"#....",
				"...#.",
			},
			expected: []Obstruction{
				{0, 1},
			},
		},
		{
			name: "circle 2b",
			input: []string{
				".....",
				".>..#",
				".....",
				"#....",
				"...#.",
			},
			expected: []Obstruction{
				{0, 1},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			labMap, err := ConstructMap(tc.input)
			if err != nil {
				t.Error("Error constructing map")
			}

			turns, err := DoWalk(labMap)
			if err != nil {
				t.Error("Error walking map")
			}

			obstructions := FindPattern(labMap, turns)
			if len(obstructions) != len(tc.expected) {
				t.Errorf("Expected %v obstructions, but got %v", len(tc.expected), len(obstructions))
			}

			for i, obstruction := range obstructions {
				for j, value := range obstruction {
					if value != tc.expected[i][j] {
						t.Errorf("(%v, %v) => Expected %v but got %v", i, j, tc.expected[i][j], value)
					}
				}
			}
		})
	}
}

func TestWalkExampleProblem(t *testing.T) {
	example := []string{
		"....#.....",
		".........#",
		"..........",
		"..#.......",
		".......#..",
		"..........",
		".#..^.....",
		"........#.",
		"#.........",
		"......#...",
	}

	expected := [][]int{
		{6, 3},
		{7, 6},
		{7, 7},
		{8, 1},
		{8, 3},
		{9, 7},
	}

	labMap, err := ConstructMap(example)
	if err != nil {
		t.Error("Error constructing map")
	}

	turns, err := DoWalk(labMap)
	if err != nil {
		t.Error("Error walking map")
	}

	obstructions := FindPattern(labMap, turns)

	if len(obstructions) != len(expected) {
		t.Errorf("Expected %v obstructions, but got %v", len(expected), len(obstructions))
	}

	for i, obstruction := range obstructions {
		if obstruction[0] != expected[i][0] || obstruction[1] != expected[i][1] {
			t.Errorf("%v => Expected (%v, %v) but got (%v, %v)", i,
				expected[i][0], expected[i][1],
				obstruction[0], obstruction[1])
		}
	}
}
