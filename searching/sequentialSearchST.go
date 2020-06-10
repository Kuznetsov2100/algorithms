package searching

import (
	"errors"
)

// SequentialSearchST struct represents an (unordered) symbol table of generic key-value pairs.
// This implementation uses a singly linked list and sequential search.
// The put and delete operations take O(n).
// The get and contains operations takes O(n) time in the worst case.
// The size, and is-empty operations take O(1) time. Construction takes O(1) time.
type SequentialSearchST struct {
	n     int
	first *stnode
}

type stnode struct {
	key  Key
	val  Value
	next *stnode
}

// NewSequentialSearchST initializes an empty symbol table.
func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{}
}

// Size returns the number of key-value pairs in this symbol table.
func (st *SequentialSearchST) Size() int {
	return st.n
}

// IsEmpty returns true if this symbol table is empty.
func (st *SequentialSearchST) IsEmpty() bool {
	return st.Size() == 0
}

// Contains returns true if this symbol table contains the specified key.
func (st *SequentialSearchST) Contains(key Key) (bool, error) {
	if key == nil {
		return false, errors.New("arugment to Contains() is nil key")
	}
	val, _ := st.Get(key)
	return val != nil, nil
}

// Get returns the value associated with the specified key in this symbol table.
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

// Put Inserts the specified key-value pair into the symbol table,
// overwriting the old value with the new value if the symbol table already contains the specified key.
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

// Delete removes the specified key and its associated value from this symbol table (if the key is in this symbol table).
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

// Keys return all of the keys in the symbol table.
func (st *SequentialSearchST) Keys() (keys []Key) {
	for x := st.first; x != nil; x = x.next {
		keys = append(keys, x.key)
	}
	return keys
}
