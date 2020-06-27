package digraph

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"github.com/handane123/algorithms/dataStructure/bag"
	"github.com/handane123/algorithms/stdin"
)

// EdgeWeightedGraph struct represents an edge-weighted graph of vertices named 0 through V – 1,
// where each undirected edge is of type Edge and has a real-valued weight.
// This implementation uses an adjacency-lists representation, which is a vertex-indexed array
// of Bag objects. It uses O(E + V) space, where E is the number of edges and V is the number
// of vertices. All instance methods take O(1) time. (Though, iterating over the edges returned
// by adj(int) takes time proportional to the degree of the vertex.) Constructing an empty
// edge-weighted graph with V vertices takes O(V) time; constructing a edge-weighted graph
// with E edges and V vertices takes O(E + V) time.
type EdgeWeightedGraph struct {
	v   int
	e   int
	adj []*bag.Bag
}

// NewEdgeWeightedGraphV initializes an empty edge-weighted graph with V vertices and 0 edges.
func NewEdgeWeightedGraphV(V int) *EdgeWeightedGraph {
	if V < 0 {
		panic("Number of vertices must be non negative")
	}
	adj := make([]*bag.Bag, V)
	for v := 0; v < V; v++ {
		adj[v] = bag.New()
	}

	return &EdgeWeightedGraph{
		v:   V,
		e:   0,
		adj: adj,
	}
}

// NewEdgeWeightedGraphVE initializes a random edge-weighted graph with V vertices and E edges.
func NewEdgeWeightedGraphVE(V, E int) *EdgeWeightedGraph {
	wg := NewEdgeWeightedGraphV(V)
	if E < 0 {
		panic("Number of edges must be non negative")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < E; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		weight := math.Round(100*rand.Float64()) / 100.0
		e := NewEdge(v, w, weight)
		wg.AddEdge(e)
	}
	return wg
}

// NewEdgeWeightedGraphIn initializes an edge-weighted graph from an input stream.
func NewEdgeWeightedGraphIn(in *stdin.In) *EdgeWeightedGraph {
	if in == nil {
		panic("argument is nil")
	}

	V := in.ReadInt()
	if V < 0 {
		panic("number of vertices must be non negative")
	}
	adj := make([]*bag.Bag, V)
	for v := 0; v < V; v++ {
		adj[v] = bag.New()
	}

	E := in.ReadInt()
	if E < 0 {
		panic("number of edges must be non negative")
	}

	wg := &EdgeWeightedGraph{v: V, e: 0, adj: adj}
	for i := 0; i < E; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		wg.validateVertex(v)
		wg.validateVertex(w)
		weight := in.ReadFloat64()
		e := NewEdge(v, w, weight)
		wg.AddEdge(e)

	}
	return wg
}

// V returns the number of vertices in this edge-weighted graph.
func (wg *EdgeWeightedGraph) V() int {
	return wg.v
}

// E returns the number of edges in this edge-weighted graph.
func (wg *EdgeWeightedGraph) E() int {
	return wg.e
}

// AddEdge adds the undirected edge e to this edge-weighted graph.
func (wg *EdgeWeightedGraph) AddEdge(e *Edge) {
	v := e.Either()
	w := e.Other(v)
	wg.validateVertex(v)
	wg.validateVertex(w)
	wg.adj[v].Add(e)
	wg.adj[w].Add(e)
	wg.e++
}

// Adj returns the edges incident on vertex v.
func (wg *EdgeWeightedGraph) Adj(v int) (edges []*Edge) {
	wg.validateVertex(v)
	for _, x := range wg.adj[v].Values() {
		edges = append(edges, x.(*Edge))
	}
	return edges
}

// Degree returns the degree of vertex v.
func (wg *EdgeWeightedGraph) Degree(v int) int {
	wg.validateVertex(v)
	return wg.adj[v].Size()
}

// Edges returns all edges in this edge-weighted graph.
func (wg *EdgeWeightedGraph) Edges() (edges []*Edge) {
	for v := 0; v < wg.v; v++ {
		selfLoops := 0
		for _, e := range wg.adj[v].Values() {
			edge := e.(*Edge)
			if val := edge.Other(v); val > v {
				edges = append(edges, edge)
			} else if val == v { // add only one copy of each self loop (self loops will be consecutive)
				if selfLoops%2 == 0 {
					edges = append(edges, edge)
				}
				selfLoops++
			}
		}
	}
	return edges
}

// String returns a string representation of the edge-weighted graph.
func (wg *EdgeWeightedGraph) String() string {
	var s strings.Builder
	fmt.Fprintf(&s, "%d %d\n", wg.v, wg.e)
	for i := 0; i < wg.v; i++ {
		fmt.Fprint(&s, i, ": ")
		for _, e := range wg.adj[i].Values() {
			edge := e.(*Edge)
			fmt.Fprint(&s, edge, "  ")
		}
		fmt.Fprintf(&s, "\n")
	}
	return s.String()
}

func (wg *EdgeWeightedGraph) validateVertex(v int) {
	if v < 0 || v >= wg.v {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", wg.v-1))
	}
}
