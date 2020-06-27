package digraph

import "github.com/handane123/algorithms/dataStructure/stack/arraystack"

// EdgeWeightedDirectedCycle struct represents a data type for determining whether
// an edge-weighted digraph has a directed cycle. The hasCycle operation determines whether
// the edge-weighted digraph has a directed cycle and, if so, the cycle operation returns one.
// This implementation uses depth-first search. The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the edge-weighted digraph).
type EdgeWeightedDirectedCycle struct {
	marked  []bool            // marked[v] = has vertex v been marked?
	edgeTo  []*DirectedEdge   // edgeTo[v] = previous edge on path to v
	onStack []bool            // onStack[v] = is vertex on the stack?
	cycle   *arraystack.Stack // directed cycle (or nil if no such cycle)
}

// NewEdgeWeightedDirectedCycle determines whether the edge-weighted digraph G has a directed cycle and, if so, finds such a cycle.
func NewEdgeWeightedDirectedCycle(G *EdgeWeightedDigraph) *EdgeWeightedDirectedCycle {
	c := &EdgeWeightedDirectedCycle{
		marked:  make([]bool, G.V()),
		edgeTo:  make([]*DirectedEdge, G.V()),
		onStack: make([]bool, G.V()),
	}

	for v := 0; v < G.V(); v++ {
		if !c.marked[v] {
			c.dfs(G, v)
		}
	}
	return c
}

// HasCycle returns true if the digraph G has a cycle.
func (c *EdgeWeightedDirectedCycle) HasCycle() bool {
	return c.cycle != nil
}

// GetCycle returns a directed cycle if the digraph has a directed cycle, and nil otherwise.
func (c *EdgeWeightedDirectedCycle) GetCycle() (cy []*DirectedEdge) {
	if !c.HasCycle() {
		return nil
	}
	for _, val := range c.cycle.Values() {
		cy = append(cy, val.(*DirectedEdge))
	}
	return cy
}

func (c *EdgeWeightedDirectedCycle) dfs(G *EdgeWeightedDigraph, v int) {
	c.onStack[v] = true
	c.marked[v] = true
	for _, e := range G.Adj(v) {
		w := e.To()
		// short circuit if cycle already found
		if c.cycle != nil {
			return
		}
		if !c.marked[w] { // found new vertex, so recursive
			c.edgeTo[w] = e
			c.dfs(G, w)
		} else if c.onStack[w] { // trace back directed cycle
			c.cycle = arraystack.New()
			f := e
			for f.From() != w {
				c.cycle.Push(f)
				f = c.edgeTo[f.From()]
			}
			c.cycle.Push(f)
			return
		}
	}
	c.onStack[v] = false
}
