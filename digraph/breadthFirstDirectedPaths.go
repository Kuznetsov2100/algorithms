package digraph

import (
	"fmt"
	"math"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

const infinity = math.MaxInt64

// BreadthFirstDirectedPaths struct represents a data type for finding shortest
// paths (number of edges) from a source vertex s (or a set of source vertices)
// to every other vertex in a digraph.
// This implementation uses breadth-first search. The constructor takes O(V + E) time in the worst case,
// where V is the number of vertices and E is the number of edges.
// Each instance method takes O(1) time. It uses O(V) extra space (not including the graph).
type BreadthFirstDirectedPaths struct {
	marked []bool // marked[v] = is there an s-v path
	edgeTo []int  // edgeTo[v] = previous edge on shortest s-v path
	distTo []int  // distTo[v] = number of edges shortest s-v path
}

// NewBreadthFirstDirectedPaths computes the shortest path between the source vertex s and every other vertex in the graph G.
func NewBreadthFirstDirectedPaths(G *Digraph, s int) *BreadthFirstDirectedPaths {
	bfp := &BreadthFirstDirectedPaths{
		marked: make([]bool, G.V()),
		edgeTo: make([]int, G.V()),
		distTo: make([]int, G.V()),
	}
	for v := 0; v < G.V(); v++ {
		bfp.distTo[v] = infinity
	}
	bfp.validateVertex(s)
	bfp.bfs(G, s)
	return bfp
}

func NewBreadthFirstDirectedPathSources(G *Digraph, sources []int) *BreadthFirstDirectedPaths {
	bfp := &BreadthFirstDirectedPaths{
		marked: make([]bool, G.V()),
		edgeTo: make([]int, G.V()),
		distTo: make([]int, G.V()),
	}
	for v := 0; v < G.V(); v++ {
		bfp.distTo[v] = infinity
	}
	bfp.validateVertices(sources)
	bfp.bfsSources(G, sources)
	return bfp
}

func (bfp *BreadthFirstDirectedPaths) bfsSources(G *Digraph, sources []int) {
	q := arrayqueue.New()
	for _, s := range sources {
		bfp.marked[s] = true
		bfp.distTo[s] = 0
		q.Enqueue(s)
	}

	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		v := val.(int)
		for _, w := range G.Adj(v) {
			if !bfp.marked[w] {
				bfp.edgeTo[w] = v
				bfp.distTo[w] = bfp.distTo[v] + 1
				bfp.marked[w] = true
				q.Enqueue(w)
			}
		}
	}
}

func (bfp *BreadthFirstDirectedPaths) bfs(G *Digraph, s int) {
	q := arrayqueue.New()
	bfp.distTo[s] = 0
	bfp.marked[s] = true
	q.Enqueue(s)

	for !q.IsEmpty() {
		val, _ := q.Dequeue()
		v := val.(int)
		for _, w := range G.Adj(v) {
			if !bfp.marked[w] {
				bfp.edgeTo[w] = v
				bfp.distTo[w] = bfp.distTo[v] + 1
				bfp.marked[w] = true
				q.Enqueue(w)
			}
		}
	}
}

// HasPathTo returns true if there is a path between the source vertex s (or sources) and vertex v.
func (bfp *BreadthFirstDirectedPaths) HasPathTo(v int) bool {
	bfp.validateVertex(v)
	return bfp.marked[v]
}

// DistTo returns the number of edges in a shortest path between the source vertex s (or sources) and vertex v
func (bfp *BreadthFirstDirectedPaths) DistTo(v int) int {
	bfp.validateVertex(v)
	return bfp.distTo[v]
}

// PathTo returns a shortest path between the source vertex s (or sources) and v, or nil if no such path.
func (bfp *BreadthFirstDirectedPaths) PathTo(v int) (p []int) {
	bfp.validateVertex(v)
	if !bfp.HasPathTo(v) {
		return nil
	}
	path := arraystack.New()
	x := v
	for ; bfp.distTo[x] != 0; x = bfp.edgeTo[x] {
		path.Push(x)
	}
	path.Push(x)
	for _, val := range path.Values() {
		p = append(p, val.(int))
	}
	return p
}

func (bfp *BreadthFirstDirectedPaths) validateVertex(v int) {
	length := len(bfp.marked)
	if v < 0 || v >= length {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", length-1))
	}
}

func (bfp *BreadthFirstDirectedPaths) validateVertices(vertices []int) {
	if vertices == nil {
		panic("argument is nil")
	}
	for _, v := range vertices {
		bfp.validateVertex(v)
	}
}
