package guard

import (
	"testing"
)

func areEqual(a, b LabMap) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestConstructLabMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected LabMap
		err      error
	}{
		{
			"all elements",
			[]string{".#X^>V<"},
			LabMap{[]int{EMPTY, OBSTACLE, VISITED, GUARD_UP, GUARD_RIGHT, GUARD_DOWN, GUARD_LEFT}},
			nil,
		},
		{
			"unknown char",
			[]string{".h"},
			LabMap{[]int{EMPTY}},
			ErrInvalidCharacter,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			myMap, err := ConstructMap(tc.input)

			if err != nil && err != ErrInvalidCharacter {
				t.Errorf("Unknown error: %v", err)
			}

			if err == nil && !areEqual(tc.expected, myMap) {
				t.Errorf("Map not the same. Expected %v but got %v", tc.expected, myMap)
			}

		})
	}
}
