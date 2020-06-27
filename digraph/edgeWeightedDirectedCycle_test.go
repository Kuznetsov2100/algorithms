package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestEdgeWeightedDirectedCycle(t *testing.T) {
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

	finder := NewEdgeWeightedDirectedCycle(G)

	assert.True(finder.HasCycle())
	cycle := []*DirectedEdge{
		NewDirectedEdge(7, 3, 0.39),
		NewDirectedEdge(3, 6, 0.52),
		NewDirectedEdge(6, 4, 0.93),
		NewDirectedEdge(4, 7, 0.37),
	}
	assert.Equal(cycle, finder.GetCycle())

	tinyEWD1 := "2\n" +
		"1\n" +
		"0 1 0.35\n"

	buf1 := strings.NewReader(tinyEWD1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	G1 := NewEdgeWeightedDigraphIn(in1)

	finder1 := NewEdgeWeightedDirectedCycle(G1)
	assert.False(finder1.HasCycle())
	assert.Equal([]*DirectedEdge(nil), finder1.GetCycle())
}
