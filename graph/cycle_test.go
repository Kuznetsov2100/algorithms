package graph

import (
	"bufio"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestCycle(t *testing.T) {
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

	finder := NewCycle(g)
	assert.True(finder.HasCycle())
	assert.Equal([]int{3, 4, 5, 3}, finder.Cycles())

}

func TestCycle_NewCycle(t *testing.T) {
	assert := assert.New(t)
	tinyG := "1\n" + "1\n" + "0 0\n"
	buf := strings.NewReader(tinyG)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	in := &stdin.In{Scanner: scanner}
	g, err := NewGraphIn(in)
	assert.Nil(err)

	finder := NewCycle(g)
	assert.True(finder.HasCycle())
	assert.Equal([]int{0, 0}, finder.Cycles())

	tinyG1 := "2\n" + "2\n" + "0 1\n" + "1 0\n"
	buf1 := strings.NewReader(tinyG1)
	scanner1 := bufio.NewScanner(buf1)
	scanner1.Split(bufio.ScanWords)
	in1 := &stdin.In{Scanner: scanner1}
	g1, err1 := NewGraphIn(in1)
	assert.Nil(err1)

	finder1 := NewCycle(g1)
	assert.True(finder1.HasCycle())
	assert.Equal([]int{0, 1, 0}, finder1.Cycles())
}
