package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDepthFirstOrder(t *testing.T) {
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
	g, err := NewDigraphIn(in)
	assert.Nil(err)

	dfo := NewDepthFirstOrder(g)
	pre := []int{0, 5, 4, 3, 2, 1, 6, 12, 11, 7, 10, 8, 9}
	post := []int{5, 4, 0, 1, 2, 3, 11, 12, 10, 9, 8, 7, 6}
	for v := 0; v < g.V(); v++ {
		assert.Equal(pre[v], dfo.Pre(v))
		assert.Equal(post[v], dfo.Post(v))
	}
	assert.Panics(func() { dfo.Pre(13) })

	assert.Equal([]int{0, 5, 4, 3, 2, 1, 6, 9, 11, 12, 10, 8, 7}, dfo.PreOrder())
	assert.Equal([]int{2, 3, 4, 5, 1, 0, 12, 11, 10, 9, 8, 6, 7}, dfo.PostOrder())
	assert.Equal([]int{7, 6, 8, 9, 10, 11, 12, 0, 1, 5, 4, 3, 2}, dfo.ReversePost())

}
