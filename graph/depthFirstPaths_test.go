package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDepthFirstPaths(t *testing.T) {
	assert := assert.New(t)
	tinyCG := "8\n" +
		"9\n" +
		"0 5\n" +
		"2 4\n" +
		"2 3\n" +
		"1 2\n" +
		"0 1\n" +
		"3 4\n" +
		"3 5\n" +
		"0 2\n" +
		"6 7\n"
	buf := strings.NewReader(tinyCG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	g, err := NewGraphIn(in)
	assert.Nil(err)
	dfp := NewDepthFirstPaths(g, 0)
	assert.Equal(dfp.PathTo(0), []int{0})
	assert.Equal(dfp.PathTo(1), []int{0, 2, 1})
	assert.Equal(dfp.PathTo(2), []int{0, 2})
	assert.Equal(dfp.PathTo(3), []int{0, 2, 3})
	assert.Equal(dfp.PathTo(4), []int{0, 2, 3, 4})
	assert.Equal(dfp.PathTo(5), []int{0, 2, 3, 5})
	assert.Equal(dfp.PathTo(6), []int(nil))
	assert.Equal(dfp.PathTo(7), []int(nil))

	assert.Panics(func() { NewDepthFirstPaths(g, 8) })
}
