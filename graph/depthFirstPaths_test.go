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
	assert.Equal([]int{0}, dfp.PathTo(0))
	assert.Equal([]int{0, 2, 1}, dfp.PathTo(1))
	assert.Equal([]int{0, 2}, dfp.PathTo(2))
	assert.Equal([]int{0, 2, 3}, dfp.PathTo(3))
	assert.Equal([]int{0, 2, 3, 4}, dfp.PathTo(4))
	assert.Equal([]int{0, 2, 3, 5}, dfp.PathTo(5))
	assert.Equal([]int(nil), dfp.PathTo(6))
	assert.Equal([]int(nil), dfp.PathTo(7))

	assert.Panics(func() { NewDepthFirstPaths(g, 8) })
}
