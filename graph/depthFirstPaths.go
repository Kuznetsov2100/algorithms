package graph

import (
	"fmt"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

type DepthFirstPaths struct {
	marked []bool
	edgeTo []int
	source int
}

func NewDepthFirstPaths(G *Graph, s int) *DepthFirstPaths {
	dfp := &DepthFirstPaths{
		marked: make([]bool, G.V()),
		edgeTo: make([]int, G.V()),
		source: s}
	dfp.validateVertex(s)
	dfp.dfs(G, s)
	return dfp
}

func (dfp *DepthFirstPaths) dfs(G *Graph, v int) {
	dfp.marked[v] = true
	for _, w := range G.Adj(v) {
		if !dfp.marked[w] {
			dfp.edgeTo[w] = v
			dfp.dfs(G, w)
		}
	}
}

func (dfp *DepthFirstPaths) HasPathTo(v int) bool {
	dfp.validateVertex(v)
	return dfp.marked[v]
}

func (dfp *DepthFirstPaths) PathTo(v int) (p []int) {
	dfp.validateVertex(v)
	if !dfp.HasPathTo(v) {
		return nil
	}
	path := arraystack.New()
	for x := v; x != dfp.source; x = dfp.edgeTo[x] {
		path.Push(x)
	}
	path.Push(dfp.source)
	for _, val := range path.Values() {
		p = append(p, val.(int))
	}
	return p
}

func (dfp *DepthFirstPaths) validateVertex(v int) {
	length := len(dfp.marked)
	if v < 0 || v >= length {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", length-1))
	}
}
