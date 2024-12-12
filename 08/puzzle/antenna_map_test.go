package puzzle

import (
	"fmt"
	"testing"
)

func TestAntennaMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    []string
		expected AntennaMap
		num      int
	}{
		{
			name: "two antinodes",
			input: []string{
				"...a..a.....",
			},
			expected: AntennaMap{
				"#..a..a..#..",
			},
			num: 2,
		},
		{
			name: "example",
			input: []string{
				"............",
				"........0...",
				".....0......",
				".......0....",
				"....0.......",
				"......A.....",
				"............",
				"............",
				"........A...",
				".........A..",
				"............",
				"............",
			},
			expected: AntennaMap{
				"......#....#",
				"...#....0...",
				"....#0....#.",
				"..#....0....",
				"....0....#..",
				".#....A.....",
				"...#........",
				"#......#....",
				"........A...",
				".........A..",
				"..........#.",
				"..........#.",
			},
			num: 14,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			theMap := CreateMap(tc.input)
			theMap.Filter()

			fmt.Println(theMap)
			if !theMap.IsSameAs(tc.expected) {
				t.Errorf("%v: expected different antinodes", tc.name)
			}

			if theMap.CountAntiNodes() != tc.num {
				t.Errorf("%v: expected %v but got %v", tc.name, tc.num, theMap.CountAntiNodes())
			}
		})
	}
}
