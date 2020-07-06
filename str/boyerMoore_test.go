package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoyerMoore(t *testing.T) {
	assert := assert.New(t)

	pattern := "ab"
	text := "da"
	pattern1 := "abracadabra"
	text1 := "abacadabrabracabracadabrabrabracad"

	bm1 := NewBoyerMoore(pattern)
	assert.Equal(len(text), bm1.Search(text))
	assert.Equal(0, bm1.Search("abc"))

	bm2 := NewBoyerMooreR([]byte(pattern), 256)
	assert.Equal(len([]byte(text)), bm2.SearchByte([]byte(text)))
	assert.Equal(0, bm2.SearchByte([]byte("abc")))

	bm3 := NewBoyerMoore(pattern1)
	assert.Equal(14, bm3.Search(text1))
	assert.Equal(3, bm3.Search("def"))

	bm4 := NewBoyerMooreR([]byte(pattern1), 256)
	assert.Equal(14, bm4.SearchByte([]byte(text1)))
	assert.Equal(3, bm4.SearchByte([]byte("def")))
}
