package puzzle

import "testing"

func TestClawGame(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected Game
		cost     int
	}{
		{
			name: "1",
			input: []string{
				"Button A: X+94, Y+34",
				"Button B: X+22, Y+67",
				"Prize: X=8400, Y=5400",
			},
			expected: Game{
				A:     []int{94, 34},
				B:     []int{22, 67},
				Prize: []int{8400, 5400},
			},
			cost: 80*3 + 40*1,
		},
		{
			name: "2",
			input: []string{
				"Button A: X+26, Y+66",
				"Button B: X+67, Y+21",
				"Prize: X=12748, Y=12176",
			},
			expected: Game{
				A:     []int{26, 66},
				B:     []int{67, 21},
				Prize: []int{12748, 12176},
			},
			cost: -1, // Solution not possible
		},
		{
			name: "3",
			input: []string{
				"Button A: X+17, Y+86",
				"Button B: X+84, Y+37",
				"Prize: X=7870, Y=6450",
			},
			expected: Game{
				A:     []int{17, 86},
				B:     []int{84, 37},
				Prize: []int{7870, 6450},
			},
			cost: 38*3 + 86*1,
		},
		{
			name: "4",
			input: []string{
				"Button A: X+69, Y+23",
				"Button B: X+27, Y+71",
				"Prize: X=18641, Y=10279",
			},
			expected: Game{
				A:     []int{67, 23},
				B:     []int{27, 71},
				Prize: []int{18641, 10279},
			},
			cost: -1, // Solution not possible
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			games := GetInput(tc.input)
			if len(games) != 1 {
				t.Fatalf("Could not parse game")
			}

			cost := games[0].Cost()
			if cost != tc.cost {
				t.Errorf("Expected cost of %v but got %v", tc.cost, cost)
			}
		})
	}
}
