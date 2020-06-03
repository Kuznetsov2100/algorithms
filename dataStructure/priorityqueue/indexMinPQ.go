package priorityqueue

import "errors"

// IndexMinPQ struct represents an indexed priority queue of generic keys.
// This implementation uses a binary heap along with an array
// to associate keys with integers in the given range.
type IndexMinPQ struct {
	maxN int   // maximum number of elements on PQ
	n    int   // number of elements on PQ
	pq   []int // binary heap using 1-based indexing
	qp   []int // inverse of pq,  qp[pq[i]] = pq[qp[i]] = i
	keys []Key // keys[i] = priority of i
}

// NewIndexMinPQ initializes an empty indexed priority queue with indices between 0 and maxN - 1.
func NewIndexMinPQ(maxN int) *IndexMinPQ {
	if maxN < 0 {
		panic("maxN should be a non negative integer!")
	}
	qp := make([]int, maxN+1)
	for i := range qp {
		qp[i] = -1
	}
	return &IndexMinPQ{
		maxN: maxN,
		n:    0,
		pq:   make([]int, maxN+1),
		qp:   qp,
		keys: make([]Key, maxN+1),
	}
}

// IsEmpty returns true if this priority queue is empty.
func (queue *IndexMinPQ) IsEmpty() bool {
	return queue.n == 0
}

// Contains check whether the priority queue has an index i
func (queue *IndexMinPQ) Contains(i int) bool {
	queue.validateIndex(i)
	return queue.qp[i] != -1
}

// Size returns the number of keys on this priority queue.
func (queue *IndexMinPQ) Size() int {
	return queue.n
}

// Insert associates key with index i.
func (queue *IndexMinPQ) Insert(i int, key Key) error {
	queue.validateIndex(i)
	if queue.Contains(i) {
		return errors.New("index is already in the priority queue")
	}
	queue.n++
	queue.qp[i] = queue.n
	queue.pq[queue.n] = i
	queue.keys[i] = key
	queue.swim(queue.n)
	return nil
}

// MinKey returns a minimum key.
func (queue *IndexMinPQ) MinKey() (Key, error) {
	if queue.IsEmpty() {
		return nil, errors.New("priority queue underflow")
	}
	return queue.keys[queue.pq[1]], nil
}

// MinIndex returns an index associated with a minimum key.
func (queue *IndexMinPQ) MinIndex() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New("priority queue underflow")
	}
	return queue.pq[1], nil
}

// KeyOf returns the key associated with index i.
func (queue *IndexMinPQ) KeyOf(i int) (Key, error) {
	queue.validateIndex(i)
	if !queue.Contains(i) {
		return nil, errors.New("index is not in the priority queue")
	}
	return queue.keys[i], nil
}

// DelMin removes a minimum key and returns its associated index.
func (queue *IndexMinPQ) DelMin() (int, error) {
	if queue.IsEmpty() {
		return -1, errors.New("priority queue underflow")
	}
	min := queue.pq[1]
	queue.swap(1, queue.n)
	queue.n--
	queue.sink(1)
	queue.qp[min] = -1
	queue.keys[min] = nil
	return min, nil
}

// ChangeKey change the key associated with index i to the specified value.
func (queue *IndexMinPQ) ChangeKey(i int, key Key) error {
	queue.validateIndex(i)
	if !queue.Contains(i) {
		return errors.New("index is not in the priority queue")
	}
	queue.keys[i] = key
	queue.swim(queue.qp[i])
	queue.sink(queue.qp[i])
	return nil
}

// DecreaseKey decrease the key associated with index i to the specified value.
func (queue *IndexMinPQ) DecreaseKey(i int, key Key) error {
	queue.validateIndex(i)
	if !queue.Contains(i) {
		return errors.New("index is not in the priority queue")
	}
	if queue.keys[i].CompareTo(key) == 0 {
		return errors.New("calling DecreaseKey() with a key equal to the key in the priority queue")
	}
	if queue.keys[i].CompareTo(key) < 0 {
		return errors.New("calling DecreaseKey() with a key strictly greater than the key in the priority queue")
	}
	queue.keys[i] = key
	queue.swim(queue.qp[i])
	return nil
}

// IncreaseKey increase the key associated with index i to the specified value.
func (queue *IndexMinPQ) IncreaseKey(i int, key Key) error {
	queue.validateIndex(i)
	if !queue.Contains(i) {
		return errors.New("index is not in the priority queue")
	}
	if queue.keys[i].CompareTo(key) == 0 {
		return errors.New("calling IncreaseKey() with a key equal to the key in the priority queue")
	}
	if queue.keys[i].CompareTo(key) > 0 {
		return errors.New("calling IncreaseKey() with a key strictly less than the key in the priority queue")
	}
	queue.keys[i] = key
	queue.sink(queue.qp[i])
	return nil
}

// Delete remove the key associated with index i.
func (queue *IndexMinPQ) Delete(i int) error {
	queue.validateIndex(i)
	if !queue.Contains(i) {
		return errors.New("index is not in the priority queue")
	}
	index := queue.qp[i]
	queue.swap(index, queue.n)
	queue.n--
	queue.swim(index)
	queue.sink(index)
	queue.keys[i] = nil
	queue.qp[i] = -1
	return nil
}

func (queue *IndexMinPQ) swap(i, j int) {
	queue.pq[i], queue.pq[j] = queue.pq[j], queue.pq[i]
	queue.qp[queue.pq[i]] = i
	queue.qp[queue.pq[j]] = j
}

func (queue *IndexMinPQ) swim(k int) {
	for k > 1 && queue.greater(k/2, k) {
		queue.swap(k, k/2)
		k = k / 2
	}
}

func (queue *IndexMinPQ) sink(k int) {
	for 2*k <= queue.n {
		j := 2 * k
		if j < queue.n && queue.greater(j, j+1) {
			j++
		}
		if !queue.greater(k, j) {
			break
		}
		queue.swap(k, j)
		k = j
	}
}

func (queue *IndexMinPQ) greater(i, j int) bool {
	return queue.keys[queue.pq[i]].CompareTo(queue.keys[queue.pq[j]]) > 0
}

func (queue *IndexMinPQ) validateIndex(i int) {
	if i < 0 || i >= queue.maxN {
		panic("index ilegal")
	}
}
