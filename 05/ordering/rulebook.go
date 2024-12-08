package ordering

import (
	"fmt"
	"regexp"
	"strconv"
)

type Rule [2]int
type Rulebook []Rule

func ConstructRuleBook(lines []string) (Rulebook, error) {
	pattern := `(\d+)\|(\d+)`

	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, fmt.Errorf("error compiling regex: %v", err)
	}

	var rulebook Rulebook
	for _, line := range lines {
		matches := rx.FindAllStringSubmatch(line, -1)
		for _, match := range matches {

			a, err := strconv.Atoi(match[1])
			if err != nil {
				return nil, fmt.Errorf("can't convert [%v] to int: %v", match[1], err)
			}

			b, err := strconv.Atoi(match[2])
			if err != nil {
				return nil, fmt.Errorf("can't convert [%v] to int: %v", match[2], err)
			}

			rulebook = append(rulebook, Rule{a, b})
		}
	}

	return rulebook, nil
}
