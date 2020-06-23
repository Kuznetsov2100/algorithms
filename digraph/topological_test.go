package digraph

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/handane123/algorithms/stdin"
	"github.com/stretchr/testify/assert"
)

func TestToplogical(t *testing.T) {
	assert := assert.New(t)

	jobs := "Algorithms/Theoretical CS/Databases/Scientific Computing\n" +
		"Introduction to CS/Advanced Programming/Algorithms\n" +
		"Advanced Programming/Scientific Computing\n" +
		"Scientific Computing/Computational Biology\n" +
		"Theoretical CS/Computational Biology/Artificial Intelligence\n" +
		"Linear Algebra/Theoretical CS\n" +
		"Calculus/Linear Algebra\n" +
		"Artificial Intelligence/Neural Networks/Robotics/Machine Learning\n" +
		"Machine Learning/Neural Networks\n"

	content := []byte(jobs)
	tmpfile, err := ioutil.TempFile("", "jobs.*.txt")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		t.Error(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Error(err)
	}

	sg := NewSymbolDigraph(tmpfile.Name(), "/")
	order := []string{
		"Calculus",
		"Linear Algebra",
		"Introduction to CS",
		"Advanced Programming",
		"Algorithms",
		"Theoretical CS",
		"Artificial Intelligence",
		"Robotics",
		"Machine Learning",
		"Neural Networks",
		"Databases",
		"Scientific Computing",
		"Computational Biology",
	}
	rank := []int{4, 5, 10, 11, 2, 3, 12, 6, 1, 0, 9, 7, 8}
	topological := NewTopological(sg.Digraph())
	assert.True(topological.HasOrder())
	for i, v := range topological.Order() {
		assert.Equal(order[i], sg.NameOf(v))
		assert.Equal(rank[i], topological.Rank(i))
	}
}

func TestToplogical_Rank(t *testing.T) {
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

	top := NewTopological(g)
	assert.False(top.HasOrder())
	assert.Equal(-1, top.Rank(0))
	assert.Panics(func() { top.Rank(13) })
}
