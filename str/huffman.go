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
	R   int                  // alphabet size of extended ASCII
	in  *binaryin.BinaryIn   // input stream
	out *binaryout.BinaryOut // output stream
}

// NewHuffman constructs the Huffman struct
func NewHuffman(r io.Reader, w io.Writer) *Huffman {
	return &Huffman{R: 256, in: binaryin.NewBinaryIn(r), out: binaryout.NewBinaryOut(w)}
}

// Compress reads a sequence of 8-bit bytes from input stream;compresses them using Huffman codes
// with an 8-bit alphabet; and writes the results to output stream.
func (hf *Huffman) Compress() {
	// read the input
	input, err := hf.in.ReadString()
	if err != nil {
		panic(err)
	}

	// tabulate frequency counts
	freq := make([]int, hf.R)
	for i := range input {
		freq[input[i]]++
	}

	// build Huffman trie
	root := hf.buildTrie(freq)
	st := make([]string, hf.R)

	// build code table
	hf.buildCode(st, root, "")
	// print trie for decoder
	hf.writeTrie(root)

	// print number of bytes in original uncompressed message
	if err := hf.out.WriteInt(len(input)); err != nil {
		panic(err)
	}

	// use Huffman code to encode input
	for i := range input {
		code := st[input[i]]
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
	// close output stream
	hf.out.Close()
}

// Expand reads a sequence of bits that represents a Huffman-compressed message from input stream;
// expands them; and writes the results to output stream.
func (hf *Huffman) Expand() {
	// read in Huffman trie from input stream
	root := hf.readTrie()
	// number of bytes to write
	length, err := hf.in.ReadInt()
	if err != nil {
		panic(err)
	}
	// decode using the Huffman trie
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
		err := hf.out.WriteByte(x.ch)
		if err != nil {
			panic(err)
		}

	}
	hf.out.Close()
}

// build the Huffman trie given frequencies
func (hf *Huffman) buildTrie(freq []int) *hnode {
	// initialze priority queue with singleton trees
	pq := priorityqueue.NewMinPQ()
	for b := 0; b < hf.R; b++ {
		if freq[b] > 0 {
			pq.Insert(newhnode(byte(b), freq[b], nil, nil))
		}
	}

	// merge two smallest trees
	// Nodes with low frequencies end up far down in the trie,
	// and nodes with high frequencies end up near the root of the trie.
	// The frequency in the root equals the number of characters in the input
	for pq.Size() > 1 {
		kl, _ := pq.DelMin()
		left := kl.(*hnode)
		kr, _ := pq.DelMin()
		right := kr.(*hnode)
		parent := newhnode('0', left.freq+right.freq, left, right)
		pq.Insert(parent)
	}
	tree, _ := pq.DelMin()
	return tree.(*hnode)
}

// make a lookup table from symbols and their encodings
func (hf *Huffman) buildCode(st []string, x *hnode, s string) {
	if !x.isLeaf() {
		hf.buildCode(st, x.left, s+"0")
		hf.buildCode(st, x.right, s+"1")
	} else {
		st[x.ch] = s
	}
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
	}
	return newhnode('0', -1, hf.readTrie(), hf.readTrie())
}

// write bitstring-encoded trie to output stream
func (hf *Huffman) writeTrie(x *hnode) {
	if x.isLeaf() {
		hf.out.WriteBit(true)
		err := hf.out.WriteByte(x.ch)
		if err != nil {
			panic(err)
		}
		return
	}
	hf.out.WriteBit(false)
	hf.writeTrie(x.left)
	hf.writeTrie(x.right)
}

// // Huffman trie node
type hnode struct {
	ch    byte
	freq  int
	left  *hnode
	right *hnode
}

func newhnode(ch byte, freq int, left, right *hnode) *hnode {
	return &hnode{ch: ch, freq: freq, left: left, right: right}
}

// // is the node a leaf node?
func (node *hnode) isLeaf() bool {
	return node.left == nil && node.right == nil
}

// // compare, based on frequency
func (node *hnode) CompareTo(key priorityqueue.Key) int {
	that := key.(*hnode)
	if node.freq < that.freq {
		return -1
	} else if node.freq > that.freq {
		return 1
	}
	return 0
}
