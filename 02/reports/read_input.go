package reports

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadInput(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var reports [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) < 2 {
			return nil, fmt.Errorf("invalid line format: %s", line)
		}

		report := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)

			if err != nil {
				return nil, fmt.Errorf("invalid number: %d", num)
			}

			report[i] = num
		}

		reports = append(reports, report)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return reports, nil
}
