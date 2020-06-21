package digraph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestKosarajuSharirSCC(t *testing.T) {
	assert := assert.New(t)
	tinyDG := "13\n" +
		"22\n" +
		"4  2\n" +
		"2  3\n" +
		"3  2\n" +
		"6  0\n" +
		"0  1\n" +
		"2  0\n" +
		"11 12\n" +
		"12  9\n" +
		"9 10\n" +
		"9 11\n" +
		"7  9\n" +
		"10 12\n" +
		"11  4\n" +
		"4  3\n" +
		"3  5\n" +
		"6  8\n" +
		"8  6\n" +
		"5  4\n" +
		"0  5\n" +
		"6  4\n" +
		"6  9\n" +
		"7  6\n"

	buf := strings.NewReader(tinyDG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	g, err := NewDigraphIn(in)
	assert.Nil(err)

	cc := NewKosarajuSharirSCC(g)

	m := cc.Count()
	components := make([]*arrayqueue.Queue, m)
	for i := 0; i < m; i++ {
		components[i] = arrayqueue.New()
	}

	for v := 0; v < g.V(); v++ {
		components[cc.Id(v)].Enqueue(v)
	}

	assert.Equal([]interface{}{1}, components[0].Values())
	assert.Equal([]interface{}{0, 2, 3, 4, 5}, components[1].Values())
	assert.Equal([]interface{}{9, 10, 11, 12}, components[2].Values())
	assert.Equal([]interface{}{6, 8}, components[3].Values())
	assert.Equal([]interface{}{7}, components[4].Values())

	assert.True(cc.StronglyConnected(0, 2))
	assert.Panics(func() { cc.Id(13) })
}
