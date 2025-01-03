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
			case OBSTRUCTION:
				fmt.Print("O")
			case VISITED:
				fmt.Print("X")
			case VISITED_VERTICAL:
				fmt.Print("|")
			case VISITED_HORIZONTAL:
				fmt.Print("-")
			case VISITED_BOTH:
				fmt.Print("+")
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
