package priorityqueue

import "errors"

// MinPQ represents a priority queue of generic keys.
// This implementation uses a binary heap.
// The insert and delete-the-minimum operations take O(log n) amortized time,
// where n is the number of elements in the priority queue.
type MinPQ struct {
	n    int
	item []Key
}

// NewMinPQ initialize a min priority queue
func NewMinPQ() *MinPQ {
	return &MinPQ{
		n:    0,
		item: make([]Key, 1),
	}
}

// IsEmpty check whether the priority queue is empty
func (pq *MinPQ) IsEmpty() bool {
	return pq.n == 0
}

// Size returns the number of keys on this priority queue
func (pq *MinPQ) Size() int {
	return pq.n
}

// Insert adds a new key to this priority queue
func (pq *MinPQ) Insert(x Key) {
	pq.n++
	// pq.item has type []Key, which is a slice,
	// so we can use append() to add element,
	// and can omit resize array operation
	pq.item = append(pq.item, x)
	pq.swim(pq.n)

}

// DelMin removes and returns a smallest key on this priority queue
func (pq *MinPQ) DelMin() (Key, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue underflow")
	}
	// pq.item[1] is the minimum key in the min priority queue
	min := pq.item[1]
	pq.swap(1, pq.n)
	pq.item = pq.item[:len(pq.item)-1]
	pq.n--
	pq.sink(1)
	return min, nil
}

func (pq *MinPQ) less(i, j int) bool {
	return pq.item[i].CompareTo(pq.item[j]) < 0
}

// Bottom-up reheapify
func (pq *MinPQ) swim(k int) {
	// In a heap, the parent of the node in position k is in position k/2
	for k > 1 && pq.less(k, k/2) {
		pq.swap(k, k/2)
		k = k / 2
	}
}

// Top-down heapify
func (pq *MinPQ) sink(k int) {
	// the two children of the node in position k are in positions 2k and 2k + 1
	for 2*k <= pq.n {
		j := 2 * k
		if j < pq.n && pq.less(j+1, j) {
			j++
		}
		if !pq.less(j, k) {
			break
		}
		pq.swap(k, j)
		k = j
	}
}

func (pq *MinPQ) swap(i, j int) {
	pq.item[i], pq.item[j] = pq.item[j], pq.item[i]
}
