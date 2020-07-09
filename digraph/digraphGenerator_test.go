package digraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigraphGenerator_Complete(t *testing.T) {
	generator := NewDigraphGenerator()
	assert := assert.New(t)

	G := generator.Complete(7)
	assert.Equal(7, G.V())
	assert.Equal(42, G.E())
}

func TestDigraphGenerator_DAG(t *testing.T) {
	generator := NewDigraphGenerator()
	assert := assert.New(t)

	// E > V*(V-1)/2
	G1, err1 := generator.Dag(3, 6)
	assert.Nil(G1)
	assert.EqualError(err1, "too many edges")

	// E < 0
	G2, err2 := generator.Dag(3, -1)
	assert.Nil(G2)
	assert.EqualError(err2, "too few edges")

	// Dag
	G3, err3 := generator.Dag(7, 4)
	assert.Nil(err3)
	finder := NewDirectedCycle(G3)
	assert.False(finder.HasCycle())
}

func TestDigraphGenerator_Cycle(t *testing.T) {
	generator := NewDigraphGenerator()
	assert := assert.New(t)

	G := generator.CycleDigraph(7)
	finder := NewDirectedCycle(G)
	assert.True(finder.HasCycle())
}

func TestDigraphGenerator_EulerianCycleDigraph(t *testing.T) {
	generator := NewDigraphGenerator()
	assert := assert.New(t)

	V, E := 5, 3
	g, err := generator.EulerianCycleDigraph(V, E)
	assert.Nil(err)
	ec := NewDirectedEulerianCycle(g)
	assert.True(ec.HasEulerianCycle())

	// E < 0
	g1, err1 := generator.EulerianCycleDigraph(5, -1)
	assert.Nil(g1)
	assert.EqualError(err1, "an Eulerian cycle must at least one edge")

	// V < 0
	g2, err2 := generator.EulerianCycleDigraph(-3, 4)
	assert.Nil(g2)
	assert.EqualError(err2, "an Eulerian cycle must at least one vertex")
}

func TestDigraphGenerator_EulerianPathDigraph(t *testing.T) {
	generator := NewDigraphGenerator()
	assert := assert.New(t)

	V, E := 7, 4
	g, err := generator.EulerianPathDigraph(V, E)
	assert.Nil(err)
	ec := NewDirectedEulerianPath(g)
	assert.True(ec.HasEulerianPath())

	// E < 0
	g1, err1 := generator.EulerianPathDigraph(5, -1)
	assert.Nil(g1)
	assert.EqualError(err1, "negative number of edges")

	// V < 0
	g2, err2 := generator.EulerianPathDigraph(-3, 4)
	assert.Nil(g2)
	assert.EqualError(err2, "an Eulerian path must have at least one vertex")

}

func TestDigraphGenerator_Path(t *testing.T) {
	generator := NewDigraphGenerator()
	assert := assert.New(t)

	G := generator.PathDigraph(7)
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
	generator := NewDigraphGenerator()
	assert := assert.New(t)

	V := 15
	g := generator.BinaryTree(V)
	finder := NewDirectedCycle(g)
	assert.False(finder.HasCycle())
	assert.Equal(V/2*2, g.E())

}

func TestDigraphGenerator_StrongDigraph(t *testing.T) {
	generator := NewDigraphGenerator()
	assert := assert.New(t)

	// c >= V || c <= 0
	G1, err1 := generator.StrongDigraph(7, 4, 9)
	assert.Nil(G1)
	assert.EqualError(err1, "number of components must be between 1 and V")

	// E <=2 *(V-c)
	G2, err2 := generator.StrongDigraph(7, 2, 4)
	assert.Nil(G2)
	assert.EqualError(err2, "number of edges must be at least 2*(V-c)")

	// E > V*(V-1)/2
	G3, err3 := generator.StrongDigraph(4, 7, 2)
	assert.Nil(G3)
	assert.EqualError(err3, "too many edges")

	// strongDigraph
	G4, err4 := generator.StrongDigraph(9, 20, 4)
	assert.Nil(err4)

	cc := NewKosarajuSharirSCC(G4)
	assert.LessOrEqual(cc.Count(), 4)
}
