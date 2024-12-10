package puzzle

import (
	"fmt"
	"regexp"
	"strconv"
)

func GetEquations(lines []string) Equations {
	var equations Equations

	pattern := `(\d+):|(\d+)`

	rx, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	for _, line := range lines {
		matches := rx.FindAllStringSubmatch(line, -1)
		equation := Equation{}
		for i, match := range matches {
			if i == 0 {
				num, err := strconv.ParseInt(match[1], 10, 64)
				if err == nil {
					equation.answer = num
				}
			} else {
				num, err := strconv.ParseInt(match[2], 10, 64)
				if err == nil {
					equation.coefficients = append(equation.coefficients, num)
				}
			}
		}
		equations = append(equations, equation)
	}

	return equations
}
