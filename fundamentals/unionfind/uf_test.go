package unionfind

import (
	"bufio"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	assert := assert.New(t)

	tinyUF := "4 3\n" +
		"3 8\n" +
		"6 5\n" +
		"9 4\n" +
		"2 1\n" +
		"8 9\n" +
		"5 0\n" +
		"7 2\n" +
		"6 1\n" +
		"1 0\n" +
		"6 7\n"
	buf := strings.NewReader(tinyUF)
	scanner := bufio.NewScanner(buf)

	uf := NewUF(10)
	var listA [][]int
	for scanner.Scan() {
		numbers := strings.Split(scanner.Text(), " ")
		p, _ := strconv.Atoi(numbers[0])
		q, _ := strconv.Atoi(numbers[1])
		if uf.Find(p) == uf.Find(q) {
			continue
		}
		uf.Union(p, q)
		listA = append(listA, []int{p, q})
	}
	listB := [][]int{
		{4, 3},
		{3, 8},
		{6, 5},
		{9, 4},
		{2, 1},
		{5, 0},
		{7, 2},
		{6, 1},
	}
	assert.Equal(listA, listB)
	assert.Equal(2, uf.Count())
	assert.Panics(func() { uf.Find(11) })
	assert.Panics(func() { NewUF(-1) })

	uf.Union(8, 9)
	assert.Equal(2, uf.Count())
}
