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

func searchForward(needle, haystack string) int {
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

func SearchForWord(needle string, haystack []string) int {
	rneedle := reverse(needle)
	count := 0
	for _, line := range haystack {
		count += searchForward(needle, line)
		count += searchForward(rneedle, line)
	}
	return count
}
