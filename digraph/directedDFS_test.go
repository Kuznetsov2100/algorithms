package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDirectedDFS(t *testing.T) {
	assert := assert.New(t)

	tinyDG := "13\n" +
		"22\n" +
		"4  2\n" +
		"2  3\n" +
		"3  2\n" +
		"6  0\n" +
		"0  1\n" +
		"2  0\n" +
		"11 12\n" +
		"12  9\n" +
		"9 10\n" +
		"9 11\n" +
		"7  9\n" +
		"10 12\n" +
		"11  4\n" +
		"4  3\n" +
		"3  5\n" +
		"6  8\n" +
		"8  6\n" +
		"5  4\n" +
		"0  5\n" +
		"6  4\n" +
		"6  9\n" +
		"7  6\n"

	buf := strings.NewReader(tinyDG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	G, err := NewDigraphIn(in)
	assert.Nil(err)

	dfs := NewDirectedDFS(G, 0)
	assert.Equal(6, dfs.Count())
	for v := 0; v < G.V(); v++ {
		if v <= 5 {
			assert.True(dfs.IsMarked(v))
		} else {
			assert.False(dfs.IsMarked(v))
		}
	}

	dfs1 := NewDirectedDFSources(G, []int{1, 9})
	assert.Equal(10, dfs1.Count())
	for v := 0; v < G.V(); v++ {
		if v <= 5 || v >= 9 {
			assert.True(dfs1.IsMarked(v))
		} else {
			assert.False(dfs1.IsMarked(v))
		}
	}

	assert.Panics(func() { dfs1.IsMarked(13) })
	assert.Panics(func() { NewDirectedDFSources(G, nil) })
}
