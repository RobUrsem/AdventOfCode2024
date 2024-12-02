package numbers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filePath string) ([]int, []int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left []int
	var right []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", parts[0])
		}

		num2, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, nil, fmt.Errorf("invalid number: %s", parts[1])
		}

		left = append(left, num1)
		right = append(right, num2)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil

}
