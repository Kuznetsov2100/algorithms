package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDirectedEulerianPath(t *testing.T) {
	assert := assert.New(t)

	graphdata1 := "3\n" +
		"3\n" +
		"0 1\n" +
		"1 2\n" +
		"2 0\n"
	buf1 := strings.NewReader(graphdata1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	g1, err1 := NewDigraphIn(in1)
	assert.Nil(err1)
	ec1 := NewDirectedEulerianPath(g1)
	assert.True(ec1.HasEulerianPath())
	assert.Equal([]int{0, 1, 2, 0}, ec1.Path())

	// digraph with zero edges has a degenerate Eulerian path []int{0}
	graphdata3 := "1\n" +
		"0\n"
	buf3 := strings.NewReader(graphdata3)
	scanner3 := bufio.NewScanner(buf3)
	scanner3.Split(bufio.ScanWords)
	in3 := &stdin.In{Scanner: scanner3}
	g3, err3 := NewDigraphIn(in3)
	assert.Nil(err3)
	ec3 := NewDirectedEulerianPath(g3)
	assert.True(ec3.HasEulerianPath())
	assert.Equal([]int{0}, ec3.Path())

	// deficit > 1
	graphdata4 := "5\n" +
		"6\n" +
		"1 2\n" +
		"1 0\n" +
		"2 0\n" +
		"0 3\n" +
		"1 3\n" +
		"3 4\n"
	buf4 := strings.NewReader(graphdata4)
	scanner4 := bufio.NewScanner(buf4)
	scanner4.Split(bufio.ScanWords)
	in4 := &stdin.In{Scanner: scanner4}
	g4, err4 := NewDigraphIn(in4)
	assert.Nil(err4)
	ec4 := NewDirectedEulerianPath(g4)
	assert.False(ec4.HasEulerianPath())
	assert.Equal([]int(nil), ec4.Path())
}
