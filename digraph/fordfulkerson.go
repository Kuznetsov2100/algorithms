package digraph

import (
	"fmt"
	"math"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
)

// FordFulkerson struct represents a data type for computing a maximum st-flow and minimum st-cut
// in a flow network. This implementation uses the Ford-Fulkerson algorithm with the shortest
// augmenting path heuristic. The constructor takes O(E V (E + V)) time, where V is the number
// of vertices and E is the number of edges. In practice, the algorithm will run much faster.
// The inCut() and value() methods take O(1) time. It uses O(V) extra space (not including the network).
type FordFulkerson struct {
	v      int         // number of vertices
	marked []bool      // marked[v] = true iff s->v path in residual graph
	edgeTo []*FlowEdge // edgeTo[v] = last edge on shortest residual s->v path
	value  float64     // current value of max flow
}

// NewFordFulkerson compute a maximum flow and minimum cut in the network G from vertex s to vertex t.
func NewFordFulkerson(G *FlowNetwork, s, t int) *FordFulkerson {
	ford := &FordFulkerson{v: G.V()}
	ford.validate(s)
	ford.validate(t)
	if s == t {
		panic("source equals sink")
	}
	if !ford.isFeasible(G, s, t) {
		panic("initial flow is infeasible")
	}

	// while there exists an augmenting path, use it
	ford.value = ford.excess(G, t)
	for ford.hasAugmentingPath(G, s, t) {
		// compute bottleneck capacity
		bottle := math.Inf(1)
		for v := t; v != s; v = ford.edgeTo[v].Other(v) {
			bottle = math.Min(bottle, ford.edgeTo[v].ResidualCapacityTo(v))
		}
		// augment flow
		for v := t; v != s; v = ford.edgeTo[v].Other(v) {
			// increase flow on forward edge(not full) or
			// decrease flow on backward edge(not empty)
			ford.edgeTo[v].AddResidualFlowTo(v, bottle)
		}
		ford.value += bottle
	}

	return ford
}

// Value returns the value of the maximum flow.
func (ford *FordFulkerson) Value() float64 {
	return ford.value
}

// InCut returns true if the specified vertex is on the s side of the mincut.
func (ford *FordFulkerson) InCut(v int) bool {
	ford.validate(v)
	return ford.marked[v]
}

// is there an augmenting path?
// if so, upon termination edgeTo[] will contain a parent-link representation of such a path
// this implementation finds a shortest augmenting path (fewest number of edges),
// which performs well both in theory and in practice
func (ford *FordFulkerson) hasAugmentingPath(G *FlowNetwork, s, t int) bool {
	ford.edgeTo = make([]*FlowEdge, G.V())
	ford.marked = make([]bool, G.V())

	queue := arrayqueue.New()
	queue.Enqueue(s)
	ford.marked[s] = true
	for !queue.IsEmpty() && !ford.marked[t] { // bfs search
		val, _ := queue.Dequeue()
		v := val.(int)
		for _, e := range G.Adj(v) {
			w := e.Other(v)
			if !ford.marked[w] && e.ResidualCapacityTo(w) > 0 {
				// found path from s to w in the residual network
				// ResidualCapacity(w) > 0 means either backward edge not empty or forward edge not full
				ford.edgeTo[w] = e
				ford.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}
	// ford.marked[t] == true: there is an augmenting path
	// ford.marked[t] == false: there is no more augmenting path,
	// which means all paths from s to t are blocked by either a
	// full forward edge or an empty backward edge.
	return ford.marked[t]
}

// return excess flow at vertex v
func (ford *FordFulkerson) excess(G *FlowNetwork, v int) float64 {
	excess := 0.0
	for _, e := range G.Adj(v) {
		if v == e.From() {
			excess -= e.Flow()
		} else {
			excess += e.Flow()
		}
	}
	return excess
}

func (ford *FordFulkerson) isFeasible(G *FlowNetwork, s, t int) bool {
	EPSILON := 1e-11

	// check that net flow into a vertex equals zero, except at source and sink
	if math.Abs(ford.value+ford.excess(G, s)) > EPSILON {
		fmt.Println("excess at source = ", ford.excess(G, s))
		fmt.Println("Max flow     = ", ford.value)
		return false
	}

	if math.Abs(ford.value-ford.excess(G, t)) > EPSILON {
		fmt.Println("excess at sink = ", ford.excess(G, t))
		fmt.Println("Max flow     = ", ford.value)
		return false
	}

	for v := 0; v < G.V(); v++ {
		if v == s || v == t {
			continue
		}
		if math.Abs(ford.excess(G, v)) > EPSILON {
			fmt.Println("Net flow out of ", v, " doesn't equal zero")
			return false
		}
	}
	return true
}

func (ford *FordFulkerson) validate(v int) {
	if v < 0 || v >= ford.v {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", ford.v-1))
	}
}
