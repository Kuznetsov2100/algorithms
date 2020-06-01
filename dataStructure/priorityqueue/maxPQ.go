package priorityqueue

import "errors"

// MaxPQ represents a priority queue of generic keys.
// This implementation uses a binary heap.
// The insert and delete-the-maximum operations take O(log n) amortized time,
// where n is the number of elements in the priority queue.
type MaxPQ struct {
	n    int
	item []Key
}

// NewMaxPQ initialize a max priority queue
func NewMaxPQ() *MaxPQ {
	return &MaxPQ{
		n:    0,
		item: make([]Key, 1),
	}
}

// IsEmpty check whether the priority queue is empty
func (pq *MaxPQ) IsEmpty() bool {
	return pq.n == 0
}

// Size returns the number of keys on this priority queue
func (pq *MaxPQ) Size() int {
	return pq.n
}

// Insert adds a new key to this priority queue
func (pq *MaxPQ) Insert(x Key) {
	pq.n++
	// pq.item has type []Key, which is a slice,
	// so we can use append() to add element,
	// and can omit resize array operation
	pq.item = append(pq.item, x)
	pq.swim(pq.n)
}

// DelMax removes and returns a largest key on this priority queue
func (pq *MaxPQ) DelMax() (Key, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue underflow")
	}
	max := pq.item[1]
	pq.swap(1, pq.n)
	pq.n--
	pq.sink(1)
	return max, nil
}

func (pq *MaxPQ) less(i, j int) bool {
	return pq.item[i].CompareTo(pq.item[j]) < 0
}

// Bottom-up reheapify
func (pq *MaxPQ) swim(k int) {
	// In a heap, the parent of the node in position k is in position k/2
	for k > 1 && pq.less(k/2, k) {
		pq.swap(k/2, k)
		k = k / 2
	}
}

// Top-down heapify
func (pq *MaxPQ) sink(k int) {
	// the two children of the node in position k are in positions 2k and 2k + 1
	for 2*k <= pq.n {
		j := 2 * k
		if j < pq.n && pq.less(j, j+1) {
			j++
		}

		if !pq.less(k, j) {
			break
		}
		pq.swap(k, j)
		k = j
	}
}

func (pq *MaxPQ) swap(i, j int) {
	pq.item[i], pq.item[j] = pq.item[j], pq.item[i]
}
