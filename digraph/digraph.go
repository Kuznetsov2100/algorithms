package digraph

// A self-loop is an edge that connects a vertex to itself.
// Two edges are parallel if they connect the same ordered pair of vertices.
// The outdegree of a vertex is the number of edges pointing from it.
// The indegree of a vertex is the number of edges pointing to it.
// A subgraph is a subset of a digraph's edges (and associated vertices) that constitutes a digraph.
// A directed path in a digraph is a sequence of vertices in which there is a (directed) edge pointing from each vertex in the sequence to its successor in the sequence, with no repeated edges.
// A directed path is simple if it has no repeated vertices.
// A directed cycle is a directed path (with at least one edge) whose first and last vertices are the same.
// A directed cycle is simple if it has no repeated vertices (other than the requisite repetition of the first and last vertices).
// The length of a path or a cycle is its number of edges.
// We say that a vertex w is reachable from a vertex v if there exists a directed path from v to w.
// We say that two vertices v and w are strongly connected if they are mutually reachable: there is a directed path from v to w and a directed path from w to v.
// A digraph is strongly connected if there is a directed path from every vertex to every other vertex.
// A digraph that is not strongly connected consists of a set of strongly connected components, which are maximal strongly connected subgraphs.
// A directed acyclic graph (or DAG) is a digraph with no directed cycles.

import (
	"fmt"
	"strings"

	"github.com/handane123/algorithms/dataStructure/bag"
	"github.com/handane123/algorithms/stdin"
	"github.com/pkg/errors"
)

// Digraph struct represents a directed graph of vertices named 0 through V - 1.
// It supports the following two primary operations: add an edge to the digraph,
// iterate over all of the vertices adjacent from a given vertex.
// It also provides methods for returning the indegree or outdegree of a vertex,
// the number of vertices V in the digraph, the number of edges E in the digraph,
// and the reverse digraph. Parallel edges and self-loops are permitted.
// This implementation uses an adjacency-lists representation,
// which is a vertex-indexed array of Bag objects. It uses O(E + V) space,
// where E is the number of edges and V is the number of vertices. All instance methods take O(1) time.
// (Though, iterating over the vertices returned by adj(int) takes time proportional to the
// outdegree of the vertex.) Constructing an empty digraph with V vertices takes O(V) time;
// constructing a digraph with E edges and V vertices takes O(E + V) time.
type Digraph struct {
	v        int
	e        int
	adj      []*bag.Bag
	indegree []int
}

// NewDigraph initializes an empty graph with V vertices and 0 edges.
func NewDigraph(V int) *Digraph {
	if V < 0 {
		panic("Number of vertices in a Digraph must be non negative")
	}

	adj := make([]*bag.Bag, V)
	for v := 0; v < V; v++ {
		adj[v] = bag.New()
	}
	return &Digraph{
		v:        V,
		e:        0,
		adj:      adj,
		indegree: make([]int, V),
	}
}

// NewDigraphIn initializes a graph from the specified input stream.
func NewDigraphIn(in *stdin.In) (*Digraph, error) {
	if in == nil {
		return nil, errors.New("argument is nil")
	}

	V := in.ReadInt()
	if V < 0 {
		return nil, errors.New("number of vertices in a Digraph must be non negative")
	}
	E := in.ReadInt()
	if E < 0 {
		return nil, errors.New("number of edges in a Digraph must be non negative")
	}
	adj := make([]*bag.Bag, V)
	for v := 0; v < V; v++ {
		adj[v] = bag.New()
	}
	dg := &Digraph{
		e:        0,
		v:        V,
		adj:      adj,
		indegree: make([]int, V),
	}
	for i := 0; i < E; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		dg.validateVertex(v)
		dg.validateVertex(w)
		dg.AddEdge(v, w)
	}
	return dg, nil
}

// AddEdge adds the directed edge v-w to this digraph.
func (dg *Digraph) AddEdge(v, w int) {
	dg.validateVertex(v)
	dg.validateVertex(w)
	dg.adj[v].Add(w)
	dg.indegree[w]++
	dg.e++
}

// OutDegree returns the number of directed edges incident from vertex v.
func (dg *Digraph) OutDegree(v int) int {
	dg.validateVertex(v)
	return dg.adj[v].Size()
}

// InDegree returns the number of directed edges incident to vertex v.
func (dg *Digraph) InDegree(v int) int {
	dg.validateVertex(v)
	return dg.indegree[v]
}

// V returns the number of vertices in this digraph.
func (dg *Digraph) V() int {
	return dg.v
}

// E returns the number of edges in this digraph.
func (dg *Digraph) E() int {
	return dg.e
}

// Adj returns the vertices adjacent to vertex v.
func (dg *Digraph) Adj(v int) (vertices []int) {
	dg.validateVertex(v)
	for _, x := range dg.adj[v].Values() {
		vertices = append(vertices, x.(int))
	}
	return vertices
}

// Reverse returns the reverse of the digraph.
func (dg *Digraph) Reverse() *Digraph {
	reverse := NewDigraph(dg.v)
	for v := 0; v < dg.v; v++ {
		for _, val := range dg.adj[v].Values() {
			w := val.(int)
			reverse.AddEdge(w, v)
		}
	}
	return reverse
}

// String returns a string representation of this digraph.
func (dg *Digraph) String() string {
	var s strings.Builder
	fmt.Fprintf(&s, "%d vertices, %d edges \n", dg.v, dg.e)
	for i := 0; i < dg.v; i++ {
		var adjs strings.Builder
		for _, w := range dg.adj[i].Values() {
			fmt.Fprintf(&adjs, " %d", w.(int))
		}
		fmt.Fprintf(&s, "%d: %s\n", i, adjs.String())
	}
	return s.String()
}

func (dg *Digraph) validateVertex(v int) {
	if v < 0 || v >= dg.v {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", dg.v-1))
	}
}
