package searching

import "errors"

type LinearProbingHashST struct {
	n    int
	m    int
	keys []HashKey
	vals []Value
}

func NewLinearProbingHashST(capacity int) *LinearProbingHashST {
	if capacity < 0 {
		panic("capacity should be non negative value")
	}
	if capacity == 0 {
		capacity = 4
	}
	return &LinearProbingHashST{
		m:    capacity,
		n:    0,
		keys: make([]HashKey, capacity),
		vals: make([]Value, capacity),
	}
}

func (lp *LinearProbingHashST) Size() int {
	return lp.n
}

func (lp *LinearProbingHashST) IsEmpty() bool {
	return lp.Size() == 0
}

func (lp *LinearProbingHashST) Contains(key HashKey) (bool, error) {
	if key == nil {
		return false, errors.New("argument to Contains() is nil key")
	}
	val, _ := lp.Get(key)
	return val != nil, nil
}

func (lp *LinearProbingHashST) hash(key HashKey) int {
	return (key.HashCode() & 0x7fffffff) % lp.m
}

func (lp *LinearProbingHashST) resize(capacity int) {
	temp := NewLinearProbingHashST(capacity)
	for i := 0; i < lp.m; i++ {
		if lp.keys[i] != nil {
			//nolint:errcheck
			temp.Put(lp.keys[i], lp.vals[i])
		}
	}
	lp.keys = temp.keys
	lp.vals = temp.vals
	lp.m = temp.m
}

func (lp *LinearProbingHashST) Put(key HashKey, val Value) error {
	if key == nil {
		return errors.New("first argument to Put() is nil key")
	}
	if val == nil {
		//nolint:errcheck
		lp.Delete(key)
	}
	if lp.n >= lp.m/2 {
		lp.resize(2 * lp.m)
	}
	i := lp.hash(key)
	for ; lp.keys[i] != nil; i = (i + 1) % lp.m {
		if lp.keys[i].CompareTo(key) == 0 {
			lp.vals[i] = val
			return nil
		}
	}
	lp.keys[i] = key
	lp.vals[i] = val
	lp.n++
	return nil
}

func (lp *LinearProbingHashST) Get(key HashKey) (Value, error) {
	if key == nil {
		return nil, errors.New("argument to Get() is nil key")
	}
	for i := lp.hash(key); lp.keys[i] != nil; i = (i + 1) % lp.m {
		if lp.keys[i].CompareTo(key) == 0 {
			return lp.vals[i], nil
		}
	}
	return nil, nil
}

func (lp *LinearProbingHashST) Delete(key HashKey) error {
	if key == nil {
		return errors.New("argument to Delete() is nil key")
	}
	if ok, _ := lp.Contains(key); !ok {
		return nil
	}

	i := lp.hash(key)
	for lp.keys[i].CompareTo(key) != 0 {
		i = (i + 1) % lp.m
	}
	lp.keys[i] = nil
	lp.vals[i] = nil
	i = (i + 1) % lp.m
	for lp.keys[i] != nil {
		keyToRehash := lp.keys[i]
		valToRehash := lp.vals[i]
		lp.keys[i] = nil
		lp.vals[i] = nil
		lp.n--
		//nolint:errcheck
		lp.Put(keyToRehash, valToRehash)
		i = (i + 1) % lp.m
	}
	lp.n--
	if lp.n > 0 && lp.n <= lp.m/8 {
		lp.resize(lp.m / 2)
	}
	return nil
}

func (lp *LinearProbingHashST) Keys() (keys []HashKey) {
	for i := 0; i < lp.m; i++ {
		if lp.keys[i] != nil {
			keys = append(keys, lp.keys[i])
		}
	}
	return keys
}
