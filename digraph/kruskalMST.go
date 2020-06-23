package digraph

import (
	"github.com/handane123/algorithms/dataStructure/priorityqueue"
	"github.com/handane123/algorithms/dataStructure/queue/arrayqueue"
	"github.com/handane123/algorithms/fundamentals/unionfind"
)

type KruskalMST struct {
	weight float64
	mst    *arrayqueue.Queue
}

func NewKruskalMST(G *EdgeWeightedGraph) *KruskalMST {
	kr := &KruskalMST{
		mst: arrayqueue.New(),
	}
	pq := priorityqueue.NewMinPQ()
	for _, e := range G.Edges() {
		pq.Insert(e)
	}

	uf := unionfind.NewUF(G.V())
	for !pq.IsEmpty() && kr.mst.Size() < (G.V()-1) {
		e, _ := pq.DelMin()
		edge := e.(*Edge)
		v := edge.Either()
		w := edge.Other(v)
		if uf.Find(v) != uf.Find(w) {
			uf.Union(v, w)
			kr.mst.Enqueue(e)
			kr.weight += edge.Weight()
		}
	}

	return kr
}

func (kr *KruskalMST) Edges() (edges []*Edge) {
	for _, x := range kr.mst.Values() {
		edges = append(edges, x.(*Edge))
	}
	return edges
}

func (kr *KruskalMST) Weight() float64 {
	return kr.weight
}
