package puzzle

import (
	"regexp"
	"strconv"
)

func getNumber(num string) int {
	value, _ := strconv.Atoi(num)
	return value
}

func GetInput(lines []string) Robots {
	rx := regexp.MustCompile(`p=(-?\d+),(-?\d+)\s+v=(-?\d+),(-?\d+)`)

	robots := Robots{}
	for _, line := range lines {
		match := rx.FindStringSubmatch(line)
		if match != nil {
			robots = append(robots, NewRobot(match[1], match[2], match[3], match[4]))
		}
	}

	return robots
}
