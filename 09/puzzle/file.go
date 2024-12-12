package puzzle

import (
	"fmt"
	"strconv"
	"strings"
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

func getFirstSpace(s string) int {
	for i := 0; i < len(s); i++ {
		if s[i] == '.' {
			return i
		}
	}
	return -1
}

func getLastNumber(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if strings.ContainsAny(string(s[i]), "0123456789") {
			return i
		}
	}
	return -1
}

func replaceNthChar(s string, n int, newChar rune) string {
	runes := []rune(s)
	if n >= 0 && n < len(runes) {
		runes[n] = newChar
	}
	return string(runes)
}

func EmptySpacesAtTheEnd(disk string, expected int) bool {
	start := len(disk) - 1
	end := start - expected
	for i := start; i >= end; i-- {
		if disk[i] != '.' {
			return false
		}
	}
	//--- make sure the next one isn't a space
	return disk[end-1] != '.'
}

func MoveBlock(disk string) (string, bool) {
	//--- find last number
	lastPos := getLastNumber(disk)

	//--- find first space
	firstPos := getFirstSpace(disk)

	//--- If index of last number < index of first space => done
	if lastPos < firstPos {
		return disk, true
	}

	//--- replace first space with last number
	//--- replace last number with a space
	disk = replaceNthChar(disk, firstPos, rune(disk[lastPos]))
	disk = replaceNthChar(disk, lastPos, '.')

	return disk, false
}

func Compress(expanded string) string {
	disk := expanded
	done := false

	for {
		disk, done = MoveBlock(disk)
		// fmt.Printf("%v\n", disk)
		if done {
			break
		}
	}

	return disk
}

func Checksum(disk string) int {
	checksum := 0

	for i, c := range disk {
		if c != '.' {
			v, _ := strconv.Atoi(string(c))
			checksum += i * v
		}
	}

	return checksum
}

/*
--------------------------------------------------------------------------
12345 => input string
1,3,5 => length of blocks
2,4 => length of spaces
==> Total of 6 spaces
Files:

	idx, len
	0,   1
	5,   2
	3,   3
	5,   3
	.,   6

steps:
12345 => disk[(0,1), (-1,2), (1,3), (-1,4), (2,5)]   = 0..111....22222
fill first spaces min (disk[1][1], disk[4][1]) => 2
disk[(0,1), (2,2), (1,3), (-1,4), (2,3), (-1,2)]     = 022111....222..
discard trailing spaces
disk[(0,1), (2,2), (1,3), (-1,4), (2,3)]             = 022111....222
fill second space min (disk[3][1], disk[4][1]) => 3
disk[(0,1), (2,2), (1,3), (2,3), (-1,1), (-1,5)]     = 022111222......
discard trailing spaces
disk[(0,1), (2,2), (1,3), (2,3)]                     = 022111222
checksum:                                              012345678
0*0 + 2*1+2*2 + 1*3+1*4+1*5 + 2*6+2*7+2*8 = 60
--------------------------------------------------------------------------
*/
type Block struct {
	id     int
	length int
}

func (b Block) IsSpace() bool {
	return b.id < 0
}

func Analyze(disk string) []Block {
	blocks := []Block{}

	for i := 0; i < len(disk); i++ {
		v, _ := strconv.Atoi(string(disk[i]))
		if i%2 == 0 {
			blocks = append(blocks, Block{i / 2, v})
		} else {
			blocks = append(blocks, Block{-1, v})
		}
	}

	return blocks
}

func findFirstSpaceBlock(b []Block) (Block, int, bool) {
	for i, block := range b {
		if block.IsSpace() {
			return block, i, true
		}
	}
	return Block{}, -1, false
}

func FillSpace(b []Block) []Block {
	lastBlockIdx := len(b) - 1
	firstSpaceBlock, firstSpaceIdx, found := findFirstSpaceBlock(b)
	if found {
		numToMove := min(firstSpaceBlock.length, b[lastBlockIdx].length)
		if numToMove >= firstSpaceBlock.length {
			//--- move from last block to first space
			b[firstSpaceIdx].id = b[lastBlockIdx].id
			// b[firstSpaceIdx][1] should be numToMove
			//--- reduce last block
			b[lastBlockIdx].length -= numToMove
		}
	}

	return b
}

func HasFragmentedSpace(b []Block) bool {
	for _, block := range b {
		if block.IsSpace() {
			return true
		}
	}
	return false
}

func DiscardTrailingSpaces(blocks []Block) []Block {
	//--- remove trailing spaces
	for len(blocks) > 0 && blocks[len(blocks)-1].IsSpace() {
		blocks = blocks[:len(blocks)-1]
	}

	//--- Remove trailing empty blocks
	for len(blocks) > 0 && blocks[len(blocks)-1].length == 0 {
		blocks = blocks[:len(blocks)-1]
	}

	return blocks
}

func FastCompress(b []Block) []Block {
	for HasFragmentedSpace(b) {
		b = FillSpace(b)
		b = DiscardTrailingSpaces(b)
	}
	return b
}
