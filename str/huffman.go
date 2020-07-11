package str

import (
	"io"

	"github.com/handane123/algorithms/dataStructure/priorityqueue"
	"github.com/handane123/algorithms/io/binaryin"
	"github.com/handane123/algorithms/io/binaryout"
)

// Huffman struct provides static methods for compressing and expanding a binary input
// using Huffman codes over the 8-bit extended ASCII alphabet.
type Huffman struct {
	R   int
	in  *binaryin.BinaryIn
	out *binaryout.BinaryOut
}

func NewHuffman(r io.Reader, w io.Writer) *Huffman {

	return &Huffman{R: 256, in: binaryin.NewBinaryIn(r), out: binaryout.NewBinaryOut(w)}

}

// Compress reads a sequence of 8-bit bytes from standard input;
// compresses them using Huffman codes with an 8-bit alphabet;
// and writes the results to standard output.
func (hf *Huffman) Compress() {
	s, err := hf.in.ReadString()
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
	err = hf.out.WriteInt(len(s))
	if err != nil {
		panic(err)
	}

	for i := range s {
		code := st[s[i]]
		for j := range code {
			if code[j] == '0' {
				hf.out.WriteBit(false)
			} else if code[j] == '1' {
				hf.out.WriteBit(true)
			} else {
				panic("illegal state")
			}
		}
	}
	hf.out.Close()
}

func (hf *Huffman) Expand() {
	root := hf.readTrie()
	length, err := hf.in.ReadInt()
	if err != nil {
		panic(err)
	}
	for i := 0; i < length; i++ {
		x := root
		for !x.isLeaf() {
			bit, err := hf.in.ReadBool()
			if err != nil {
				panic(err)
			}
			if bit {
				x = x.right
			} else {
				x = x.left
			}
		}
		err := hf.out.WriteBitR(int(x.ch), 8)
		if err != nil {
			panic(err)
		}

	}
	hf.out.Close()
}

func (hf *Huffman) readTrie() *hnode {
	isLeaf, err := hf.in.ReadBool()
	if err != nil {
		panic(err)
	}
	if isLeaf {
		b, err1 := hf.in.ReadByte()
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
	if x.isLeaf() {
		hf.out.WriteBit(true)
		err := hf.out.WriteBitR(int(x.ch), 8)
		if err != nil {
			panic(err)
		}
		return
	}
	hf.out.WriteBit(false)
	hf.writeTrie(x.left)
	hf.writeTrie(x.right)
}

type hnode struct {
	ch    byte
	freq  int
	left  *hnode
	right *hnode
}

func newhnode(ch byte, freq int, left *hnode, right *hnode) *hnode {
	return &hnode{ch: ch, freq: freq, left: left, right: right}
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
