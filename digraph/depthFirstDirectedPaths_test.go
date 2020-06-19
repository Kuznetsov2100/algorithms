package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDepthFirstPaths(t *testing.T) {
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
	dfp := NewDepthFirstDirectedPaths(g, 0)
	assert.Equal([]int{0}, dfp.PathTo(0))
	assert.Equal([]int{0, 1}, dfp.PathTo(1))
	assert.Equal([]int{0, 5, 4, 3, 2}, dfp.PathTo(2))
	assert.Equal([]int{0, 5, 4, 3}, dfp.PathTo(3))
	assert.Equal([]int{0, 5, 4}, dfp.PathTo(4))
	assert.Equal([]int{0, 5}, dfp.PathTo(5))
	assert.Equal([]int(nil), dfp.PathTo(6))
	assert.Equal([]int(nil), dfp.PathTo(7))
	assert.Equal([]int(nil), dfp.PathTo(8))
	assert.Equal([]int(nil), dfp.PathTo(9))
	assert.Equal([]int(nil), dfp.PathTo(10))
	assert.Equal([]int(nil), dfp.PathTo(11))
	assert.Equal([]int(nil), dfp.PathTo(12))

	assert.Panics(func() { NewDepthFirstDirectedPaths(g, 13) })
}
