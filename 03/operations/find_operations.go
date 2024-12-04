package operations

import (
	"fmt"
	"regexp"
	"strconv"
)

func findMultiplications(text string) []Operation {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`

	rx, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return nil
	}

	operations := []Operation{}
	matches := rx.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Printf("can't convert %v to int: %v\n", match[1], err)
			break
		}

		b, err := strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("can't convert %v to int: %v\n", match[2], err)
			break
		}

		operations = append(operations, Operation{
			Params:        []int{a, b},
			OperationType: Multiply,
		})
	}

	return operations
}

func FindOperations(text string) ([]Operation, error) {
	operations := findMultiplications(text)
	return operations, nil
}
