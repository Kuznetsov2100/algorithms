package digraph

import "fmt"

type Topological struct {
	order []int
	rank  []int
}

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

func (top *Topological) HasOrder() bool {
	return top.order != nil
}

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
