package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestDijkstraAllPairsSP(t *testing.T) {
	assert := assert.New(t)

	tinyEWD := "8\n" +
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

	buf := strings.NewReader(tinyEWD)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	G := NewEdgeWeightedDigraphIn(in)

	spt := NewDijkstraAllPairsSP(G)

	path0to1 := []*DirectedEdge{
		NewDirectedEdge(0, 4, 0.38),
		NewDirectedEdge(4, 5, 0.35),
		NewDirectedEdge(5, 1, 0.32),
	}

	assert.Equal(path0to1, spt.Path(0, 1))
	assert.True(spt.HasPath(0, 1))
	assert.InEpsilon(1.05, spt.Dist(0, 1), 1e-9)
	assert.Panics(func() { spt.HasPath(9, 12) })

}
