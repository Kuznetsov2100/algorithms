package digraph

import "fmt"

type DirectedDFS struct {
	marked []bool
	count  int
}

func NewDirectedDFS(G *Digraph, s int) *DirectedDFS {
	d := &DirectedDFS{marked: make([]bool, G.V()), count: 0}
	d.validateVertex(s)
	d.dfs(G, s)
	return d
}

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

func (d *DirectedDFS) IsMarked(v int) bool {
	d.validateVertex(v)
	return d.marked[v]
}

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
