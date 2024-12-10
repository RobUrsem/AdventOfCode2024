package puzzle

import "testing"

func TestGetEquations(t *testing.T) {
	testCases := []struct {
		input    string
		expected Equation
	}{
		{
			"190: 10 19",
			Equation{
				answer:       190,
				coefficients: []int64{10, 19},
			},
		},
		{
			"3267: 81 40 27",
			Equation{
				answer:       3267,
				coefficients: []int64{81, 40, 27},
			},
		},
		{
			"83: 17 5",
			Equation{
				answer:       83,
				coefficients: []int64{17, 5},
			},
		},
		{
			"156: 15 6",
			Equation{
				answer:       156,
				coefficients: []int64{15, 6},
			},
		},
		{
			"7290: 6 8 6 15",
			Equation{
				answer:       7290,
				coefficients: []int64{6, 8, 6, 15},
			},
		},
		{
			"161011: 16 10 13",
			Equation{
				answer:       161011,
				coefficients: []int64{16, 10, 13},
			},
		},
		{
			"192: 17 8 14",
			Equation{
				answer:       192,
				coefficients: []int64{17, 8, 14},
			},
		},
		{
			"21037: 9 7 18 13",
			Equation{
				answer:       21037,
				coefficients: []int64{9, 7, 18, 13},
			},
		},
		{
			"292: 11 6 16 20",
			Equation{
				answer:       292,
				coefficients: []int64{11, 6, 16, 20},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			equations := GetEquations([]string{tc.input})

			if len(equations) != 1 {
				t.Errorf("Expected 1 equation, got %v", len(equations))
			}

			if !AreEqual(tc.expected, equations[0]) {
				t.Errorf("Equations are not the same")
			}
		})
	}
}
