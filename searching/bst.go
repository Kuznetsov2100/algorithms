package searching

import (
	"errors"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
)

// BST struct represents an ordered symbol table of generic key-value pairs.
// This implementation uses an (unbalanced) binary search tree.
// The put, contains, remove, minimum, maximum, ceiling, floor, select, and rank operations each take O(n) time in the worst case,
// where n is the number of key-value pairs. The size and is-empty operations take O(1) time.
// The keys method takes O(n) time in the worst case.
// Construction takes O(1) time.
type BST struct {
	root *node
}

// node represents the node in a binary search tree.
type node struct {
	key   Key
	val   Value
	left  *node
	right *node
	size  int
}

// NewBST initializes an empty symbol table.
func NewBST() *BST {
	return &BST{}
}

// IsEmpty returns true if this symbol table is empty.
func (b *BST) IsEmpty() bool {
	return b.Size() == 0
}

// Size returns the number of key-value pairs in this symbol table.
func (b *BST) Size() int {
	return b.size(b.root)
}

func (b *BST) size(x *node) int {
	if x == nil {
		return 0
	}
	return x.size
}

// Contains check whether the symbol table contain the given key.
func (b *BST) Contains(key Key) (bool, error) {
	if key == nil {
		return false, errors.New("argument to Contains() is nil")
	}
	val, _ := b.Get(key)
	return val != nil, nil
}

// Get returns the value associated with the given key in this symbol table.
func (b *BST) Get(key Key) (Value, error) {
	if key == nil {
		return nil, errors.New("calls get() with a nil key")
	}

	return b.get(b.root, key), nil
}

func (b *BST) get(x *node, key Key) Value {
	if x == nil {
		return nil
	}
	if cmp := key.CompareTo(x.key); cmp < 0 {
		return b.get(x.left, key)
	} else if cmp > 0 {

		return b.get(x.right, key)
	} else {
		return x.val
	}
}

// Put inserts the specified key-value pair into the symbol table,
// overwriting the old value with the new value if the symbol table already contains the specified key.
func (b *BST) Put(key Key, val Value) error {
	if key == nil {
		return errors.New("calls Put() with a nil key")
	}
	if val == nil {
		//nolint:errcheck
		b.Delete(key)
		return nil
	}
	b.root = b.put(b.root, key, val)
	return nil
}

func (b *BST) put(x *node, key Key, val Value) *node {
	if x == nil {
		return &node{
			key:  key,
			val:  val,
			size: 1,
		}
	}
	if cmp := key.CompareTo(x.key); cmp < 0 {
		x.left = b.put(x.left, key, val)
	} else if cmp > 0 {
		x.right = b.put(x.right, key, val)
	} else {
		x.val = val
	}
	x.size = 1 + b.size(x.left) + b.size(x.right)
	return x
}

// DeleteMin removes the smallest key and associated value from the symbol table.
func (b *BST) DeleteMin() error {
	if b.IsEmpty() {
		return errors.New("symbol table underflow")
	}
	b.root = b.deleteMin(b.root)
	return nil
}

// DeleteMin removes the smallest key and associated value from this symbol table.
func (b *BST) deleteMin(x *node) *node {
	if x.left == nil {
		return x.right
	}
	x.left = b.deleteMin(x.left)
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

// DeleteMax removes the largest key and associated value from this symbol table.
func (b *BST) DeleteMax() error {
	if b.IsEmpty() {
		return errors.New("symbol table underflow")
	}
	b.root = b.deleteMax(b.root)
	return nil
}

func (b *BST) deleteMax(x *node) *node {
	if x.right == nil {
		return x.left
	}
	x.right = b.deleteMax(x.right)
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

// Delete removes the specified key and associated value from this symbol table (if the key is in the symbol table).
func (b *BST) Delete(key Key) error {
	if key == nil {
		return errors.New("calls Delete() with a nil key")
	}
	b.root = b.delete(b.root, key)
	return nil
}

func (b *BST) delete(x *node, key Key) *node {
	if x == nil {
		return nil
	}
	if cmp := key.CompareTo(x.key); cmp < 0 {
		x.left = b.delete(x.left, key)
	} else if cmp > 0 {
		x.right = b.delete(x.right, key)
	} else {
		if x.right == nil {
			return x.left
		}
		if x.left == nil {
			return x.right
		}
		var t *node = x
		x = b.min(t.right)
		x.right = b.deleteMin(t.right)
		x.left = t.left
	}
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

// Min returns the smallest key in this symbol table.
func (b *BST) Min() (Key, error) {
	if b.IsEmpty() {
		return nil, errors.New("calls Min() with empty symbol table")
	}
	return b.min(b.root).key, nil
}

func (b *BST) min(x *node) *node {
	if x.left == nil {
		return x
	}
	return b.min(x.left)
}

// Max returns the largest key in this symbol table.
func (b *BST) Max() (Key, error) {
	if b.IsEmpty() {
		return nil, errors.New("calls Max() with empty symbol table")
	}
	return b.max(b.root).key, nil
}

func (b *BST) max(x *node) *node {
	if x.right == nil {
		return x
	}
	return b.max(x.right)
}

// Floor returns the largest key in this symbol table less than or equal to key.
func (b *BST) Floor(key Key) (Key, error) {
	if key == nil {
		return nil, errors.New("argument to Floor() is nil")
	}
	if b.IsEmpty() {
		return nil, errors.New("calls Floor() with empty symbol table")
	}
	if x := b.floor(b.root, key); x == nil {
		return nil, errors.New("argument to Floor() is too small")
	} else {
		return x.key, nil
	}
}

func (b *BST) floor(x *node, key Key) *node {
	if x == nil {
		return nil
	}
	if cmp := key.CompareTo(x.key); cmp == 0 {
		return x
	} else if cmp < 0 {
		return b.floor(x.left, key)
	}
	if t := b.floor(x.right, key); t != nil {
		return t
	} else {
		return x
	}
}

// Floor2 returns the largest key in this symbol table less than or equal to key.
func (b *BST) Floor2(key Key) (Key, error) {
	if key == nil {
		return nil, errors.New("argument to Floor2() is nil key")
	}
	if b.IsEmpty() {
		return nil, errors.New("calls Floor() with empty symbol table")
	}
	if x := b.floor2(b.root, key, nil); x == nil {
		return nil, errors.New("argument to Floor2() is too small")
	} else {
		return x, nil
	}
}

func (b *BST) floor2(x *node, key Key, best Key) Key {
	if x == nil {
		return best
	}
	if cmp := key.CompareTo(x.key); cmp < 0 {
		return b.floor2(x.left, key, best)
	} else if cmp > 0 {
		return b.floor2(x.right, key, x.key)
	} else {
		return x.key
	}
}

// Ceiling returns the smallest key in this symbol table greater than or equal to key.
func (b *BST) Ceiling(key Key) (Key, error) {
	if key == nil {
		return nil, errors.New("argument to Ceiling() is nil key")
	}
	if b.IsEmpty() {
		return nil, errors.New("calls Ceiling() with empty symbol table")
	}
	if x := b.ceiling(b.root, key); x == nil {
		return nil, errors.New("argument to Ceiling() is too large")
	} else {
		return x.key, nil
	}
}

func (b *BST) ceiling(x *node, key Key) *node {
	if x == nil {
		return nil
	}
	if cmp := key.CompareTo(x.key); cmp == 0 {
		return x
	} else if cmp < 0 {
		if t := b.ceiling(x.left, key); t != nil {
			return t
		} else {
			return x
		}
	}
	return b.ceiling(x.right, key)
}

// Select return the key in the symbol table of a given rank.
func (b *BST) Select(rank int) (Key, error) {
	if rank < 0 || rank >= b.Size() {
		return nil, errors.New("argument to Select() is invalid")
	}
	return b.select1(b.root, rank), nil
}

func (b *BST) select1(x *node, rank int) Key {
	if leftSize := b.size(x.left); leftSize > rank {
		return b.select1(x.left, rank)
	} else if leftSize < rank {
		return b.select1(x.right, rank-leftSize-1)
	} else {
		return x.key
	}
}

// Rank returns the number of keys in this symbol table strictly less than key.
func (b *BST) Rank(key Key) (int, error) {
	if key == nil {
		return -1, errors.New("argument to Rank() is nil")
	}
	return b.rank(key, b.root), nil
}

func (b *BST) rank(key Key, x *node) int {
	if x == nil {
		return 0
	}
	if cmp := key.CompareTo(x.key); cmp < 0 {
		return b.rank(key, x.left)
	} else if cmp > 0 {
		return 1 + b.size(x.left) + b.rank(key, x.right)
	} else {
		return b.size(x.left)
	}
}

// SizeOf returns the number of keys in this symbol table in the specified range.
func (b *BST) SizeOf(lo, hi Key) (int, error) {
	if lo == nil {
		return -1, errors.New("first argument to SizeOf() is nil")
	}
	if hi == nil {
		return -1, errors.New("second argument to SizeOf() is nil")
	}
	if lo.CompareTo(hi) > 0 {
		return 0, nil
	}
	if ok, _ := b.Contains(hi); ok {
		high, _ := b.Rank(hi)
		low, _ := b.Rank(lo)
		return high - low + 1, nil
	}
	high, _ := b.Rank(hi)
	low, _ := b.Rank(lo)
	return high - low, nil
}

// Keys return all of the keys in the symbol table.
func (b *BST) Keys() []Key {
	min, _ := b.Min()
	max, _ := b.Max()
	return b.keys(min, max)
}

func (b *BST) keys(lo, hi Key) (keys []Key) {
	queue := arrayqueue.New()
	b.keysOf(b.root, queue, lo, hi)
	for _, x := range queue.Values() {
		keys = append(keys, x.(Key))
	}
	return keys
}

func (b *BST) keysOf(x *node, queue *arrayqueue.Queue, lo, hi Key) {
	if x == nil {
		return
	}
	cmplo := lo.CompareTo(x.key)
	cmphi := hi.CompareTo(x.key)
	if cmplo < 0 {
		b.keysOf(x.left, queue, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0 {
		queue.Enqueue(x.key)
	}
	if cmphi > 0 {
		b.keysOf(x.right, queue, lo, hi)
	}
}
