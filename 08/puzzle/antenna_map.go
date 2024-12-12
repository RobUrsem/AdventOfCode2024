package puzzle

import "strings"

type Location []int
type Locations []Location

type AntennaPositions map[string]Locations
type AntennaMap []string

const (
	ANTINODE byte = '#'
	EMPTY    byte = '.'
)

func CreateMap(lines []string) AntennaMap {
	var theMap AntennaMap

	for _, line := range lines {
		theMap = append(theMap, line)
	}

	return theMap
}

func (m AntennaMap) getAntennaPositions() AntennaPositions {
	allowed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	pos := make(AntennaPositions)
	for i, line := range m {
		for j, char := range line {
			if strings.ContainsRune(allowed, char) {
				pos[string(char)] = append(pos[string(char)], Location{i, j})
			}
		}
	}
	return pos
}

func (m AntennaMap) OutsideMap(r, c int) bool {
	return r < 0 || c < 0 || r >= len(m) || c >= len(m[0])
}

func (m AntennaMap) DetermineAntinodes() AntennaPositions {
	pos := m.getAntennaPositions()

	antiNodes := make(AntennaPositions)
	for char, locs := range pos {
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j < len(locs); j++ {
				dx := locs[j][0] - locs[i][0]
				dy := locs[j][1] - locs[i][1]
				if !m.OutsideMap(locs[i][0]-dx, locs[i][1]-dy) {
					antiNodes[char] = append(antiNodes[char], Location{locs[i][0] - dx, locs[i][1] - dy})
				}
				if !m.OutsideMap(locs[j][0]+dx, locs[j][1]+dy) {
					antiNodes[char] = append(antiNodes[char], Location{locs[j][0] + dx, locs[j][1] + dy})
				}
			}
		}
	}
	return antiNodes
}

func replaceNthChar(s string, n int, newChar rune) string {
	runes := []rune(s)
	if n >= 0 && n < len(runes) {
		runes[n] = newChar
	}
	return string(runes)
}

func (a AntennaMap) Filter() {
	antiNodes := a.DetermineAntinodes()
	for _, locs := range antiNodes {
		for _, loc := range locs {
			current := a[loc[0]][loc[1]]
			if current == EMPTY {
				a[loc[0]] = replaceNthChar(a[loc[0]], loc[1], rune(ANTINODE))
			}
		}
	}
}

func (a AntennaMap) String() string {
	return strings.Join(a, "\n")
}

func (a AntennaMap) CountAntiNodes() int {
	count := 0
	for _, line := range a {
		for _, char := range line {
			if char == rune(ANTINODE) {
				count++
			}
		}
	}
	return count
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
