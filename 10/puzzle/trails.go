package puzzle

import (
	"advent/shared"
)

func FindTrailHeads(input []string) shared.Locations {
	return FindLocations(input, '0')
}

func FindSummits(input []string) shared.Locations {
	return FindLocations(input, '9')
}

func FindLocations(input []string, what rune) shared.Locations {
	trailHeads := shared.Locations{}
	for r, line := range input {
		for c, char := range line {
			if char == what {
				trailHeads = append(trailHeads, shared.NewLocation(r, c))
			}
		}
	}
	return trailHeads
}

func Outside(input []string, loc shared.Location) bool {
	return loc.R < 0 || loc.C < 0 || loc.R >= len(input) || loc.C >= len(input[0])
}

func GetNextSteps(input []string, loc shared.Location) shared.Locations {
	height := input[loc.R][loc.C]
	possible := shared.Locations{}

	for dr := -1; dr == 1; dr += 2 {
		for dc := -1; dc == 1; dc += 2 {
			p := shared.NewLocation(loc.R+dr, loc.C+dc)
			if !Outside(input, p) && input[p.R][p.C] == height+1 {
				possible = append(possible, p)
			}
		}
	}

	return possible
}

func Venture(input []string, trail shared.Locations, summits *shared.Locations) {
	options := shared.Locations{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for _, option := range options {
		Walk(input, trail, option, summits)
	}

	*summits = shared.RemoveDuplicates(*summits)
}

func Walk(input []string, trail shared.Locations, dir shared.Location, summits *shared.Locations) bool {
	from := trail[len(trail)-1]
	heightFrom := input[from.R][from.C]
	to := shared.NewLocation(from.R+dir.R, from.C+dir.C)
	if Outside(input, to) {
		return false
	}
	heightTo := input[to.R][to.C]

	//--- Have we reached the top?
	if heightTo == '9' && heightTo-heightFrom == 1 {
		// fmt.Printf("%v -> %v ", trail, to)
		// fmt.Println("-> At summit")
		*summits = append(*summits, to)
		return true
	}

	//--- Is the next position 1 higher?
	if heightTo-heightFrom == 1 {
		trail = append(trail, to)
		copy := trail.Copy()

		options := shared.Locations{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
		for _, option := range options {
			Walk(input, copy, option, summits)
		}
		// } else {
		// 	fmt.Println("-> Path ends")
	}
	return false
}
