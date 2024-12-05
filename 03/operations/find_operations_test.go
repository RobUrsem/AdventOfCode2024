package operations

import "testing"

func areEqual(a, b Operation) bool {
	if len(a.Params) != len(b.Params) {
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
		{"mul(44,46)", []Operation{
			{Params: []int{44, 46}, OperationType: Multiply},
		}},
		{"mul(123,4)", []Operation{
			{Params: []int{123, 4}, OperationType: Multiply},
		}},
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
				{Params: []int{8, 5}, OperationType: Multiply},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			operations, err := FindOperations(tc.input)
			if err != nil {
				t.Errorf("input [%v] gave error %v", tc.input, err)
			}
			for i, operation := range operations {
				if !areEqual(operation, tc.expected[i]) {
					t.Errorf("expected %v but got %v", tc.expected[i], operation)
				}
			}
		})
	}
}
