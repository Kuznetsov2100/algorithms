package graph

// A self-loop is an edge that connects a vertex to itself.
// Two edges are parallel if they connect the same pair of vertices.
// When an edge connects two vertices, we say that the vertices are adjacent to one another and that the edge is incident on both vertices.
// The degree of a vertex is the number of edges incident on it.
// A subgraph is a subset of a graph's edges (and associated vertices) that constitutes a graph.
// A path in a graph is a sequence of vertices connected by edges, with no repeated edges.
// A simple path is a path with no repeated vertices.
// A cycle is a path (with at least one edge) whose first and last vertices are the same.
// A simple cycle is a cycle with no repeated vertices (other than the requisite repetition of the first and last vertices).
// The length of a path or a cycle is its number of edges.
// We say that one vertex is connected to another if there exists a path that contains both of them.
// A graph is connected if there is a path from every vertex to every other vertex.
// A graph that is not connected consists of a set of connected components, which are maximal connected subgraphs.
// An acyclic graph is a graph with no cycles.
// A tree is an acyclic connected graph.
// A forest is a disjoint set of trees.
// A spanning tree of a connected graph is a subgraph that contains all of that graph's vertices and is a single tree. A spanning forest of a graph is the union of the spanning trees of its connected components.
// A bipartite graph is a graph whose vertices we can divide into two sets such that all edges connect a vertex in one set with a vertex in the other set.

import (
	"errors"
	"fmt"
	"strings"

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

// Degree returns the degree of vertex v.
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
func (g *Graph) String() string {
	var s strings.Builder
	fmt.Fprintf(&s, "%d vertices, %d edges \n", g.v, g.e)
	for i := 0; i < g.v; i++ {
		var adjs strings.Builder
		for _, w := range g.Adj(i) {
			fmt.Fprintf(&adjs, " %d", w)
		}
		fmt.Fprintf(&s, "%d: %s\n", i, adjs.String())
	}
	return s.String()
}
