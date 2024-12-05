package search

import "strings"

func SearchForWord(needle string, haystack []string) int {
	count := 0
	for _, line := range haystack {
		if strings.HasPrefix(line, needle) {
			count++
		}
	}
	return count
}
