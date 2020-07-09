package digraph

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/handane123/algorithms/dataStructure/bag"
	"github.com/handane123/algorithms/io/stdin"
)

// EdgeWeightedDigraph struct represents a edge-weighted digraph of vertices named 0 through V - 1,
// where each directed edge is of type DirectedEdge and has a real-valued weight.
// This implementation uses an adjacency-lists representation, which is a vertex-indexed array
// of Bag objects. It uses O(E + V) space, where E is the number of edges and V is the number
// of vertices. All instance methods take O(1) time. (Though, iterating over the edges returned
// by adj(int) takes time proportional to the outdegree of the vertex.) Constructing an empty
// edge-weighted digraph with V vertices takes O(V) time; constructing an edge-weighted digraph
// with E edges and V vertices takes O(E + V) time.
type EdgeWeightedDigraph struct {
	v        int
	e        int
	adj      []*bag.Bag
	indegree []int
}

// NewEdgeWeightedDigraphV initializes an empty edge-weighted digraph with V vertices and 0 edges.
func NewEdgeWeightedDigraphV(V int) *EdgeWeightedDigraph {
	if V < 0 {
		panic("number of vertices in a digraph must be non negative")
	}
	adj := make([]*bag.Bag, V)
	for v := 0; v < V; v++ {
		adj[v] = bag.New()
	}
	return &EdgeWeightedDigraph{
		v:        V,
		e:        0,
		adj:      adj,
		indegree: make([]int, V),
	}
}

// NewEdgeWeightedDigraphVE initializes a random edge-weighted digraph with V vertices and E edges.
func NewEdgeWeightedDigraphVE(V, E int) *EdgeWeightedDigraph {
	wd := NewEdgeWeightedDigraphV(V)
	if E < 0 {
		panic("number of edges in a digraph must be non negative")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < E; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		weight := 0.01 * float64(rand.Intn(100))
		edge := NewDirectedEdge(v, w, weight)
		wd.AddEdge(edge)
	}
	return wd
}

// NewEdgeWeightedDigraphIn initializes an edge-weighted digraph from an input stream.
func NewEdgeWeightedDigraphIn(in *stdin.In) *EdgeWeightedDigraph {
	if in == nil {
		panic("argument is nil")
	}

	V := in.ReadInt()
	if V < 0 {
		panic("number of vertices in a digraph must be non negative")
	}
	adj := make([]*bag.Bag, V)
	for v := 0; v < V; v++ {
		adj[v] = bag.New()
	}

	E := in.ReadInt()
	if E < 0 {
		panic("number of edges must be non negative")
	}

	wd := &EdgeWeightedDigraph{v: V, e: 0, adj: adj, indegree: make([]int, V)}
	for i := 0; i < E; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		wd.validateVertex(v)
		wd.validateVertex(w)
		weight := in.ReadFloat64()
		e := NewDirectedEdge(v, w, weight)
		wd.AddEdge(e)
	}
	return wd
}

// V returns the number of vertices in this edge-weighted digraph.
func (wd *EdgeWeightedDigraph) V() int {
	return wd.v
}

// E returns the number of edges in this edge-weighted digraph.
func (wd *EdgeWeightedDigraph) E() int {
	return wd.e
}

// Adj returns the directed edges incident from vertex v.
func (wd *EdgeWeightedDigraph) Adj(v int) (edges []*DirectedEdge) {
	wd.validateVertex(v)
	for _, x := range wd.adj[v].Values() {
		edges = append(edges, x.(*DirectedEdge))
	}
	return edges
}

// OutDegree returns the number of directed edges incident from vertex v.
func (wd *EdgeWeightedDigraph) OutDegree(v int) int {
	wd.validateVertex(v)
	return wd.adj[v].Size()
}

// InDegree returns the number of directed edges incident to vertex v.
func (wd *EdgeWeightedDigraph) InDegree(v int) int {
	wd.validateVertex(v)
	return wd.indegree[v]
}

// Edges returns all directed edges in this edge-weighted digraph.
func (wd *EdgeWeightedDigraph) Edges() (edges []*DirectedEdge) {
	for v := 0; v < wd.v; v++ {
		for _, e := range wd.adj[v].Values() {
			edges = append(edges, e.(*DirectedEdge))
		}
	}
	return edges
}

// AddEdge adds the directed edge e to this edge-weighted digraph.
func (wd *EdgeWeightedDigraph) AddEdge(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	wd.validateVertex(v)
	wd.validateVertex(w)
	wd.adj[v].Add(e)
	wd.indegree[w]++
	wd.e++
}

// String returns a string representation of the edge-weighted digraph.
func (wd *EdgeWeightedDigraph) String() string {
	var s strings.Builder
	fmt.Fprintf(&s, "%d %d\n", wd.v, wd.e)
	for i := 0; i < wd.v; i++ {
		fmt.Fprint(&s, i, ": ")
		for _, e := range wd.adj[i].Values() {
			edge := e.(*DirectedEdge)
			fmt.Fprint(&s, edge, "  ")
		}
		fmt.Fprintf(&s, "\n")
	}
	return s.String()
}

func (wd *EdgeWeightedDigraph) validateVertex(v int) {
	if v < 0 || v >= wd.v {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", wd.v-1))
	}
}
