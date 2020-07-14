package digraph

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/handane123/algorithms/dataStructure/bag"
	"github.com/handane123/algorithms/io/stdin"
)

type FlowNetwork struct {
	v   int
	e   int
	adj []*bag.Bag
}

func NewFlowNetwork(V int) *FlowNetwork {
	if V < 0 {
		panic("number of vertices in a graph must be non-negative")
	}
	fn := &FlowNetwork{v: V, e: 0, adj: make([]*bag.Bag, V)}
	for v := 0; v < V; v++ {
		fn.adj[v] = bag.New()
	}
	return fn
}

func NewFlowNetworkVE(V, E int) *FlowNetwork {
	fn := NewFlowNetwork(V)
	if E < 0 {
		panic("number of edges must be non-negative")
	}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < E; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		capacity := float64(rand.Intn(100))
		fn.AddEdge(NewFlowEdge(v, w, capacity))
	}
	return fn
}

func NewFlowNetworkIn(in *stdin.In) *FlowNetwork {
	fn := NewFlowNetwork(in.ReadInt())
	E := in.ReadInt()
	if E < 0 {
		panic("number of edges must be non-negative")
	}
	for i := 0; i < E; i++ {
		v := in.ReadInt()
		w := in.ReadInt()
		fn.validateVertex(v)
		fn.validateVertex(w)
		capacity := in.ReadFloat64()
		fn.AddEdge(NewFlowEdge(v, w, capacity))
	}
	return fn
}

func (fn *FlowNetwork) AddEdge(e *FlowEdge) {
	v := e.From()
	w := e.To()
	fn.validateVertex(v)
	fn.validateVertex(w)
	fn.adj[v].Add(e)
	fn.adj[w].Add(e)
	fn.e++
}

func (fn *FlowNetwork) V() int {
	return fn.v
}

func (fn *FlowNetwork) E() int {
	return fn.e
}

func (fn *FlowNetwork) Adj(v int) (edges []*FlowEdge) {
	fn.validateVertex(v)
	for _, val := range fn.adj[v].Values() {
		edges = append(edges, val.(*FlowEdge))
	}
	return edges
}

func (fn *FlowNetwork) Edges() (edges []*FlowEdge) {
	for v := 0; v < fn.v; v++ {
		for _, e := range fn.adj[v].Values() {
			edge := e.(*FlowEdge)
			if edge.To() != v {
				edges = append(edges, edge)
			}
		}
	}
	return edges
}

func (fn *FlowNetwork) String() string {
	var s strings.Builder
	fmt.Fprintf(&s, "%d %d\n", fn.v, fn.e)
	for i := 0; i < fn.v; i++ {
		fmt.Fprint(&s, i, ": ")
		for _, e := range fn.adj[i].Values() {
			edge := e.(*FlowEdge)
			if edge.To() != i {
				fmt.Fprint(&s, edge, "  ")
			}
		}
		fmt.Fprintf(&s, "\n")
	}
	return s.String()
}

func (fn *FlowNetwork) validateVertex(v int) {
	if v < 0 || v >= fn.v {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", fn.v-1))
	}
}
