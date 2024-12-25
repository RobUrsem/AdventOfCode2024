package puzzle

import (
	"fmt"
	"testing"
)

func AreEqual(a, b Computer) bool {
	if a.A != b.A {
		return false
	}
	if a.B != b.B {
		return false
	}
	if a.C != b.C {
		return false
	}

	if len(a.output) != len(b.output) {
		return false
	}

	for i := range a.output {
		if a.output[i] != b.output[i] {
			return false
		}
	}
	return true
}

func TestComputer(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected Computer
	}{
		{
			"ex 1",
			[]string{
				"Register A: 0",
				"Register B: 0",
				"Register C: 9",
				"",
				"Program: 2,6",
			},
			Computer{
				A:      0,
				B:      1,
				C:      9,
				output: []int{},
			},
		},
		{
			"ex 2",
			[]string{
				"Register A: 10",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 5,0,5,1,5,4",
			},
			Computer{
				A:      10,
				B:      0,
				C:      0,
				output: []int{0, 1, 2},
			},
		},
		{
			"ex 3",
			[]string{
				"Register A: 2024",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,1,5,4,3,0",
			},
			Computer{
				A:      0,
				B:      0,
				C:      0,
				output: []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0},
			},
		},
		{
			"1",
			[]string{
				"Register A: 729",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,1,5,4,3,0",
			},
			Computer{
				A:      0,
				B:      0,
				C:      0,
				output: []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0},
			},
		},
		{
			"Full",
			[]string{
				"Register A: 25358015",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4,1,1,7,5,0,3,4,7,1,6,5,5,3,0",
			},
			Computer{
				A:      0,
				B:      7,
				C:      1,
				output: []int{2, 7, 2, 5, 1, 2, 7, 3, 7},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			computer := NewComputer(tc.input)
			computer.Run()

			output := computer.Output()
			fmt.Printf("Output: [%v]\n", output)

			if !AreEqual(computer, tc.expected) {
				t.Errorf("Expected output: \n%v\n, but got \n%v\n", tc.expected, computer)
			}
		})
	}
}
