package operations

import "testing"

func TestExecuteOperations(t *testing.T) {
	testCases := []struct {
		name     string
		input    []Operation
		expected int
	}{
		{
			name: "mul(2,3)",
			input: []Operation{
				{Params: []int{2, 3}, OperationType: Multiply},
			},
			expected: 6,
		},
		{
			name: "mul(2,3);mul(44,46)",
			input: []Operation{
				{Params: []int{2, 3}, OperationType: Multiply},
				{Params: []int{44, 46}, OperationType: Multiply},
			},
			expected: 2030,
		},
		{
			name: "mul(123,456)",
			input: []Operation{
				{Params: []int{123, 456}, OperationType: Multiply},
			},
			expected: 56088,
		},
		{
			name: "large example",
			input: []Operation{
				{Params: []int{2, 4}, OperationType: Multiply},
				{Params: []int{5, 5}, OperationType: Multiply},
				{Params: []int{11, 8}, OperationType: Multiply},
				{Params: []int{8, 5}, OperationType: Multiply},
			},
			expected: 161,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			total := 0
			for _, op := range tc.input {
				total += Execute(op)
			}

			if total != tc.expected {
				t.Errorf("%v: expected %v but got %v", tc.name, tc.expected, total)
			}
		})
	}
}
