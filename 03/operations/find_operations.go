package operations

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func handleMultiplication(str_a string, str_b string) (Operation, error) {
	a, err := strconv.Atoi(str_a)
	if err != nil {
		return Operation{}, fmt.Errorf("can't convert [%v] to int: %v", str_a, err)
	}

	b, err := strconv.Atoi(str_b)
	if err != nil {
		return Operation{}, fmt.Errorf("can't convert [%v] to int: %v", str_b, err)
	}

	return Operation{
		Params:        []int{a, b},
		OperationType: Multiply,
	}, nil
}

func handleEnable() (Operation, error) {
	return Operation{
		Params:        nil,
		OperationType: Enable,
	}, nil
}

func handleDisable() (Operation, error) {
	return Operation{
		Params:        nil,
		OperationType: Disable,
	}, nil
}

func FindOperations(text string) []Operation {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`

	rx, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	operations := []Operation{}
	op := Operation{}

	matches := rx.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		operatorType := match[0]
		switch {
		case strings.HasPrefix(operatorType, "mul("):
			op, err = handleMultiplication(match[1], match[2])
		case strings.HasPrefix(operatorType, "do("):
			op, err = handleEnable()
		case strings.HasPrefix(operatorType, "don't("):
			op, err = handleDisable()
		}

		if err == nil {
			operations = append(operations, op)
		}
	}

	return operations
}
