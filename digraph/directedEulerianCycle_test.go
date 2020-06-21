package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestEulerianCycle(t *testing.T) {
	assert := assert.New(t)

	graphdata1 := "3\n" +
		"3\n" +
		"0 1\n" +
		"1 2\n" +
		"2 0\n"
	buf1 := strings.NewReader(graphdata1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	g1, err1 := NewDigraphIn(in1)
	assert.Nil(err1)
	ec1 := NewDirectedEulerianCycle(g1)
	assert.True(ec1.HasEulerianCycle())
	assert.Equal([]int{0, 1, 2, 0}, ec1.GetCycle())

	// graph with zero edge
	G2 := NewDigraph(3)
	ec2 := NewDirectedEulerianCycle(G2)
	assert.False(ec2.HasEulerianCycle())

	// indegree != outdegree
	graphdata3 := "3\n" +
		"2\n" +
		"0 1\n" +
		"0 2\n"
	buf3 := strings.NewReader(graphdata3)
	scanner3 := bufio.NewScanner(buf3)
	scanner3.Split(bufio.ScanWords)
	in3 := &stdin.In{Scanner: scanner3}
	g3, err3 := NewDigraphIn(in3)
	assert.Nil(err3)
	ec3 := NewDirectedEulerianCycle(g3)
	assert.False(ec3.HasEulerianCycle())
}
