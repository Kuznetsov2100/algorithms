package graph

import "github.com/handane123/algorithms/dataStructure/stack/arraystack"

// Cycle struct represents a data type for determining whether an undirected graph has a simple cycle.
// This implementation uses depth-first search. The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges. (The depth-first search part takes only O(V) time;
// however, checking for self-loops and parallel edges takes O(V + E) time in the worst case.)
// Each instance method takes O(1) time. It uses O(V) extra space (not including the graph).
type Cycle struct {
	marked []bool
	edgeTo []int
	cycle  *arraystack.Stack
}

// NewCycle determines whether the undirected graph G has a cycle and, if so, finds such a cycle.
func NewCycle(G *Graph) *Cycle {
	c := &Cycle{}
	if c.hasSelfLoop(G) {
		return c
	}
	if c.hasParallelEdges(G) {
		return c
	}
	c.marked = make([]bool, G.V())
	c.edgeTo = make([]int, G.V())
	for v := 0; v < G.V(); v++ {
		if !c.marked[v] {
			c.dfs(G, -1, v)
		}
	}
	return c
}

// HasCycle returns true if the graph G has a cycle.
func (c *Cycle) HasCycle() bool {
	return c.cycle != nil
}

// GetCycle returns a cycle in the graph G.
func (c *Cycle) GetCycle() (cy []int) {
	for _, val := range c.cycle.Values() {
		cy = append(cy, val.(int))
	}
	return cy
}

func (c *Cycle) dfs(G *Graph, u, v int) {
	c.marked[v] = true
	for _, w := range G.Adj(v) {
		if c.cycle != nil {
			return
		}
		if !c.marked[w] {
			c.edgeTo[w] = v
			c.dfs(G, v, w)
		} else if w != u {
			c.cycle = arraystack.New()
			for x := v; x != w; x = c.edgeTo[x] {
				c.cycle.Push(x)
			}
			c.cycle.Push(w)
			c.cycle.Push(v)
		}
	}
}

func (c *Cycle) hasParallelEdges(G *Graph) bool {
	c.marked = make([]bool, G.V())

	for v := 0; v < G.V(); v++ {
		for _, w := range G.Adj(v) {
			if c.marked[w] {
				c.cycle = arraystack.New()
				c.cycle.Push(v)
				c.cycle.Push(w)
				c.cycle.Push(v)
				return true
			}
			c.marked[w] = true
		}

		for _, w := range G.Adj(v) {
			c.marked[w] = false
		}
	}
	return false
}

func (c *Cycle) hasSelfLoop(G *Graph) bool {
	for v := 0; v < G.V(); v++ {
		for _, w := range G.Adj(v) {
			if v == w {
				c.cycle = arraystack.New()
				c.cycle.Push(v)
				c.cycle.Push(v)
				return true
			}
		}
	}
	return false
}
