package search

import (
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

func searchDiagonal(needle string, haystack []string) int {
	numRows := len(haystack)
	numColumns := len(haystack[0])
	needleLength := len(needle)

	count := 0

	//--- Downward NW-SE
	for r := 0; r <= numRows-needleLength; r++ {
		for c := 0; c <= numColumns-needleLength; c++ {
			numMatches := 0
			for offset := 0; offset < needleLength; offset++ {
				if haystack[r+offset][c+offset] == byte(needle[numMatches]) {
					numMatches++
					if numMatches == needleLength {
						count++
						numMatches = 0
					}
				}
			}
		}
	}

	//--- Downward NE-SW
	for r := 0; r <= numRows-needleLength; r++ {
		for c := numColumns - 1; c >= needleLength-1; c-- {
			numMatches := 0
			for offset := 0; offset < needleLength; offset++ {
				if haystack[r+offset][c-offset] == byte(needle[numMatches]) {
					numMatches++
					if numMatches == needleLength {
						count++
						numMatches = 0
					}
				}
			}
		}
	}

	//--- Upward SW-NE
	for r := numRows - 1; r >= needleLength-1; r-- {
		for c := 0; c <= numColumns-needleLength; c++ {
			numMatches := 0
			for offset := 0; offset < needleLength; offset++ {
				if haystack[r-offset][c+offset] == byte(needle[numMatches]) {
					numMatches++
					if numMatches == needleLength {
						count++
						numMatches = 0
					}
				}
			}
		}
	}

	//--- Upward SE-NW
	for r := numRows - 1; r >= needleLength-1; r-- {
		for c := numColumns - 1; c >= needleLength-1; c-- {
			numMatches := 0
			for offset := 0; offset < needleLength; offset++ {
				if haystack[r-offset][c-offset] == byte(needle[numMatches]) {
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
