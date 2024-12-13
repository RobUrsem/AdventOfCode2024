package puzzle

import (
	"strconv"
)

const (
	SPACE = -1
)

/*
--------------------------------------------------------------------------
12345 => input string
1,3,5 => length of blocks
2,4 => length of spaces
==> Total of 6 spaces
Files:

	i	idx, len
	0	 0,   1
	1   -1,   2
	2	 1,   3
	3   -1,   4
	4 	 2,   5

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
			blocks = append(blocks, Block{SPACE, v})
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
			b[lastBlockIdx].id = SPACE
			b[lastBlockIdx].length -= numToMove
			newBlock := Block{SPACE, splitSize}
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

func MergeBlocks(blocks []Block) []Block {
	for i := 0; i < len(blocks)-1; i++ {
		if blocks[i].id == blocks[i+1].id {
			blocks[i].length += blocks[i+1].length
			blocks = append(blocks[:i+1], blocks[i+2:]...)
			i--
		}
	}
	return blocks
}

func FastCompress(b []Block) []Block {
	for HasFragmentedSpace(b) {
		b = FillSpace(b)
		b = DiscardTrailingSpaces(b)
		b = MergeBlocks(b)
	}
	return b
}

func FastChecksum(b []Block) int {
	idx := 0
	count := 0
	for _, block := range b {
		for i := 0; i < block.length; i++ {
			if !block.IsSpace() {
				count += idx * block.id
			}
			idx++
		}
	}
	return count
}

/*
--------------------------------------------------------------------------
This time, attempt to move whole files to the leftmost span of free space
blocks that could fit the file. Attempt to move each file exactly once in
order of decreasing file ID number starting with the file with the highest
file ID number. If there is no span of free space to the left of a file
that is large enough to fit the file, the file does not move.

00...111...2...333.44.5555.6666.777.888899      (0,2),(-1,3),(1,3),(-1,3),(2,1),(-1,3),(3,3),(-1,1),(4,2),(-1,1),(5,4),(-1,1),(6,4),(-1,1),(7,3),(-1,1),(8,4),(9,2)
0099.111...2...333.44.5555.6666.777.8888    9   (0,2),(9,2),(-1,1),(1,3),(-1,3),(2,1),(-1,3),(3,3),(-1,1),(4,2),(-1,1),(5,4),(-1,1),(6,4),(-1,1),(7,3),(-1,1),(8,4)
8 is bigger than any space block to the left
0099.1117772...333.44.5555.6666.....8888    7   (0,2),(9,2),(-1,1),(1,3),(7,3),(2,1),(-1,3),(3,3),(-1,1),(4,2),(-1,1),(5,4),(-1,1),(6,4),(-1,1),(-1,1),(8,4)
6 is bigger than any space block to the left
5 is bigger than any space block to the left
0099.111777244.333....5555.6666.....8888    4   (0,2),(9,2),(-1,1),(1,3),(7,3),(2,1),(-1,3),(3,3),(-1,1),(4,2),(-1,1),(5,4),(-1,1),(6,4),(-1,1),(-1,1),(8,4)
3 is bigger than any space block to the left
00992111777.44.333....5555.6666.....8888    2   (0,2),(9,2),(2,1),(1,3),(7,3),(-1,1),(4,2),(-1,1),(3,3),(-1,4),(5,4),(-1,1),(6,4),(-1,5),(8,4)

- find block with highest id
- find first space that fits that block
- if found -> move block
- until no block can be moved

12345
(0,1), (-1,2), (1,3), (-1,4), (2,5)     => 0..111....22222
cannot be defragged                        012345678901234
      0*0 + 1*3+1*4+1*5 + 2*10+2*11+2*12+2*13+2*14
--------------------------------------------------------------------------
*/

func FindIdxToMove(disk []Block, fileId int) int {
	for idx := len(disk) - 1; idx >= 0; idx-- {
		if disk[idx].id == fileId {
			return idx
		}
	}
	return -1
}

func Defrag(disk []Block) []Block {
	//--- Find highest file id
	fileId := disk[len(disk)-1].id
	if disk[len(disk)-1].IsSpace() {
		fileId = disk[len(disk)-2].id
	}

	for ; fileId >= 0; fileId-- {
		idxToMove := FindIdxToMove(disk, fileId)
		for spaceIdx := 0; spaceIdx < idxToMove; spaceIdx++ {
			if disk[spaceIdx].IsSpace() && disk[spaceIdx].length >= disk[idxToMove].length {
				if disk[idxToMove].length == disk[spaceIdx].length {
					// fmt.Printf("Moving block at index %v %v to index %v %v\n", idxToMove, disk[idxToMove], spaceIdx, disk[spaceIdx])
					disk[spaceIdx].id = disk[idxToMove].id
					disk[idxToMove].id = SPACE
				} else { // Space > Needed
					// fmt.Printf("Moving block at index %v %v to index %v %v and splitting space block\n", idxToMove, disk[idxToMove], spaceIdx, disk[spaceIdx])
					splitSize := disk[spaceIdx].length - disk[idxToMove].length
					disk[spaceIdx].id = disk[idxToMove].id
					disk[spaceIdx].length = disk[idxToMove].length
					disk[idxToMove].id = SPACE
					newBlock := Block{SPACE, splitSize}
					disk = append(disk[:spaceIdx+1], append([]Block{newBlock}, disk[spaceIdx+1:]...)...)
				}
				disk = MergeBlocks(disk)
				disk = DiscardTrailingSpaces(disk)
				break
			}
		}
	}

	return disk
}
