package puzzle

import (
	"fmt"
	"testing"
)

func TestMemory(t *testing.T) {
	testCases := []struct {
		name  string
		input []string
	}{
		{
			"example",
			[]string{
				"5,4",
				"4,2",
				"4,5",
				"3,0",
				"2,1",
				"6,3",
				"2,4",
				"1,5",
				"0,6",
				"3,3",
				"2,6",
				"5,1",
				"1,2",
				"5,5",
				"2,5",
				"6,5",
				"1,4",
				"0,4",
				"6,4",
				"1,1",
				"6,1",
				"1,0",
				"0,5",
				"1,6",
				"2,0",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mem := MakeMemory(7, 7)
			mem.LoadBytes(tc.input)
			mem.Simulate(12)
			_, cost := mem.SolveMaze()
			mem.Print()
			fmt.Printf("Cost to the maze: %v\n", cost)

			for step := 13; step < len(mem.bytes); step++ {
				mem.NextByte()
				_, cost = mem.SolveMaze()
				if cost < 0 {
					fmt.Printf("Solve impossible at step %v last block %v\n", step, mem.bytes[step-1])
					break
				}
			}
		})
	}
}
