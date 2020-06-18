package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphGenerator_Simple(t *testing.T) {
	assert := assert.New(t)
	g, err := Simple(3, 4)
	assert.Nil(g)
	assert.Error(err)

	g1, err1 := Simple(3, -1)
	assert.Nil(g1)
	assert.Error(err1)

	V := 20
	E := 30
	g2, err2 := Simple(V, E)
	assert.Nil(err2)

	// a simple graph has no self-loop and no parallel edges.
	finder := NewCycle(g2)
	assert.False(finder.hasParallelEdges(g2))
	assert.False(finder.hasSelfLoop(g2))
}

func TestGraphGenerator_SimpleP(t *testing.T) {
	assert := assert.New(t)
	g, err := SimpleP(3, 1.2)
	assert.Nil(g)
	assert.Error(err)

	V := 5
	p := 0.769283
	g1, err1 := SimpleP(V, p)
	assert.Nil(err1)

	finder := NewCycle(g1)
	assert.False(finder.hasParallelEdges(g1))
	assert.False(finder.hasSelfLoop(g1))
}

func TestGraphGenerator_Complete(t *testing.T) {
	assert := assert.New(t)
	V := 3
	g := Complete(V)
	for i := 0; i < V; i++ {
		assert.Equal(2, g.Degree(i))
	}
}

func TestGraphGenerator_Bipartite(t *testing.T) {
	assert := assert.New(t)
	V1, V2, E := 10, 5, 8
	G, err := BipartiteGraph(V1, V2, E)
	assert.Nil(err)
	b := NewBipartite(G)
	assert.True(b.IsBipartite())

	// too many edges
	G1, err1 := BipartiteGraph(7, 3, 30)
	assert.Nil(G1)
	assert.EqualError(err1, "too many edges")

	// too few edges
	G2, err2 := BipartiteGraph(7, 3, -1)
	assert.Nil(G2)
	assert.EqualError(err2, "too few edges")
}

func TestGraphGenerator_BipartiteP(t *testing.T) {
	assert := assert.New(t)
	V1, V2, p := 10, 5, 0.5732
	G, err := BipartiteP(V1, V2, p)
	assert.Nil(err)
	b := NewBipartite(G)
	assert.True(b.IsBipartite())

	// invalid probability
	G1, err1 := BipartiteP(10, 5, 1.2)
	assert.Nil(G1)
	assert.EqualError(err1, "probability must be between 0 and 1")
}

func TestGraphGenerator_CompleteBipartite(t *testing.T) {
	assert := assert.New(t)

	V1, V2 := 5, 3
	G := CompleteBipartite(V1, V2)

	b := NewBipartite(G)
	assert.True(b.IsBipartite())
	assert.Equal(V1+V2, G.V())
	assert.Equal(V1*V2, G.E())
}

func TestGraphGenerator_CycleGraph(t *testing.T) {
	assert := assert.New(t)

	V := 7
	g := CycleGraph(V)
	finder := NewCycle(g)
	assert.True(finder.HasCycle())
}

func TestGraphGenerator_Path(t *testing.T) {
	assert := assert.New(t)

	V := 3
	g := PathGraph(V)

	degree1, degree2 := 0, 0
	for v := 0; v < g.V(); v++ {
		if g.Degree(v) == 1 {
			degree1++
		} else if g.Degree(v) == 2 {
			degree2++
		}
	}
	assert.Equal(2, degree1)
	assert.Equal(1, degree2)
}

func TestGraphGenerator_EulerianCycleGraph(t *testing.T) {
	assert := assert.New(t)

	V, E := 5, 3
	g, err := EulerianCycleGraph(V, E)
	assert.Nil(err)
	ec := NewEulerianCycle(g)
	assert.True(ec.HasEulerianCycle())

	// E < 0
	g1, err1 := EulerianCycleGraph(5, -1)
	assert.Nil(g1)
	assert.EqualError(err1, "an Eulerian cycle must have at least one edge")

	// V < 0
	g2, err2 := EulerianCycleGraph(-3, 4)
	assert.Nil(g2)
	assert.EqualError(err2, "an Eulerian cycle must have at least one vertex")
}

func TestGraphGenerator_EulerianPathGraph(t *testing.T) {
	assert := assert.New(t)

	V, E := 7, 4
	g, err := EulerianPathGraph(V, E)
	assert.Nil(err)
	ec := NewEulerianPath(g)
	assert.True(ec.HasEulerianPath())

	// E < 0
	g1, err1 := EulerianPathGraph(5, -1)
	assert.Nil(g1)
	assert.EqualError(err1, "negative number of edges")

	// V < 0
	g2, err2 := EulerianPathGraph(-3, 4)
	assert.Nil(g2)
	assert.EqualError(err2, "an Eulerian path must have at least one vertex")

}
