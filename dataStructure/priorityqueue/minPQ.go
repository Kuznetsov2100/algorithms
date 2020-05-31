package priorityqueue

import "errors"

type MinPQ struct {
	n    int
	item []Key
}

func NewPQ() *MinPQ {
	return &MinPQ{
		n:    0,
		item: make([]Key, 1),
	}
}

func (pq *MinPQ) IsEmpty() bool {
	return pq.n == 0
}

func (pq *MinPQ) Size() int {
	return pq.n
}

func (pq *MinPQ) Insert(x Key) {
	pq.n++
	pq.item = append(pq.item, x)
	pq.swim(pq.n)

}

func (pq *MinPQ) DelMin() (Key, error) {
	if pq.IsEmpty() {
		return nil, errors.New("priority queue underflow")
	}
	min := pq.item[1]
	pq.swap(1, pq.n)
	pq.n--
	pq.sink(1)
	return min, nil
}

func (pq *MinPQ) less(i, j int) bool {
	return pq.item[i].CompareTo(pq.item[j]) < 0
}

func (pq *MinPQ) swim(k int) {
	for k > 1 && pq.less(k, k/2) {
		pq.swap(k, k/2)
		k = k / 2
	}
}

func (pq *MinPQ) sink(k int) {
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
