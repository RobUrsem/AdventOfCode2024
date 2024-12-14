package puzzle

import (
	"fmt"
	"math"
)

/*
If the stone is engraved with the number 0,
it is replaced by a stone engraved with the number 1.
*/
func applyRuleOne(stones []int, i int) bool {
	if stones[i] == 0 {
		stones[i] = 1
		// fmt.Print("1 ")
		return true
	}

	return false
}

func altApplyRuleOne(stone int) (bool, []int) {
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
func applyRuleTwo(stones *[]int, i *int) bool {
	value := (*stones)[*i]
	numdigits := numDigits(value)
	if numdigits%2 == 0 {
		half := numdigits / 2
		factor := math.Pow10(half)
		(*stones)[*i] = int(math.Trunc(float64(value) / factor))
		right := []int{value - (*stones)[*i]*int(factor)}

		*stones = append((*stones)[:*i+1], append(right, (*stones)[*i+1:]...)...)
		*i++
		// fmt.Print("2 ")
		return true
	}
	return false
}

func altApplyRuleTwo(stone int) (bool, []int) {
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
func applyRuleThree(stones []int, i int) bool {
	stones[i] *= 2024
	// fmt.Print("3 ")
	return true
}

func altApplyRuleThree(stone int) (bool, []int) {
	return true, []int{stone * 2024}
}

func Blink(stones []int) []int {
	for i := 0; i < len(stones); i++ {
		if !applyRuleOne(stones, i) {
			if !applyRuleTwo(&stones, &i) {
				applyRuleThree(stones, i)
			}
		}
	}

	fmt.Printf(" -> %v\n", len(stones))
	return stones
}

func RecursiveBlink(stones []int) int {
	for i := 0; i < len(stones); i++ {

	}
	return 42
}

// https://cp-algorithms.com/
