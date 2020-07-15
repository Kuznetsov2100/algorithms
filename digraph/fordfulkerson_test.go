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

	// tinyfn1 := "6\n" +
	// 	"8\n" +
	// 	"0 1 2.0\n" +
	// 	"0 2 3.0\n" +
	// 	"1 3 3.0\n" +
	// 	"1 4 1.0\n" +
	// 	"2 3 1.0\n" +
	// 	"2 4 1.0\n" +
	// 	"3 5 2.0\n" +
	// 	"4 5 3.0\n"
	// buf1 := strings.NewReader(tinyfn1)
	// s1 := bufio.NewScanner(buf1)
	// s1.Split(bufio.ScanWords)
	// in1 := &stdin.In{Scanner: s1}
	// G1 := NewFlowNetworkIn(in1)
}
