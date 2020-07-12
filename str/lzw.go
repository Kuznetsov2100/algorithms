package str

import (
	"io"

	"github.com/handane123/algorithms/io/binaryin"
	"github.com/handane123/algorithms/io/binaryout"
)

type LZW struct {
	R   int
	L   int
	W   int
	in  *binaryin.BinaryIn
	out *binaryout.BinaryOut
}

func NewLZW(r io.Reader, w io.Writer) *LZW {
	return &LZW{R: 256, L: 4096, W: 12,
		in:  binaryin.NewBinaryIn(r),
		out: binaryout.NewBinaryOut(w),
	}
}

func (lzw *LZW) Compress() {
	input, err := lzw.in.ReadString()
	if err != nil {
		panic(err)
	}
	st := NewTST()
	for i := 0; i < lzw.R; i++ {
		st.Put(string(i), i)
	}
	code := lzw.R + 1

	for len(input) > 0 {
		s := st.LongestPrefixOf(input)
		val := st.Get(s)
		//nolint:errcheck
		lzw.out.WriteBitR(val.(int), lzw.W)
		t := len(s)
		if t < len(input) && code < lzw.L {
			st.Put(input[:t+1], code)
			code++
		}
		input = input[t:]
	}
	//nolint:errcheck
	lzw.out.WriteBitR(lzw.R, lzw.W)
	lzw.out.Close()
}

func (lzw *LZW) Expand() {
	st := make([]string, lzw.L)
	var i int

	for i = 0; i < lzw.R; i++ {
		st[i] = string(i)
	}
	st[i] = ""
	i++

	codeword, err := lzw.in.ReadIntR(lzw.W)
	if err != nil {
		panic(err)
	}
	if codeword == lzw.R {
		return
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
			s = val + string(val[0])
		}
		if i < lzw.L {
			st[i] = val + string(s[0])
			i++
		}
		val = s
	}
	lzw.out.Close()
}
