package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDepthFirstSearch(t *testing.T) {
	assert := assert.New(t)
	tinyG := "13\n" +
		"13\n" +
		"0 5\n" +
		"4 3\n" +
		"0 1\n" +
		"9 12\n" +
		"6 4\n" +
		"5 4\n" +
		"0 2\n" +
		"11 12\n" +
		"9 10\n" +
		"0 6\n" +
		"7 8\n" +
		"9 11\n" +
		"5 3\n"
	buf := strings.NewReader(tinyG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	g, err := NewGraphIn(in)
	assert.Nil(err)
	search := NewDepthFirstSearch(g, 0)
	assert.Equal(true, search.IsMarked(0))
	assert.Equal(true, search.IsMarked(1))
	assert.Equal(true, search.IsMarked(2))
	assert.Equal(true, search.IsMarked(3))
	assert.Equal(true, search.IsMarked(4))
	assert.Equal(true, search.IsMarked(5))
	assert.Equal(true, search.IsMarked(6))
	assert.Equal(false, search.IsMarked(7))
	assert.Equal(false, search.IsMarked(8))
	assert.Equal(false, search.IsMarked(9))
	assert.Equal(false, search.IsMarked(10))
	assert.Equal(false, search.IsMarked(11))
	assert.Equal(false, search.IsMarked(12))
	assert.NotEqual(search.Count(), g.V())
	assert.Panics(func() { NewDepthFirstSearch(g, 13) })
}
