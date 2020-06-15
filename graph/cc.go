package graph

import "fmt"

// The CC struct represents a data type for determining the connected components in an undirected graph.
// This implementation uses depth-first search. The constructor takes O(V + E) time,
// where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the graph).
type CC struct {
	marked []bool // marked[v] = has vertex v been marked?
	id     []int  // id[v] = id of connected component containing v
	size   []int  // size[id] = number of vertices in given component
	count  int    // number of connected components
}

// NewCC computes the connected components of the undirected graph G
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

// Id returns the component id of the connected component containing vertex v.
func (cc *CC) Id(v int) int {
	cc.validateVertex(v)
	return cc.id[v]
}

// Size returns the component id of the connected component containing vertex v.
func (cc *CC) Size(v int) int {
	cc.validateVertex(v)
	return cc.size[cc.id[v]]
}

// Count returns the number of connected components in the graph G.
func (cc *CC) Count() int {
	return cc.count
}

// Connected returns true if vertices v and w are in the same connected component.
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
