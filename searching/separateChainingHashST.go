package searching

import (
	"github.com/pkg/errors"
)

// HashKey describes the requirements for the type using SeparateChainingHashST in this package.
type HashKey interface {
	HashCode() int
	Key
}

// StringHashKey implements the HashKey interface.
type StringHashKey string

// HashCode returns hash code of type StringHashKey
func (s StringHashKey) HashCode() int {
	R, hash := 31, 0
	for _, runeValue := range s {
		hash = R*hash + int(runeValue)
	}
	return hash
}

// CompareTo compares two StringHashKey type
func (s StringHashKey) CompareTo(k Key) int {
	t := k.(StringHashKey)
	if s < t {
		return -1
	} else if s > t {
		return 1
	} else {
		return 0
	}
}

// SeparateChainingHashST struct represents a symbol table of generic key-value pairs.
// This implementation uses a separate chaining hash table.
// It requires that the key type implements the CompareTo() and HashCode() methods.
// The expected time per put, contains, or remove operation is constant, subject to the uniform hashing assumption.
// The size, and is-empty operations take constant time. Construction takes constant time.
type SeparateChainingHashST struct {
	n  int
	m  int
	st []*SequentialSearchST
}

// NewSeparateChainingHashST initializes an empty symbol table with capaciity chains.
func NewSeparateChainingHashST(capacity int) *SeparateChainingHashST {
	if capacity < 0 {
		panic("capacity should be non negative value")
	}
	if capacity == 0 {
		capacity = 4
	}
	st := make([]*SequentialSearchST, capacity)
	for i := range st {
		st[i] = NewSequentialSearchST()
	}
	return &SeparateChainingHashST{m: capacity, st: st, n: 0}
}

// resize the hash table to have the given number of chains,
// rehashing all of the keys
func (sc *SeparateChainingHashST) resize(chains int) {
	temp := NewSeparateChainingHashST(chains)
	for i := 0; i < sc.m; i++ {
		for _, key := range sc.st[i].Keys() {

			val, _ := sc.st[i].Get(key)
			//nolint:errcheck
			temp.Put(key.(HashKey), val)
		}
	}
	sc.m = temp.m
	sc.n = temp.n
	sc.st = temp.st
}

// hash value between 0 and m-1
func (sc *SeparateChainingHashST) hash(key HashKey) int {
	// key.HashCode() & 0x7fffffff to keep the value always non negative(aka:bitwise mask)
	return (key.HashCode() & 0x7fffffff) % sc.m
}

// Size returns the number of key-value pairs in this symbol table.
func (sc *SeparateChainingHashST) Size() int {
	return sc.n
}

// IsEmpty returns true if this symbol table is empty.
func (sc *SeparateChainingHashST) IsEmpty() bool {
	return sc.Size() == 0
}

// Contains returns true if this symbol table contains the specified key.
func (sc *SeparateChainingHashST) Contains(key HashKey) (bool, error) {
	if key == nil {
		return false, errors.New("argument to Contains() is nil key")
	}
	val, _ := sc.Get(key)
	return val != nil, nil
}

// Get returns the value associated with the specified key in this symbol table.
func (sc *SeparateChainingHashST) Get(key HashKey) (Value, error) {
	if key == nil {
		return nil, errors.New("argument to Get() is nil key")
	}
	val, _ := sc.st[sc.hash(key)].Get(key)
	return val, nil
}

// Put Inserts the specified key-value pair into the symbol table,
// overwriting the old value with the new value if the symbol table already contains the specified key.
func (sc *SeparateChainingHashST) Put(key HashKey, val Value) error {
	if key == nil {
		return errors.New("argument to Put() is nil key")
	}
	if val == nil {
		//nolint:errcheck
		sc.st[sc.hash(key)].Delete(key)
		return nil
	}
	// double table size if average length of list >= 10
	if sc.n >= 10*sc.m {
		sc.resize(2 * sc.m)
	}
	i := sc.hash(key)
	if ok, _ := sc.st[i].Contains(key); !ok {
		sc.n++
	}
	//nolint:errcheck
	sc.st[i].Put(key, val)
	return nil
}

// Delete removes the specified key and its associated value from this symbol table (if the key is in this symbol table).
func (sc *SeparateChainingHashST) Delete(key HashKey) error {
	if key == nil {
		return errors.New("argument to Delete() is nil key")
	}
	i := sc.hash(key)
	if ok, _ := sc.st[i].Contains(key); ok {
		sc.n--
	}
	//nolint:errcheck
	sc.st[i].Delete(key)
	// halve table size if average length of list <= 2
	if sc.m > 4 && sc.n <= 2*sc.m {
		sc.resize(sc.m / 2)
	}
	return nil
}

// Keys return all of the keys in the symbol table.
func (sc *SeparateChainingHashST) Keys() (keys []HashKey) {
	for i := 0; i < sc.m; i++ {
		for _, x := range sc.st[i].Keys() {
			keys = append(keys, x.(HashKey))
		}
	}
	return keys
}
