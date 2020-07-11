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
	assert.Equal("PJ\"CCT\xa8@\x00\x00\x01\x8f\x96\x8f\x94", b.String())

	var b1 bytes.Buffer
	huffman1 := NewHuffman(bytes.NewBuffer(nil), &b1)
	assert.Panics(func() { huffman1.Compress() })

}

func TestHuffmanExpand(t *testing.T) {
	assert := assert.New(t)
	var b bytes.Buffer
	buf := bytes.NewBufferString("PJ\"CCT\xa8@\x00\x00\x01\x8f\x96\x8f\x94")
	huffman := NewHuffman(buf, &b)
	huffman.Expand()
	assert.Equal("ABRACADABRA!", b.String())
}
