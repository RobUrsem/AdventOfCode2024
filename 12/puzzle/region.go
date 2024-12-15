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

func CalcCost(regions Regions) int {
	return 42
}
