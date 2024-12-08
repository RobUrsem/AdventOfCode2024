package ordering

import "testing"

func areEqual(a, b Update) bool {
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

func TestFixUpdates(t *testing.T) {
	rulebook := Rulebook{
		{47, 53},
		{97, 13},
		{97, 61},
		{97, 47},
		{75, 29},
		{61, 13},
		{75, 53},
		{29, 13},
		{97, 29},
		{53, 29},
		{61, 53},
		{97, 53},
		{61, 29},
		{47, 13},
		{75, 47},
		{97, 75},
		{47, 61},
		{75, 61},
		{47, 29},
		{75, 13},
		{53, 13},
	}

	testCases := []struct {
		name     string
		updates  Update
		expected Update
	}{
		{"4", Update{75, 97, 47, 61, 53}, Update{97, 75, 47, 61, 53}},
		{"5", Update{61, 13, 29}, Update{61, 29, 13}},
		{"6", Update{97, 13, 75, 29, 47}, Update{97, 75, 47, 29, 13}},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			fixed := fixUpdate(testCase.updates, rulebook)

			if !areEqual(fixed, testCase.expected) {
				t.Errorf("Case %v, expected %v but got %v", testCase.name, testCase.expected, fixed)
			}
		})
	}
}
