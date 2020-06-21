package digraph

import "fmt"

// KosarajuSharirSCC struct represents a data type for determining the strong components in a digraph.
// The id operation determines in which strong component a given vertex lies;
// the areStronglyConnected operation determines whether two vertices are in the same strong component;
// and the count operation determines the number of strong components.
// The component identifier of a component is one of the vertices in the strong component:
// two vertices have the same component identifier if and only if they are in the same strong component.
// This implementation uses the Kosaraju-Sharir algorithm. The constructor takes O(V + E) time,
// where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the digraph).
type KosarajuSharirSCC struct {
	marked []bool
	id     []int
	count  int
}

// NewKosarajuSharirSCC computes the strong components of the digraph G.
func NewKosarajuSharirSCC(G *Digraph) *KosarajuSharirSCC {
	scc := &KosarajuSharirSCC{
		marked: make([]bool, G.V()),
		id:     make([]int, G.V()),
	}

	dfo := NewDepthFirstOrder(G.Reverse())
	for _, v := range dfo.ReversePost() {
		if !scc.marked[v] {
			scc.dfs(G, v)
			scc.count++
		}
	}
	return scc
}

func (scc *KosarajuSharirSCC) dfs(G *Digraph, v int) {
	scc.marked[v] = true
	scc.id[v] = scc.count
	for _, w := range G.Adj(v) {
		if !scc.marked[w] {
			scc.dfs(G, w)
		}
	}
}

// Count returns the number of strong components.
func (scc *KosarajuSharirSCC) Count() int {
	return scc.count
}

// StronglyConnected returns true if vertices v and w in the same strong component
func (scc *KosarajuSharirSCC) StronglyConnected(v, w int) bool {
	scc.validateVertex(v)
	scc.validateVertex(w)
	return scc.id[v] == scc.id[w]
}

// Id returns the component id of the strong component containing vertex v.
func (scc *KosarajuSharirSCC) Id(v int) int {
	scc.validateVertex(v)
	return scc.id[v]
}

func (scc *KosarajuSharirSCC) validateVertex(v int) {
	V := len(scc.marked)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
