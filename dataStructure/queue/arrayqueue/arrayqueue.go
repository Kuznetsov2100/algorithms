package arrayqueue

import "errors"

type Queue struct {
	item []interface{}
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) IsEmpty() bool {
	return len(q.item) == 0
}

func (q *Queue) Size() int {
	return len(q.item)
}

func (q *Queue) Enqueue(element interface{}) {
	q.item = append(q.item, element)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("No element in queue")
	}
	item := q.item[0]
	q.item = q.item[1:]
	return item, nil
}

func (q *Queue) Peek() (interface{}, error) {
	if q.IsEmpty() {
		return -1, errors.New("No element in queue")
	}
	return q.item[0], nil
}
