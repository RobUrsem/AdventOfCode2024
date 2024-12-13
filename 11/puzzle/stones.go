package puzzle

import "math"

/*
If the stone is engraved with the number 0,
it is replaced by a stone engraved with the number 1.
*/
func applyRuleOne(stones []int, i int) bool {
	if stones[i] == 0 {
		stones[i] = 1
		return true
	}

	return false
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
		return true
	}
	return false
}

/*
If none of the other rules apply, the stone is replaced
by a new stone; the old stone's number multiplied by 2024
is engraved on the new stone.
*/
func applyRuleThree(stones []int, i int) bool {
	stones[i] *= 2024
	return true
}

func Blink(stones []int) []int {
	for i := 0; i < len(stones); i++ {
		if !applyRuleOne(stones, i) {
			if !applyRuleTwo(&stones, &i) {
				applyRuleThree(stones, i)
			}
		}
	}

	return stones
}
