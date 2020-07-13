package str

import (
	"fmt"
	"io"

	"github.com/handane123/algorithms/io/binaryin"
	"github.com/handane123/algorithms/io/binaryout"
)

// LZW struct provides static methods for compressing and expanding a binary input
// using LZW compression over the 8-bit extended ASCII alphabet with 12-bit codewords.
type LZW struct {
	asciiR  int // number of input chars:256
	numberL int // number of codewords = 2^W = 4096
	widthW  int // codeword width:12
	in      *binaryin.BinaryIn
	out     *binaryout.BinaryOut
}

// NewLZW constructs the LZW struct
func NewLZW(r io.Reader, w io.Writer) *LZW {
	return &LZW{asciiR: 256, numberL: 4096, widthW: 12,
		in:  binaryin.NewBinaryIn(r),
		out: binaryout.NewBinaryOut(w),
	}
}

// Compress reads a sequence of 8-bit bytes from input stream;
// compresses them using LZW compression with 12-bit codewords;
// and writes the results to output stream.
func (lzw *LZW) Compress() {
	input, err := lzw.in.ReadString()
	if err != nil {
		panic(err)
	}
	st := NewTST()
	for i := 0; i < lzw.asciiR; i++ {
		st.Put(fmt.Sprintf("%c", i), i)
	}
	code := lzw.asciiR + 1 // R is codeword for EOF

	for len(input) > 0 {
		s := st.LongestPrefixOf(input) // Find max prefix match s.
		//nolint:errcheck
		lzw.out.WriteBitR(st.Get(s).(int), lzw.widthW) // Print s's encoding.
		t := len(s)
		if t < len(input) && code < lzw.numberL {
			st.Put(input[:t+1], code) // Add s+c to symbol table.
			code++
		}
		input = input[t:] // Scan past s in input.
	}
	//nolint:errcheck
	lzw.out.WriteBitR(lzw.asciiR, lzw.widthW) // write EOF  0x10,0x00
	lzw.out.Close()
}

// Expand reads a sequence of bit encoded using LZW compression with 12-bit codewords from input stream;
// expands them; and writes the results to output stream.
func (lzw *LZW) Expand() {
	st := make([]string, lzw.numberL)
	var i int // next available codeword value

	// initialize symbol table with all 1-character strings
	for i = 0; i < lzw.asciiR; i++ {
		st[i] = fmt.Sprintf("%c", i)
	}
	st[i] = "" // (unused) lookahead for EOF
	i++

	codeword, err := lzw.in.ReadIntR(lzw.widthW)
	if err != nil {
		panic(err)
	}
	if codeword == lzw.asciiR {
		return // expanded message is empty string
	}

	val := st[codeword]
	for {
		//nolint:errcheck
		lzw.out.WriteString(val) // write curent substring
		codeword, err = lzw.in.ReadIntR(lzw.widthW)
		if err != nil {
			panic(err)
		}
		if codeword == lzw.asciiR {
			break
		}
		s := st[codeword]  // get next codeword
		if i == codeword { // if lookahead is invalid
			s = val + string(val[0]) // special case hack, make codeword from last one
		}
		if i < lzw.numberL {
			st[i] = val + string(s[0]) // add new entry to code table
			i++                        // update current codeword
		}
		val = s
	}
	lzw.out.Close()
}
