package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/io/stdin"
	"github.com/stretchr/testify/assert"
)

func TestFordFulkerson(t *testing.T) {
	assert := assert.New(t)

	tinyfn := "6\n" +
		"8\n" +
		"0 1 2.0\n" +
		"0 2 3.0\n" +
		"1 3 3.0\n" +
		"1 4 1.0\n" +
		"2 3 1.0\n" +
		"2 4 1.0\n" +
		"3 5 2.0\n" +
		"4 5 3.0\n"
	buf := strings.NewReader(tinyfn)
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: s}
	G := NewFlowNetworkIn(in)

	maxflow := NewFordFulkerson(G, 0, G.V()-1)
	assert.True(maxflow.InCut(0))
	assert.True(maxflow.InCut(2))
	assert.Equal(4.0, maxflow.Value())
	assert.Panics(func() { maxflow.InCut(9) })
	assert.PanicsWithValue("source equals sink", func() { NewFordFulkerson(G, G.V()-1, G.V()-1) })

	// excess at source
	tinyfn1 := "6\n" +
		"8\n" +
		"0 1 2.0 1.9\n" +
		"0 2 3.0 1.0\n" +
		"1 3 3.0 0.0\n" +
		"1 4 1.0 0.0\n" +
		"2 3 1.0 0.0\n" +
		"2 4 1.0 0.0\n" +
		"3 5 2.0 0.0\n" +
		"4 5 3.0 0.0\n"
	buf1 := strings.NewReader(tinyfn1)
	s1 := bufio.NewScanner(buf1)
	s1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: s1}
	G1 := NewFlowNetworkInF(in1)
	assert.PanicsWithValue("initial flow is infeasible", func() { NewFordFulkerson(G1, 0, G1.V()-1) })

	// excess at sink
	tinyfn2 := "6\n" +
		"8\n" +
		"0 1 2.0 0.0\n" +
		"0 2 3.0 0.0\n" +
		"1 3 3.0 0.0\n" +
		"1 4 1.0 0.0\n" +
		"2 3 1.0 0.0\n" +
		"2 4 1.0 0.0\n" +
		"3 5 2.0 1.0\n" +
		"4 5 3.0 0.8\n"
	buf2 := strings.NewReader(tinyfn2)
	s2 := bufio.NewScanner(buf2)
	s2.Split(bufio.ScanWords)
	in2 := &stdin.In{Scanner: s2}
	G2 := NewFlowNetworkInF(in2)
	assert.PanicsWithValue("initial flow is infeasible", func() { NewFordFulkerson(G2, 0, G2.V()-1) })

	// Net flow out of vertex 1 doesn't equal zero
	tinyfn3 := "6\n" +
		"8\n" +
		"0 1 2.0 0.0\n" +
		"0 2 3.0 0.0\n" +
		"1 3 3.0 0.1\n" +
		"1 4 1.0 0.8\n" +
		"2 3 1.0 0.0\n" +
		"2 4 1.0 0.0\n" +
		"3 5 2.0 0.0\n" +
		"4 5 3.0 0.0\n"
	buf3 := strings.NewReader(tinyfn3)
	s3 := bufio.NewScanner(buf3)
	s3.Split(bufio.ScanWords)
	in3 := &stdin.In{Scanner: s3}
	G3 := NewFlowNetworkInF(in3)
	assert.PanicsWithValue("initial flow is infeasible", func() { NewFordFulkerson(G3, 0, G3.V()-1) })
}
