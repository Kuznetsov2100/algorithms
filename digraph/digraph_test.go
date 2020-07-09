package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDigraph_NewDigraph(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() { NewDigraph(-1) })
}

func TestDigraph_NewDigraphIn(t *testing.T) {
	assert := assert.New(t)
	buf := strings.NewReader("2\n" + "1\n" + "0 1\n")
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: s}
	g, err := NewDigraphIn(in)
	assert.Nil(err)

	assert.Equal(1, g.E())
	assert.Equal(2, g.V())

	g1, err1 := NewDigraphIn(nil)
	assert.Nil(g1)
	assert.Error(err1)

	buf1 := strings.NewReader("-1\n" + "1\n" + "0 1\n")
	s1 := bufio.NewScanner(buf1)
	s1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: s1}
	g2, err2 := NewDigraphIn(in1)
	assert.Nil(g2)
	assert.Error(err2)

	buf2 := strings.NewReader("2\n" + "-2\n" + "0 1\n")
	s2 := bufio.NewScanner(buf2)
	s2.Split(bufio.ScanWords)
	in2 := &stdin.In{Scanner: s2}
	g3, err3 := NewDigraphIn(in2)
	assert.Nil(g3)
	assert.Error(err3)
}

func TestDigraph_OutDegree(t *testing.T) {
	assert := assert.New(t)
	g := NewDigraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)

	assert.Equal(2, g.OutDegree(0))
	assert.Panics(func() { g.OutDegree(3) })
}

func TestDigraph_InDegree(t *testing.T) {
	assert := assert.New(t)
	g := NewDigraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)

	assert.Equal(1, g.InDegree(1))
	assert.Equal(1, g.InDegree(2))
}

func TestDigraph_Reverse(t *testing.T) {
	assert := assert.New(t)
	g := NewDigraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)

	reverse := g.Reverse()
	assert.Equal(2, reverse.InDegree(0))
	assert.Equal(1, reverse.OutDegree(1))
	assert.Equal(1, reverse.OutDegree(2))
}

func TestDigraph_Adj(t *testing.T) {
	assert := assert.New(t)
	g := NewDigraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	assert.Equal([]int{2, 1}, g.Adj(0))
}

func TestDigraph_String(t *testing.T) {
	assert := assert.New(t)
	g := NewDigraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	assert.Equal("3 vertices, 2 edges \n"+"0:  2 1\n"+"1: \n"+"2: \n", g.String())
}
