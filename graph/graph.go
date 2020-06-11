package graph

import (
	"errors"
	"fmt"

	"github.com/handane123/algorithms/stdin"

	linkedbag "github.com/handane123/algorithms/dataStructure/bag"
)

type Graph struct {
	v   int
	e   int
	adj []*linkedbag.Bag
}

func NewGraph(V int) *Graph {
	if V < 0 {
		panic("Number of vertices must be non negative")
	}
	adj := make([]*linkedbag.Bag, V)
	for i := 0; i < V; i++ {
		adj[i] = linkedbag.New()
	}
	return &Graph{v: V, e: 0, adj: adj}
}

func NewGraphIn(in *stdin.In) (*Graph, error) {
	if in == nil {
		return nil, errors.New("argument is nil")
	}
	V := in.ReadInt()
	if V < 0 {
		return nil, errors.New("number of vertices in a Graph must be non negative")
	}
	E := in.ReadInt()
	if E < 0 {
		return nil, errors.New("number of edges in a Graph must be non negative")
	}
	adj := make([]*linkedbag.Bag, V)
	for i := 0; i < V; i++ {
		adj[i] = linkedbag.New()
	}
	g := &Graph{e: E, v: V, adj: adj}
	for i := 0; i < E; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		g.validateVertex(v)
		g.validateVertex(w)
		g.AddEdge(v, w)
	}
	return g, nil
}

func (g *Graph) validateVertex(v int) {
	if v < 0 || v >= g.v {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", g.v-1))
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.validateVertex(v)
	g.validateVertex(w)
	g.e++
	g.adj[v].Add(w)
	g.adj[w].Add(v)
}

func (g *Graph) V() int {
	return g.v
}

func (g *Graph) E() int {
	return g.e
}

func (g *Graph) Degree(v int) int {
	g.validateVertex(v)
	return g.adj[v].Size()
}

func (g *Graph) Adj(v int) (vertices []int) {
	for _, x := range g.adj[v].Values() {
		vertices = append(vertices, x.(int))
	}
	return vertices
}

func (g *Graph) String() string {
	s := fmt.Sprintf("%d vertices, %d edges \n", g.v, g.e)
	for i := 0; i < g.v; i++ {
		adjs := ""
		for _, w := range g.Adj(i) {
			adjs = adjs + fmt.Sprintf(" %d", w)
		}
		s += fmt.Sprintf("%d: %s\n", i, adjs)
	}
	return s
}
