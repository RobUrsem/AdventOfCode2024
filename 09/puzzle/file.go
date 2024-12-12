package puzzle

import (
	"fmt"
	"strconv"
)

func GetFileLengths(input string) []int {
	lengths := []int{}
	for i := 0; i < len(input); i += 2 {
		v, _ := strconv.Atoi(string(input[i]))
		lengths = append(lengths, v)
	}

	return lengths
}

func Expand(a string) string {
	var output string

	x := true
	count := 0
	for i := 0; i < len(a); i++ {
		v, _ := strconv.Atoi(string(a[i]))
		if x {
			for j := 0; j < v; j++ {
				output += fmt.Sprint(count)
			}
			count++
		} else {
			for j := 0; j < v; j++ {
				output += "."
			}
		}
		x = !x
	}

	return output
}
