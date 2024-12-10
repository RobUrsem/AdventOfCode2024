package puzzle

import "testing"

func TestAntennaMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected AntennaMap
	}{
		{
			name: "two antinodes",
			input: []string{
				"...a..a.....",
			},
			expected: AntennaMap{
				"#..a..a..#..",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			theMap := CreateMap(tc.input)
			theMap.DetermineAntinodes()

			if !theMap.IsSameAs(tc.expected) {
				t.Errorf("%v: expected different antinodes", tc.name)
			}
		})
	}
}
