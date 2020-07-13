package str

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLZWCompress(t *testing.T) {
	assert := assert.New(t)
	var b bytes.Buffer
	buf := bytes.NewBufferString("ABRACADABRABRABRA")
	lzw := NewLZW(buf, &b)
	lzw.Compress()
	assert.Equal(
		[]byte{0x04, 0x10, 0x42, 0x05, 0x20, 0x41, 0x04, 0x30, 0x41, 0x04, 0x41, 0x01, 0x10, 0x31, 0x02, 0x10,
			0x80, 0x41, 0x10, 0x00},
		b.Bytes())

	var b1 bytes.Buffer
	lzw1 := NewLZW(bytes.NewBuffer(nil), &b1)
	assert.Panics(func() { lzw1.Compress() })

	var b2 bytes.Buffer
	buf2 := bytes.NewBufferString("ABABABA")
	lzw2 := NewLZW(buf2, &b2)
	lzw2.Compress()
	assert.Equal([]byte{0x04, 0x10, 0x42, 0x10, 0x11, 0x03, 0x10, 0x00}, b2.Bytes())

}

func TestLZWExpand(t *testing.T) {
	assert := assert.New(t)
	var b bytes.Buffer
	buf := bytes.NewBuffer([]byte{0x04, 0x10, 0x42, 0x05, 0x20, 0x41, 0x04, 0x30, 0x41, 0x04, 0x41, 0x01, 0x10, 0x31, 0x02, 0x10,
		0x80, 0x41, 0x10, 0x00})
	lzw := NewLZW(buf, &b)
	lzw.Expand()
	assert.Equal("ABRACADABRABRABRA", b.String())

	var b3 bytes.Buffer
	lzw3 := NewLZW(bytes.NewBuffer([]byte{0x04}), &b3)
	assert.Panics(func() { lzw3.Expand() })

	var b4 bytes.Buffer
	lzw4 := NewLZW(bytes.NewBuffer([]byte{0x04, 0x10}), &b4)
	assert.Panics(func() { lzw4.Expand() })

	// codeword == lzw.asciiR  expanded meesage is empty string
	var b5 bytes.Buffer
	lzw5 := NewLZW(bytes.NewBuffer([]byte{0x10, 0x00}), &b5)
	lzw5.Expand()
	assert.Equal(0, b5.Len())

	var b6 bytes.Buffer
	buf6 := bytes.NewBuffer([]byte{0x04, 0x10, 0x42, 0x10, 0x11, 0x03, 0x10, 0x00})
	lzw6 := NewLZW(buf6, &b6)
	lzw6.Expand()
	assert.Equal("ABABABA", b6.String())
}
