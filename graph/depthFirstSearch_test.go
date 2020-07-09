package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
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
	for v := 0; v < g.V(); v++ {
		if v <= 6 {
			assert.True(search.IsMarked(v))
		} else {
			assert.False(search.IsMarked(v))
		}
	}

	assert.NotEqual(search.Count(), g.V())
	assert.Panics(func() { NewDepthFirstSearch(g, 13) })
}
