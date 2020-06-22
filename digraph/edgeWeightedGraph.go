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

type EdgeWeightedGraph struct {
	v   int
	e   int
	adj []*bag.Bag
}

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

func NewEdgeWeightedGraphVE(V, E int) *EdgeWeightedGraph {
	if E < 0 {
		panic("Number of edges must be non negative")
	}
	rand.Seed(time.Now().UnixNano())
	wg := &EdgeWeightedGraph{v: V, e: 0}
	for i := 0; i < E; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		weight := math.Round(100*rand.Float64()) / 100.0
		e := NewEdge(v, w, weight)
		wg.AddEdge(e)
	}
	return wg
}

func NewEdgeWeightedGraphIn(in *stdin.In) *EdgeWeightedGraph {
	if in == nil {
		panic("argument is nil")
	}

	V := in.ReadInt()
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

func (wg *EdgeWeightedGraph) V() int {
	return wg.v
}

func (wg *EdgeWeightedGraph) E() int {
	return wg.e
}

func (wg *EdgeWeightedGraph) AddEdge(e *Edge) {
	v := e.Either()
	w := e.Other(v)
	wg.validateVertex(v)
	wg.validateVertex(w)
	wg.adj[v].Add(e)
	wg.adj[w].Add(e)
	wg.e++
}

func (wg *EdgeWeightedGraph) Adj(v int) (edges []*Edge) {
	for _, x := range wg.adj[v].Values() {
		edges = append(edges, x.(*Edge))
	}
	return edges
}

func (wg *EdgeWeightedGraph) Degree(v int) int {
	wg.validateVertex(v)
	return wg.adj[v].Size()
}

func (wg *EdgeWeightedGraph) Edges() (edges []*Edge) {
	for v := 0; v < wg.v; v++ {
		selfLoops := 0
		for _, e := range wg.adj[v].Values() {
			edge := e.(*Edge)
			if val := edge.Other(v); val > v {
				edges = append(edges, edge)
			} else if val == v {
				if selfLoops%2 == 0 {
					edges = append(edges, edge)
				}
				selfLoops++
			}
		}
	}
	return edges
}

func (wg *EdgeWeightedGraph) String() string {
	var s strings.Builder
	fmt.Fprintf(&s, "%d %d\n", wg.v, wg.e)
	for i := 0; i < wg.v; i++ {
		fmt.Fprint(&s, i, ": ")
		for _, e := range wg.adj[i].Values() {
			edge := e.(*Edge)
			fmt.Fprint(&s, edge, " ")
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
