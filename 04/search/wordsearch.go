package search

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func reverse(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

func searchHorizontal(needle, haystack string) int {
	count := 0
	offset := 0
	for {
		index := strings.Index(haystack[offset:], needle)
		if index == -1 {
			break
		}

		count++
		offset = offset + index + 1
	}

	return count
}

func searchVertical(needle string, haystack []string) int {
	numRows := len(haystack)
	numColumns := len(haystack[0])
	needleLength := len(needle)

	count := 0

	for c := 0; c < numColumns; c++ {

		//--- Downward
		for r := 0; r <= numRows-needleLength; r++ {
			numMatches := 0
			for offset := 0; offset < needleLength; offset++ {
				if haystack[r+offset][c] == byte(needle[numMatches]) {
					numMatches++
					if numMatches == needleLength {
						count++
						numMatches = 0
					}
				}
			}
		}

		//--- Upward
		for r := numRows - 1; r >= needleLength-1; r-- {
			numMatches := 0
			for offset := 0; offset < needleLength; offset++ {
				if haystack[r-offset][c] == byte(needle[numMatches]) {
					numMatches++
					if numMatches == needleLength {
						count++
						numMatches = 0
					}
				}
			}
		}
	}

	return count
}

type BoundaryData struct {
	Start, Step, Stop int
}
type SearchData struct {
	row, col    BoundaryData
	findOnlyOne bool
}

func outsideGrid(row, col int, haystack []string) bool {
	if row < 0 || col < 0 {
		return true
	}

	numRows := len(haystack)
	if row >= numRows {
		return true
	}

	numColumns := len(haystack[0])
	return col >= numColumns
}

func searchDirection(data SearchData, needle string, haystack []string) int {
	needleLength := len(needle)

	count := 0

	if outsideGrid(data.row.Start, data.col.Start, haystack) ||
		outsideGrid(data.row.Stop, data.col.Stop, haystack) {
		return count
	}

	for r := data.row.Start; r != data.row.Stop; r += data.row.Step {
		for c := data.col.Start; c != data.col.Stop; c += data.col.Step {
			numMatches := 0
			for offset := 0; offset < needleLength; offset++ {
				row := r + offset*data.row.Step
				col := c + offset*data.col.Step
				if outsideGrid(row, col, haystack) {
					if data.findOnlyOne {
						return 0
					}
					break
				}
				if haystack[row][col] == needle[numMatches] {
					numMatches++
					if numMatches == needleLength {
						count++
						if data.findOnlyOne {
							return count
						}
						break
					}
				} else {
					if data.findOnlyOne {
						return 0
					}
					break
				}
			}
		}
	}

	return count
}

func searchDiagonal(needle string, haystack []string) int {
	numRows := len(haystack)
	numColumns := len(haystack[0])
	needleLength := len(needle)

	//--- Downward NW-SE
	count := searchDirection(SearchData{
		row: BoundaryData{Start: 0, Step: 1, Stop: numRows - needleLength + 1},
		col: BoundaryData{Start: 0, Step: 1, Stop: numColumns - needleLength + 1},
	}, needle, haystack)

	//--- Downward NE-SW
	count += searchDirection(SearchData{
		row: BoundaryData{Start: 0, Step: 1, Stop: numRows - needleLength + 1},
		col: BoundaryData{Start: numColumns - 1, Step: -1, Stop: needleLength - 2},
	}, needle, haystack)

	//--- Upward SW-NE
	count += searchDirection(SearchData{
		row: BoundaryData{Start: numRows - 1, Step: -1, Stop: needleLength - 2},
		col: BoundaryData{Start: 0, Step: 1, Stop: numColumns - needleLength + 1},
	}, needle, haystack)

	//--- Upward SE-NW
	count += searchDirection(SearchData{
		row: BoundaryData{Start: numRows - 1, Step: -1, Stop: needleLength - 2},
		col: BoundaryData{Start: numColumns - 1, Step: -1, Stop: needleLength - 2},
	}, needle, haystack)

	return count
}

func SearchForWord(needle string, haystack []string) int {
	count := 0

	rneedle := reverse(needle)
	for _, line := range haystack {
		count += searchHorizontal(needle, line)
		count += searchHorizontal(rneedle, line)
	}

	count += searchVertical(needle, haystack)

	count += searchDiagonal(needle, haystack)

	return count
}

type NeedleData struct {
	needle     string
	centerChar byte
	halfLength int
}

func testForCross(needleData NeedleData, haystack []string, row, col int) bool {
	//--- Downward NW-SE
	count := searchDirection(SearchData{
		row:         BoundaryData{Start: row - needleData.halfLength, Step: 1, Stop: row + needleData.halfLength},
		col:         BoundaryData{Start: col - needleData.halfLength, Step: 1, Stop: col + needleData.halfLength},
		findOnlyOne: true,
	}, needleData.needle, haystack)

	//--- Downward NE-SW
	count += searchDirection(SearchData{
		row:         BoundaryData{Start: row - needleData.halfLength, Step: 1, Stop: row + needleData.halfLength},
		col:         BoundaryData{Start: col + needleData.halfLength, Step: -1, Stop: col - needleData.halfLength},
		findOnlyOne: true,
	}, needleData.needle, haystack)

	//--- Upward SW-NE
	count += searchDirection(SearchData{
		row:         BoundaryData{Start: row + needleData.halfLength, Step: -1, Stop: row - needleData.halfLength},
		col:         BoundaryData{Start: col - needleData.halfLength, Step: 1, Stop: col + needleData.halfLength},
		findOnlyOne: true,
	}, needleData.needle, haystack)

	//--- Upward SE-NW
	count += searchDirection(SearchData{
		row:         BoundaryData{Start: row + needleData.halfLength, Step: -1, Stop: row - needleData.halfLength},
		col:         BoundaryData{Start: col + needleData.halfLength, Step: -1, Stop: col - needleData.halfLength},
		findOnlyOne: true,
	}, needleData.needle, haystack)

	return count == 2
}

func SearchForCross(needle string, haystack []string) (int, error) {
	numRows := len(haystack)
	numColumns := len(haystack[0])
	needleLength := len(needle)

	if needleLength%2 == 0 {
		return 0, fmt.Errorf("needle length needs to be odd")
	}

	needleData := NeedleData{
		needle:     needle,
		halfLength: (needleLength - 1) / 2,
		centerChar: needle[(needleLength-1)/2],
	}

	count := 0
	for row := needleData.halfLength; row < numRows-needleData.halfLength; row++ {
		for col := needleData.halfLength; col < numColumns-needleData.halfLength; col++ {
			if haystack[row][col] == needleData.centerChar {
				if testForCross(needleData, haystack, row, col) {
					count++
				}
			}
		}
	}

	return count, nil
}
