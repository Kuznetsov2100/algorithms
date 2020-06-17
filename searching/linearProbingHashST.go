package searching

import "github.com/pkg/errors"

// LinearProbingHashST struct represents a symbol table of generic key-value pairs.
// This implementation uses a linear probing hash table.
// It requires that the key type implements the CompareTo() and HashCode() methods.
// The expected time per put, contains, or remove operation is constant,
// subject to the uniform hashing assumption. The size, and is-empty operations take constant time. Construction takes constant time.
type LinearProbingHashST struct {
	n    int       // number of key-value pairs in the symbol table
	m    int       // size of linear probing table
	keys []HashKey // the keys
	vals []Value   // the values
}

// NewLinearProbingHashST initializes an empty symbol table with the specified initial capacity.
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

// Size returns the number of key-value pairs in this symbol table.
func (lp *LinearProbingHashST) Size() int {
	return lp.n
}

// IsEmpty returns true if this symbol table is empty.
func (lp *LinearProbingHashST) IsEmpty() bool {
	return lp.Size() == 0
}

// Contains returns true if this symbol table contains the specified key.
func (lp *LinearProbingHashST) Contains(key HashKey) (bool, error) {
	if key == nil {
		return false, errors.New("argument to Contains() is nil key")
	}
	val, _ := lp.Get(key)
	return val != nil, nil
}

// hash value between 0 and m-1
func (lp *LinearProbingHashST) hash(key HashKey) int {
	// key.HashCode() & 0x7fffffff to keep the value always non negative(aka:bitwise mask)
	return (key.HashCode() & 0x7fffffff) % lp.m
}

// resize the hash table to have the given number of chains,
// rehashing all of the keys
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

// Put Inserts the specified key-value pair into the symbol table,
// overwriting the old value with the new value if the symbol table already contains the specified key.
func (lp *LinearProbingHashST) Put(key HashKey, val Value) error {
	if key == nil {
		return errors.New("first argument to Put() is nil key")
	}
	if val == nil {
		//nolint:errcheck
		lp.Delete(key)
	}
	// double table size if 50% full
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

// Get returns the value associated with the specified key in this symbol table.
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

// Delete removes the specified key and its associated value from this symbol table (if the key is in this symbol table).
func (lp *LinearProbingHashST) Delete(key HashKey) error {
	if key == nil {
		return errors.New("argument to Delete() is nil key")
	}
	if ok, _ := lp.Contains(key); !ok {
		return nil
	}
	// find position i of key
	i := lp.hash(key)
	for lp.keys[i].CompareTo(key) != 0 {
		i = (i + 1) % lp.m
	}
	lp.keys[i] = nil
	lp.vals[i] = nil
	// rehash all keys in same cluster
	i = (i + 1) % lp.m
	for lp.keys[i] != nil {
		// delete keys[i] an vals[i] and reinsert
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
	// halves size of array if it's 12.5% full or less
	if lp.n > 0 && lp.n <= lp.m/8 {
		lp.resize(lp.m / 2)
	}
	return nil
}

// Keys return all of the keys in the symbol table.
func (lp *LinearProbingHashST) Keys() (keys []HashKey) {
	for i := 0; i < lp.m; i++ {
		if lp.keys[i] != nil {
			keys = append(keys, lp.keys[i])
		}
	}
	return keys
}
