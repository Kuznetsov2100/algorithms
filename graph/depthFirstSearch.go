package graph

import "fmt"

type DepthFirstSearch struct {
	marked []bool
	count  int
}

func NewDepthFirstSearch(G *Graph, s int) *DepthFirstSearch {
	search := &DepthFirstSearch{marked: make([]bool, G.V())}
	search.validateVertex(s)
	search.dfs(G, s)
	return search
}

func (search *DepthFirstSearch) dfs(G *Graph, v int) {
	search.count++
	search.marked[v] = true
	for _, w := range G.Adj(v) {
		if !search.marked[w] {
			search.dfs(G, w)
		}
	}
}
func (search *DepthFirstSearch) IsMarked(v int) bool {
	search.validateVertex(v)
	return search.marked[v]
}

func (search *DepthFirstSearch) validateVertex(v int) {
	length := len(search.marked)
	if v < 0 || v >= length {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", length-1))
	}
}

func (search *DepthFirstSearch) Count() int {
	return search.count
}
