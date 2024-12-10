package puzzle

import "testing"

func TestGetEquations(t *testing.T) {
	testCases := []struct {
		input    string
		expected Equation
		valid    bool
	}{
		{
			"190: 10 19",
			Equation{
				Answer:       190,
				Coefficients: []int64{10, 19},
			},
			true,
		},
		{
			"3267: 81 40 27",
			Equation{
				Answer:       3267,
				Coefficients: []int64{81, 40, 27},
			},
			true,
		},
		{
			"83: 17 5",
			Equation{
				Answer:       83,
				Coefficients: []int64{17, 5},
			},
			false,
		},
		{
			"156: 15 6",
			Equation{
				Answer:       156,
				Coefficients: []int64{15, 6},
			},
			false,
		},
		{
			"7290: 6 8 6 15",
			Equation{
				Answer:       7290,
				Coefficients: []int64{6, 8, 6, 15},
			},
			false,
		},
		{
			"161011: 16 10 13",
			Equation{
				Answer:       161011,
				Coefficients: []int64{16, 10, 13},
			},
			false,
		},
		{
			"192: 17 8 14",
			Equation{
				Answer:       192,
				Coefficients: []int64{17, 8, 14},
			},
			false,
		},
		{
			"21037: 9 7 18 13",
			Equation{
				Answer:       21037,
				Coefficients: []int64{9, 7, 18, 13},
			},
			false,
		},
		{
			"292: 11 6 16 20",
			Equation{
				Answer:       292,
				Coefficients: []int64{11, 6, 16, 20},
			},
			true,
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

			solved := SolveEquation(equations[0])

			if solved.Valid != tc.valid {
				t.Errorf("Expected %v to be %v but got %v", tc.input, tc.valid, solved.Valid)
			}
		})
	}
}
