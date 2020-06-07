package linkedqueue

import "errors"

// The Queue represents a first-in-first-out (FIFO) queue of generic items,
// This implementation using a singly linked list.
type Queue struct {
	first *node
	last  *node
	n     int
}

type node struct {
	item interface{}
	next *node
}

// New initialize a Queue.
func New() *Queue {
	return &Queue{}
}

// IsEmpty check whether the queue is empty.
func (q *Queue) IsEmpty() bool {
	return q.first == nil
}

// Size returns the number of items in this queue.
func (q *Queue) Size() int {
	return q.n
}

// Enqueue adds the item to this queue.
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

// Dequeue removes and returns the item on this queue that was least recently added.
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

// Peek returns the item least recently added to this queue.
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("queue undeflow")
	}
	return q.first.item, nil
}

func (q *Queue) Values() (v []interface{}) {
	for x := q.first; x != nil; x = x.next {
		v = append(v, x.item)
	}
	return v
}
