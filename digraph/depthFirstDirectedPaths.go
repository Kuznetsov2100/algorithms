package digraph

import (
	"fmt"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// DepthFirstDirectedPaths struct represents a data type for finding directed paths
// from a source vertex s to every other vertex in the digraph.
// This implementation uses depth-first search. The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the digraph).
type DepthFirstDirectedPaths struct {
	marked []bool // marked[v] = is there an s-v path?
	edgeTo []int  // edgeTo[v] = last edge on s-v path
	source int    // source vertex
}

// NewDepthFirstPaths computes a directed path from s to every other vertex in digraph G.
func NewDepthFirstDirectedPaths(G *Digraph, s int) *DepthFirstDirectedPaths {
	dfp := &DepthFirstDirectedPaths{
		marked: make([]bool, G.V()),
		edgeTo: make([]int, G.V()),
		source: s}
	dfp.validateVertex(s)
	dfp.dfs(G, s)
	return dfp
}

func (dfp *DepthFirstDirectedPaths) dfs(G *Digraph, v int) {
	dfp.marked[v] = true
	for _, w := range G.Adj(v) {
		if !dfp.marked[w] {
			dfp.edgeTo[w] = v
			dfp.dfs(G, w)
		}
	}
}

// HasPathTo returns true if there is a directed path from the source vertex s to vertex v
func (dfp *DepthFirstDirectedPaths) HasPathTo(v int) bool {
	dfp.validateVertex(v)
	return dfp.marked[v]
}

// PathTo returns a directed path from the source vertex s to vertex v, or nil if no such path.
func (dfp *DepthFirstDirectedPaths) PathTo(v int) (p []int) {
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

func (dfp *DepthFirstDirectedPaths) validateVertex(v int) {
	length := len(dfp.marked)
	if v < 0 || v >= length {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", length-1))
	}
}
