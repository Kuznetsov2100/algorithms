package searching

import "errors"

// Value type is a generic type.
type Value interface{}

// BinarySearchST struct represents an ordered symbol table of generic key-value pairs.
// A symbol table implements the associative array abstraction:
// when associating a value with a key that is already in the symbol table,
// the convention is to replace the old value with the new value.
// Unlike Map, BinarySearchST uses the convention that values cannot be null—setting the value
// associated with a key to null is equivalent to deleting the key from the symbol table.
type BinarySearchST struct {
	keys []Key
	vals []Value
	n    int
}

// NewBinarySearchST initializes an empty symbol table.
func NewBinarySearchST() *BinarySearchST {
	return &BinarySearchST{
		keys: make([]Key, 2),
		vals: make([]Value, 2),
		n:    0,
	}
}

// Size returns the number of key-value pairs in this symbol table.
func (bst *BinarySearchST) Size() int {
	return bst.n
}

// IsEmpty returns true if this symbol table is empty.
func (bst *BinarySearchST) IsEmpty() bool {
	return bst.n == 0
}

// Contains check whether the symbol table contain the given key.
func (bst *BinarySearchST) Contains(key Key) (bool, error) {
	if key == nil {
		return false, errors.New("argument to Contains() is nil")
	}
	k, _ := bst.Get(key)
	return k != nil, nil
}

// Get returns the value associated with the given key in this symbol table.
func (bst *BinarySearchST) Get(key Key) (Value, error) {
	if key == nil {
		return nil, errors.New("argument to Get() is nil")
	}
	if bst.IsEmpty() {
		return nil, nil
	}
	if i, _ := bst.Rank(key); i < bst.n && bst.keys[i].CompareTo(key) == 0 {
		return bst.vals[i], nil // the key is in the symbol table
	}
	return nil, nil
}

// Rank returns the number of keys in this symbol table strictly less than key.
func (bst *BinarySearchST) Rank(key Key) (int, error) {
	if key == nil {
		return -1, errors.New("argument to Rank() is nil")
	}
	lo, hi := 0, bst.n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if cmp := key.CompareTo(bst.keys[mid]); cmp < 0 {
			hi = mid - 1
		} else if cmp > 0 {
			lo = mid + 1
		} else {
			return mid, nil
		}
	}
	return lo, nil
}

// Put inserts the specified key-value pair into the symbol table, overwriting the old value with the new value if the symbol table already contains the specified key.
func (bst *BinarySearchST) Put(key Key, val Value) error {
	if key == nil {
		return errors.New("argument to Put() is nil")
	}
	if val == nil {
		//nolint:errcheck
		bst.Delete(key)
		return nil
	}
	i, _ := bst.Rank(key)
	if i < bst.n && bst.keys[i].CompareTo(key) == 0 {
		bst.vals[i] = val // the key is in the symbol table
		return nil
	}
	// the key is not in the symbol table
	bst.keys = append(bst.keys[:i+1], bst.keys[i:]...)
	bst.vals = append(bst.vals[:i+1], bst.vals[i:]...)
	bst.keys[i], bst.vals[i] = key, val
	bst.n++
	return nil
}

// Delete removes the specified key and associated value from this symbol table (if the key is in the symbol table).
func (bst *BinarySearchST) Delete(key Key) error {
	if key == nil {
		return errors.New("argument to Delete() is nil")
	}
	if bst.IsEmpty() {
		return nil
	}
	i, _ := bst.Rank(key)
	if i == bst.n || bst.keys[i].CompareTo(key) != 0 {
		return nil // the key is not in the symbol table
	}
	bst.keys = append(bst.keys[:i], bst.keys[i+1:]...)
	bst.vals = append(bst.vals[:i], bst.vals[i+1:]...)
	bst.n--
	return nil
}

// DeleteMin removes the smallest key and associated value from this symbol table.
func (bst *BinarySearchST) DeleteMin() error {
	if bst.IsEmpty() {
		return errors.New("symbol table underflow")
	}
	k, _ := bst.Min()
	//nolint:errcheck
	bst.Delete(k)
	return nil
}

// DeleteMax removes the largest key and associated value from this symbol table.
func (bst *BinarySearchST) DeleteMax() error {
	if bst.IsEmpty() {
		return errors.New("symbol table underflow")
	}
	k, _ := bst.Max()
	//nolint:errcheck
	bst.Delete(k)
	return nil
}

// Min returns the smallest key in this symbol table.
func (bst *BinarySearchST) Min() (Key, error) {
	if bst.IsEmpty() {
		return nil, errors.New("called Min() with empty symbol table")
	}
	return bst.keys[0], nil
}

// Max returns the largest key in this symbol table.
func (bst *BinarySearchST) Max() (Key, error) {
	if bst.IsEmpty() {
		return nil, errors.New("called Max() with empty symbol table")
	}
	return bst.keys[bst.n-1], nil
}

// Select return the kth smallest key in this symbol table.
func (bst *BinarySearchST) Select(k int) (Key, error) {
	if k < 0 || k >= bst.Size() {
		return nil, errors.New("called Select() with invalid argument")
	}
	return bst.keys[k], nil
}

// Floor returns the largest key in this symbol table less than or equal to key.
func (bst *BinarySearchST) Floor(key Key) (Key, error) {
	if key == nil {
		return nil, errors.New("argument to Floor() is nil")
	}
	i, _ := bst.Rank(key)
	if i < bst.n && key.CompareTo(bst.keys[i]) == 0 {
		return key, nil // the key is in the symbol table
	}
	if i == 0 {
		return nil, nil // all keys in the symbol are strictly greater than the key.
	}
	return bst.keys[i-1], nil
}

// Ceiling returns the smallest key in this symbol table greater than or equal to key.
func (bst *BinarySearchST) Ceiling(key Key) (Key, error) {
	if key == nil {
		return nil, errors.New("argument to Ceiling() is ni")
	}
	i, _ := bst.Rank(key)
	if i == bst.n {
		return nil, nil // all keys in the symbol table are strictly less than the key.
	}
	return bst.keys[i], nil
}

// SizeOf returns the number of keys in this symbol table in the specified range.
func (bst *BinarySearchST) SizeOf(lo, hi Key) (int, error) {
	if lo == nil {
		return -1, errors.New("first argument to SizeOf() is nil")
	}
	if hi == nil {
		return -1, errors.New("second argument to SizeOf() is nil")
	}
	if lo.CompareTo(hi) > 0 {
		return 0, nil
	}
	if ok, _ := bst.Contains(hi); ok {
		high, _ := bst.Rank(hi)
		low, _ := bst.Rank(lo)
		return high - low + 1, nil
	}
	high, _ := bst.Rank(hi)
	low, _ := bst.Rank(lo)
	return high - low, nil
}

// Keys return all of the keys in the symbol table.
func (bst *BinarySearchST) Keys() []Key {
	return bst.keys[:bst.n]
}
