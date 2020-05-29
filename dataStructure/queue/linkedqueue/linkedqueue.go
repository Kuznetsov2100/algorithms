package linkedqueue

import "errors"

type Queue struct {
	first *node
	last  *node
	n     int
}

type node struct {
	item interface{}
	next *node
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

func (q *Queue) Size() int {
	return q.n
}

func (q *Queue) Enqueue(element interface{}) {
	var oldlast *node = q.last
	q.last = &node{}
	q.last.item = element
	q.last.next = nil
	if q.IsEmpty() {
		q.first = q.last
	} else {
		oldlast.next = q.last
	}
	q.n++
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue underflow")
	}
	element := q.first.item
	q.first = q.first.next
	q.n--
	if q.IsEmpty() {
		q.last = nil
	}
	return element, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue undeflow")
	}
	return q.first.item, nil
}
