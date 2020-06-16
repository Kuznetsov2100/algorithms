package graph

import (
	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

type EulerianPath struct {
	path *arraystack.Stack
}

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

func NewEulerianPath(G *Graph) *EulerianPath {
	oddDegreeVertices := 0
	s := nonIsolatedVertex(G)
	for v := 0; v < G.V(); v++ {
		if G.Degree(v)%2 != 0 {
			oddDegreeVertices++
			s = v
		}
	}
	if oddDegreeVertices > 2 {
		return nil
	}
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
				if selfLoops%2 == 0 {
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

	ep := &EulerianPath{path: arraystack.New()}
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
		ep.path.Push(v)
	}

	if ep.path.Size() != G.E()+1 {
		ep.path = nil
	}
	return ep
}

func (ep *EulerianPath) Path() (p []int) {
	for _, val := range ep.path.Values() {
		p = append(p, val.(int))
	}
	return p
}

func (ep *EulerianPath) HasEulerianPath() bool {
	return ep.path != nil
}
