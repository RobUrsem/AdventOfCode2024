package puzzle

import (
	"math"

	lru "github.com/hashicorp/golang-lru/v2"
)

type Rule func(int) (bool, []int)

/*
If the stone is engraved with the number 0,
it is replaced by a stone engraved with the number 1.
*/
func ruleOne(stone int) (bool, []int) {
	if stone == 0 {
		return true, []int{1}
	}
	return false, nil
}

func numDigits(n int) int {
	if n < 0 {
		n = -n // Handle negative numbers
	}

	if n == 0 {
		return 1 // Special case for 0
	}

	return int(math.Log10(float64(n))) + 1
}

/*
If the stone is engraved with a number that has an even
number of digits, it is replaced by two stones. The left
half of the digits are engraved on the new left stone, and
the right half of the digits are engraved on the new right
stone. (The new numbers don't keep extra leading zeroes:
1000 would become stones 10 and 0.)
*/
func ruleTwo(stone int) (bool, []int) {
	numdigits := numDigits(stone)
	if numdigits%2 == 0 {
		half := numdigits / 2
		factor := math.Pow10(half)
		left := int(math.Trunc(float64(stone) / factor))
		right := stone - left*int(factor)
		return true, []int{left, right}
	}
	return false, nil
}

/*
If none of the other rules apply, the stone is replaced
by a new stone; the old stone's number multiplied by 2024
is engraved on the new stone.
*/
func ruleThree(stone int) (bool, []int) {
	return true, []int{stone * 2024}
}

// Memoization, we remember what we calculated before
// Outside class since we update the class contents
// not pretty but it works
var memo, _ = lru.New[int, []int](128)

/*
The Stone Counter stores how many times we have each
type of stone. Thus we only have to process each
unique type of stone once and just multiply the outcome
of each process with the number of times the stone
occurred.

When we combine this with the memoization, where we
actually don't process stone types we processed in the
past, the method becomes very fast
*/
type StoneCounter map[int]int

func MakeStoneCounter(arr []int) StoneCounter {
	s := make(StoneCounter)
	for _, num := range arr {
		s[num]++
	}
	return s
}

func (s StoneCounter) Update(arr []int, count int) {
	for _, v := range arr {
		s[v] += count
	}
}

func (s StoneCounter) Total() int {
	total := 0
	for _, count := range s {
		total += count
	}
	return total
}

func (s *StoneCounter) Blink() {
	rules := []Rule{ruleOne, ruleTwo, ruleThree}
	nextStones := StoneCounter{}

	for stone, count := range *s {
		if cache, ok := memo.Get(stone); ok {
			nextStones.Update(cache, count)
		} else {
			for _, rule := range rules {
				if used, result := rule(stone); used {
					memo.Add(stone, result)
					nextStones.Update(result, count)
					break
				}
			}
		}
	}

	*s = nextStones
}
