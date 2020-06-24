package digraph

import (
	"github.com/handane123/algorithms/dataStructure/priorityqueue"
	"github.com/handane123/algorithms/fundamentals/unionfind"
)

type KruskalMST struct {
	weight float64
	mst    []*Edge
}

func NewKruskalMST(G *EdgeWeightedGraph) *KruskalMST {
	kr := &KruskalMST{}
	pq := priorityqueue.NewMinPQ()
	for _, e := range G.Edges() {
		pq.Insert(e)
	}

	uf := unionfind.NewUF(G.V())
	for !pq.IsEmpty() && len(kr.mst) < (G.V()-1) {
		e, _ := pq.DelMin()
		edge := e.(*Edge)
		v := edge.Either()
		w := edge.Other(v)
		if uf.Find(v) != uf.Find(w) {
			uf.Union(v, w)
			kr.mst = append(kr.mst, edge)
			kr.weight += edge.Weight()
		}
	}

	return kr
}

func (kr *KruskalMST) Edges() (edges []*Edge) {
	edges = make([]*Edge, len(kr.mst))
	copy(edges, kr.mst)
	return edges
}

func (kr *KruskalMST) Weight() float64 {
	return kr.weight
}
