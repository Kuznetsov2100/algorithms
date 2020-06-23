package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestLazyPrimMST(t *testing.T) {
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

	buf := strings.NewReader(tinyEWG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	G := NewEdgeWeightedGraphIn(in)

	mst := NewLazyPrimMST(G)
	assert.Equal(1.81000, mst.Weight())

	edges := []*Edge{
		NewEdge(0, 7, 0.16000),
		NewEdge(1, 7, 0.19000),
		NewEdge(0, 2, 0.26000),
		NewEdge(2, 3, 0.17000),
		NewEdge(5, 7, 0.28000),
		NewEdge(4, 5, 0.35000),
		NewEdge(6, 2, 0.40000),
	}
	assert.Equal(edges, mst.Edges())
}