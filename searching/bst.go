package searching

import (
	"errors"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
)

type BST struct {
	root *node
}

type node struct {
	key   Key
	val   Value
	left  *node
	right *node
	size  int
}

func NewBST() *BST {
	return &BST{}
}

func (b *BST) IsEmpty() bool {
	return b.Size() == 0
}

func (b *BST) Size() int {
	return b.size(b.root)
}

func (b *BST) size(x *node) int {
	if x == nil {
		return 0
	}
	return x.size
}

func (b *BST) Contains(key Key) (bool, error) {
	if key == nil {
		return false, errors.New("argument to Contains() is nil")
	}
	return b.Get(key) != nil, nil
}

func (b *BST) Get(key Key) Value {
	val, _ := b.get(b.root, key)
	return val
}

func (b *BST) get(x *node, key Key) (Value, error) {
	if key == nil {
		return nil, errors.New("calls get() with a nil key")
	}
	if x == nil {
		return nil, nil
	}
	if cmp := key.CompareTo(x.key); cmp < 0 {
		val, _ := b.get(x.left, key)
		return val, nil
	} else if cmp > 0 {
		val, _ := b.get(x.right, key)
		return val, nil
	} else {
		return x.val, nil
	}
}

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

func (b *BST) DeleteMin() error {
	if b.IsEmpty() {
		return errors.New("symbol table underflow")
	}
	b.root = b.deleteMin(b.root)
	return nil
}

func (b *BST) deleteMin(x *node) *node {
	if x.left == nil {
		return x.right
	}
	x.left = b.deleteMin(x.left)
	x.size = b.size(x.left) + b.size(x.right) + 1
	return x
}

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

func (b *BST) Floor2(key Key) (Key, error) {
	if x := b.floor2(b.root, key, nil); x == nil {
		return nil, errors.New("argument to Floor2() is nil")
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

func (b *BST) Ceiling(key Key) (Key, error) {
	if key == nil {
		return nil, errors.New("argument to Ceiling() is nil key")
	}
	if b.IsEmpty() {
		return nil, errors.New("calls Ceiling() with empty symbol table")
	}
	if x := b.ceiling(b.root, key); x == nil {
		return nil, errors.New("argument to ceiling() is nil")
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

func (b *BST) Select(rank int) (Key, error) {
	if rank < 0 || rank >= b.Size() {
		return nil, errors.New("argument to Select() is invalid")
	}
	return b.select1(b.root, rank), nil
}

func (b *BST) select1(x *node, rank int) Key {
	if x == nil {
		return nil
	}
	if leftSize := b.size(x.left); leftSize < rank {
		return b.select1(x.left, rank)
	} else if leftSize > rank {
		return b.select1(x.right, rank-leftSize-1)
	} else {
		return x.key
	}
}

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

func (b *BST) Keys() []Key {
	min, _ := b.Min()
	max, _ := b.Max()
	return b.keys(min, max)
}

func (b *BST) keys(lo, hi Key) (keys []Key) {
	queue := arrayqueue.New()
	b.nodeKeys(b.root, queue, lo, hi)
	for !queue.IsEmpty() {
		val, _ := queue.Dequeue()
		keys = append(keys, val.(Key))
	}
	return keys
}

func (b *BST) nodeKeys(x *node, queue *arrayqueue.Queue, lo, hi Key) {
	if x == nil {
		return
	}
	cmplo := lo.CompareTo(x.key)
	cmphi := hi.CompareTo(x.key)
	if cmplo < 0 {
		b.nodeKeys(x.left, queue, lo, hi)
	}
	if cmplo <= 0 && cmphi >= 0 {
		queue.Enqueue(x.key)
	}
	if cmphi > 0 {
		b.nodeKeys(x.right, queue, lo, hi)
	}
}
