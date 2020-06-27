package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestBipartite(t *testing.T) {
	assert := assert.New(t)
	bipartitegraphdata := "15\n" +
		"10\n" +
		"0 11\n" +
		"1 13\n" +
		"1 10\n" +
		"1 11\n" +
		"3 13\n" +
		"3 5\n" +
		"4 11\n" +
		"6 13\n" +
		"9 13\n" +
		"11 14\n"
	buf := strings.NewReader(bipartitegraphdata)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	g, err := NewGraphIn(in)
	assert.Nil(err)

	b := NewBipartite(g)
	assert.Equal([]int(nil), b.OddCycle())
	assert.True(b.IsBipartite())
	assert.False(b.Color(0))
	assert.False(b.Color(1))
	assert.False(b.Color(2))
	assert.False(b.Color(3))
	assert.False(b.Color(4))
	assert.True(b.Color(5))
	assert.False(b.Color(6))
	assert.False(b.Color(7))
	assert.False(b.Color(8))
	assert.False(b.Color(9))
	assert.True(b.Color(10))
	assert.True(b.Color(11))
	assert.False(b.Color(12))
	assert.True(b.Color(13))
	assert.False(b.Color(14))
	assert.Panics(func() { b.Color(15) })

	nobipartitegraphdata := "15\n" +
		"10\n" +
		"0 4\n" +
		"0 9\n" +
		"0 9\n" +
		"0 10\n" +
		"1 13\n" +
		"3 10\n" +
		"3 14\n" +
		"4 10\n" +
		"4 9\n" +
		"8 14\n"
	buf1 := strings.NewReader(nobipartitegraphdata)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	g1, err1 := NewGraphIn(in1)
	assert.Nil(err1)

	b1 := NewBipartite(g1)
	assert.False(b1.IsBipartite())
	assert.Equal([]int{0, 10, 4, 0}, b1.OddCycle())
	assert.PanicsWithValue("graph is not  bipartite", func() { b1.Color(0) })
}
