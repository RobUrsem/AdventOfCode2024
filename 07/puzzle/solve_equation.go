package puzzle

func generatePermutations(n int) [][]int {
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

func SolveEquation(equation Equation) Equation {
	numSpaces := len(equation.Coefficients) - 1
	permutations := generatePermutations(numSpaces)

	for _, permutation := range permutations {
		answer := equation.Coefficients[0]
		for i := 1; i < len(equation.Coefficients); i++ {
			switch permutation[i-1] {
			case ADD:
				answer = answer + equation.Coefficients[i]
			case MUL:
				answer = answer * equation.Coefficients[i]
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
