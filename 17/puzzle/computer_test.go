package puzzle

import (
	"fmt"
	"testing"
	"time"
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

	if len(a.Output) != len(b.Output) {
		return false
	}

	for i := range a.Output {
		if a.Output[i] != b.Output[i] {
			return false
		}
	}
	return true
}

func AreEqualSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
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
				Output: []int{},
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
				Output: []int{0, 1, 2},
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
				Output: []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0},
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
				Output: []int{4, 6, 3, 5, 6, 3, 5, 2, 1, 0},
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
				Output: []int{2, 7, 2, 5, 1, 2, 7, 3, 7},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			computer := NewComputer(tc.input)
			computer.Run()

			output := computer.Print()
			fmt.Printf("Output: [%v]\n", output)

			if !AreEqual(computer, tc.expected) {
				t.Errorf("Expected output: \n%v\n, but got \n%v\n", tc.expected, computer)
			}
		})
	}
}

func TestReverseComputerFull(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			"ex 1",
			[]string{
				"Register A: 2024",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 0,3,5,4,3,0",
			},
			117440,
		},
		{
			"Full",
			[]string{
				"Register A: 25358015",
				"Register B: 0",
				"Register C: 0",
				"",
				"Program: 2,4, 1,1, 7,5, 0,3, 4,7, 1,6, 5,5, 3,0",
			},
			247839002892474,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			computer := NewComputer(tc.input)

			begin := time.Now()
			computer.RunReverse()
			elapsed := time.Since(begin)

			fmt.Printf("Value of register A: %v\n", computer.CorrectedA)
			fmt.Printf("Elapsed: %v\n", elapsed)

			if computer.CorrectedA != tc.expected {
				t.Errorf("Incorrect: got %v but expected %v", computer.CorrectedA, tc.expected)
			}
		})
	}
}

func TestSingle(t *testing.T) {
	input := []string{
		"Register A: 5",
		"Register B: 0",
		"Register C: 0",
		"",
		"Program: 2,4, 1,1, 7,5, 0,3, 4,7, 1,6, 5,5, 3,0",
	}

	computer := NewComputer(input)
	options := []int{
		247839002892474,
	}
	for i := range options {
		computer.Reset(options[i], 0, 0)
		computer.Run()
		fmt.Printf("%d -> %v\n", options[i], computer.Print())
	}

}
