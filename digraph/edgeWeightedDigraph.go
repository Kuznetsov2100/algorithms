package digraph

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/handane123/algorithms/dataStructure/bag"
	"github.com/handane123/algorithms/stdin"
)

type EdgeWeightedDigraph struct {
	v        int
	e        int
	adj      []*bag.Bag
	indegree []int
}

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

// NewEdgeWeightedGraphIn initializes an edge-weighted graph from an input stream.
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

func (wd *EdgeWeightedDigraph) V() int {
	return wd.v
}

func (wd *EdgeWeightedDigraph) E() int {
	return wd.e
}

// Adj returns the edges incident on vertex v.
func (wd *EdgeWeightedDigraph) Adj(v int) (edges []*DirectedEdge) {
	for _, x := range wd.adj[v].Values() {
		edges = append(edges, x.(*DirectedEdge))
	}
	return edges
}

func (wd *EdgeWeightedDigraph) OutDegree(v int) int {
	wd.validateVertex(v)
	return wd.adj[v].Size()
}

func (wd *EdgeWeightedDigraph) InDegree(v int) int {
	wd.validateVertex(v)
	return wd.indegree[v]
}

func (wd *EdgeWeightedDigraph) Edges() (edges []*DirectedEdge) {
	for v := 0; v < wd.v; v++ {
		for _, e := range wd.adj[v].Values() {
			edges = append(edges, e.(*DirectedEdge))
		}
	}
	return edges
}

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
