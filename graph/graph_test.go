package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestGraph_NewGraph(t *testing.T) {
	assert := assert.New(t)
	assert.Panics(func() { NewGraph(-1) })
}

func TestGraph_NewGraphIn(t *testing.T) {
	assert := assert.New(t)
	buf := strings.NewReader("2\n" + "1\n" + "0 1\n")
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: s}
	g, err := NewGraphIn(in)
	assert.Nil(err)

	assert.Equal(1, g.E())
	assert.Equal(2, g.V())

	g1, err1 := NewGraphIn(nil)
	assert.Nil(g1)
	assert.Error(err1)

	buf1 := strings.NewReader("-1\n" + "1\n" + "0 1\n")
	s1 := bufio.NewScanner(buf1)
	s1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: s1}
	g2, err2 := NewGraphIn(in1)
	assert.Nil(g2)
	assert.Error(err2)

	buf2 := strings.NewReader("2\n" + "-2\n" + "0 1\n")
	s2 := bufio.NewScanner(buf2)
	s2.Split(bufio.ScanWords)
	in2 := &stdin.In{Scanner: s2}
	g3, err3 := NewGraphIn(in2)
	assert.Nil(g3)
	assert.Error(err3)
}

func TestGraph_Degree(t *testing.T) {
	assert := assert.New(t)
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)

	assert.Equal(2, g.Degree(0))
	assert.Equal(2, g.E())
	assert.Equal(3, g.V())

	assert.Panics(func() { g.AddEdge(0, 4) })
}

func TestGraph_Adj(t *testing.T) {
	assert := assert.New(t)
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	assert.Equal([]int{2, 1}, g.Adj(0))
}

func TestGraph_String(t *testing.T) {
	assert := assert.New(t)
	g := NewGraph(3)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	assert.Equal("3 vertices, 2 edges \n"+"0:  2 1\n"+"1:  0\n"+"2:  0\n", g.String())
}
