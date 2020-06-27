package digraph

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// DirectedEulerianPath struct represents a data type for finding an Eulerian path in a digraph.
// An Eulerian path is a path (not necessarily simple) that uses every edge in the digraph exactly once.
// This implementation uses a nonrecursive depth-first search. The constructor take O(E + V) time
// in the worst case, where E is the number of edges and V is the number of vertices. It uses O(V) extra
// space (not including the digraph).
type DirectedEulerianPath struct {
	path *arraystack.Stack
}

// NewDirectedEulerianPath computes an Eulerian path in the specified digraph, if one exists.
func NewDirectedEulerianPath(G *Digraph) *DirectedEulerianPath {
	dep := &DirectedEulerianPath{}

	// find vertex from which to start potential Eulerian path:
	// a vertex v with outdegree(v) > indegree(v) if it exits;
	// otherwise a vertex with outdegree(v) > 0
	deficit := 0
	s := nonIsolatedVertex(G)
	for v := 0; v < G.V(); v++ {
		if G.OutDegree(v) > G.InDegree(v) {
			deficit += G.OutDegree(v) - G.InDegree(v)
			s = v
		}
	}

	// digraph can't have an Eulerian path
	// (this condition is needed)
	if deficit > 1 {
		return dep
	}

	// special case for digraph with zero edges (has a degenerate Eulerian path)
	if s == -1 {
		s = 0
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

	stack := arraystack.New()
	stack.Push(s)

	// greedily add to putative cycle, depth-first search style
	dep.path = arraystack.New()
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		v := val.(int)
		for adj[v].Next() {
			stack.Push(v)
			v = adj[v].Value().(int)
		}
		// push vertex with no more available edges to path
		dep.path.Push(v)
	}

	// check if all edges have been used
	if dep.path.Size() != G.E()+1 {
		dep.path = nil
	}

	return dep
}

// Path returns the sequence of vertices on an Eulerian path.
func (dep *DirectedEulerianPath) Path() (p []int) {
	if !dep.HasEulerianPath() {
		return nil
	}
	for _, val := range dep.path.Values() {
		p = append(p, val.(int))
	}
	return p
}

// HasEulerianPath returns true if the digraph has an Eulerian path.
func (dep *DirectedEulerianPath) HasEulerianPath() bool {
	return dep.path != nil
}
