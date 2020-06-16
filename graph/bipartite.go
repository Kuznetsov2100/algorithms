package graph

import (
	"fmt"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// Bipartite struct represents a data type for determining whether an undirected graph is bipartite
// or whether it has an odd-length cycle.
// This implementation uses depth-first search. The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges.
// Each instance method takes O(1) time. It uses O(V) extra space (not including the graph).
type Bipartite struct {
	isBipartite bool
	color       []bool
	marked      []bool
	edgeTo      []int
	cycle       *arraystack.Stack
}

// NewBipartite determines whether an undirected graph is bipartite and finds either a bipartition or an odd-length cycle.
func NewBipartite(G *Graph) *Bipartite {
	b := &Bipartite{
		isBipartite: true,
		color:       make([]bool, G.V()),
		marked:      make([]bool, G.V()),
		edgeTo:      make([]int, G.V()),
	}

	for v := 0; v < G.V(); v++ {
		if !b.marked[v] {
			b.dfs(G, v)
		}
	}
	return b
}

func (b *Bipartite) dfs(G *Graph, v int) {
	b.marked[v] = true
	for _, w := range G.Adj(v) {
		if b.cycle != nil {
			return
		}
		if !b.marked[w] {
			b.edgeTo[w] = v
			b.color[w] = !b.color[v]
			b.dfs(G, w)
		} else if b.color[w] == b.color[v] {
			b.isBipartite = false
			b.cycle = arraystack.New()
			b.cycle.Push(w)
			for x := v; x != w; x = b.edgeTo[x] {
				b.cycle.Push(x)
			}
			b.cycle.Push(w)
		}
	}
}

// IsBipartite returns true if the graph is bipartite.
func (b *Bipartite) IsBipartite() bool {
	return b.isBipartite
}

// Color returns the side of the bipartite that vertex v is on.
func (b *Bipartite) Color(v int) bool {
	b.validateVertex(v)
	if !b.isBipartite {
		panic("graph is not  bipartite")
	}
	return b.color[v]
}

// OddCycle returns an odd-length cycle if the graph is not bipartite, and nil otherwise.
func (b *Bipartite) OddCycle() (cy []int) {
	for _, val := range b.cycle.Values() {
		cy = append(cy, val.(int))
	}
	return cy
}

func (b *Bipartite) validateVertex(v int) {
	length := len(b.marked)
	if v < 0 || v >= length {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", length-1))
	}
}
