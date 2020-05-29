package arrayqueue

import "errors"

// The Queue represents a first-in-first-out (FIFO) queue of generic items.
// This implementation uses a slice.
type Queue struct {
	item []interface{}
}

// Initialize a Queue.
func New() *Queue {
	return &Queue{}
}

// Is this queue empty?
func (q *Queue) IsEmpty() bool {
	return len(q.item) == 0
}

// Returns the number of items in this queue.
func (q *Queue) Size() int {
	return len(q.item)
}

// Adds the item to this queue.
func (q *Queue) Enqueue(element interface{}) {
	q.item = append(q.item, element)
}

// Removes and returns the item on this queue that was least recently added.
func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("No element in queue")
	}
	item := q.item[0]
	q.item = q.item[1:]
	return item, nil
}

// Returns the item least recently added to this queue.
func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("No element in queue")
	}
	return q.item[0], nil
}
