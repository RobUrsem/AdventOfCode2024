package puzzle

import (
	"advent/shared"
	"fmt"
	"testing"
)

func TestTrails(t *testing.T) {
	testCases := []struct {
		name       string
		input      []string
		trailHeads shared.Locations
		summits    shared.Locations
		trails     []shared.Locations
		scores     []int
		ratings    []int
	}{
		{
			name: "1",
			input: []string{
				"...0...",
				"...1...",
				"...2...",
				"6543456",
				"7.....7",
				"8.....8",
				"9.....9",
			},
			trailHeads: shared.Locations{
				shared.NewLocation(0, 3),
			},
			summits: shared.Locations{
				shared.NewLocation(6, 0),
				shared.NewLocation(6, 6),
			},
			trails: []shared.Locations{
				{
					shared.NewLocation(0, 3),
					shared.NewLocation(1, 3),
					shared.NewLocation(2, 3),
					shared.NewLocation(3, 3), // split point
					shared.NewLocation(3, 2),
					shared.NewLocation(3, 1),
					shared.NewLocation(3, 0),
					shared.NewLocation(4, 0),
					shared.NewLocation(5, 0),
					shared.NewLocation(6, 0),
				},
				{
					shared.NewLocation(0, 3),
					shared.NewLocation(1, 3),
					shared.NewLocation(2, 3),
					shared.NewLocation(3, 3),
					shared.NewLocation(3, 4),
					shared.NewLocation(3, 5),
					shared.NewLocation(3, 6),
					shared.NewLocation(4, 6),
					shared.NewLocation(5, 6),
					shared.NewLocation(6, 6),
				},
			},
			scores:  []int{2},
			ratings: []int{2},
		},
		{
			name: "2",
			input: []string{
				"10..9..",
				"2...8..",
				"3...7..",
				"4567654",
				"...8..3",
				"...9..2",
				".....01",
			},
			trailHeads: shared.Locations{
				shared.NewLocation(0, 1),
				shared.NewLocation(6, 5),
			},
			summits: shared.Locations{
				shared.NewLocation(0, 4),
				shared.NewLocation(5, 3),
			},
			trails: []shared.Locations{
				{
					shared.NewLocation(0, 1),
					shared.NewLocation(0, 0),
					shared.NewLocation(1, 0),
					shared.NewLocation(2, 0),
					shared.NewLocation(3, 0),
					shared.NewLocation(3, 1),
					shared.NewLocation(3, 2),
					shared.NewLocation(3, 3),
					shared.NewLocation(4, 3),
					shared.NewLocation(5, 3),
				},
				{
					shared.NewLocation(6, 5),
					shared.NewLocation(6, 6),
					shared.NewLocation(5, 6),
					shared.NewLocation(4, 6),
					shared.NewLocation(3, 6),
					shared.NewLocation(3, 5),
					shared.NewLocation(3, 4),
					shared.NewLocation(3, 3),
					shared.NewLocation(4, 3),
					shared.NewLocation(5, 3),
				},
				{
					shared.NewLocation(6, 5),
					shared.NewLocation(6, 6),
					shared.NewLocation(5, 6),
					shared.NewLocation(4, 6),
					shared.NewLocation(3, 6),
					shared.NewLocation(3, 5),
					shared.NewLocation(3, 4),
					shared.NewLocation(2, 4),
					shared.NewLocation(1, 4),
					shared.NewLocation(0, 4),
				},
			},
			scores:  []int{1, 2},
			ratings: []int{1, 2},
		},
		{
			name: "3",
			input: []string{
				"..90..9",
				"...1.98",
				"...2..7",
				"6543456",
				"765.987",
				"876....",
				"987....",
			},
			trailHeads: shared.Locations{
				shared.NewLocation(0, 3),
			},
			summits: shared.Locations{
				shared.NewLocation(0, 2),
				shared.NewLocation(0, 6),
				shared.NewLocation(1, 5),
				shared.NewLocation(4, 4),
				shared.NewLocation(6, 0),
			},
			trails: []shared.Locations{
				{
					shared.NewLocation(0, 3),
					shared.NewLocation(1, 3),
					shared.NewLocation(2, 3),
					shared.NewLocation(3, 3),
					shared.NewLocation(3, 3),
				},
				{
					shared.NewLocation(0, 3),
				},
			},
			scores:  []int{4},
			ratings: []int{13},
		},
		{
			name: "4",
			input: []string{
				"89010123",
				"78121874",
				"87430965",
				"96549874",
				"45678903",
				"32019012",
				"01329801",
				"10456732",
			},
			trailHeads: shared.Locations{
				shared.NewLocation(0, 2),
				shared.NewLocation(0, 4),
				shared.NewLocation(2, 4),
				shared.NewLocation(4, 6),
				shared.NewLocation(5, 2),
				shared.NewLocation(5, 5),
				shared.NewLocation(6, 0),
				shared.NewLocation(6, 6),
				shared.NewLocation(7, 1),
			},
			summits: shared.Locations{
				shared.NewLocation(0, 1),
				shared.NewLocation(2, 5),
				shared.NewLocation(3, 0),
				shared.NewLocation(3, 4),
				shared.NewLocation(4, 5),
				shared.NewLocation(5, 4),
				shared.NewLocation(6, 4),
			},
			trails: []shared.Locations{
				{
					shared.NewLocation(0, 1),
					shared.NewLocation(0, 3),
				},
				{
					shared.NewLocation(0, 3),
					shared.NewLocation(0, 3),
				},
			},
			scores:  []int{5, 6, 5, 3, 1, 3, 5, 3, 5},
			ratings: []int{20, 24, 10, 4, 1, 4, 5, 8, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			trailHeads := FindTrailHeads(tc.input)

			if !shared.AreEqual(trailHeads, tc.trailHeads) {
				t.Errorf("Expected %v trail heads but got %v", tc.trailHeads, trailHeads)
			}

			summits := FindSummits(tc.input)
			if !shared.AreEqual(summits, tc.summits) {
				t.Errorf("Expected %v summits but got %v", tc.summits, summits)
			}

			result := []struct {
				actual   int
				expected int
			}{{0, 0}, {0, 0}}
			for i, th := range trailHeads {
				summitLocations := shared.Locations{}
				rating := Venture(tc.input, shared.Locations{th}, &summitLocations)
				score := len(summitLocations)
				result[0].actual += score
				result[0].expected += tc.scores[i]
				result[1].actual += rating
				result[1].expected += tc.ratings[i]
				fmt.Printf("Summit locations: %v\n", summitLocations)
				if score != tc.scores[i] {
					t.Errorf("Expected score %v for trail head %v but got %v", tc.scores[i], th, score)
				}
			}

			if result[0].actual != result[0].expected {
				t.Errorf("Expected total score of %v but got %v", result[0].expected, result[0].actual)
			}

			if result[1].actual != result[1].expected {
				t.Errorf("Expected total rating of %v but got %v", result[1].expected, result[1].actual)
			}
		})
	}
}
