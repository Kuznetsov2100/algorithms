package graph

import (
	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// EulerianCycle struct represents a data type for finding an Eulerian cycle or path in a graph.
// An Eulerian cycle is a cycle (not necessarily simple) that uses every edge in the graph exactly once.
// This implementation uses a nonrecursive depth-first search.
// The constructor takes O(E + V) time in the worst case,
// where E is the number of edges and V is the number of vertices Each instance method takes O(1) time.
// It uses O(E + V) extra space in the worst case (not including the graph).
type EulerianCycle struct {
	cycle *arraystack.Stack
}

// an undirected edge, with a field to indicate whether the edge has already been used
type ecEdge struct {
	v      int
	w      int
	isUsed bool
}

func newecEdge(v, w int) *ecEdge {
	return &ecEdge{
		v:      v,
		w:      w,
		isUsed: false,
	}
}

// returns the other vertex of the edge
func (e *ecEdge) other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	if vertex == e.w {
		return e.v
	}
	panic("Illegal endpoint")
}

// NewEuerianCycle computes an Eulerian cycle in the specified graph, if one exists
func NewEulerianCycle(G *Graph) *EulerianCycle {
	// must have at least one edge
	if G.E() == 0 {
		return nil
	}
	// necessary condition: all vertices have even degree
	// (this test is needed or it might find an Eulerian path instead of cycle)
	for v := 0; v < G.V(); v++ {
		if G.Degree(v)%2 != 0 {
			return nil
		}
	}

	// create local view of adjacency lists, to iterate one vertex at a time
	// the helper Edge data type is used to avoid exploring both copies of an edge v-w
	adj := make([]*arrayqueue.Queue, G.V())
	for v := 0; v < G.V(); v++ {
		adj[v] = arrayqueue.New()
	}
	for v := 0; v < G.V(); v++ {
		selfLoops := 0
		for _, w := range G.Adj(v) {
			// careful with self loops
			if v == w {
				if selfLoops%2 == 0 {
					e := newecEdge(v, w)
					adj[v].Enqueue(e)
					adj[w].Enqueue(e)
				}
				selfLoops++
			} else if v < w {
				e := newecEdge(v, w)
				adj[v].Enqueue(e)
				adj[w].Enqueue(e)
			}
		}
	}

	// initialize stack with any non-isolated vertex
	s := nonIsolatedVertex(G)
	stack := arraystack.New()
	stack.Push(s)

	// greedily search through edges in iterative DFS style
	ec := &EulerianCycle{cycle: arraystack.New()}
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		v := val.(int)
		for !adj[v].IsEmpty() {
			e, _ := adj[v].Dequeue()
			edge := e.(*ecEdge)
			if edge.isUsed {
				continue
			}
			edge.isUsed = true
			stack.Push(v)
			v = edge.other(v)
		}
		// push vertex with no more leaving edges to cycle
		ec.cycle.Push(v)
	}
	if ec.cycle.Size() != G.E()+1 {
		ec.cycle = nil
	}
	return ec
}

// GetCycle returns the sequence of vertices on an Eulerian cycle.
func (ec *EulerianCycle) GetCycle() (cy []int) {
	for _, val := range ec.cycle.Values() {
		cy = append(cy, val.(int))
	}
	return cy
}

// HasEulerianCycle returns true if the graph has an Eulerian cycle.
func (ec *EulerianCycle) HasEulerianCycle() bool {
	return ec.cycle != nil
}

// returns any non-isolated vertex; -1 if no such vertex
func nonIsolatedVertex(G *Graph) int {
	for v := 0; v < G.V(); v++ {
		if G.Degree(v) > 0 {
			return v
		}
	}
	return -1
}
