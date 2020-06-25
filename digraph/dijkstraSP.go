package digraph

import (
	"fmt"
	"math"

	"github.com/handane123/algorithms/dataStructure/priorityqueue"
	"github.com/handane123/algorithms/dataStructure/stack/arraystack"
)

type DijkstraSP struct {
	distTo []double
	edgeTo []*DirectedEdge
	pq     *priorityqueue.IndexMinPQ
}

func NewDijkstraSP(G *EdgeWeightedDigraph, s int) *DijkstraSP {
	for _, e := range G.Edges() {
		if e.Weight() < 0 {
			panic(fmt.Sprintln("edge ", e, " has negative weight"))
		}
	}

	sp := &DijkstraSP{
		distTo: make([]double, G.V()),
		edgeTo: make([]*DirectedEdge, G.V()),
		pq:     priorityqueue.NewIndexMinPQ(G.V()),
	}
	sp.validateVertex(s)
	for v := 0; v < G.V(); v++ {
		sp.distTo[v] = math.MaxFloat64
	}

	sp.distTo[s] = 0.0
	//nolint:errcheck
	sp.pq.Insert(s, sp.distTo[s])
	for !sp.pq.IsEmpty() {
		v, _ := sp.pq.DelMin()
		for _, e := range G.Adj(v) {
			sp.relax(e)
		}
	}
	return sp
}

func (sp *DijkstraSP) relax(e *DirectedEdge) {
	v := e.From()
	w := e.To()
	if sp.distTo[w] > sp.distTo[v]+double(e.Weight()) {
		sp.distTo[w] = sp.distTo[v] + double(e.Weight())
		sp.edgeTo[w] = e
		if sp.pq.Contains(w) {
			//nolint:errcheck
			sp.pq.DecreaseKey(w, sp.distTo[w])
		} else {
			//nolint:errcheck
			sp.pq.Insert(w, sp.distTo[w])
		}
	}
}

func (sp *DijkstraSP) DistTo(v int) float64 {
	sp.validateVertex(v)
	return float64(sp.distTo[v])
}

func (sp *DijkstraSP) HasPathTo(v int) bool {
	sp.validateVertex(v)
	return sp.distTo[v] < math.MaxFloat64
}

func (sp *DijkstraSP) PathTo(v int) (edges []*DirectedEdge) {
	sp.validateVertex(v)
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

func (sp *DijkstraSP) validateVertex(v int) {
	V := len(sp.distTo)
	if v < 0 || v >= V {
		panic(fmt.Sprintln("vertex ", v, " is not between 0 and ", V-1))
	}
}
