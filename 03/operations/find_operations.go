package operations

import (
	"fmt"
	"regexp"
	"strconv"
)

func handleMultiplication(match []string) (Operation, error) {
	a, err := strconv.Atoi(match[1])
	if err != nil {
		return Operation{}, fmt.Errorf("%v: can't convert [%v] to int: %v\n", match[0], match[1], err)
	}

	b, err := strconv.Atoi(match[2])
	if err != nil {
		return Operation{}, fmt.E("%v: can't convert [%v] to int: %v\n", match[0], match[2], err)
	}

	return Operation{
		Params:        []int{a, b},
		OperationType: Multiply,
	}, nil
}

func handleEnable(match []string) (Operation, error) {
	return Operation{
		Params:        []int{},
		OperationType: On,
	}, nil
}

func handleDisable(match []string) (Operation, error) {
	return Operation{
		Params:        []int{},
		OperationType: Off,
	}, nil
}

func findMultiplications(text string) []Operation {
	pattern := `(mul)\((\d{1,3}),(\d{1,3})\)|(do)\(\)|(don't)\(\)`

	rx, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	ignoreOperation := false
	operations := []Operation{}
	matches := rx.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		if ignoreOperation {
			continue
		}

		op := Operation{}

		switch match[0] {
		case "mul":
			op, err = handleMultiplication(match)
		case "do":
			op, err = handleEnable(match)
		case "don't":
			op, err = handleDisable(match)
		}

		if err != nil {
			operations = append(operations, op)
		}
	}

	return operations
}

func FindOperations(text string) ([]Operation, error) {
	operations := findMultiplications(text)
	return operations, nil
}
