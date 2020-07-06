package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRabinKarp(t *testing.T) {
	assert := assert.New(t)

	pattern := "ab"
	text := "da"
	pattern1 := "abracadabra"
	text1 := "abacadabrabracabracadabrabrabracad"

	rk1 := NewRabinKarp(pattern)
	assert.Equal(len(text), rk1.Search(text))
	assert.Equal(0, rk1.Search("abc"))

	rk2 := NewRabinKarp(pattern1)
	assert.Equal(14, rk2.Search(text1))
	assert.Equal(3, rk2.Search("def"))
	assert.False(rk2.check(text1, 9))

}
