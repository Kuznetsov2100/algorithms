package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestEdgeWeightedDigraph(t *testing.T) {
	assert := assert.New(t)

	tinyEWG := "8\n" +
		"15\n" +
		"4 5 0.35\n" +
		"5 4 0.35\n" +
		"4 7 0.37\n" +
		"5 7 0.28\n" +
		"7 5 0.28\n" +
		"5 1 0.32\n" +
		"0 4 0.38\n" +
		"0 2 0.26\n" +
		"7 3 0.39\n" +
		"1 3 0.29\n" +
		"2 7 0.34\n" +
		"6 2 0.40\n" +
		"3 6 0.52\n" +
		"6 0 0.58\n" +
		"6 4 0.93\n"

	assert.Panics(func() { NewEdgeWeightedDigraphIn(nil) })
	buf := strings.NewReader(tinyEWG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	G := NewEdgeWeightedDigraphIn(in)

	assert.Equal(8, G.V())
	assert.Equal(15, G.E())

	toString := "8 15\n" +
		"0: 0->2  0.26  0->4  0.38  \n" +
		"1: 1->3  0.29  \n" +
		"2: 2->7  0.34  \n" +
		"3: 3->6  0.52  \n" +
		"4: 4->7  0.37  4->5  0.35  \n" +
		"5: 5->1  0.32  5->7  0.28  5->4  0.35  \n" +
		"6: 6->4  0.93  6->0  0.58  6->2  0.40  \n" +
		"7: 7->3  0.39  7->5  0.28  \n"
	assert.Equal(toString, G.String())

	assert.Equal(2, G.OutDegree(0))
	assert.Equal(1, G.Indegree(0))
	assert.Panics(func() { G.OutDegree(8) })

	edges := []*DirectedEdge{
		NewDirectedEdge(0, 2, 0.26),
		NewDirectedEdge(0, 4, 0.38),
	}
	assert.Equal(edges, G.Adj(0))

	tinyEWG1 := "8\n" +
		"-1\n"
	buf1 := strings.NewReader(tinyEWG1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	assert.PanicsWithValue("number of edges must be non negative", func() { NewEdgeWeightedDigraphIn(in1) })

	// NewEdgeWeightedGraphV
	assert.PanicsWithValue("number of vertices in a digraph must be non negative", func() { NewEdgeWeightedDigraphV(-1) })
	G2 := NewEdgeWeightedGraphV(3)
	assert.Equal(3, G2.V())
	assert.Equal(0, G2.E())

	//NewEdgeWeightedGraphVE
	assert.PanicsWithValue("number of edges in a digraph must be non negative", func() { NewEdgeWeightedDigraphVE(3, -1) })
	G3 := NewEdgeWeightedDigraphVE(3, 3)
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
	G4 := NewEdgeWeightedDigraphIn(in2)
	edges2 := []*DirectedEdge{
		NewDirectedEdge(0, 1, 0.37),
		NewDirectedEdge(0, 0, 0.35),
	}
	assert.Equal(edges2, G4.Edges())

	tinyEWG3 := "-2\n" +
		"2\n" +
		"0 0 0.35\n" +
		"0 1 0.37\n"
	buf3 := strings.NewReader(tinyEWG3)
	scanner3 := bufio.NewScanner(buf3)
	scanner3.Split(bufio.ScanWords)
	in3 := &stdin.In{Scanner: scanner3}
	assert.PanicsWithValue("number of vertices in a digraph must be non negative", func() { NewEdgeWeightedDigraphIn(in3) })
}
