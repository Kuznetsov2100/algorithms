package graph

import "fmt"

type CC struct {
	marked []bool
	id     []int
	size   []int
	count  int
}

func NewCC(G *Graph) *CC {
	cc := &CC{
		marked: make([]bool, G.V()),
		id:     make([]int, G.V()),
		size:   make([]int, G.V()),
	}
	for v := 0; v < G.V(); v++ {
		if !cc.marked[v] {
			cc.dfs(G, v)
			cc.count++
		}
	}
	return cc
}

func (cc *CC) dfs(G *Graph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	cc.size[cc.count]++
	for _, w := range G.Adj(v) {
		if !cc.marked[w] {
			cc.dfs(G, w)
		}
	}
}

func (cc *CC) Id(v int) int {
	cc.validateVertex(v)
	return cc.id[v]
}

func (cc *CC) Size(v int) int {
	cc.validateVertex(v)
	return cc.size[cc.id[v]]
}

func (cc *CC) Count() int {
	return cc.count
}

func (cc *CC) Connected(v, w int) bool {
	cc.validateVertex(v)
	cc.validateVertex(w)
	return cc.id[v] == cc.id[w]
}

func (cc *CC) validateVertex(v int) {
	length := len(cc.marked)
	if v < 0 || v >= length {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", length-1))
	}
}
