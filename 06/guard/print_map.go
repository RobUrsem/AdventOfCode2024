package guard

import "fmt"

func PrintMap(labMap LabMap) {
	for _, row := range labMap {
		for _, col := range row {
			switch col {
			case EMPTY:
				fmt.Print(".")
			case OBSTACLE:
				fmt.Print("#")
			case VISITED:
				fmt.Print("X")
			case GUARD_UP:
				fmt.Print("^")
			case GUARD_LEFT:
				fmt.Print("<")
			case GUARD_DOWN:
				fmt.Print("V")
			case GUARD_RIGHT:
				fmt.Print(">")
			}
		}
		fmt.Println()
	}
}
