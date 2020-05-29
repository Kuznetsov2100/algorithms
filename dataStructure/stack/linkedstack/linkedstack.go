package linkedstack

import "errors"

type Stack struct {
	first *node
	n     int
}

type node struct {
	item interface{}
	next *node
}

func New() *Stack {
	return &Stack{}
}

func (ls *Stack) IsEmpty() bool {
	return ls.first == nil
}

func (ls *Stack) Size() int {
	return ls.n
}

func (ls *Stack) Push(element interface{}) {
	var oldfirst *node = ls.first
	ls.first = &node{}
	ls.first.item = element
	ls.first.next = oldfirst
	ls.n++
}

func (ls *Stack) Pop() (interface{}, error) {
	if ls.IsEmpty() {
		return -1, errors.New("Stack underflow")
	}
	item := ls.first.item
	ls.first = ls.first.next
	ls.n--
	return item, nil
}

func (ls *Stack) Peek() (interface{}, error) {
	if ls.IsEmpty() {
		return -1, errors.New("Stack underflow")
	}
	return ls.first.item, nil
}
