package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphGenerator_Simple(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)
	g, err := generator.Simple(3, 4)
	assert.Nil(g)
	assert.Error(err)

	g1, err1 := generator.Simple(3, -1)
	assert.Nil(g1)
	assert.Error(err1)

	V := 20
	E := 30
	g2, err2 := generator.Simple(V, E)
	assert.Nil(err2)

	// a simple graph has no self-loop and no parallel edges.
	finder := NewCycle(g2)
	assert.False(finder.hasParallelEdges(g2))
	assert.False(finder.hasSelfLoop(g2))
}

func TestGraphGenerator_SimpleP(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)
	g, err := generator.SimpleP(3, 1.2)
	assert.Nil(g)
	assert.Error(err)

	V := 5
	p := 0.769283
	g1, err1 := generator.SimpleP(V, p)
	assert.Nil(err1)

	finder := NewCycle(g1)
	assert.False(finder.hasParallelEdges(g1))
	assert.False(finder.hasSelfLoop(g1))
}

func TestGraphGenerator_Complete(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)
	V := 3
	g := generator.Complete(V)
	for i := 0; i < V; i++ {
		assert.Equal(2, g.Degree(i))
	}
}

func TestGraphGenerator_Bipartite(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)
	V1, V2, E := 10, 5, 8
	G, err := generator.BipartiteGraph(V1, V2, E)
	assert.Nil(err)
	b := NewBipartite(G)
	assert.True(b.IsBipartite())

	// too many edges
	G1, err1 := generator.BipartiteGraph(7, 3, 30)
	assert.Nil(G1)
	assert.EqualError(err1, "too many edges")

	// too few edges
	G2, err2 := generator.BipartiteGraph(7, 3, -1)
	assert.Nil(G2)
	assert.EqualError(err2, "too few edges")
}

func TestGraphGenerator_BipartiteP(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)
	V1, V2, p := 10, 5, 0.5732
	G, err := generator.BipartiteP(V1, V2, p)
	assert.Nil(err)
	b := NewBipartite(G)
	assert.True(b.IsBipartite())

	// invalid probability
	G1, err1 := generator.BipartiteP(10, 5, 1.2)
	assert.Nil(G1)
	assert.EqualError(err1, "probability must be between 0 and 1")
}

func TestGraphGenerator_CompleteBipartite(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	V1, V2 := 5, 3
	G := generator.CompleteBipartite(V1, V2)

	b := NewBipartite(G)
	assert.True(b.IsBipartite())
	assert.Equal(V1+V2, G.V())
	assert.Equal(V1*V2, G.E())
}

func TestGraphGenerator_CycleGraph(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	V := 7
	g := generator.CycleGraph(V)
	finder := NewCycle(g)
	assert.True(finder.HasCycle())
}

func TestGraphGenerator_Path(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	V := 3
	g := generator.PathGraph(V)

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

func TestGraphGenerator_BinaryTree(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	V := 15
	g := generator.BinaryTree(V)
	finder := NewCycle(g)
	search := NewDepthFirstSearch(g, 0)
	assert.False(finder.HasCycle())
	assert.Equal(V/2*2, g.E())
	assert.Equal(g.V(), search.Count())
}

func TestGraphGenerator_EulerianCycleGraph(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	V, E := 5, 3
	g, err := generator.EulerianCycleGraph(V, E)
	assert.Nil(err)
	ec := NewEulerianCycle(g)
	assert.True(ec.HasEulerianCycle())

	// E < 0
	g1, err1 := generator.EulerianCycleGraph(5, -1)
	assert.Nil(g1)
	assert.EqualError(err1, "an Eulerian cycle must have at least one edge")

	// V < 0
	g2, err2 := generator.EulerianCycleGraph(-3, 4)
	assert.Nil(g2)
	assert.EqualError(err2, "an Eulerian cycle must have at least one vertex")
}

func TestGraphGenerator_EulerianPathGraph(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	V, E := 7, 4
	g, err := generator.EulerianPathGraph(V, E)
	assert.Nil(err)
	ec := NewEulerianPath(g)
	assert.True(ec.HasEulerianPath())

	// E < 0
	g1, err1 := generator.EulerianPathGraph(5, -1)
	assert.Nil(g1)
	assert.EqualError(err1, "negative number of edges")

	// V < 0
	g2, err2 := generator.EulerianPathGraph(-3, 4)
	assert.Nil(g2)
	assert.EqualError(err2, "an Eulerian path must have at least one vertex")

}

func TestGraphGenerator_Wheel(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	g1, err1 := generator.Wheel(1)
	assert.Nil(g1)
	assert.Error(err1)

	V := 8
	g2, err2 := generator.Wheel(V)
	assert.Nil(err2)
	assert.Equal(2*(V-1), g2.E())
}

func TestGraphGenerator_Star(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	g1, err1 := generator.Star(0)
	assert.Nil(g1)
	assert.Error(err1)

	V := 8
	g2, err2 := generator.Star(V)
	assert.Nil(err2)
	degree1count := 0
	degreenMinus1Count := 0
	for v := 0; v < g2.V(); v++ {
		if g2.Degree(v) == 1 {
			degree1count++
		} else if g2.Degree(v) == V-1 {
			degreenMinus1Count++
		}
	}
	assert.Equal(V-1, degree1count)
	assert.Equal(1, degreenMinus1Count)
}

func TestGraphGenerator_Regular(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	g1, err1 := generator.Regular(5, 3)
	assert.Nil(g1)
	assert.Error(err1)

	V, k := 5, 2
	g2, err2 := generator.Regular(V, k)
	assert.Nil(err2)
	for v := 0; v < g2.V()-1; v++ {
		assert.Equal(g2.Degree(v), g2.Degree(v+1))
	}
}

func TestGraphGenerator_Tree(t *testing.T) {
	generator := NewGraphGenerator()
	assert := assert.New(t)

	g1 := generator.Tree(1)
	finder := NewCycle(g1)
	search := NewDepthFirstSearch(g1, 0)
	assert.False(finder.HasCycle())
	assert.Equal(g1.V(), search.Count())

	g2 := generator.Tree(15)
	finder2 := NewCycle(g2)
	search2 := NewDepthFirstSearch(g2, 0)
	assert.False(finder2.HasCycle())
	assert.Equal(g2.V(), search2.Count())
}

func TestVkey(t *testing.T) {
	assert := assert.New(t)

	var a vkey = 1
	var b vkey = 1
	assert.Equal(0, a.CompareTo(b))
}
