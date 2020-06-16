package graph

import (
	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

type EulerianCycle struct {
	cycle *arraystack.Stack
}

type pEdge struct {
	v      int
	w      int
	isUsed bool
}

func newpEdge(v, w int) *pEdge {
	return &pEdge{
		v:      v,
		w:      w,
		isUsed: false,
	}
}

func (e *pEdge) other(vertex int) int {
	if vertex == e.v {
		return e.w
	}
	if vertex == e.w {
		return e.v
	}
	panic("Illegal endpoint")
}

func NewEulerianCycle(G *Graph) *EulerianCycle {
	if G.E() == 0 {
		return nil
	}
	for v := 0; v < G.V(); v++ {
		if G.Degree(v)%2 != 0 {
			return nil
		}
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
					e := newpEdge(v, w)
					adj[v].Enqueue(e)
					adj[w].Enqueue(e)
				}
				selfLoops++
			} else if v < w {
				e := newpEdge(v, w)
				adj[v].Enqueue(e)
				adj[w].Enqueue(e)
			}
		}
	}
	ec := &EulerianCycle{}
	s := nonIsolatedVertex(G)
	stack := arraystack.New()
	stack.Push(s)

	ec.cycle = arraystack.New()
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		v := val.(int)
		for !adj[v].IsEmpty() {
			e, _ := adj[v].Dequeue()
			edge := e.(*pEdge)
			if edge.isUsed {
				continue
			}
			edge.isUsed = true
			stack.Push(v)
			v = edge.other(v)
		}
		ec.cycle.Push(v)
	}
	if ec.cycle.Size() != G.E()+1 {
		ec.cycle = nil
	}
	return ec
}

func (ec *EulerianCycle) GetCycle() (cy []int) {
	for _, val := range ec.cycle.Values() {
		cy = append(cy, val.(int))
	}
	return cy
}

func (ec *EulerianCycle) HasEulerianCycle() bool {
	return ec.cycle != nil
}

func nonIsolatedVertex(G *Graph) int {
	for v := 0; v < G.V(); v++ {
		if G.Degree(v) > 0 {
			return v
		}
	}
	return -1
}
