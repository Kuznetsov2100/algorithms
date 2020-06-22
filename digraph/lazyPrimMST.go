package digraph

import (
	"github.com/handane123/algorithms/dataStructure/priorityqueue"
	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
)

type LazyPrimMST struct {
	weight float64
	mst    *arrayqueue.Queue
	marked []bool
	pq     *priorityqueue.MinPQ
}

func NewLazyPrimMST(G *EdgeWeightedGraph) *LazyPrimMST {
	lp := &LazyPrimMST{
		mst:    arrayqueue.New(),
		pq:     priorityqueue.NewMinPQ(),
		marked: make([]bool, G.V()),
	}
	for v := 0; v < G.V(); v++ {
		if !lp.marked[v] {
			lp.prim(G, v)
		}
	}
	return lp
}

func (lp *LazyPrimMST) prim(G *EdgeWeightedGraph, s int) {
	lp.scan(G, s)
	for !lp.pq.IsEmpty() {
		e, _ := lp.pq.DelMin()
		edge := e.(*Edge)
		v := edge.Either()
		w := edge.Other(v)
		if lp.marked[v] && lp.marked[w] {
			continue
		}
		lp.mst.Enqueue(edge)
		lp.weight += edge.Weight()
		if !lp.marked[v] {
			lp.scan(G, v)
		}
		if !lp.marked[w] {
			lp.scan(G, w)
		}
	}
}

func (lp *LazyPrimMST) scan(G *EdgeWeightedGraph, v int) {
	lp.marked[v] = true
	for _, e := range G.Adj(v) {
		if !lp.marked[e.Other(v)] {
			lp.pq.Insert(e)
		}
	}
}

func (lp *LazyPrimMST) Edges() (edges []*Edge) {
	for _, x := range lp.mst.Values() {
		edges = append(edges, x.(*Edge))
	}
	return edges
}

func (lp *LazyPrimMST) Weight() float64 {
	return lp.weight
}
