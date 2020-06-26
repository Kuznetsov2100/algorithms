package digraph

import (
	"bufio"
	"math"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestAcyclicLP(t *testing.T) {
	assert := assert.New(t)

	tinyEWDAG := "8\n" +
		"13\n" +
		"5 4 0.35\n" +
		"4 7 0.37\n" +
		"5 7 0.28\n" +
		"5 1 0.32\n" +
		"4 0 0.38\n" +
		"0 2 0.26\n" +
		"3 7 0.39\n" +
		"1 3 0.29\n" +
		"7 2 0.34\n" +
		"6 2 0.40\n" +
		"3 6 0.52\n" +
		"6 0 0.58\n" +
		"6 4 0.93\n"

	buf := strings.NewReader(tinyEWDAG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	G := NewEdgeWeightedDigraphIn(in)

	lp, err := NewAcyclicLP(G, 1)

	assert.Nil(err)
	assert.True(lp.HasPathTo(0))
	assert.Panics(func() { lp.HasPathTo(9) })

	pathto0 := []*DirectedEdge{
		NewDirectedEdge(1, 3, 0.29),
		NewDirectedEdge(3, 6, 0.52),
		NewDirectedEdge(6, 4, 0.93),
		NewDirectedEdge(4, 0, 0.38),
	}
	pathto1 := []*DirectedEdge(nil)
	pathto2 := []*DirectedEdge{
		NewDirectedEdge(1, 3, 0.29),
		NewDirectedEdge(3, 6, 0.52),
		NewDirectedEdge(6, 4, 0.93),
		NewDirectedEdge(4, 7, 0.37),
		NewDirectedEdge(7, 2, 0.34),
	}
	pathto3 := []*DirectedEdge{
		NewDirectedEdge(1, 3, 0.29),
	}
	pathto4 := []*DirectedEdge{
		NewDirectedEdge(1, 3, 0.29),
		NewDirectedEdge(3, 6, 0.52),
		NewDirectedEdge(6, 4, 0.93),
	}
	pathto5 := []*DirectedEdge(nil)
	pathto6 := []*DirectedEdge{
		NewDirectedEdge(1, 3, 0.29),
		NewDirectedEdge(3, 6, 0.52),
	}
	pathto7 := []*DirectedEdge{
		NewDirectedEdge(1, 3, 0.29),
		NewDirectedEdge(3, 6, 0.52),
		NewDirectedEdge(6, 4, 0.93),
		NewDirectedEdge(4, 7, 0.37),
	}
	assert.Equal(pathto0, lp.PathTo(0))
	assert.InEpsilon(2.12, lp.DistTo(0), 1e-9)
	assert.Equal(pathto1, lp.PathTo(1))
	assert.Equal(0.00, lp.DistTo(1))
	assert.Equal(pathto2, lp.PathTo(2))
	assert.InEpsilon(2.45, lp.DistTo(2), 1e-9)
	assert.Equal(pathto3, lp.PathTo(3))
	assert.InEpsilon(0.29, lp.DistTo(3), 1e-9)
	assert.Equal(pathto4, lp.PathTo(4))
	assert.InEpsilon(1.74, lp.DistTo(4), 1e-9)
	assert.Equal(pathto5, lp.PathTo(5))
	assert.Equal(math.Inf(-1), lp.DistTo(5))
	assert.Equal(pathto6, lp.PathTo(6))
	assert.InEpsilon(0.81, lp.DistTo(6), 1e-9)
	assert.Equal(pathto7, lp.PathTo(7))
	assert.InEpsilon(2.11, lp.DistTo(7), 1e-9)

	tinyEWD1 := "8\n" +
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

	buf1 := strings.NewReader(tinyEWD1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	G1 := NewEdgeWeightedDigraphIn(in1)

	sp1, err1 := NewAcyclicLP(G1, 0)

	assert.Nil(sp1)
	assert.EqualError(err1, "digraph is not acyclic")
}
