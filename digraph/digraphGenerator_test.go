package digraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigraphGenerator_Cycle(t *testing.T) {
	assert := assert.New(t)

	G := CycleDigraph(7)
	finder := NewDirectedCycle(G)
	assert.True(finder.HasCycle())
}

func TestDigraphGenerator_EulerianCycleDigraph(t *testing.T) {
	assert := assert.New(t)

	V, E := 5, 3
	g, err := EulerianCycleDigraph(V, E)
	assert.Nil(err)
	ec := NewDirectedEulerianCycle(g)
	assert.True(ec.HasEulerianCycle())

	// E < 0
	g1, err1 := EulerianCycleDigraph(5, -1)
	assert.Nil(g1)
	assert.EqualError(err1, "an Eulerian cycle must at least one edge")

	// V < 0
	g2, err2 := EulerianCycleDigraph(-3, 4)
	assert.Nil(g2)
	assert.EqualError(err2, "an Eulerian cycle must at least one vertex")
}

func TestDigraphGenerator_EulerianPathDigraph(t *testing.T) {
	assert := assert.New(t)

	V, E := 7, 4
	g, err := EulerianPathDigraph(V, E)
	assert.Nil(err)
	ec := NewDirectedEulerianPath(g)
	assert.True(ec.HasEulerianPath())

	// E < 0
	g1, err1 := EulerianPathDigraph(5, -1)
	assert.Nil(g1)
	assert.EqualError(err1, "negative number of edges")

	// V < 0
	g2, err2 := EulerianPathDigraph(-3, 4)
	assert.Nil(g2)
	assert.EqualError(err2, "an Eulerian path must have at least one vertex")

}

func TestDigraphGenerator_Path(t *testing.T) {
	assert := assert.New(t)

	G := PathDigraph(7)
	count := 0
	for v := 0; v < G.V(); v++ {
		if G.InDegree(v) == 0 {
			assert.Equal(1, G.OutDegree(v))
		}
		if G.OutDegree(v) == 0 {
			assert.Equal(1, G.InDegree(v))
		}
		if G.InDegree(v) == 1 && G.OutDegree(v) == 1 {
			count++
		}
	}
	assert.Equal(count, G.V()-2)
}

func TestDigraphGenerator_BinaryTree(t *testing.T) {
	assert := assert.New(t)

	V := 15
	g := BinaryTree(V)
	finder := NewDirectedCycle(g)
	assert.False(finder.HasCycle())
	assert.Equal(V/2*2, g.E())

}
