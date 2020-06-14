package graph

import (
	"fmt"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// DepthFirstPaths struct represents a data type for finding paths
// from a source vertex s to every other vertex in an undirected graph.
// This implementation uses depth-first search.
// The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges.
// Each instance method takes O(1) time. It uses O(V) extra space (not including the graph).
type DepthFirstPaths struct {
	marked []bool
	edgeTo []int
	source int
}

// NewDepthFirstPaths computes a path between s and every other vertex in graph G.
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

// HasPathTo returns true if there is a path between the source vertex s and vertex v
func (dfp *DepthFirstPaths) HasPathTo(v int) bool {
	dfp.validateVertex(v)
	return dfp.marked[v]
}

// PathTo returns a path between the source vertex s and vertex v, or nil if no such path.
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
