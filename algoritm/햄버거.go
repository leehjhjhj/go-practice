import (
    "fmt"
)

type Stack struct {
    elements []int
}

func (s *Stack) Push(element int) {
    s.elements = append(s.elements, element)
}

func (s *Stack) Pop() (int, error) {
    if len(s.elements) == 0 {
        return 0, fmt.Errorf("스택이 비었습니다")
    }
    elements := s.elements[len(s.elements)-1]
    s.elements = s.elements[:len(s.elements)-1]
    return element, nil
}

func solution(ingredient []int) int {
    return 0
}