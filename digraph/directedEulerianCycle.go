package digraph

import (
	"github.com/emirpasic/gods/lists/arraylist"

	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

type DirectedEulerianCycle struct {
	cycle *arraystack.Stack
}

func NewDirectedEulerianCycle(G *Digraph) *DirectedEulerianCycle {
	dec := &DirectedEulerianCycle{}
	if G.E() == 0 {
		return dec
	}

	for v := 0; v < G.V(); v++ {
		if G.OutDegree(v) != G.InDegree(v) {
			return dec
		}
	}

	adj := make([]arraylist.Iterator, G.V())
	for v := 0; v < G.V(); v++ {
		vertices := arraylist.New()
		for _, w := range G.Adj(v) {
			vertices.Add(w)
		}
		adj[v] = vertices.Iterator()
	}

	s := nonIsolatedVertex(G)
	stack := arraystack.New()
	stack.Push(s)

	dec.cycle = arraystack.New()
	for !stack.IsEmpty() {
		val, _ := stack.Pop()
		v := val.(int)
		for adj[v].Next() {
			stack.Push(v)
			v = adj[v].Value().(int)
		}
		dec.cycle.Push(v)
	}

	if dec.cycle.Size() != G.E()+1 {
		dec.cycle = nil
	}

	return dec
}

// GetCycle returns the sequence of vertices on an Eulerian cycle.
func (dec *DirectedEulerianCycle) GetCycle() (cy []int) {
	for _, val := range dec.cycle.Values() {
		cy = append(cy, val.(int))
	}
	return cy
}

// HasEulerianCycle returns true if the graph has an Eulerian cycle.
func (dec *DirectedEulerianCycle) HasEulerianCycle() bool {
	return dec.cycle != nil
}

func nonIsolatedVertex(G *Digraph) int {
	s := -1
	for v := 0; v < G.V(); v++ {
		if G.OutDegree(v) > 0 {
			s = v
			break
		}
	}
	return s
}
