package digraph

import "fmt"

// Topological struct represents a data type for determining a topological order
// of a directed acyclic graph (DAG). A digraph has a topological order if and only if it is a DAG.
// The hasOrder operation determines whether the digraph has a topological order, and if so,
// the order operation returns one.
// This implementation uses depth-first search. The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the digraph).
type Topological struct {
	order []int // topological order
	rank  []int // rank[v] = rank of vertex v in order
}

// NewTopological determines whether the digraph G has a topological order and, if so,
// finds such a topological order.
func NewTopological(G *Digraph) *Topological {
	top := &Topological{rank: make([]int, G.V())}
	finder := NewDirectedCycle(G)
	if !finder.HasCycle() {
		dfo := NewDepthFirstOrder(G)
		top.order = dfo.ReversePost()
		for i, v := range top.order {
			top.rank[v] = i
		}
	}
	return top
}

func (top *Topological) Order() (order []int) {
	order = make([]int, len(top.order))
	copy(order, top.order)
	return order
}

// HasOrder returns true if the digraph have a topological order
func (top *Topological) HasOrder() bool {
	return top.order != nil
}

// Rank returns the the rank of vertex v in the topological order; -1 if the digraph is not a DAG
func (top *Topological) Rank(v int) int {
	top.validateVertex(v)
	if top.HasOrder() {
		return top.rank[v]
	}
	return -1
}

func (top *Topological) validateVertex(v int) {
	V := len(top.rank)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
