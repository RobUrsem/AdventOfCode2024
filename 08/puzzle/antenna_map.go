package puzzle

import (
	"fmt"
	"strings"
)

type Location []int
type Locations []Location

type AntennaPositions map[string]Locations
type AntennaMap []string

const antinodeBase = byte('#')

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

func (m AntennaMap) DetermineAntinodes() Locations {
	pos := m.getAntennaPositions()

	var antiNodes Locations
	for _, locs := range pos {
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j < len(locs); j++ {
				dx := locs[j][0] - locs[i][0]
				dy := locs[j][1] - locs[i][1]
				if !m.OutsideMap(locs[i][0]-dx, locs[i][1]-dy) {
					antiNodes = append(antiNodes, Location{locs[i][0] - dx, locs[i][1] - dy})
				}
				if !m.OutsideMap(locs[j][0]+dx, locs[j][1]+dy) {
					antiNodes = append(antiNodes, Location{locs[j][0] + dx, locs[j][1] + dy})
				}
			}
		}
	}

	return removeDuplicates(antiNodes)
}

func removeDuplicates(locations Locations) Locations {
	uniqueMap := make(map[string]struct{})
	var uniqueLocations Locations

	for _, loc := range locations {
		locKey := fmt.Sprint(loc)
		if _, exists := uniqueMap[locKey]; !exists {
			uniqueMap[locKey] = struct{}{}
			uniqueLocations = append(uniqueLocations, loc)
		}
	}

	if len(locations) != len(uniqueLocations) {
		fmt.Printf("Orig: %v, unique %v\n", len(locations), len(uniqueLocations))
	}

	return uniqueLocations
}

func replaceNthChar(s string, n int, newChar rune) string {
	runes := []rune(s)
	if n >= 0 && n < len(runes) {
		runes[n] = newChar
	}
	return string(runes)
}

func (a AntennaMap) IsAntenna(r, c int) bool {
	antenna := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	return strings.ContainsAny(antenna, string(a[r][c]))
}

func (a AntennaMap) AddAntiNode(r, c int) {
	current := a[r][c]
	if current == '.' || a.IsAntenna(r, c) {
		a[r] = replaceNthChar(a[r], c, rune(antinodeBase))
	} else {
		newAntiNode := antinodeBase + byte(current-antinodeBase+1)
		a[r] = replaceNthChar(a[r], c, rune(newAntiNode))
	}
}

func (a AntennaMap) Filter() {
	antiNodes := a.DetermineAntinodes()

	for _, loc := range antiNodes {
		a.AddAntiNode(loc[0], loc[1])
	}
}

func (a AntennaMap) String() string {
	return strings.Join(a, "\n")
}

func (a AntennaMap) CountAntiNodes() int {
	count := 0
	for _, line := range a {
		for _, char := range line {
			if char == rune(antinodeBase) {
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
