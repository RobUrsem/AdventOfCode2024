package puzzle

import (
	"fmt"
	"testing"
)

func areEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGetFiles(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		lengths    []int
		expanded   string
		compressed string
		checksum   int
	}{
		{
			name:       "1",
			input:      "12345",
			lengths:    []int{1, 3, 5},
			expanded:   "0..111....22222",
			compressed: "022111222......",
			checksum:   60, // 0*0+1*2+2*2+3*1+4*1+5*1+6*2+7*2+8*2 = 60
		},
		{
			name:       "2",
			input:      "2333133121414131402",
			lengths:    []int{2, 3, 1, 3, 2, 4, 4, 3, 4, 2},
			expanded:   "00...111...2...333.44.5555.6666.777.888899",
			compressed: "0099811188827773336446555566..............",
			checksum:   1928,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lengths := GetFileLengths(tc.input)
			if !areEqual(lengths, tc.lengths) {
				t.Errorf("Expected lengths %v but got %v", tc.lengths, lengths)
			}
			output := Expand(tc.input)
			if output != tc.expanded {
				t.Errorf("Expected string \n%v but got \n%v", tc.expanded, output)
			}

			compressed := Compress(output)
			if compressed != tc.compressed {
				t.Errorf("Expected string \n%v but got \n%v", tc.compressed, compressed)
			}

			checksum := Checksum(compressed)
			if checksum != tc.checksum {
				t.Errorf("Expected checksum %v but got %v", tc.checksum, checksum)
			}

			disk := Analyze(tc.input)
			fast := FastCompress(disk)
			for i, block := range fast {
				fmt.Printf("%02d [%v]\n", i, block)
			}
		})
	}
}
