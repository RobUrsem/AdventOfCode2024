package operations

import "testing"

func areEqual(a, b Operation) bool {
	if len(a.Params) != len(b.Params) {
		return false
	}

	if a.OperationType != b.OperationType {
		return false
	}

	for i, par := range a.Params {
		if par != b.Params[i] {
			return false
		}
	}

	return true
}

func TestFindOperations(t *testing.T) {
	testCases := []struct {
		input    string
		expected []Operation
	}{
		{"mul(4*", nil},
		{"mul(6,9!", nil},
		{"?(12,34)", nil},
		{"mul ( 2 , 4 )", nil},
		{"mul(531,132what()", nil},
		{"mul(298,605*<mul(189,109)", []Operation{
			{Params: []int{189, 109}, OperationType: Multiply},
		}},
		{"mul(44,46)", []Operation{
			{Params: []int{44, 46}, OperationType: Multiply},
		}},
		{"mul(123,4)", []Operation{
			{Params: []int{123, 4}, OperationType: Multiply},
		}},
		{"mul(1234,4)", nil},
		{"xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))",
			[]Operation{
				{Params: []int{2, 4}, OperationType: Multiply},
				{Params: []int{5, 5}, OperationType: Multiply},
				{Params: []int{11, 8}, OperationType: Multiply},
				{Params: []int{8, 5}, OperationType: Multiply},
			},
		},
		{"xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))",
			[]Operation{
				{Params: []int{2, 4}, OperationType: Multiply},
				{Params: nil, OperationType: Disable},
				{Params: []int{5, 5}, OperationType: Multiply},
				{Params: []int{11, 8}, OperationType: Multiply},
				{Params: nil, OperationType: Enable},
				{Params: []int{8, 5}, OperationType: Multiply},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			operations := FindOperations(tc.input)

			if len(operations) != len(tc.expected) {
				t.Fatalf("expected %v operations but got %v", len(tc.expected), len(operations))
			}

			for i, operation := range operations {
				if !areEqual(operation, tc.expected[i]) {
					t.Errorf("expected %v but got %v", tc.expected[i], operation)
				}
			}
		})
	}
}
