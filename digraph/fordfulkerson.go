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
	maxflow := &FordFulkerson{v: G.V()}
	maxflow.validate(s)
	maxflow.validate(t)
	if s == t {
		panic("source equals sink")
	}
	if !maxflow.isFeasible(G, s, t) {
		panic("initial flow is infeasible")
	}

	// while there exists an augmenting path, use it
	maxflow.value = maxflow.excess(G, t)
	for maxflow.hasAugmentingPath(G, s, t) {
		// compute bottleneck capacity
		bottle := math.Inf(1)
		for v := t; v != s; v = maxflow.edgeTo[v].Other(v) {
			bottle = math.Min(bottle, maxflow.edgeTo[v].ResidualCapacityTo(v))
		}
		// augment flow
		for v := t; v != s; v = maxflow.edgeTo[v].Other(v) {
			// increase flow on forward edge(not full) or
			// decrease flow on backward edge(not empty)
			maxflow.edgeTo[v].AddResidualFlowTo(v, bottle)
		}
		maxflow.value += bottle
	}

	return maxflow
}

// Value returns the value of the maximum flow.
func (maxflow *FordFulkerson) Value() float64 {
	return maxflow.value
}

// InCut returns true if the specified vertex is on the s side of the mincut.
func (maxflow *FordFulkerson) InCut(v int) bool {
	maxflow.validate(v)
	return maxflow.marked[v]
}

// is there an augmenting path?
// if so, upon termination edgeTo[] will contain a parent-link representation of such a path
// this implementation finds a shortest augmenting path (fewest number of edges),
// which performs well both in theory and in practice
func (maxflow *FordFulkerson) hasAugmentingPath(G *FlowNetwork, s, t int) bool {
	maxflow.edgeTo = make([]*FlowEdge, G.V())
	maxflow.marked = make([]bool, G.V())

	queue := arrayqueue.New()
	queue.Enqueue(s)
	maxflow.marked[s] = true
	for !queue.IsEmpty() && !maxflow.marked[t] { // bfs search
		val, _ := queue.Dequeue()
		v := val.(int)
		for _, e := range G.Adj(v) {
			w := e.Other(v)
			if !maxflow.marked[w] && e.ResidualCapacityTo(w) > 0 {
				// found path from s to w in the residual network
				// ResidualCapacity(w) > 0 means either backward edge not empty or forward edge not full
				maxflow.edgeTo[w] = e
				maxflow.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}
	// maxflow.marked[t] == true: there is an augmenting path
	// maxflow.marked[t] == false: there is no more augmenting path,
	// which means all paths from s to t are blocked by either a
	// full forward edge or an empty backward edge.
	return maxflow.marked[t]
}

// return excess flow at vertex v
func (maxflow *FordFulkerson) excess(G *FlowNetwork, v int) float64 {
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

func (maxflow *FordFulkerson) isFeasible(G *FlowNetwork, s, t int) bool {
	EPSILON := 1e-11

	// check that net flow into a vertex equals zero, except at source and sink
	if math.Abs(maxflow.value+maxflow.excess(G, s)) > EPSILON {
		fmt.Println("excess at source = ", maxflow.excess(G, s))
		fmt.Println("Max flow     = ", maxflow.value)
		return false
	}

	if math.Abs(maxflow.value-maxflow.excess(G, t)) > EPSILON {
		fmt.Println("excess at sink = ", maxflow.excess(G, t))
		fmt.Println("Max flow     = ", maxflow.value)
		return false
	}

	for v := 0; v < G.V(); v++ {
		if v == s || v == t {
			continue
		}
		if math.Abs(maxflow.excess(G, v)) > EPSILON {
			fmt.Println("Net flow out of ", v, " doesn't equal zero")
			return false
		}
	}
	return true
}

func (maxflow *FordFulkerson) validate(v int) {
	if v < 0 || v >= maxflow.v {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", maxflow.v-1))
	}
}
