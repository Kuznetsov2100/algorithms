package digraph

import (
	"bufio"
	"math"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstDirectedPaths(t *testing.T) {
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
	bfp := NewBreadthFirstDirectedPaths(g, 3)
	assert.Equal([]int{3, 2, 0}, bfp.PathTo(0))
	assert.Equal(2, bfp.DistTo(0))
	assert.Equal([]int{3, 2, 0, 1}, bfp.PathTo(1))
	assert.Equal(3, bfp.DistTo(1))
	assert.Equal([]int{3, 2}, bfp.PathTo(2))
	assert.Equal(1, bfp.DistTo(2))
	assert.Equal([]int{3}, bfp.PathTo(3))
	assert.Equal(0, bfp.DistTo(3))
	assert.Equal([]int{3, 5, 4}, bfp.PathTo(4))
	assert.Equal(2, bfp.DistTo(4))
	assert.Equal([]int{3, 5}, bfp.PathTo(5))
	assert.Equal(1, bfp.DistTo(5))
	assert.Equal([]int(nil), bfp.PathTo(6))
	assert.Equal(math.MaxInt64, bfp.DistTo(6))
	assert.Equal([]int(nil), bfp.PathTo(7))
	assert.Equal(math.MaxInt64, bfp.DistTo(7))
	assert.Equal([]int(nil), bfp.PathTo(8))
	assert.Equal(math.MaxInt64, bfp.DistTo(8))
	assert.Equal([]int(nil), bfp.PathTo(9))
	assert.Equal(math.MaxInt64, bfp.DistTo(9))
	assert.Equal([]int(nil), bfp.PathTo(10))
	assert.Equal(math.MaxInt64, bfp.DistTo(10))
	assert.Equal([]int(nil), bfp.PathTo(11))
	assert.Equal(math.MaxInt64, bfp.DistTo(11))
	assert.Equal([]int(nil), bfp.PathTo(12))
	assert.Equal(math.MaxInt64, bfp.DistTo(12))

	assert.Panics(func() { NewBreadthFirstDirectedPaths(g, 13) })

	bfp1 := NewBreadthFirstDirectedPathSources(g, []int{0, 1})
	assert.Equal([]int{0}, bfp1.PathTo(0))
	assert.Equal(0, bfp1.DistTo(0))
	assert.Equal([]int{1}, bfp1.PathTo(1))
	assert.Equal(0, bfp1.DistTo(1))
	assert.Equal([]int{0, 5, 4, 2}, bfp1.PathTo(2))
	assert.Equal(3, bfp1.DistTo(2))
	assert.Equal([]int{0, 5, 4, 3}, bfp1.PathTo(3))
	assert.Equal(3, bfp1.DistTo(3))
	assert.Equal([]int{0, 5, 4}, bfp1.PathTo(4))
	assert.Equal(2, bfp1.DistTo(4))
	assert.Equal([]int{0, 5}, bfp1.PathTo(5))
	assert.Equal(1, bfp1.DistTo(5))
	assert.Equal([]int(nil), bfp1.PathTo(6))
	assert.Equal(math.MaxInt64, bfp1.DistTo(6))
	assert.Equal([]int(nil), bfp1.PathTo(7))
	assert.Equal(math.MaxInt64, bfp1.DistTo(7))
	assert.Equal([]int(nil), bfp1.PathTo(8))
	assert.Equal(math.MaxInt64, bfp1.DistTo(8))
	assert.Equal([]int(nil), bfp1.PathTo(9))
	assert.Equal(math.MaxInt64, bfp1.DistTo(9))
	assert.Equal([]int(nil), bfp1.PathTo(10))
	assert.Equal(math.MaxInt64, bfp1.DistTo(10))
	assert.Equal([]int(nil), bfp1.PathTo(11))
	assert.Equal(math.MaxInt64, bfp1.DistTo(11))
	assert.Equal([]int(nil), bfp1.PathTo(12))
	assert.Equal(math.MaxInt64, bfp1.DistTo(12))

	assert.Panics(func() { NewBreadthFirstDirectedPathSources(g, nil) })
}
