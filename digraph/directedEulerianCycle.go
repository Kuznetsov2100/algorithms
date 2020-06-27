package digraph

import (
	"github.com/emirpasic/gods/lists/arraylist"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// DirectedEulerianCycle struct represents a data type for finding an Eulerian cycle or path in a digraph.
// An Eulerian cycle is a cycle (not necessarily simple) that uses every edge in the digraph exactly once.
// This implementation uses a nonrecursive depth-first search. The constructor takes O(E + V) time
// in the worst case, where E is the number of edges and V is the number of vertices Each instance
// method takes O(1) time. It uses O(V) extra space (not including the digraph).
type DirectedEulerianCycle struct {
	cycle *arraystack.Stack
}

// NewDirectedEulerianCycle computes an Eulerian cycle in the specified digraph, if one exists.
func NewDirectedEulerianCycle(G *Digraph) *DirectedEulerianCycle {
	dec := &DirectedEulerianCycle{}
	// must have at least one edge
	if G.E() == 0 {
		return dec
	}

	// necessary condition: indegree(v) = outdegree(v) for each vertex v
	// (without this check, DFS might return a path instead of a cycle)
	for v := 0; v < G.V(); v++ {
		if G.OutDegree(v) != G.InDegree(v) {
			return dec
		}
	}

	// create local view of adjacency lists, to iterate one vertex at a time
	adj := make([]arraylist.Iterator, G.V())
	for v := 0; v < G.V(); v++ {
		vertices := arraylist.New()
		for _, w := range G.Adj(v) {
			vertices.Add(w)
		}
		adj[v] = vertices.Iterator()
	}

	// initialize stack with any non-isolated vertex
	s := nonIsolatedVertex(G)
	stack := arraystack.New()
	stack.Push(s)

	// greedily add to putative cycle, depth-first search style
	dec.cycle = arraystack.New()
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		v := val.(int)
		for adj[v].Next() {
			stack.Push(v)
			v = adj[v].Value().(int)
		}
		dec.cycle.Push(v)
	}

	if dec.cycle.Size() != G.E()+1 {
		dec.cycle = nil
	}

	return dec
}

// GetCycle returns the sequence of vertices on an Eulerian cycle.
func (dec *DirectedEulerianCycle) GetCycle() (cy []int) {
	if !dec.HasEulerianCycle() {
		return nil
	}
	for _, val := range dec.cycle.Values() {
		cy = append(cy, val.(int))
	}
	return cy
}

// HasEulerianCycle returns true if the digraph has an Eulerian cycle.
func (dec *DirectedEulerianCycle) HasEulerianCycle() bool {
	return dec.cycle != nil
}

func nonIsolatedVertex(G *Digraph) int {
	s := -1
	for v := 0; v < G.V(); v++ {
		if G.OutDegree(v) > 0 {
			s = v
			break
		}
	}
	return s
}
