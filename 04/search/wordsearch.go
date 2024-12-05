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

func stringToArray(s string) []rune {
	return []rune(s)
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

func searchVertical(needle []rune, haystack []string) int {
	numRows := len(haystack)
	numColumns := len(haystack[0])
	needleLength := len(needle)

	count := 0

	for c := 0; c < numColumns; c++ {

		//--- Downward
		numMatches := 0
		for r := 0; r < numRows; r++ {
			if haystack[r][c] == byte(needle[numMatches]) {
				numMatches++
				if numMatches == needleLength {
					count++
					numMatches = 0
				}
			}
		}

		//--- Upward
		numMatches = 0
		for r := numRows - 1; r >= 0; r-- {
			if haystack[r][c] == byte(needle[numMatches]) {
				numMatches++
				if numMatches == needleLength {
					count++
					numMatches = 0
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

	needleArray := stringToArray(needle)
	count += searchVertical(needleArray, haystack)

	return count
}
