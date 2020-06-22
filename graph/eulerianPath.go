package graph

// Eulerian Path
// An undirected graph has Eulerian Path if following two conditions are true.
// (a) All vertices with non-zero degree are connected.
//     We don’t care about vertices with zero degree
//     because they don’t belong to Eulerian Path (we only consider all edges).
//
// (b) If zero or two vertices have odd degree and all other vertices have even degree.
//     Note that only one vertex with odd degree is not possible in an undirected graph
//     (sum of all degrees is always even in an undirected graph)

import (
	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

// EulerianPath struct represents a data type for finding an Eulerian path in a graph.
// An Eulerian path is a path (not necessarily simple) that uses every edge in the graph exactly once.
// This implementation uses a nonrecursive depth-first search.
// The constructor takes O(E + V) time in the worst case,
// where E is the number of edges and V is the number of vertices.
// Each instance method takes O(1) time. It uses O(E + V) extra space in the
// worst case (not including the graph).
type EulerianPath struct {
	path *arraystack.Stack // Eulerian path; null if no suh path
}

// an undirected edge, with a field to indicate whether the edge has already been used
type epEdge struct {
	v      int
	w      int
	isUsed bool
}

func newepEdge(v, w int) *epEdge {
	return &epEdge{
		v:      v,
		w:      w,
		isUsed: false,
	}
}

// returns the other vertex of the edge
func (e *epEdge) other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	if vertex == e.w {
		return e.v
	}
	panic("Illegal endpoint")
}

// NewEulerianPath computes an Eulerian path in the specified graph, if one exists.
func NewEulerianPath(G *Graph) *EulerianPath {
	ep := &EulerianPath{}
	oddDegreeVertices := 0
	s := nonIsolatedVertex(G)
	for v := 0; v < G.V(); v++ {
		if G.Degree(v)%2 != 0 {
			oddDegreeVertices++
			s = v
		}
	}
	if oddDegreeVertices > 2 {
		return ep
	}

	// graph has no nonisolated vertex
	if s == -1 {
		s = 0
	}
	adj := make([]*arrayqueue.Queue, G.V())
	for v := 0; v < G.V(); v++ {
		adj[v] = arrayqueue.New()
	}

	for v := 0; v < G.V(); v++ {
		selfLoops := 0
		for _, w := range G.Adj(v) {
			if v == w {
				if selfLoops%2 == 0 { // add only one copy of each self loop (self loops will be consecutive)
					e := newepEdge(v, w)
					adj[v].Enqueue(e)
					adj[w].Enqueue(e)
				}
				selfLoops++
			} else if v < w {
				e := newepEdge(v, w)
				adj[v].Enqueue(e)
				adj[w].Enqueue(e)
			}
		}
	}

	stack := arraystack.New()
	stack.Push(s)

	ep.path = arraystack.New()
	// greedily search through edges in iterative DFS style
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		v := val.(int)
		for !adj[v].IsEmpty() {
			e, _ := adj[v].Dequeue()
			edge := e.(*epEdge)
			if edge.isUsed {
				continue
			}
			edge.isUsed = true
			stack.Push(v)
			v = edge.other(v)
		}
		// push vertex with no more leaving edges to path
		ep.path.Push(v)
	}

	// check if all edges are used
	if ep.path.Size() != G.E()+1 {
		ep.path = nil
	}
	return ep
}

// Path returns the sequence of vertices on an Eulerian path.
func (ep *EulerianPath) Path() (p []int) {
	for _, val := range ep.path.Values() {
		p = append(p, val.(int))
	}
	return p
}

// HasEulerianPath returns true if the graph has an Eulerian path.
func (ep *EulerianPath) HasEulerianPath() bool {
	return ep.path != nil
}
