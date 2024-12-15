package puzzle

type Regions [][]int

func Outside(input []string, r, c int) bool {
	return r < 0 || c < 0 || r >= len(input) || c >= len(input[0])
}

func IsSame(input []string, r, c int, v rune) bool {
	return Outside(input, r, c) || input[r][c] == byte(v)
}

func Check(input []string, labels [][]int, char rune, r, c, dr, dc int) bool {
	if labels[r][c] == 0 &&
		!Outside(input, r+dr, c+dc) &&
		input[r+dr][c+dc] == byte(char) {
		labels[r][c] = labels[r+dr][c+dc]
		return true
	}

	return false
}

func Label(labels [][]int, r, c int) int {
	if r < 0 || c < 0 || r >= len(labels) || c >= len(labels[0]) {
		return 0
	}

	return labels[r][c]
}

func Plant(input []string, r, c int) byte {
	if r < 0 || c < 0 || r >= len(input) || c >= len(input[0]) {
		return 0
	}

	return input[r][c]
}

func Relabel(labels [][]int, from, to int) {
	for r, line := range labels {
		for c, char := range line {
			if char == from {
				labels[r][c] = to
			}
		}
	}
}

func Segment(input []string) Regions {
	labels := make([][]int, len(input))
	for r, line := range input {
		labels[r] = make([]int, len(line))
	}

	nextLabel := 1
	for r, line := range input {
		for c, char := range line {
			if !Check(input, labels, char, r, c, -1, 0) &&
				!Check(input, labels, char, r, c, 0, -1) {
				labels[r][c] = nextLabel
				nextLabel++
			}
		}
	}

	for r, line := range input {
		for c, char := range line {
			plant := Plant(input, r-1, c)
			label := Label(labels, r-1, c)
			if plant == byte(char) && label != labels[r][c] {
				Relabel(labels, label, labels[r][c])
			}
			plant = Plant(input, r, c-1)
			label = Label(labels, r, c-1)
			if plant == byte(char) && label != labels[r][c] {
				Relabel(labels, label, labels[r][c])
			}
		}
	}

	return labels
}

func IsEdge(regions Regions, region, r, c int) int {
	label := Label(regions, r, c)
	if label == 0 || label != region {
		return 1
	}
	return 0
}

func CalcCost(regions Regions) int {
	counts := map[int]int{}
	perimiter := map[int]int{}

	for r, line := range regions {
		for c, region := range line {
			counts[region]++
			perimiter[region] += IsEdge(regions, region, r-1, c)
			perimiter[region] += IsEdge(regions, region, r+1, c)
			perimiter[region] += IsEdge(regions, region, r, c-1)
			perimiter[region] += IsEdge(regions, region, r, c+1)
		}
	}

	total := 0
	for k, v := range counts {
		p, exists := perimiter[k]
		if exists {
			total += v * p
		}
	}

	return total
}

func isOutsideCorner(regions [][]int, region, r, c, dr, dc int) bool {
	return IsEdge(regions, region, r+dr, c)+IsEdge(regions, region, r, c+dc) == 2
}

func isInsidecorner(regions [][]int, region, r, c, dr, dc int) bool {
	numRow := len(regions)
	numCol := len(regions[0])

	diagR, diagC := r+dr, c+dc

	return diagR >= 0 && diagR < numRow &&
		diagC >= 0 && diagC < numCol &&
		IsEdge(regions, region, r+dr, c) == 0 &&
		IsEdge(regions, region, r, c+dc) == 0 &&
		Label(regions, diagR, diagC) != region
}

func CalcSides(regions Regions) int {
	counts := map[int]int{}
	sides := map[int]int{}

	directions := [][2]int{
		{-1, -1}, // NW
		{-1, 1},  // NE
		{1, -1},  // SW
		{1, 1},   // SE
	}

	for r, line := range regions {
		for c, region := range line {
			counts[region]++

			for _, dir := range directions {
				if isOutsideCorner(regions, region, r, c, dir[0], dir[1]) {
					sides[region]++
				}

				if isInsidecorner(regions, region, r, c, dir[0], dir[1]) {
					sides[region]++
				}
			}
		}
	}

	total := 0
	for k, v := range counts {
		s, exists := sides[k]
		if exists {
			total += v * s
		}
	}

	return total
}
