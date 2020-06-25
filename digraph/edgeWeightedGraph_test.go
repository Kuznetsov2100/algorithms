package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestEdgeWeightedGraph(t *testing.T) {
	assert := assert.New(t)

	tinyEWG := "8\n" +
		"16\n" +
		"4 5 0.35\n" +
		"4 7 0.37\n" +
		"5 7 0.28\n" +
		"0 7 0.16\n" +
		"1 5 0.32\n" +
		"0 4 0.38\n" +
		"2 3 0.17\n" +
		"1 7 0.19\n" +
		"0 2 0.26\n" +
		"1 2 0.36\n" +
		"1 3 0.29\n" +
		"2 7 0.34\n" +
		"6 2 0.40\n" +
		"3 6 0.52\n" +
		"6 0 0.58\n" +
		"6 4 0.93\n"

	assert.Panics(func() { NewEdgeWeightedGraphIn(nil) })
	buf := strings.NewReader(tinyEWG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	G := NewEdgeWeightedGraphIn(in)

	assert.Equal(8, G.V())
	assert.Equal(16, G.E())

	toString := "8 16\n" +
		"0: 6-0 0.58000  0-2 0.26000  0-4 0.38000  0-7 0.16000  \n" +
		"1: 1-3 0.29000  1-2 0.36000  1-7 0.19000  1-5 0.32000  \n" +
		"2: 6-2 0.40000  2-7 0.34000  1-2 0.36000  0-2 0.26000  2-3 0.17000  \n" +
		"3: 3-6 0.52000  1-3 0.29000  2-3 0.17000  \n" +
		"4: 6-4 0.93000  0-4 0.38000  4-7 0.37000  4-5 0.35000  \n" +
		"5: 1-5 0.32000  5-7 0.28000  4-5 0.35000  \n" +
		"6: 6-4 0.93000  6-0 0.58000  3-6 0.52000  6-2 0.40000  \n" +
		"7: 2-7 0.34000  1-7 0.19000  0-7 0.16000  5-7 0.28000  4-7 0.37000  \n"
	assert.Equal(toString, G.String())

	assert.Equal(4, G.Degree(0))
	assert.Panics(func() { G.Degree(8) })

	edges := []*Edge{
		NewEdge(6, 0, 0.58),
		NewEdge(0, 2, 0.26),
		NewEdge(0, 4, 0.38),
		NewEdge(0, 7, 0.16),
	}
	assert.Equal(edges, G.Adj(0))

	tinyEWG1 := "8\n" +
		"-1\n"
	buf1 := strings.NewReader(tinyEWG1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	assert.PanicsWithValue("number of edges must be non negative", func() { NewEdgeWeightedGraphIn(in1) })

	// NewEdgeWeightedDigraphIn V < 0
	tinyEWG3 := "-2\n" +
		"2\n" +
		"0 0 0.35\n" +
		"0 1 0.37\n"
	buf3 := strings.NewReader(tinyEWG3)
	scanner3 := bufio.NewScanner(buf3)
	scanner3.Split(bufio.ScanWords)
	in3 := &stdin.In{Scanner: scanner3}
	assert.PanicsWithValue("number of vertices must be non negative",
		func() { NewEdgeWeightedGraphIn(in3) })

	// NewEdgeWeightedGraphV
	assert.PanicsWithValue("Number of vertices must be non negative", func() { NewEdgeWeightedGraphV(-1) })
	G2 := NewEdgeWeightedGraphV(3)
	assert.Equal(3, G2.V())
	assert.Equal(0, G2.E())

	//NewEdgeWeightedGraphVE
	assert.PanicsWithValue("Number of edges must be non negative", func() { NewEdgeWeightedGraphVE(3, -1) })
	G3 := NewEdgeWeightedGraphVE(3, 3)
	assert.Equal(3, G3.V())
	assert.Equal(3, G3.E())

	tinyEWG2 := "2\n" +
		"2\n" +
		"0 0 0.35\n" +
		"0 1 0.37\n"
	buf2 := strings.NewReader(tinyEWG2)
	scanner2 := bufio.NewScanner(buf2)
	scanner2.Split(bufio.ScanWords)
	in2 := &stdin.In{Scanner: scanner2}
	G4 := NewEdgeWeightedGraphIn(in2)
	edges2 := []*Edge{
		NewEdge(0, 1, 0.37),
		NewEdge(0, 0, 0.35),
	}
	assert.Equal(edges2, G4.Edges())
}
