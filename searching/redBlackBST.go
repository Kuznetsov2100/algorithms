package searching

import (
	"errors"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
)

const RED bool = true
const BLACK bool = false

// The RedBlackBST struct represents an ordered symbol table of generic key-value pairs.
// This implementation uses a left-leaning red-black BST.
// The put, get, contains, remove, minimum, maximum, ceiling, floor, rank, and select operations
// each take O(log n) time in the worst case, where n is the number of key-value pairs in the symbol table.
// The size, and is-empty operations take O(1) time.
// The Keys method take O(log n + m) time, where m is the number of keys returned by the method.
// Construction takes O(1) time.
type RedBlackBST struct {
	root *rbnode
}

type rbnode struct {
	key   Key     // key
	val   Value   // associated data
	left  *rbnode // link to  left subtrees
	right *rbnode // link to right subtrees
	color bool    // color of parent link
	size  int     // subtree count
}

// NewRedBlackBST initializes an empty symbol table.
func NewRedBlackBST() *RedBlackBST {
	return &RedBlackBST{}
}

func (b *RedBlackBST) isRed(x *rbnode) bool {
	if x == nil {
		return false
	}
	return x.color
}

// Size returns the number of key-value pairs in this symbol table.
func (b *RedBlackBST) Size() int {
	return b.size(b.root)
}

// number of node in subtree rooted at x; 0 if x is null
func (b *RedBlackBST) size(x *rbnode) int {
	if x == nil {
		return 0
	}
	return x.size
}

// IsEmpty returns true if this symbol table is empty and false otherwise
func (b *RedBlackBST) IsEmpty() bool {
	return b.root == nil
}

// Get returns the value associated with the given key.
// the Get() method uses the standard BST search.
func (b *RedBlackBST) Get(key Key) (Value, error) {
	if key == nil {
		return nil, errors.New("argument to Get() is nil key")
	}
	return b.get(b.root, key), nil
}

func (b *RedBlackBST) get(x *rbnode, key Key) Value {
	for x != nil {
		if cmp := key.CompareTo(x.key); cmp < 0 {
			x = x.left
		} else if cmp > 0 {
			x = x.right
		} else {
			return x.val
		}
	}
	return nil
}

// Contains check whether the symbol table contain the given key.
func (b *RedBlackBST) Contains(key Key) bool {
	val, _ := b.Get(key)
	return val != nil
}

// Put inserts the specified key-value pair into the symbol table,
// overwriting the old value with the new value if the symbol table already contains the specified key.
func (b *RedBlackBST) Put(key Key, val Value) error {
	if key == nil {
		return errors.New("first argument to Put() is nil key")
	}
	if val == nil {
		//nolint:errcheck
		b.Delete(key)
		return nil
	}
	b.root = b.put(b.root, key, val)
	b.root.color = BLACK
	return nil
}

func (b *RedBlackBST) put(h *rbnode, key Key, val Value) *rbnode {
	// Do standard insert, with red link to parent.
	if h == nil {
		return &rbnode{key: key, val: val, color: RED, size: 1}
	}
	if cmp := key.CompareTo(h.key); cmp < 0 {
		h.left = b.put(h.left, key, val)
	} else if cmp > 0 {
		h.right = b.put(h.right, key, val)
	} else {
		h.val = val
	}
	// fix-up any right-leaning links
	// rotates left any right-leaning 3-node (or a right-leaning red link at the bottom of a temporary 4-node)
	if b.isRed(h.right) && !b.isRed(h.left) {
		h = b.rotateLeft(h)
	}
	// rotates right the top link in a temporary 4-node with two left-leaning red links
	if b.isRed(h.left) && b.isRed(h.left.left) {
		h = b.rotateRight(h)
	}
	// flips colors to pass a red link up the tree
	if b.isRed(h.left) && b.isRed(h.right) {
		b.flipColors(h)
	}
	h.size = b.size(h.left) + b.size(h.right) + 1
	return h
}

// DeleteMin removes the smallest key and associated value from the symbol table.
func (b *RedBlackBST) DeleteMin() error {
	if b.IsEmpty() {
		return errors.New("RedBlackBST underflow")
	}
	// if both children of root are black, set root to red
	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.color = RED
	}
	b.root = b.deleteMin(b.root)
	if !b.IsEmpty() {
		b.root.color = BLACK
	}
	return nil
}

func (b *RedBlackBST) deleteMin(h *rbnode) *rbnode {
	if h.left == nil {
		return nil
	}
	if !b.isRed(h.left) && !b.isRed(h.left.left) {
		h = b.moveRedLeft(h)
	}
	h.left = b.deleteMin(h.left)
	return b.balance(h)
}

// DeleteMax removes the largest key and associated value from this symbol table.
func (b *RedBlackBST) DeleteMax() error {
	if b.IsEmpty() {
		return errors.New("RedBlackBST underflow")
	}
	// if both children of root are black, set root to red
	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.color = RED
	}
	b.root = b.deleteMax(b.root)
	if !b.IsEmpty() {
		b.root.color = BLACK
	}
	return nil
}

func (b *RedBlackBST) deleteMax(h *rbnode) *rbnode {
	if b.isRed(h.left) {
		h = b.rotateRight(h)
	}
	if h.right == nil {
		return nil
	}
	if !b.isRed(h.right) && !b.isRed(h.right.left) {
		h = b.moveRedRight(h)
	}
	h.right = b.deleteMax(h.right)
	return b.balance(h)
}

// Delete removes the specified key and associated value from this symbol table (if the key is in the symbol table).
func (b *RedBlackBST) Delete(key Key) error {
	if key == nil {
		return errors.New("argument to Delete() is nil key")
	}
	if !b.Contains(key) {
		return nil
	}
	// if both children of root are black, set root to red
	if !b.isRed(b.root.left) && !b.isRed(b.root.right) {
		b.root.color = RED
	}
	b.root = b.delete(b.root, key)
	if !b.IsEmpty() {
		b.root.color = BLACK
	}
	return nil
}

func (b *RedBlackBST) delete(h *rbnode, key Key) *rbnode {
	if key.CompareTo(h.key) < 0 {
		if !b.isRed(h.left) && !b.isRed(h.left.left) {
			h = b.moveRedLeft(h)
		}
		h.left = b.delete(h.left, key)
	} else {
		if b.isRed(h.left) {
			h = b.rotateRight(h)
		}
		if key.CompareTo(h.key) == 0 && h.right == nil {
			return nil
		}
		if !b.isRed(h.right) && !b.isRed(h.right.left) {
			h = b.moveRedRight(h)
		}
		if key.CompareTo(h.key) == 0 {
			x := b.min(h.right)
			h.key = x.key
			h.val = x.val
			h.right = b.deleteMin(h.right)
		} else {
			h.right = b.delete(h.right, key)
		}
	}
	return b.balance(h)
}

// make a left-leaning link lean to the right
func (b *RedBlackBST) rotateRight(h *rbnode) *rbnode {
	var x *rbnode = h.left
	h.left = x.right
	x.right = h
	x.color = x.right.color
	x.right.color = RED
	x.size = h.size
	h.size = b.size(h.left) + b.size(h.right) + 1
	return x
}

// make a right-leaning link lean to the left
func (b *RedBlackBST) rotateLeft(h *rbnode) *rbnode {
	var x *rbnode = h.right
	h.right = x.left
	x.left = h
	x.color = x.left.color
	x.left.color = RED
	x.size = h.size
	h.size = b.size(h.left) + b.size(h.right) + 1
	return x
}

// flip the colors of a node and its two children
func (b *RedBlackBST) flipColors(h *rbnode) {
	h.color = !h.color
	h.left.color = !h.left.color
	h.right.color = !h.right.color
}

// Assuming that h is red and both h.left and h.left.left
// are black, make h.left or one of its children red.
func (b *RedBlackBST) moveRedLeft(h *rbnode) *rbnode {
	b.flipColors(h)
	if b.isRed(h.right.left) {
		h.right = b.rotateRight(h.right)
		h = b.rotateLeft(h)
		b.flipColors(h)
	}
	return h
}

// Assuming that h is red and both h.right and h.right.left
// are black, make h.right or one of its children red.
func (b *RedBlackBST) moveRedRight(h *rbnode) *rbnode {
	b.flipColors(h)
	if b.isRed(h.left.left) {
		h = b.rotateRight(h)
		b.flipColors(h)
	}
	return h
}

// restore red-black tree invariant
func (b *RedBlackBST) balance(h *rbnode) *rbnode {
	if b.isRed(h.right) {
		h = b.rotateLeft(h)
	}
	if b.isRed(h.left) && b.isRed(h.left.left) {
		h = b.rotateRight(h)
	}
	if b.isRed(h.left) && b.isRed(h.right) {
		b.flipColors(h)
	}
	h.size = b.size(h.left) + b.size(h.right) + 1
	return h
}

// Min returns the smallest key in this symbol table.
func (b *RedBlackBST) Min() (Key, error) {
	if b.IsEmpty() {
		return nil, errors.New("calls Min() with empty symbol table")
	}
	return b.min(b.root).key, nil
}

func (b *RedBlackBST) min(x *rbnode) *rbnode {
	if x.left == nil {
		return x
	}
	return b.min(x.left)
}

// Max returns the largest key in this symbol table.
func (b *RedBlackBST) Max() (Key, error) {
	if b.IsEmpty() {
		return nil, errors.New("calls Max() with empty symbol table")
	}
	return b.max(b.root).key, nil
}

func (b *RedBlackBST) max(x *rbnode) *rbnode {
	if x.right == nil {
		return x
	}
	return b.max(x.right)
}

// Floor returns the largest key in this symbol table less than or equal to key.
func (b *RedBlackBST) Floor(key Key) (Key, error) {
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

// If a given key is less than the key at the root of a RedBlackBST,
// then the floor of key (the largest key in the RedBlackBST less than or equal to key) must be in the left subtree.
// If key is greater than the key at the root, then the floor of key could be in the right subtree,
// but only if there is a key smaller than or equal to key in the right subtree;
// if not (or if key is equal to the key at the root) then the key at the root is the floor of key.
func (b *RedBlackBST) floor(x *rbnode, key Key) *rbnode {
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

// Ceiling returns the smallest key in this symbol table greater than or equal to key.
func (b *RedBlackBST) Ceiling(key Key) (Key, error) {
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

// If a given key is greater than the key at the root of a RedBlackBST,
// then the ceiling of key (the smallest key in the RedBlackBST greater than or equal to key) must be in the right subtree.
// If key is less than the key at the root, then the ceiling of key could be in the left subtree,
// but only if there is a key greater than or equal to key in the left subtree;
// if not (or if key is equal to the key at the root) then the key at the root is the ceiling of key.
func (b *RedBlackBST) ceiling(x *rbnode, key Key) *rbnode {
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
func (b *RedBlackBST) Select(rank int) (Key, error) {
	if rank < 0 || rank >= b.Size() {
		return nil, errors.New("argument to Select() is invalid")
	}
	return b.select1(b.root, rank), nil
}

// If the number of keys leftsize in the left subtree is larger than k,
// we look (recursively) for the key of rank k in the left subtree;
// if leftsize is equal to k, we return the key at the root;
// and if leftsize is smaller than k,
// we look (recursively) for the key of rank k - leftsizes - 1 in the right subtree.
func (b *RedBlackBST) select1(x *rbnode, rank int) Key {
	if leftSize := b.size(x.left); leftSize > rank {
		return b.select1(x.left, rank)
	} else if leftSize < rank {
		return b.select1(x.right, rank-leftSize-1)
	} else {
		return x.key
	}
}

// Rank returns the number of keys in this symbol table strictly less than key.
func (b *RedBlackBST) Rank(key Key) (int, error) {
	if key == nil {
		return -1, errors.New("argument to Rank() is nil")
	}
	return b.rank(key, b.root), nil
}

// If the given key is equal to the key at the root,
// we return the number of keys t in the left subtree;
// if the given key is less than the key at the root,
// we return the rank of the key in the left subtree;
// and if the given key is larger than the key at the root,
// we return t plus one (to count the key at the root) plus the rank of the key in the right subtree.
func (b *RedBlackBST) rank(key Key, x *rbnode) int {
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
func (b *RedBlackBST) SizeOf(lo, hi Key) (int, error) {
	if lo == nil {
		return -1, errors.New("first argument to SizeOf() is nil")
	}
	if hi == nil {
		return -1, errors.New("second argument to SizeOf() is nil")
	}
	if lo.CompareTo(hi) > 0 {
		return 0, nil
	}
	if b.Contains(hi) {
		high, _ := b.Rank(hi)
		low, _ := b.Rank(lo)
		return high - low + 1, nil
	}
	high, _ := b.Rank(hi)
	low, _ := b.Rank(lo)
	return high - low, nil
}

// Keys return all of the keys in the symbol table.
func (b *RedBlackBST) Keys() []Key {
	min, _ := b.Min()
	max, _ := b.Max()
	return b.keys(min, max)
}

func (b *RedBlackBST) keys(lo, hi Key) (keys []Key) {
	queue := arrayqueue.New()
	b.keysOf(b.root, queue, lo, hi)
	for _, x := range queue.Values() {
		keys = append(keys, x.(Key))
	}
	return keys
}

func (b *RedBlackBST) keysOf(x *rbnode, queue *arrayqueue.Queue, lo, hi Key) {
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
