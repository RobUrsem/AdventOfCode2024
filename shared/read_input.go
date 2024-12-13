package shared

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func TextToIntArray(line string) ([]int, error) {
	parts := strings.Fields(line)

	report := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)

		if err != nil {
			return nil, fmt.Errorf("invalid number: %d", num)
		}

		report[i] = num
	}

	return report, nil
}
