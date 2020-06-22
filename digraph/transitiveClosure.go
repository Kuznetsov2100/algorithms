package digraph

import "fmt"

// TransitiveClosure struct represents a data type for computing the transitive closure of a digraph.
// This implementation runs depth-first search from each vertex.
// The constructor takes O(V(V + E)) in the worst case, where V is the number of vertices
// and E is the number of edges. Each instance method takes O(1) time. It uses O(V2) extra
// space (not including the digraph).
type TransitiveClosure struct {
	tc []*DirectedDFS
}

// NewTransitiveClosure computes the transitive closure of the digraph G.
func NewTransitiveClosure(G *Digraph) *TransitiveClosure {
	tc := make([]*DirectedDFS, G.V())
	for v := 0; v < G.V(); v++ {
		tc[v] = NewDirectedDFS(G, v)
	}
	return &TransitiveClosure{tc: tc}
}

// Reachable returns true if there is a directed path from vertex v to vertex w in the digraph
func (t *TransitiveClosure) Reachable(v, w int) bool {
	t.validateVertex(v)
	t.validateVertex(w)
	return t.tc[v].IsMarked(w)
}

func (t *TransitiveClosure) validateVertex(v int) {
	V := len(t.tc)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
