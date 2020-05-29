package linkedstack

import "errors"

// The Stack represents a last-in-first-out (LIFO) stack of generic items.
// This implementation uses a singly linked list.
type Stack struct {
	first *node
	n     int
}

type node struct {
	item interface{}
	next *node
}

// Initialize a stack.
func New() *Stack {
	return &Stack{}
}

// Is this stack empty?
func (ls *Stack) IsEmpty() bool {
	return ls.first == nil
}

// Returns the number of items in the stack.
func (ls *Stack) Size() int {
	return ls.n
}

// Adds the item to this stack.
func (ls *Stack) Push(element interface{}) {
	var oldfirst *node = ls.first
	ls.first = &node{}
	ls.first.item = element
	ls.first.next = oldfirst
	ls.n++
}

// Removes and returns the item most recently added to this stack.
func (ls *Stack) Pop() (interface{}, error) {
	if ls.IsEmpty() {
		return -1, errors.New("Stack underflow")
	}
	item := ls.first.item
	ls.first = ls.first.next
	ls.n--
	return item, nil
}

// Returns (but does not remove) the item most recently added to this stack.
func (ls *Stack) Peek() (interface{}, error) {
	if ls.IsEmpty() {
		return -1, errors.New("Stack underflow")
	}
	return ls.first.item, nil
}
