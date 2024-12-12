package puzzle

import "testing"

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
		name    string
		input   string
		lengths []int
		print   string
	}{
		{
			name:    "1",
			input:   "12345",
			lengths: []int{1, 3, 5},
			print:   "0..111....22222",
		},
		{
			name:    "2",
			input:   "2333133121414131402",
			lengths: []int{2, 3, 1, 3, 2, 4, 4, 3, 4, 2},
			print:   "00...111...2...333.44.5555.6666.777.888899",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lengths := GetFileLengths(tc.input)
			if !areEqual(lengths, tc.lengths) {
				t.Errorf("Expected lengths %v but got %v", tc.lengths, lengths)
			}
			output := Expand(tc.input)
			if output != tc.print {
				t.Errorf("Expected string \n%v but got \n%v", tc.print, output)
			}
		})
	}
}
