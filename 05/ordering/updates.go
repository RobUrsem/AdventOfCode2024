package ordering

import (
	"fmt"
	"strconv"
	"strings"
)

type Update []int

func convertToList(input string) (Update, error) {
	var numbers Update
	parts := strings.Split(input, ",")
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err != nil {
			fmt.Printf("Error converting '%s' to int: %v\n", part, err)
			return nil, err
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func GetUpdates(lines []string) ([]Update, error) {
	var updates []Update

	for _, line := range lines {
		if strings.Contains(line, ",") {
			update, err := convertToList(line)
			if err != nil {
				return nil, err
			}
			updates = append(updates, update)
		}
	}

	return updates, nil
}
