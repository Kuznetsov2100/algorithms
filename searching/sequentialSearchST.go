package searching

import (
	"errors"
)

type SequentialSearchST struct {
	n     int
	first *stnode
}

type stnode struct {
	key  Key
	val  Value
	next *stnode
}

func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{}
}

func (st *SequentialSearchST) Size() int {
	return st.n
}

func (st *SequentialSearchST) IsEmpty() bool {
	return st.Size() == 0
}

func (st *SequentialSearchST) Contains(key Key) (bool, error) {
	if key == nil {
		return false, errors.New("arugment to Contains() is nil key")
	}
	val, _ := st.Get(key)
	return val != nil, nil
}

func (st *SequentialSearchST) Get(key Key) (Value, error) {
	if key == nil {
		return nil, errors.New("argument to Get() is nil key")
	}
	for x := st.first; x != nil; x = x.next {
		if key.CompareTo(x.key) == 0 {
			return x.val, nil
		}
	}
	return nil, nil
}

func (st *SequentialSearchST) Put(key Key, val Value) error {
	if key == nil {
		return errors.New("first argument to Put() is nil key")
	}
	if val == nil {
		//nolint:errcheck
		st.Delete(key)
		return nil
	}
	for x := st.first; x != nil; x = x.next {
		if key.CompareTo(x.key) == 0 {
			x.val = val
			return nil
		}
	}
	st.first = &stnode{key: key, val: val, next: st.first}
	st.n++
	return nil
}

func (st *SequentialSearchST) Delete(key Key) error {
	if key == nil {
		return errors.New("argument to Delete() is nil key")
	}
	st.first = st.delete(st.first, key)
	return nil
}

func (st *SequentialSearchST) delete(x *stnode, key Key) *stnode {
	if x == nil {
		return nil
	}
	if key.CompareTo(x.key) == 0 {
		st.n--
		return x.next
	}
	x.next = st.delete(x.next, key)
	return x
}

func (st *SequentialSearchST) Keys() (keys []Key) {
	for x := st.first; x != nil; x = x.next {
		keys = append(keys, x.key)
	}
	return keys
}
