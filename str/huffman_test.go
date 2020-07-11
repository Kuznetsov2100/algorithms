package str

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHuffmanCompress(t *testing.T) {
	assert := assert.New(t)
	var b bytes.Buffer
	buf := bytes.NewBufferString("ABRACADABRA!")
	huffman := NewHuffman(buf, &b)
	huffman.Compress()
	assert.Equal(
		[]byte{0x50, 0x4a, 0x22, 0x43, 0x43, 0x54, 0xa8, 0x40, 0x00, 0x00, 0x01, 0x8f, 0x96, 0x8f, 0x94},
		b.Bytes())
	var b1 bytes.Buffer
	huffman1 := NewHuffman(bytes.NewBuffer(nil), &b1)
	assert.Panics(func() { huffman1.Compress() })

}

func TestHuffmanExpand(t *testing.T) {
	assert := assert.New(t)
	var b bytes.Buffer
	buf := bytes.NewBuffer([]byte{0x50, 0x4a, 0x22, 0x43, 0x43, 0x54, 0xa8, 0x40, 0x00, 0x00, 0x01, 0x8f, 0x96, 0x8f, 0x94})
	huffman := NewHuffman(buf, &b)
	huffman.Expand()
	assert.Equal("ABRACADABRA!", b.String())
}
