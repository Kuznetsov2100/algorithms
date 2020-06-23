package digraph

import (
	"math"

	"github.com/handane123/algorithms/dataStructure/priorityqueue"
)

type double float64

func (d double) CompareTo(k priorityqueue.Key) int {
	that := k.(double)
	if d < that {
		return -1
	} else if d > that {
		return 1
	} else {
		return 0
	}
}

type PrimMST struct {
	edgeTo []*Edge
	distTo []double
	marked []bool
	pq     *priorityqueue.IndexMinPQ
}

func NewPrimMST(G *EdgeWeightedGraph) *PrimMST {
	mst := &PrimMST{
		edgeTo: make([]*Edge, G.V()),
		distTo: make([]double, G.V()),
		marked: make([]bool, G.V()),
		pq:     priorityqueue.NewIndexMinPQ(G.V()),
	}
	for i := range mst.distTo {
		mst.distTo[i] = math.MaxFloat64
	}

	for v := 0; v < G.V(); v++ {
		if !mst.marked[v] {
			mst.prim(G, v)
		}
	}
	return mst
}

func (mst *PrimMST) prim(G *EdgeWeightedGraph, s int) {
	mst.distTo[s] = 0.0
	//nolint:errcheck
	mst.pq.Insert(s, mst.distTo[s])
	for !mst.pq.IsEmpty() {
		v, _ := mst.pq.DelMin()
		mst.scan(G, v)
	}
}

func (mst *PrimMST) scan(G *EdgeWeightedGraph, v int) {
	mst.marked[v] = true
	for _, e := range G.Adj(v) {
		w := e.Other(v)
		if mst.marked[w] {
			continue
		}
		if e.Weight() < float64(mst.distTo[w]) {
			mst.distTo[w] = double(e.Weight())
			mst.edgeTo[w] = e
			if mst.pq.Contains(w) {
				//nolint:errcheck
				mst.pq.DecreaseKey(w, mst.distTo[w])
			} else {
				//nolint:errcheck
				mst.pq.Insert(w, mst.distTo[w])
			}
		}
	}
}

func (mst *PrimMST) Edges() (edges []*Edge) {
	for v := 0; v < len(mst.edgeTo); v++ {
		if e := mst.edgeTo[v]; e != nil {
			edges = append(edges, e)
		}
	}
	return edges
}

func (mst *PrimMST) Weight() float64 {
	weight := 0.0
	for _, e := range mst.Edges() {
		weight += e.Weight()
	}
	return weight
}
