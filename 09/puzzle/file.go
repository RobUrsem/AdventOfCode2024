package puzzle

import (
	"strconv"
)

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
			b[lastBlockIdx].length -= numToMove
		} else {
			//--- split block
			splitSize := b[firstSpaceIdx].length - numToMove
			b[firstSpaceIdx].id = b[lastBlockIdx].id
			b[firstSpaceIdx].length = numToMove
			b[lastBlockIdx].id = -1
			b[lastBlockIdx].length -= numToMove
			newBlock := Block{-1, splitSize}
			b = append(b[:firstSpaceIdx+1], append([]Block{newBlock}, b[firstSpaceIdx+1:]...)...)
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

func FastChecksum(b []Block) int {
	idx := 0
	count := 0
	for _, block := range b {
		for i := 0; i < block.length; i++ {
			count += idx * block.id
			idx++
		}
	}
	return count
}
