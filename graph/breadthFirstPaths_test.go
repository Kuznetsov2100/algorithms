package graph

import (
	"bufio"
	"math"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestBreadthFirstPaths(t *testing.T) {
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
	bfp := NewBreadthFirstPaths(g, 0)
	assert.Equal([]int{0}, bfp.PathTo(0))
	assert.Equal(0, bfp.DistTo(0))
	assert.Equal([]int{0, 1}, bfp.PathTo(1))
	assert.Equal(1, bfp.DistTo(1))
	assert.Equal([]int{0, 2}, bfp.PathTo(2))
	assert.Equal(1, bfp.DistTo(2))
	assert.Equal([]int{0, 2, 3}, bfp.PathTo(3))
	assert.Equal(2, bfp.DistTo(3))
	assert.Equal([]int{0, 2, 4}, bfp.PathTo(4))
	assert.Equal(2, bfp.DistTo(4))
	assert.Equal([]int{0, 5}, bfp.PathTo(5))
	assert.Equal(1, bfp.DistTo(5))
	assert.Equal([]int(nil), bfp.PathTo(6))
	assert.Equal(math.MaxInt64, bfp.DistTo(6))
	assert.Equal([]int(nil), bfp.PathTo(7))
	assert.Equal(math.MaxInt64, bfp.DistTo(7))

	assert.Panics(func() { NewBreadthFirstPaths(g, 8) })
}
