package arraystack

import "errors"

// The Stack represents a last-in-first-out (LIFO) stack of generic items.
// This implementation uses a slice.
type Stack struct {
	item []interface{}
}

// New initialize a stack.
func New() *Stack {
	return &Stack{}
}

// IsEmpty check whether the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.item) == 0
}

// Size returns the number of items in the stack.
func (s *Stack) Size() int {
	return len(s.item)
}

// Push adds the item to this stack.
func (s *Stack) Push(element interface{}) {
	s.item = append(s.item, element)
}

// Pop removes and returns the item most recently added to this stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack underflow")
	}
	element := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]
	return element, nil
}

// Peek returns (but does not remove) the item most recently added to this stack.
func (s *Stack) Peek() (interface{}, error) {
	if s.IsEmpty() {
		return -1, errors.New("Stack underflow")
	}
	return s.item[len(s.item)-1], nil
}
