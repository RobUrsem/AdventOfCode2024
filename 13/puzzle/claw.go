package puzzle

import (
	"math"
	"regexp"
	"strconv"
)

type Game struct {
	A, B, Prize []int
}

type Games []Game

func getNumber(num string) int {
	value, _ := strconv.Atoi(num)
	return value
}

func GetInput(lines []string) Games {
	rxButton := regexp.MustCompile(`^Button (.):\sX\+(\d+),\sY\+(\d+)`)
	rxPrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	games := Games{}
	currentGame := Game{}
	for _, line := range lines {
		if len(line) == 0 {
			games = append(games, currentGame)
			currentGame = Game{}
		} else {
			match := rxButton.FindStringSubmatch(line)
			if match != nil {
				if match[1] == "A" {
					currentGame.A = []int{getNumber(match[2]), getNumber(match[3])}
				} else {
					currentGame.B = []int{getNumber(match[2]), getNumber(match[3])}
				}
			}

			match = rxPrize.FindStringSubmatch(line)
			if match != nil {
				currentGame.Prize = []int{getNumber(match[1]), getNumber(match[2])}
			}
		}
	}

	games = append(games, currentGame)
	return games
}

func isInteger(f float64) bool {
	return math.Mod(f, 1) == 0 // Check if there's no remainder
}

func (game Game) Cost() int {
	/* Solve
		Prize[0] = x A[0] + y B[0]
		Prize[1] = x A[1] + y B[1]

		where x and y are the number of times
		we press the A and the B buttons

		     A   *   x  =     B
		|A[0] B[0]| |x|   |Prize[0]|
		|A[1] B[1]| |y| = |Prize[1]|

		With matrix A : | a b |  a = A0 b = B0
		                | c d |  c = A1 d = B1
	    And vector B as | e |
		                | f |

		determinant: ad - bc
		 -1          | d -b |
		A    = 1/det |-c  a |

		find x by multiplying A^-1 * B
		x = (d * e - b * f) / det
	    y = (a * f - c * e) / det

		94 A + 22 B = 8400
		34 A + 67 B = 5400
		det = 94*+67-34*22 = 5500
		x = ()
	*/

	det := float64(game.A[0]*game.B[1] - game.A[1]*game.B[0])
	x := float64(float64(game.B[1]*game.Prize[0]-game.B[0]*game.Prize[1]) / det)
	y := float64(game.A[0]*game.Prize[1]-game.A[1]*game.Prize[0]) / det

	if !isInteger(x) || !isInteger(y) {
		// invalid solution
		return -1
	} else {
		return 3*int(x) + int(y)
	}
}
