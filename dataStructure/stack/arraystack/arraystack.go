package arraystack

import "errors"

type Stack struct {
	item []interface{}
}

func New() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return len(s.item) == 0
}

func (s *Stack) Size() int {
	return len(s.item)
}

func (s *Stack) Push(element interface{}) {
	s.item = append(s.item, element)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack underflow")
	}
	element := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]
	return element, nil
}

func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack underflow")
	}
	return s.item[len(s.item)-1], nil
}
