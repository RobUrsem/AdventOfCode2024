package guard

import (
	"errors"
)

const (
	EMPTY = iota
	OBSTACLE
	VISITED
	GUARD_UP
	GUARD_LEFT
	GUARD_DOWN
	GUARD_RIGHT
)

type LabMap [][]int

var ErrInvalidCharacter = errors.New("invalid character")

func ConstructMap(lines []string) (LabMap, error) {
	var labMap LabMap

	for _, line := range lines {
		row := make([]int, len(line))

		for j, char := range line {
			switch char {
			case '.':
				row[j] = EMPTY
			case '#':
				row[j] = OBSTACLE
			case '^':
				row[j] = GUARD_UP
			case '>':
				row[j] = GUARD_RIGHT
			case 'V':
				row[j] = GUARD_DOWN
			case '<':
				row[j] = GUARD_LEFT
			case 'X':
				row[j] = VISITED
			default:
				return nil, ErrInvalidCharacter
			}
		}

		labMap = append(labMap, row)
	}

	return labMap, nil
}
