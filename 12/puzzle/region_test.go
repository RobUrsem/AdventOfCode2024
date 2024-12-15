package puzzle

import (
	"testing"
)

func TestRegion(t *testing.T) {
	testCases := []struct {
		name    string
		input   []string
		regions Regions
		cost    int
		sides   int
	}{
		{
			name: "1",
			input: []string{
				"AAAA",
				"BBCD",
				"BBCC",
				"EEEC",
			},
			regions: Regions{
				{1, 1, 1, 1},
				{2, 2, 3, 4},
				{2, 2, 3, 3},
				{5, 5, 5, 3},
			},
			cost:  4*10 + 4*8 + 4*10 + 1*4 + 3*8,
			sides: 4*4 + 4*4 + 4*8 + 1*4 + 3*4,
		},
		{
			name: "XOXO",
			input: []string{
				"OOOOO",
				"OXOXO",
				"OOOOO",
				"OXOXO",
				"OOOOO",
			},
			regions: Regions{
				{1, 1, 1, 1, 1},
				{1, 2, 1, 3, 1},
				{1, 1, 1, 1, 1},
				{1, 4, 1, 5, 1},
				{1, 1, 1, 1, 1},
			},
			cost:  21*36 + 1*4 + 1*4 + 1*4 + 1*4,
			sides: 21*20 + 1*4 + 1*4 + 1*4 + 1*4,
		},
		{
			name: "E",
			input: []string{
				"EEEEE",
				"EXXXX",
				"EEEEE",
				"EXXXX",
				"EEEEE",
			},
			regions: Regions{
				{1, 1, 1, 1, 1},
				{1, 2, 2, 2, 2},
				{1, 1, 1, 1, 1},
				{1, 3, 3, 3, 3},
				{1, 1, 1, 1, 1},
			},
			cost:  17*36 + 4*10 + 4*10,
			sides: 17*12 + 4*4 + 4*4,
		},
		{
			name: "ABBA",
			input: []string{
				"AAAAAA",
				"AAABBA",
				"AAABBA",
				"ABBAAA",
				"ABBAAA",
				"AAAAAA",
			},
			regions: Regions{
				{1, 1, 1, 1, 1, 1},
				{1, 1, 1, 2, 2, 1},
				{1, 1, 1, 2, 2, 1},
				{1, 3, 3, 1, 1, 1},
				{1, 3, 3, 1, 1, 1},
				{1, 1, 1, 1, 1, 1},
			},
			cost:  28*40 + 4*8 + 4*8,
			sides: 28*12 + 4*4 + 4*4,
		},
		{
			name: "Large",
			input: []string{
				"RRRRIICCFF",
				"RRRRIICCCF",
				"VVRRRCCFFF",
				"VVRCCCJFFF",
				"VVVVCJJCFE",
				"VVIVCCJJEE",
				"VVIIICJJEE",
				"MIIIIIJJEE",
				"MIIISIJEEE",
				"MMMISSJEEE",
			},
			regions: Regions{
				// "RR R  R  I  I  C  C  F  F",
				{1, 1, 1, 1, 2, 2, 3, 3, 4, 4},

				// "RR R  R  I  I  C  C  C  F",
				{1, 1, 1, 1, 2, 2, 3, 3, 3, 4},

				// "VV R  R  R  C  C  F  F  F",
				{5, 5, 1, 1, 1, 3, 3, 4, 4, 4},

				// "VV R  C  C  C  J  F  F  F",
				{5, 5, 1, 3, 3, 3, 9, 4, 4, 4},

				// "VV V  V  C  J  J   C  F   E",
				{5, 5, 5, 5, 3, 9, 9, 11, 4, 12},

				// "VV  I  V  C  C  J  J   E   E",
				{5, 5, 13, 5, 3, 3, 9, 9, 12, 12},

				// "VV  I   I   I  C  J  J   E   E",
				{5, 5, 13, 13, 13, 3, 9, 9, 12, 12},

				// "M I   I   I   I   I  J  J   E   E",
				{15, 13, 13, 13, 13, 13, 9, 9, 12, 12},

				// "M I   I   I   S   I  J   E   E   E",
				{15, 13, 13, 13, 17, 13, 9, 12, 12, 12},

				// "M M   M   I   S   S  J   E   E   E",
				{15, 15, 15, 13, 17, 17, 9, 12, 12, 12},
			},
			cost:  12*18 + 4*8 + 14*28 + 10*18 + 13*20 + 11*20 + 1*4 + 13*18 + 14*22 + 5*12 + 3*8,
			sides: 12*10 + 4*4 + 14*22 + 10*12 + 13*10 + 11*12 + 1*4 + 13*8 + 14*16 + 5*6 + 3*6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			regions := Segment(tc.input)
			if len(tc.regions) != len(regions) {
				t.Errorf("Expected %v regions but got %v", len(tc.regions), len(regions))
			}

			for r := range regions {
				for c := range regions[r] {
					if regions[r][c] != tc.regions[r][c] {
						t.Errorf("Regions (%v,%v) expected %v but got %v", r, c, tc.regions[r][c], regions[r][c])
					}
				}
			}

			cost := CalcCost(regions)
			if cost != tc.cost {
				t.Errorf("Expected cost of %v but got %v", tc.cost, cost)
			}

			sides := CalcSides(regions)
			if sides != tc.sides {
				t.Errorf("Expected sides of %v but got %v", tc.sides, sides)
			}
		})
	}
}
