package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestCC(t *testing.T) {
	assert := assert.New(t)
	tinyG := "13\n" +
		"13\n" +
		"0 5\n" +
		"4 3\n" +
		"0 1\n" +
		"9 12\n" +
		"6 4\n" +
		"5 4\n" +
		"0 2\n" +
		"11 12\n" +
		"9 10\n" +
		"0 6\n" +
		"7 8\n" +
		"9 11\n" +
		"5 3\n"
	buf := strings.NewReader(tinyG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	g, err := NewGraphIn(in)
	assert.Nil(err)

	cc := NewCC(g)

	m := cc.Count()
	components := make([]*arrayqueue.Queue, m)
	for i := 0; i < m; i++ {
		components[i] = arrayqueue.New()
	}

	for v := 0; v < g.V(); v++ {
		components[cc.Id(v)].Enqueue(v)
	}

	assert.Equal([]interface{}{0, 1, 2, 3, 4, 5, 6}, components[0].Values())
	assert.Equal([]interface{}{7, 8}, components[1].Values())
	assert.Equal([]interface{}{9, 10, 11, 12}, components[2].Values())
	assert.Equal(7, cc.Size(0))
	assert.True(cc.Connected(0, 1))

	assert.Panics(func() { cc.Size(13) })
}
