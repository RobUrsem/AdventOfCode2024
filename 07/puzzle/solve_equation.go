package puzzle

import (
	"fmt"
	"strconv"
)

func generatePermutationsForTwoOperations(n int) [][]int {
	var result [][]int

	// Total number of permutations is 2^n (since we have 2 operations to choose from in each of the n positions)
	totalPermutations := 1 << n

	// Iterate over all possible numbers from 0 to 2^n - 1 (binary representation of permutations)
	for i := 0; i < totalPermutations; i++ {
		var permutation []int
		for j := 0; j < n; j++ {
			// Use bitwise AND to decide which operation to use
			if (i>>j)&1 == 0 {
				permutation = append(permutation, ADD)
			} else {
				permutation = append(permutation, MUL)
			}
		}
		result = append(result, permutation)
	}

	return result
}

func intPow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp%2 != 0 {
			result *= base
		}
		base *= base
		exp /= 2
	}
	return result
}

func generatePermutations(n int, operations []int) [][]int {
	var result [][]int

	totalPermutations := 1
	for i := 0; i < n; i++ {
		totalPermutations *= len(operations)
	}

	for i := 0; i < totalPermutations; i++ {
		var permutation []int
		for j := 0; j < n; j++ {
			index := (i / intPow(len(operations), j)) % len(operations)
			permutation = append(permutation, operations[index])
		}
		result = append(result, permutation)
	}

	return result
}

func SolveEquation(equation Equation, operations []int) Equation {
	numSpaces := len(equation.Coefficients) - 1
	permutations := generatePermutations(numSpaces, operations)

	for _, permutation := range permutations {
		answer := equation.Coefficients[0]
		for i := 1; i < len(equation.Coefficients); i++ {
			switch permutation[i-1] {
			case ADD:
				answer = answer + equation.Coefficients[i]
			case MUL:
				answer = answer * equation.Coefficients[i]
			case CAT:
				answerStr := strconv.FormatInt(answer, 10)
				coeffStr := strconv.Itoa(int(equation.Coefficients[i]))
				answerStr = answerStr + coeffStr
				answerTmp, err := strconv.ParseInt(answerStr, 10, 64)
				if err != nil {
					fmt.Printf("Cannot convert [%v] to int64\n", answerStr)
				} else {
					answer = answerTmp
				}
			}
		}

		if answer == equation.Answer {
			equation.Valid = true
			equation.Operators = permutation
			return equation
		}
	}

	return equation
}
