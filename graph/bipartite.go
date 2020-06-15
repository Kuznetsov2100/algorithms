package graph

import (
	"fmt"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

type Bipartite struct {
	isBipartite bool
	color       []bool
	marked      []bool
	edgeTo      []int
	cycle       *arraystack.Stack
}

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

func (b *Bipartite) IsBipartite() bool {
	return b.isBipartite
}

func (b *Bipartite) Color(v int) bool {
	b.validateVertex(v)
	if !b.isBipartite {
		panic("graph is not  bipartite")
	}
	return b.color[v]
}

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
