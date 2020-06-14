package graph

import "fmt"

// DepthFirstSearch struct represents a data type for determining the vertices connected to a given source vertex s in an undirected graph.
// This implementation uses depth-first search.The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges.
// Each instance method takes O(1) time. It uses O(V) extra space (not including the graph).
type DepthFirstSearch struct {
	marked []bool
	count  int
}

// NewDepthFirstSearch computes the vertices in graph G that are connected to the source vertex s.
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

// IsMarked returns true if there is a path between the source vertex s and vertex v
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

// Count returns the number of vertices connected to the source vertex s.
func (search *DepthFirstSearch) Count() int {
	return search.count
}
