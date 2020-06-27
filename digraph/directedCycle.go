package digraph

import "github.com/handane123/algorithms/dataStructure/stack/arraystack"

// DirectedCycle struct represents a data type for determining whether a digraph has a directed cycle.
// The hasCycle operation determines whether the digraph has a simple directed cycle and,
// if so, the cycle operation returns one.
// This implementation uses depth-first search. The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the digraph).
type DirectedCycle struct {
	marked  []bool
	edgeTo  []int
	onStack []bool
	cycle   *arraystack.Stack
}

// NewDirectedCycle determines whether the digraph G has a directed cycle and, if so, finds such a cycle.
func NewDirectedCycle(G *Digraph) *DirectedCycle {
	c := &DirectedCycle{
		marked:  make([]bool, G.V()),
		edgeTo:  make([]int, G.V()),
		onStack: make([]bool, G.V()),
	}

	for v := 0; v < G.V(); v++ {
		if !c.marked[v] && c.cycle == nil {
			c.dfs(G, v)
		}
	}
	return c
}

// HasCycle returns true if the digraph G has a cycle.
func (c *DirectedCycle) HasCycle() bool {
	return c.cycle != nil
}

// GetCycle returns a directed cycle if the digraph has a directed cycle, and nil otherwise.
func (c *DirectedCycle) GetCycle() (cy []int) {
	if !c.HasCycle() {
		return nil
	}
	for _, val := range c.cycle.Values() {
		cy = append(cy, val.(int))
	}
	return cy
}

func (c *DirectedCycle) dfs(G *Digraph, v int) {
	c.onStack[v] = true
	c.marked[v] = true
	for _, w := range G.Adj(v) {
		// short circuit if cycle already found
		if c.cycle != nil {
			return
		}
		if !c.marked[w] {
			c.edgeTo[w] = v
			c.dfs(G, w)
		} else if c.onStack[w] { // trace back directed cycle
			c.cycle = arraystack.New()
			for x := v; x != w; x = c.edgeTo[x] {
				c.cycle.Push(x)
			}
			c.cycle.Push(w)
			c.cycle.Push(v)
		}
	}
	c.onStack[v] = false
}
