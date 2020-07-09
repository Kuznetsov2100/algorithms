package digraph

import (
	"bufio"
	"math"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDijkstraSP(t *testing.T) {
	assert := assert.New(t)

	tinyEWD := "9\n" +
		"16\n" +
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
		"6 4 0.93\n" +
		"8 7 0.12\n"

	buf := strings.NewReader(tinyEWD)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	G := NewEdgeWeightedDigraphIn(in)

	sp := NewDijkstraSP(G, 5)
	assert.True(sp.HasPathTo(7))
	assert.Panics(func() { sp.HasPathTo(9) })
	pathto0 := []*DirectedEdge{
		NewDirectedEdge(5, 1, 0.32),
		NewDirectedEdge(1, 3, 0.29),
		NewDirectedEdge(3, 6, 0.52),
		NewDirectedEdge(6, 0, 0.58),
	}

	pathto1 := []*DirectedEdge{
		NewDirectedEdge(5, 1, 0.32),
	}
	pathto2 := []*DirectedEdge{
		NewDirectedEdge(5, 1, 0.32),
		NewDirectedEdge(1, 3, 0.29),
		NewDirectedEdge(3, 6, 0.52),
		NewDirectedEdge(6, 2, 0.40),
	}
	pathto3 := []*DirectedEdge{
		NewDirectedEdge(5, 1, 0.32),
		NewDirectedEdge(1, 3, 0.29),
	}
	pathto4 := []*DirectedEdge{
		NewDirectedEdge(5, 4, 0.35),
	}
	pathto5 := []*DirectedEdge(nil)
	pathto6 := []*DirectedEdge{
		NewDirectedEdge(5, 1, 0.32),
		NewDirectedEdge(1, 3, 0.29),
		NewDirectedEdge(3, 6, 0.52),
	}
	pathto7 := []*DirectedEdge{
		NewDirectedEdge(5, 7, 0.28),
	}
	pathto8 := []*DirectedEdge(nil)
	assert.Equal(pathto0, sp.PathTo(0))
	assert.InEpsilon(1.71, sp.DistTo(0), 1e-9)
	assert.Equal(pathto1, sp.PathTo(1))
	assert.InEpsilon(0.32, sp.DistTo(1), 1e-9)
	assert.Equal(pathto2, sp.PathTo(2))
	assert.InEpsilon(1.53, sp.DistTo(2), 1e-9)
	assert.Equal(pathto3, sp.PathTo(3))
	assert.InEpsilon(0.61, sp.DistTo(3), 1e-9)
	assert.Equal(pathto4, sp.PathTo(4))
	assert.InEpsilon(0.35, sp.DistTo(4), 1e-9)
	assert.Equal(pathto5, sp.PathTo(5))
	assert.Equal(0.00, sp.DistTo(5))
	assert.Equal(pathto6, sp.PathTo(6))
	assert.InEpsilon(1.13, sp.DistTo(6), 1e-9)
	assert.Equal(pathto7, sp.PathTo(7))
	assert.InEpsilon(0.28, sp.DistTo(7), 1e-9)
	assert.Equal(pathto8, sp.PathTo(8))
	assert.Equal(math.MaxFloat64, sp.DistTo(8))

	tinyEWD1 := "8\n" +
		"15\n" +
		"4 5 0.35\n" +
		"5 4 0.35\n" +
		"4 7 0.37\n" +
		"5 7 0.28\n" +
		"7 5 0.28\n" +
		"5 1 0.32\n" +
		"0 4 0.38\n" +
		"0 2 -0.28\n" +
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

	assert.PanicsWithValue("edge  0->2 -0.28  has negative weight\n", func() { NewDijkstraSP(G1, 0) })
}
