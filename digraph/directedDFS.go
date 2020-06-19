package digraph

import "fmt"

// DirectedDFS struct represents a data type for determining the vertices
// reachable from a given source vertex s (or set of source vertices) in a digraph.
// This implementation uses depth-first search.
// The constructor takes time proportional to V + E (in the worst case),
// where V is the number of vertices and E is the number of edges. Each instance method takes O(1) time.
// It uses O(V) extra space (not including the digraph).
type DirectedDFS struct {
	marked []bool
	count  int
}

// NewDirectedDFS computes the vertices in digraph G that are reachable from the source vertex s.
func NewDirectedDFS(G *Digraph, s int) *DirectedDFS {
	d := &DirectedDFS{marked: make([]bool, G.V()), count: 0}
	d.validateVertex(s)
	d.dfs(G, s)
	return d
}

// NewDirectedDFSources computes the vertices in digraph G that are connected to any of the source vertices sources.
func NewDirectedDFSources(G *Digraph, sources []int) *DirectedDFS {
	d := &DirectedDFS{marked: make([]bool, G.V()), count: 0}
	d.validateVertices(sources)
	for _, v := range sources {
		if !d.marked[v] {
			d.dfs(G, v)
		}
	}
	return d
}

func (d *DirectedDFS) dfs(G *Digraph, v int) {
	d.count++
	d.marked[v] = true
	for _, w := range G.Adj(v) {
		if !d.marked[w] {
			d.dfs(G, w)
		}
	}
}

// IsMarked returns true if there is a directed path from the source vertex (or any of the source vertices) and vertex v
func (d *DirectedDFS) IsMarked(v int) bool {
	d.validateVertex(v)
	return d.marked[v]
}

// Count returns the number of vertices reachable from the source vertex (or source vertices).
func (d *DirectedDFS) Count() int {
	return d.count
}

func (d *DirectedDFS) validateVertex(v int) {
	V := len(d.marked)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}

func (d *DirectedDFS) validateVertices(vertices []int) {
	if vertices == nil {
		panic("argument is nil")
	}
	for _, v := range vertices {
		d.validateVertex(v)
	}
}
