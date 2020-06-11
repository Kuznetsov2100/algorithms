package graph

import (
	"errors"
	"fmt"

	"github.com/handane123/algorithms/dataStructure/bag"
	"github.com/handane123/algorithms/stdin"
)

// Graph struct represents an undirected graph of vertices named 0 through V – 1.
// This implementation uses an adjacency-lists representation,
// which is a vertex-indexed array of Bag objects. It uses O(E + V) space,
// where E is the number of edges and V is the number of vertices.
// All instance methods take O(1) time.
// (Though, iterating over the vertices returned by adj(int)
// takes time proportional to the degree of the vertex.)
// Constructing an empty graph with V vertices takes O(V) time;
// constructing a graph with E edges and V vertices takes O(E + V) time.
type Graph struct {
	v   int
	e   int
	adj []*bag.Bag
}

// NewGraph initializes an empty graph with V vertices and 0 edges.
func NewGraph(V int) *Graph {
	if V < 0 {
		panic("Number of vertices must be non negative")
	}
	adj := make([]*bag.Bag, V)
	for i := 0; i < V; i++ {
		adj[i] = bag.New()
	}
	return &Graph{v: V, e: 0, adj: adj}
}

// NewGraphIn initializes a graph from the specified input stream.
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
	adj := make([]*bag.Bag, V)
	for i := 0; i < V; i++ {
		adj[i] = bag.New()
	}
	g := &Graph{e: 0, v: V, adj: adj}
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

// V returns the number of vertices in this graph.
func (g *Graph) V() int {
	return g.v
}

// E returns the number of edges in this graph.
func (g *Graph) E() int {
	return g.e
}

// AddEdge adds the undirected edge v-w to this graph.
func (g *Graph) AddEdge(v, w int) {
	g.validateVertex(v)
	g.validateVertex(w)
	g.e++
	g.adj[v].Add(w)
	g.adj[w].Add(v)
}

// returns the degree of vertex v.
func (g *Graph) Degree(v int) int {
	g.validateVertex(v)
	return g.adj[v].Size()
}

// Adj returns the vertices adjacent to vertex v.
func (g *Graph) Adj(v int) (vertices []int) {
	for _, x := range g.adj[v].Values() {
		vertices = append(vertices, x.(int))
	}
	return vertices
}

// ToString returns a string representation of this graph.
func (g *Graph) ToString() string {
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
