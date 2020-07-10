package str

import (
	"sync"

	"github.com/handane123/algorithms/dataStructure/priorityqueue"
	"github.com/handane123/algorithms/io/binarystdin"
	"github.com/handane123/algorithms/io/binarystdout"
)

var huffman *Huffman
var once sync.Once

type Huffman struct {
	R int
}

func NewHuffman() *Huffman {
	once.Do(func() {
		huffman = &Huffman{R: 256}
	})
	return huffman
}

type hnode struct {
	ch    byte
	freq  int
	left  *hnode
	right *hnode
}

func newhnode(ch byte, freq int, left *hnode, right *hnode) *hnode {
	return &hnode{
		ch:    ch,
		freq:  freq,
		left:  left,
		right: right,
	}
}

func (node *hnode) isLeaf() bool {
	return node.left == nil && node.right == nil
}

func (node *hnode) CompareTo(key priorityqueue.Key) int {
	that := key.(*hnode)
	if node.freq < that.freq {
		return -1
	} else if node.freq > that.freq {
		return 1
	}
	return 0
}

func (hf *Huffman) Compress() {
	binarystdin := binarystdin.NewBinaryStdIn()
	binarystdout := binarystdout.NewBinaryStdOut()
	s, err := binarystdin.ReadString()
	if err != nil {
		panic(err)
	}
	freq := make([]int, hf.R)
	for i := range s {
		freq[s[i]]++
	}
	root := hf.buildTrie(freq)
	st := make([]string, hf.R)
	hf.buildCode(st, root, "")
	hf.writeTrie(root)
	err = binarystdout.WriteInt(len(s))
	if err != nil {
		panic(err)
	}

	for i := range s {
		code := st[s[i]]
		for j := range code {
			if code[j] == '0' {
				binarystdout.WriteBit(false)
			} else if code[j] == '1' {
				binarystdout.WriteBit(true)
			} else {
				panic("illegal state")
			}
		}
	}
	binarystdout.Close()
}

func (hf *Huffman) Expand() {
	root := hf.readTrie()
	binarystdin := binarystdin.NewBinaryStdIn()
	binarystdout := binarystdout.NewBinaryStdOut()
	length, err := binarystdin.ReadInt()
	if err != nil {
		panic(err)
	}
	for i := 0; i < length; i++ {
		x := root
		for !x.isLeaf() {
			bit, err := binarystdin.ReadBool()
			if err != nil {
				panic(err)
			}
			if bit {
				x = x.right
			} else {
				x = x.left
			}
		}
		err := binarystdout.WriteBitR(int(x.ch), 8)
		if err != nil {
			panic(err)
		}

	}
	binarystdout.Close()
}

func (hf *Huffman) readTrie() *hnode {
	binarystdin := binarystdin.NewBinaryStdIn()
	isLeaf, err := binarystdin.ReadBool()
	if err != nil {
		panic(err)
	}
	if isLeaf {
		b, err1 := binarystdin.ReadByte()
		if err1 != nil {
			panic(err1)
		}
		return newhnode(b, -1, nil, nil)
	} else {
		return newhnode('0', -1, hf.readTrie(), hf.readTrie())
	}
}

func (hf *Huffman) buildTrie(freq []int) *hnode {
	pq := priorityqueue.NewMinPQ()
	for b := 0; b < hf.R; b++ {
		if freq[b] > 0 {
			pq.Insert(newhnode(byte(b), freq[b], nil, nil))
		}
	}

	for pq.Size() > 1 {
		kl, _ := pq.DelMin()
		left := kl.(*hnode)
		kr, _ := pq.DelMin()
		right := kr.(*hnode)
		parent := newhnode(0, left.freq+right.freq, left, right)
		pq.Insert(parent)
	}
	k, _ := pq.DelMin()
	return k.(*hnode)
}

func (hf *Huffman) buildCode(st []string, x *hnode, s string) {
	if !x.isLeaf() {
		hf.buildCode(st, x.left, s+"0")
		hf.buildCode(st, x.right, s+"1")
	} else {
		st[x.ch] = s
	}
}

func (hf *Huffman) writeTrie(x *hnode) {
	binarystdout := binarystdout.NewBinaryStdOut()
	if x.isLeaf() {
		binarystdout.WriteBit(true)
		err := binarystdout.WriteBitR(int(x.ch), 8)
		if err != nil {
			panic(err)
		}
		return
	}
	binarystdout.WriteBit(false)
	hf.writeTrie(x.left)
	hf.writeTrie(x.right)
}
