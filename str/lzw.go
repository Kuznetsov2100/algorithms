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
	R   int // number of input chars:256
	L   int // number of codewords = 2^W = 4096
	W   int // codeword width:12
	in  *binaryin.BinaryIn
	out *binaryout.BinaryOut
}

// NewLZW constructs the LZW struct
func NewLZW(r io.Reader, w io.Writer) *LZW {
	return &LZW{R: 256, L: 4096, W: 12,
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
	for i := 0; i < lzw.R; i++ {
		st.Put(fmt.Sprintf("%c", i), i)
	}
	code := lzw.R + 1 // R is codeword for EOF

	for len(input) > 0 {
		s := st.LongestPrefixOf(input) // Find max prefix match s.
		//nolint:errcheck
		lzw.out.WriteBitR(st.Get(s).(int), lzw.W) // Print s's encoding.
		t := len(s)
		if t < len(input) && code < lzw.L {
			st.Put(input[:t+1], code) // Add s to symbol table.
			code++
		}
		input = input[t:] // Scan past s in input.
	}
	//nolint:errcheck
	lzw.out.WriteBitR(lzw.R, lzw.W)
	lzw.out.Close()
}

// Expand reads a sequence of bit encoded using LZW compression with 12-bit codewords from input stream;
// expands them; and writes the results to output stream.
func (lzw *LZW) Expand() {
	st := make([]string, lzw.L)
	var i int // next available codeword value

	// initialize symbol table with all 1-character strings
	for i = 0; i < lzw.R; i++ {
		st[i] = fmt.Sprintf("%c", i)
	}
	st[i] = "" // (unused) lookahead for EOF
	i++

	codeword, err := lzw.in.ReadIntR(lzw.W)
	if err != nil {
		panic(err)
	}
	if codeword == lzw.R {
		return // expanded message is empty string
	}

	val := st[codeword]
	for {
		//nolint:errcheck
		lzw.out.WriteString(val)
		codeword, err = lzw.in.ReadIntR(lzw.W)
		if err != nil {
			panic(err)
		}
		if codeword == lzw.R {
			break
		}
		s := st[codeword]
		if i == codeword {
			s = val + string(val[0]) // special case hack
		}
		if i < lzw.L {
			st[i] = val + string(s[0])
			i++
		}
		val = s
	}
	lzw.out.Close()
}
