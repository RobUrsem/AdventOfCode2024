package puzzle

import (
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
			name:     "1",
			input:    "12345",
			checksum: 60,
		},
		{
			name:     "2",
			input:    "2333133121414131402",
			checksum: 1928,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			disk := Analyze(tc.input)
			fast := FastCompress(disk)
			checksum := FastChecksum(fast)

			if checksum != tc.checksum {
				t.Errorf("Expected checksum %v but got %v", tc.checksum, checksum)
			}
		})
	}
}
