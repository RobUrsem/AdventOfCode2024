package shared

import "fmt"

type Stack struct {
	Items []int
}

func (s *Stack) Push(data int) {
	s.Items = append(s.Items, data)
}

func (s *Stack) Pop() {
	if s.IsEmpty() {
		return
	}
	s.Items = s.Items[:len(s.Items)-1]
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("stack is empty")
	}
	return s.Items[len(s.Items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	if len(s.Items) == 0 {
		return true
	}
	return false
}

func (s *Stack) Print() {
	for _, item := range s.Items {
		fmt.Print(item, " ")
	}
	fmt.Println()
}
