package digraph

import (
	"fmt"
	"strings"

	"github.com/handane123/algorithms/dataStructure/bag"
	"github.com/handane123/algorithms/stdin"
	"github.com/pkg/errors"
)

type Digraph struct {
	v        int
	e        int
	adj      []*bag.Bag
	indegree []int
}

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

func (dg *Digraph) AddEdge(v, w int) {
	dg.validateVertex(v)
	dg.validateVertex(w)
	dg.adj[v].Add(w)
	dg.indegree[w]++
	dg.e++
}

func (dg *Digraph) OutDegree(v int) int {
	dg.validateVertex(v)
	return dg.adj[v].Size()
}

func (dg *Digraph) InDegree(v int) int {
	dg.validateVertex(v)
	return dg.indegree[v]
}

func (dg *Digraph) V() int {
	return dg.v
}

func (dg *Digraph) E() int {
	return dg.e
}

func (dg *Digraph) Adj(v int) (vertices []int) {
	for _, x := range dg.adj[v].Values() {
		vertices = append(vertices, x.(int))
	}
	return vertices
}

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

// String returns a string representation of this graph.
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
