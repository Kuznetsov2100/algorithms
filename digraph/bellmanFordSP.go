package digraph

import (
	"fmt"
	"math"

	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

type BellmanFordSP struct {
	distTo  []float64
	edgeTo  []*DirectedEdge
	onQueue []bool
	queue   *arrayqueue.Queue
	cost    int
	cycle   []*DirectedEdge
}

func NewBellmanFordSP(G *EdgeWeightedDigraph, s int) *BellmanFordSP {
	sp := &BellmanFordSP{
		distTo:  make([]float64, G.V()),
		edgeTo:  make([]*DirectedEdge, G.V()),
		onQueue: make([]bool, G.V()),
	}
	for v := 0; v < G.V(); v++ {
		sp.distTo[v] = math.Inf(1)
	}
	sp.distTo[s] = 0.0

	sp.queue = arrayqueue.New()
	sp.queue.Enqueue(s)
	sp.onQueue[s] = true
	for !sp.queue.IsEmpty() && !sp.HasNegativeCycle() {
		val, _ := sp.queue.Dequeue()
		v := val.(int)
		sp.onQueue[v] = false
		sp.relax(G, v)
	}
	return sp
}

func (sp *BellmanFordSP) relax(G *EdgeWeightedDigraph, v int) {
	for _, e := range G.Adj(v) {
		w := e.To()
		if sp.distTo[w] > sp.distTo[v]+e.Weight() {
			sp.distTo[w] = sp.distTo[v] + e.Weight()
			sp.edgeTo[w] = e
			if !sp.onQueue[w] {
				sp.queue.Enqueue(w)
				sp.onQueue[w] = true
			}
		}
		if sp.cost++; sp.cost%G.V() == 0 {
			sp.FindNegativeCycle()
			if sp.HasNegativeCycle() {
				return
			}
		}
	}
}

func (sp *BellmanFordSP) HasNegativeCycle() bool {
	return sp.cycle != nil
}

func (sp *BellmanFordSP) FindNegativeCycle() {
	V := len(sp.edgeTo)
	spt := NewEdgeWeightedDigraphV(V)
	for v := 0; v < V; v++ {
		if sp.edgeTo[v] != nil {
			spt.AddEdge(sp.edgeTo[v])
		}
	}
	finder := NewEdgeWeightedDirectedCycle(spt)
	sp.cycle = finder.GetCycle()
}

func (sp *BellmanFordSP) NegativeCycle() (edges []*DirectedEdge) {
	if !sp.HasNegativeCycle() {
		return nil
	}
	edges = make([]*DirectedEdge, len(sp.cycle))
	copy(edges, sp.cycle)
	return edges
}

func (sp *BellmanFordSP) DistTo(v int) float64 {
	sp.validateVertex(v)
	return sp.distTo[v]
}

func (sp *BellmanFordSP) HasPathTo(v int) bool {
	sp.validateVertex(v)
	return sp.distTo[v] < math.Inf(1)
}

func (sp *BellmanFordSP) PathTo(v int) (edges []*DirectedEdge) {
	sp.validateVertex(v)
	if sp.HasNegativeCycle() {
		panic("Negative cost cycle exists")
	}
	if !sp.HasPathTo(v) {
		return nil
	}

	path := arraystack.New()
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()] {
		path.Push(e)
	}

	for _, e := range path.Values() {
		edges = append(edges, e.(*DirectedEdge))
	}
	return edges
}

func (sp *BellmanFordSP) validateVertex(v int) {
	V := len(sp.distTo)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
