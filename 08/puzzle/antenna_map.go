package puzzle

type AntennaMap []string

const (
	ANTINODE rune = '#'
	EMPTY    rune = '.'
)

func CreateMap(lines []string) AntennaMap {
	var theMap AntennaMap

	for _, line := range lines {
		theMap = append(theMap, line)
	}

	return theMap
}

func (m AntennaMap) DetermineAntinodes() {

}

func (a AntennaMap) IsSameAs(b AntennaMap) bool {
	if len(a) != len(b) {
		return false
	}

	for i, line := range a {
		if line != b[i] {
			return false
		}
	}

	return true
}
