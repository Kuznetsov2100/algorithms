package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKMP(t *testing.T) {
	assert := assert.New(t)

	pattern := "ab"
	text := "da"

	kmp := NewKMP(pattern)
	assert.Equal(len(text), kmp.Search(text))
	assert.Equal(0, kmp.Search("abc"))

	kmp1 := NewKMPR([]byte(pattern), 256)
	assert.Equal(len([]byte(text)), kmp1.SearchByte([]byte(text)))
	assert.Equal(0, kmp1.SearchByte([]byte("abc")))
}
