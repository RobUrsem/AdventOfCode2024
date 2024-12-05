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
			expected: 2 * 3,
		},
		{
			name: "mul(2,3);mul(44,46)",
			input: []Operation{
				{Params: []int{2, 3}, OperationType: Multiply},
				{Params: []int{44, 46}, OperationType: Multiply},
			},
			expected: 2*3 + 44*46,
		},
		{
			name: "mul(123,456)",
			input: []Operation{
				{Params: []int{123, 456}, OperationType: Multiply},
			},
			expected: 123 * 456,
		},
		{
			name: "large example",
			input: []Operation{
				{Params: []int{2, 4}, OperationType: Multiply},
				{Params: []int{5, 5}, OperationType: Multiply},
				{Params: []int{11, 8}, OperationType: Multiply},
				{Params: []int{8, 5}, OperationType: Multiply},
			},
			expected: 2*4 + 5*5 + 11*8 + 8*5,
		},
		{
			name: "with disable and enable",
			input: []Operation{
				{Params: []int{2, 4}, OperationType: Multiply},
				{Params: nil, OperationType: Disable},
				{Params: []int{5, 5}, OperationType: Multiply},
				{Params: []int{11, 8}, OperationType: Multiply},
				{Params: nil, OperationType: Enable},
				{Params: []int{8, 5}, OperationType: Multiply},
			},
			expected: 2*4 + 8*5,
		},
		{
			name: "with multiple disables and enable",
			input: []Operation{
				{Params: []int{2, 4}, OperationType: Multiply},
				{Params: nil, OperationType: Disable},
				{Params: nil, OperationType: Disable},
				{Params: []int{5, 5}, OperationType: Multiply},
				{Params: nil, OperationType: Enable},
				{Params: []int{11, 8}, OperationType: Multiply},
				{Params: nil, OperationType: Enable},
				{Params: []int{8, 5}, OperationType: Multiply},
			},
			expected: 2*4 + 11*8 + 8*5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			total := ExecuteOperations(tc.input)

			if total != tc.expected {
				t.Errorf("%v: expected %v but got %v", tc.name, tc.expected, total)
			}
		})
	}
}
